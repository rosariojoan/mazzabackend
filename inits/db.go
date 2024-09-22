package inits

import (
	"context"
	"log"
	ent "mazza/ent/generated"
	"mazza/ent/generated/migrate"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

var DB *ent.Client

/* Instantiate DB and run auto migration */
func InitDB(ctx context.Context) {
	db, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=arafate dbname=mazza2 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to the db: %v", err)
	}
	// defer db.Close()

	// Auto migration
	if err := db.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	DB = db
}
