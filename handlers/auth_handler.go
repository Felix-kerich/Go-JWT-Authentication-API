package handlers

import (
	"auth_api/models"
	"auth_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService  *services.AuthService
	emailService *services.EmailService
}

func NewAuthHandler(authService *services.AuthService, emailService *services.EmailService) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		emailService: emailService,
	}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Email:    req.Email,
		Username: req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := h.authService.RegisterUser(user); err != nil {
		switch err.Error() {
		case "email already exists":
			c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		case "username already exists":
			c.JSON(http.StatusConflict, gin.H{"error": "Username is already registered"})
		case "user is deleted, please recover it first":
			c.JSON(http.StatusConflict, gin.H{"error": "User is deleted, please recover it first"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered. Please verify your email."})
}

func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token"})
		return
	}

	if err := h.authService.VerifyEmail(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
