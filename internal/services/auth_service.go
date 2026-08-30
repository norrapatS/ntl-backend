package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"ntl-test/backend/internal/models"
	"ntl-test/backend/internal/repositories"
)

type AuthService struct {
	userRepository *repositories.UserRepository
	jwtSecret      string
}

func NewAuthService(
	userRepository *repositories.UserRepository,
	jwtSecret string,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtSecret:      jwtSecret,
	}
}

type RegisterInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func (s *AuthService) Register(input RegisterInput) (*models.User, error) {
	existingUser, err := s.userRepository.FindByEmail(input.Email)

	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        uuid.New(),
		Email:     input.Email,
		Password:  string(hashedPassword),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Role:      "user",
	}

	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

type LoginInput struct {
	Email    string
	Password string
}

func (s *AuthService) Login(
	input LoginInput,
) (string, string, error) {

	user, err := s.userRepository.FindByEmail(input.Email)

	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := s.generateAccessToken(user)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateRefreshToken(user)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) generateAccessToken(
	user *models.User,
) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"type":    "access",
		"role":    user.Role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) generateRefreshToken(
	user *models.User,
) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"type":    "refresh",
		"role":    user.Role,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	user, err := s.userRepository.FindByID(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
