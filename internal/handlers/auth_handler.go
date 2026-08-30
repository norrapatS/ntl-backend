package handlers

import (
	"net/http"

	"ntl-test/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	user, err := h.authService.Register(
		services.RegisterInput{
			Email:     request.Email,
			Password:  request.Password,
			FirstName: request.FirstName,
			LastName:  request.LastName,
		},
	)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "registration successful",
		"user":    user,
	})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.authService.Login(
		services.LoginInput{
			Email:    request.Email,
			Password: request.Password,
		},
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetCookie(
		"access_token",
		accessToken,
		60*15,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refresh_token",
		refreshToken,
		60*60*24*7,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func (h *AuthHandler) AuthUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	user, err := h.authService.GetUserByID(userID.(uuid.UUID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"role" : user.Role,
		},
	})
}