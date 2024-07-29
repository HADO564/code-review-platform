package database

import (
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
	"github.com/joho/godotenv"
)

var client *supabase.Client

func init() {
	// Load environment variables
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("SUPABASE_KEY")
	apiUrl := os.Getenv("SUPABASE_URL")

	var err error
	client, err = supabase.NewClient(apiUrl, apiKey, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetClient() *supabase.Client {
	return client
}