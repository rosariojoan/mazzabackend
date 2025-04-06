package inits

import (
	"context"
	"database/sql"
	"log"
	ent "mazza/ent/generated"
	"mazza/ent/generated/migrate"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Client *ent.Client

/* Instantiate DB and run auto migration */
func InitDB(ctx context.Context) {
	// db, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=arafate dbname=mazza2 sslmode=disable")
	// os.Getenv("DATABASE_URL")
	// Open new connection
	databaseUrl := os.Getenv("DB_URL")
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatalf("failed opening connection to the db: %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	Client = ent.NewClient(ent.Driver(drv))

	// Auto migration
	if err := Client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}
}
