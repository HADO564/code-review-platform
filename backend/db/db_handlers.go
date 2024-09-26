package pkg

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"

    _ "github.com/lib/pq"
	"github.com/yourusername/code-review-platform/graph/model"
	"github.com/yourusername/code-review-platform/db"
)

// Connect to the database
func connectDB() (*sql.DB, error) {
    connStr := "postgres://username:password@localhost/dbname?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to the database: %v", err)
    }
    return db, nil
}

// Create user route
func createUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input model.NewUser

        // Decode the JSON body into a Go struct
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        // Create the new user using the functions package
        newUser := &model.User{
            Username: input.Username,
            Email:    input.Email,
        }

        err = functions.CreateUser(r.Context(), db, newUser)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
            return
        }

        // Respond with the created user in JSON format
        json.NewEncoder(w).Encode(newUser)
    }
}

// Get user by email route
func getUserByEmailHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        email := r.URL.Query().Get("email")
        if email == "" {
            http.Error(w, "Missing email parameter", http.StatusBadRequest)
            return
        }

        user, err := functions.GetUserByEmail(r.Context(), db, email)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error fetching user: %v", err), http.StatusInternalServerError)
            return
        }

        // Respond with the fetched user in JSON format
        json.NewEncoder(w).Encode(user)
    }
}
