package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"github.com/hardiksharma/clarityfin-api/internal/dto"
	"github.com/hardiksharma/clarityfin-api/pkg/response"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	userUseCase domain.UserUseCase
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(userUseCase domain.UserUseCase) *AuthHandler {
	return &AuthHandler{
		userUseCase: userUseCase,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.userUseCase.Register(req.PhoneNumber, req.Password)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil, "Registration successful")
}

// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	token, err := h.userUseCase.Login(req.PhoneNumber, req.Password)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	authResponse := dto.AuthResponse{
		Token: token,
	}

	response.Success(c, authResponse, "Login successful")
}
