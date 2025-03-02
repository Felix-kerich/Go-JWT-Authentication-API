package repositories

import (
	"auth_api/database"
	"auth_api/models"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// UserRepository struct
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository initializes UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{DB: database.DB}
}

// GetAllUsers retrieves all users from the database
func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Find(&users).Error
	return users, err
}

// GetUserByID finds a usmy-gin-apper by ID
func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	return &user, err
}

// CreateUser saves a new user
func (repo *UserRepository) CreateUser(user *models.User) error {
	err := repo.DB.Create(user).Error
	if err != nil {
		// Check for duplicate entry error
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "users.uni_users_email") {
				return errors.New("email already exists")
			}
			if strings.Contains(err.Error(), "users.uni_users_username") {
				return errors.New("username already exists")
			}
		}
		return err
	}
	return nil
}

// UpdateUser updates an existing user
func (repo *UserRepository) UpdateUser(user *models.User) error {
	return repo.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// DeleteUser removes a user from the database
func (repo *UserRepository) DeleteUser(id uint) error {
	return repo.DB.Delete(&models.User{}, id).Error
}

// GetUserByToken finds a user by verification token
func (repo *UserRepository) GetUserByToken(token string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("verification_token = ?", token).First(&user).Error
	return &user, err
}

// GetUserByEmail finds a user by email
func (repo *UserRepository) GetUserByEmail(email string) (*models.User, bool, error) {
	var user models.User
	err := repo.DB.Unscoped().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, false, err
	}
	isDeleted := !user.DeletedAt.Time.IsZero()
	return &user, isDeleted, nil
}

// Add new method for email verification
func (repo *UserRepository) UpdateEmailVerification(email string, verified bool) error {
	return repo.DB.Model(&models.User{}).Where("email = ?", email).Update("email_verified", verified).Error
}

// GetUserByUsername finds a user by username
func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// RecoverUser restores a soft-deleted user
func (repo *UserRepository) RecoverUser(id uint) error {
	result := repo.DB.Unscoped().Model(&models.User{}).Where("id = ?", id).Update("deleted_at", nil)
	if result.RowsAffected == 0 {
		return errors.New("user not found or already active")
	}
	return result.Error
}

// GetUserByIDWithDeleted finds a user by ID with deleted users included
func (repo *UserRepository) GetUserByIDWithDeleted(id uint) (*models.User, bool, error) {
	var user models.User
	err := repo.DB.Unscoped().First(&user, id).Error
	if err != nil {
		return nil, false, err
	}
	// Check if user is deleted (deleted_at is not null)
	isDeleted := !user.DeletedAt.Time.IsZero()
	return &user, isDeleted, nil
}

