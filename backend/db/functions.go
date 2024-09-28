package pkg

// Insert the new user into the database
import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/yourusername/code-review-platform/graph/model"
	"github.com/yourusername/code-review-platform/internal/supabase"
)

// CreateUser creates a new user in the Supabase database
func CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	newUser := model.NewUser{
		Username: input.Username,
		Email:    input.Email,
	}

	// Insert the new user into the database
	_, count, err := supabase.Client.From("users").Insert(newUser,false,"","","exact").Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("user was not created")
	}

	// Fetch the created user to return (including auto-generated fields)
	var createdUser model.User
	_, _, err = supabase.Client.From("users").
		Select("*","exact",false).
		Eq("email", input.Email).
		Single().
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created user: %w", err)
	}

	return &createdUser, nil
}

func GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	responseBytes, count, err := supabase.Client.From("users").
		Select("*", "exact", false).
		Eq("email", email).
		Single().
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("user not found")
	}
	
	// Unmarshal the response data into the user struct
	err = json.Unmarshal(responseBytes, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	// Now that user is populated, marshal it to JSON for printing
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data to JSON: %w", err)
	}
	
	// Print the JSON data
	fmt.Println(string(jsonData))

	// Log the createdAt and updatedAt fields
	log.Printf("User createdAt: %s, updatedAt: %s", user.CreatedAt, user.UpdatedAt)
	return &user, nil
}
