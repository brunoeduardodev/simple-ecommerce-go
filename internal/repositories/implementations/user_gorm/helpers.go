package user_gorm

import "github.com/brunoeduardodev/simple-ecommerce-go/internal/models"

func dbUserToModel(dbUser UserGormModel) models.User {
	return models.User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

func dbUsersToModel(dbUsers []UserGormModel) []models.User {
	var users []models.User
	for i := 0; i < len(dbUsers); i++ {
		users = append(users, dbUserToModel(dbUsers[i]))
	}

	return users
}
