package pkg

// Insert the new user into the database
import (
	"context"
	"fmt"

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

// GetUserByEmail retrieves a user from the Supabase database by email
func GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	_, count, err := supabase.Client.From("users").
		Select("*","exact",false).
		Eq("email", email).
		Single().
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}
