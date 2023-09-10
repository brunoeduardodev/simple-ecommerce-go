package views

import (
	"time"

	"github.com/brunoeduardodev/simple-ecommerce-go/internal/models"
)

// User Information
// @Description Information of a user
type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"string"` // Created at here
}

func DbUserToView(user *models.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

func DbUsersToView(dbUsers []*models.User) []User {
	users := []User{}

	for i := 0; i < len(dbUsers); i++ {
		users = append(users, DbUserToView(dbUsers[i]))
	}

	return users
}
