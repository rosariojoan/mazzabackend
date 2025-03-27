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

	_ "mazza/ent/generated/runtime" // to avoid import cycle error. See: entgo.io/docs/hooks
)

const defaultPort = "8000"

// const defaultPort = "8080"

var ctx = context.Background()

func main() {
	defer inits.Client.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := gin.Default()

	// Security headers
	router.Use(func(ctx *gin.Context) {
		// if ctx.Request.Host != expectedHost {
		// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid host header"})
		// 	return
		// }
		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("X-XSS-Protection", "1; mode=block")
		ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		ctx.Header("Preferrer-Policy", "strict-origin")
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Next()
	})

	router.Use(middlewares.GinContextToContextMiddleware())

	// This router group do not require the user to be logged in
	authRouter := router.Group("/account")
	{
		authRouter.POST("/signup", auth.Signup)
		authRouter.POST("/signup-invited-user", auth.SignupInvitedUser)
		authRouter.POST("/verify-invitation-token", auth.VerifyMemberInvitationToken)
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
	srv := handler.NewDefaultServer(mazza.NewSchema(inits.Client))
	srv.Use(entgql.Transactioner{TxOpener: inits.Client})

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
