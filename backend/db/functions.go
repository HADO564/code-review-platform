package pkg

import (
    "context"
    "database/sql"
    _ "github.com/lib/pq"
	"github.com/yourusername/code-review-platform/graph/model"
)

// Insert the new user into the database
func CreateUser(ctx context.Context, db *sql.DB, user *model.User) error {
    query := `INSERT INTO users (username, email) VALUES ($1, $2)`
    _, err := db.ExecContext(ctx, query, user.Username, user.Email)
    return err
}

// Fetch the user by email (after creation)
func GetUserByEmail(ctx context.Context, db *sql.DB, email string) (*model.User, error) {
    query := `SELECT id, username, email, created_at, updated_at FROM users WHERE email = $1`
    row := db.QueryRowContext(ctx, query, email)

    var user model.User
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
