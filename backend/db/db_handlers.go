package pkg

import (
    "encoding/json"
    "fmt"
    "net/http"

    _ "github.com/lib/pq"
	"github.com/yourusername/code-review-platform/graph/model"
)

// Connect to the database
// Create user route
func createUserHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input model.NewUser

        // Decode the JSON body into a Go struct
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        // Create the new user using the functions package
        newUser := &model.NewUser{
            Username: input.Username,
            Email:    input.Email,
        }

        user,err := CreateUser(r.Context(), newUser)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
            return
        }

        // Respond with the created user in JSON format
        json.NewEncoder(w).Encode(user)
    }
}

// Get user by email route
func getUserByEmailHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        email := r.URL.Query().Get("email")
        if email == "" {
            http.Error(w, "Missing email parameter", http.StatusBadRequest)
            return
        }

        user, err := GetUserByEmail(r.Context(), email)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error fetching user: %v", err), http.StatusInternalServerError)
            return
        }

        // Respond with the fetched user in JSON format
        json.NewEncoder(w).Encode(user)
    }
}
