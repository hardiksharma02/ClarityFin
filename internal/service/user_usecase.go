package service

import (
	"github.com/hardiksharma/clarityfin-api/internal/domain"
)

// userUseCase implements the UserUseCase interface
type userUseCase struct {
	userService domain.UserService
}

// NewUserUseCase creates a new instance of UserUseCase
func NewUserUseCase(userService domain.UserService) domain.UserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

// Register handles user registration
func (uc *userUseCase) Register(phoneNumber, password string) error {
	return uc.userService.Register(phoneNumber, password)
}

// Login handles user authentication and returns JWT token
func (uc *userUseCase) Login(phoneNumber, password string) (string, error) {
	user, err := uc.userService.Authenticate(phoneNumber, password)
	if err != nil {
		return "", err
	}

	// Cast to concrete type to access GenerateJWT method
	userSvc := uc.userService.(*userService)
	return userSvc.GenerateJWT(user)
}
