package main

import (
	"backend/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

const defaultPort = "8080"

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	apiKey := os.Getenv("SUPABASE_KEY")
	apiUrl := os.Getenv("SUPABASE_URL")

	client, err := supabase.NewClient(apiUrl, apiKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
