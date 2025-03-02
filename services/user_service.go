package services

import (
	"errors"

	"auth_api/models"
	"auth_api/repositories"
)

// UserService struct
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService initializes UserService
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id uint) error {
	_, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}
	return s.userRepo.DeleteUser(id)
}

// UpdateUser updates a user by ID
func (s *UserService) UpdateUserFields(id uint, updates *models.User) error {
	_, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	// Set the ID in the updates
	updates.ID = id

	return s.userRepo.UpdateUser(updates)
}

// RecoverUser recovers a deleted user
func (s *UserService) RecoverUser(id uint) error {
	// Check if user exists and is deleted
	var user models.User
	if err := s.userRepo.DB.Unscoped().First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	if user.DeletedAt.Time.IsZero() {
		return errors.New("user is not deleted")
	}

	return s.userRepo.RecoverUser(id)
}


