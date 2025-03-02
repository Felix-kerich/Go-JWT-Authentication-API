package services

import (
	"errors"
	"time"

	"auth_api/models"
	"auth_api/repositories"
	"auth_api/utils"
)

// AuthService struct
type AuthService struct {
	userRepo     *repositories.UserRepository
	emailService *EmailService
}

// NewAuthService initializes AuthService
func NewAuthService(userRepo *repositories.UserRepository, emailService *EmailService) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		emailService: emailService,
	}
}

// RegisterUser registers a new user and sends a verification email
func (s *AuthService) RegisterUser(user *models.User) error {
	// Check if email already exists
	_, isDeleted, err := s.userRepo.GetUserByEmail(user.Email)
	if err == nil {
		if isDeleted {
			return errors.New("user is deleted, please recover it first")
		}
		return errors.New("email already exists")
	}

	// Check if username already exists
	if _, err := s.userRepo.GetUserByUsername(user.Username); err == nil {
		return errors.New("username already exists")
	}

	// Hash password
	if err := user.HashPassword(); err != nil {
		return err
	}

	// Save user
	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}

	// Generate email verification JWT
	verificationToken := utils.GenerateJWT(user.Email, "email_verification", time.Hour*24)

	// Send verification email
	return s.emailService.SendVerificationEmail(user.Email, verificationToken)
}

// VerifyEmail verifies a user's email using the JWT token
func (s *AuthService) VerifyEmail(token string) error {
	email, err := utils.ValidateJWT(token, "email_verification")
	if err != nil {
		return errors.New("invalid or expired verification token")
	}

	// Retrieve user by email
	user, isDeleted, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	if isDeleted {
		return errors.New("user is deleted, please recover it first")
	}

	// Mark email as verified
	user.EmailVerified = true
	return s.userRepo.UpdateUser(user)
}

// LoginUser authenticates a user and returns a JWT token
func (s *AuthService) LoginUser(email, password string) (string, error) {
	user, isDeleted, err := s.userRepo.GetUserByEmail(email)
	if err != nil || isDeleted || !user.ComparePassword(password) {
		return "", errors.New("invalid credentials")
	}

	// Check if email is verified
	if !user.EmailVerified {
		return "", errors.New("email not verified")
	}

	// Generate JWT token for authentication
	return utils.GenerateJWT(user.Email, "access", time.Hour*24), nil
}
