package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	database "github.com/amaraad/goapp/db"
	"github.com/amaraad/goapp/graph"
	"github.com/amaraad/goapp/graph/generated"
	"github.com/amaraad/goapp/migrations"
	rest "github.com/amaraad/goapp/rest"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	database.InitDatabase()
	migrations.Migrate()
	r := gin.Default()
	// graphql example
	r.POST("/graphql/query", graphqlHandler())
	r.GET("/graphql/", playgroundHandler())
	// rest example
	r.GET("/rest/request/get", rest.GetRequest)
	r.GET("/rest/request/post", rest.PostRequest)
	r.Run(defaultPort)
}
