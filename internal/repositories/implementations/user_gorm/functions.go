package user_gorm

import (
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/models"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/repositories"
	"gorm.io/gorm"
)

type UserRepositoryGormImplementation struct {
	db *gorm.DB
}

func (r UserRepositoryGormImplementation) Create(input repositories.CreateUserInput) error {
	user := UserGormModel{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	result := r.db.Create(&user)

	return result.Error
}

func (r UserRepositoryGormImplementation) FindByEmail(email string) (*models.User, error) {
	var dbUser UserGormModel
	result := r.db.Where(&UserGormModel{Email: email}).First(&dbUser)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil
		}

		return nil, result.Error
	}

	user := dbUserToModel(dbUser)

	return &user, result.Error

}

func NewUserRepositoryGormIplementation(db *gorm.DB) UserRepositoryGormImplementation {
	return UserRepositoryGormImplementation{db: db}

}
