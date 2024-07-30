package supabase

import (
    "fmt"
    "os"
    "github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func Init() error {
    var err error
    Client, err = supabase.NewClient(os.Getenv("SUPABASE_PUBLIC_URL"), os.Getenv("SUPABASE_PUBLIC_API_KEY"), "", nil)
    if err != nil {
        return fmt.Errorf("cannot initialize client: %w", err)
    }
    return nil
}