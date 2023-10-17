package main

import (
	"event_management/database"
	"event_management/graph"
	resolver "event_management/graph/resolvers"
	"event_management/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	database.DatabaseConnection()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := graph.Config{Resolvers: &resolver.Resolver{}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.CorsMiddleware(middleware.AuthenticationMiddleware(srv)))

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
