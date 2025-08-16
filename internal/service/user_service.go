package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// userService implements the UserService interface
type userService struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo domain.UserRepository, jwtSecret string) domain.UserService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Register creates a new user with hashed password
func (s *userService) Register(phoneNumber, password string) error {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByPhoneNumber(phoneNumber)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user
	user := &domain.User{
		PhoneNumber: phoneNumber,
		Password:    string(hashedPassword),
	}

	return s.userRepo.Create(user)
}

// Authenticate validates user credentials and returns user if valid
func (s *userService) Authenticate(phoneNumber, password string) (*domain.User, error) {
	user, err := s.userRepo.FindByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetByID retrieves a user by ID
func (s *userService) GetByID(id uint) (*domain.User, error) {
	return s.userRepo.FindByID(id)
}

// GetByPhoneNumber retrieves a user by phone number
func (s *userService) GetByPhoneNumber(phoneNumber string) (*domain.User, error) {
	return s.userRepo.FindByPhoneNumber(phoneNumber)
}

// GenerateJWT generates a JWT token for a user
func (s *userService) GenerateJWT(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   user.PhoneNumber,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
