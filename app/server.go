package main

import (
	"context"
	"log"
	"mazza"
	"mazza/app/auth"
	"mazza/app/middlewares"
	"mazza/inits"

	"os"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

var ctx = context.Background()

func main() {
	defer inits.DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := gin.Default()
	router.Use(middlewares.GinContextToContextMiddleware())

	// This router group do not require the user to be logged in
	authRouter := router.Group("/account")
	{
		authRouter.POST("/signup")
		authRouter.POST("/login", auth.Login)
		authRouter.POST("/forgot-password")
		authRouter.POST("/verify-password")

	}

	gqlRouter := router.Group("/", middlewares.LoginRequired)
	{
		gqlRouter.POST("query", graphqlHandler())
		gqlRouter.GET("", playgroundHandler())
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	router.Run()
}

func init() {
	inits.LoadEnv()
	inits.InitDB(ctx)
	inits.InitFirebase()
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(mazza.NewSchema(inits.DB))
	srv.Use(entgql.Transactioner{TxOpener: inits.DB})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("Mazza GraphQL", "/query")

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
