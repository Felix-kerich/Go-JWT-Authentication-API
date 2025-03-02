package repositories

import (
	"auth_api/database"
	"auth_api/models"

	"gorm.io/gorm"
)

// AuthRepository struct
type AuthRepository struct {
	DB *gorm.DB
}

// NewAuthRepository initializes AuthRepository
func NewAuthRepository() *AuthRepository {
	return &AuthRepository{DB: database.DB}
}

// GetUserByToken finds a user by verification token
func (repo *AuthRepository) GetUserByToken(token string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("verification_token = ?", token).First(&user).Error
	return &user, err
}

// GetUserByEmail finds a user by email
func (repo *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
