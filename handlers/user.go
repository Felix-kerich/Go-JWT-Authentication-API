package handlers

import (
	
	"auth_api/repositories"
	"auth_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


var userRepo = repositories.NewUserRepository()
var userService = services.NewUserService(userRepo)

// GetUsers - Retrieve all users
func GetUsers(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID - Retrieve a single user by ID
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}



// DeleteUser - Delete a user by ID
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
