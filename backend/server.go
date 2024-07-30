package main

import (
    "log"
    "net/http"
    "os"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/joho/godotenv"
    "github.com/yourusername/code-review-platform/graph"
    "github.com/yourusername/code-review-platform/internal/supabase"
)

const defaultPort = "8080"

func main() {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: Error loading .env file")
    }

    // Initialize Supabase client
    err = supabase.Init()
    if err != nil {
        log.Fatalf("Failed to initialize Supabase client: %v", err)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}