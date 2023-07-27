package main

import (
	"common/db"
	"common/repos"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"api/graph"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbConn, err := db.SetupDatabaseConn()
	if err != nil {
		panic(err)
	}
	repo := repos.NewAlbumRepo(dbConn)

	/* init graphql server
		NewDefaultServer creates a new server and expects an executable graphql schema as the argument
		NewExecutableSchema creates the schema based on the graphql config
		graph.Config is an object where the resolvers are specified -> AlbumRepo to interact with the db
	*/
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{AlbumRepo: repo}}))

	r := gin.Default()
	// r.Use(utils.GinContextToContextMiddleware()) // middleware that extracts contet from gin.context to context.Context, in order to use gin.Context
	// c.Writer the data requested by a client, in the form of a JSON object
	// c.Request in graphql contains the graphql query
	r.POST("/query", func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }) 
	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}
