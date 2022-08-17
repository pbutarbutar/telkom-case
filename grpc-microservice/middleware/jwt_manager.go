package middleware

import (
	"fmt"
	"grpc-microservice/config"
	"time"

	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
)

type JWTManagerRepo interface {
	Generate(user *User) (string, error)
	Verify(accessToken string) (*UserClaims, error)
}

// JWTManager is a JSON web token manager
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
	Logger        *zap.Logger
}

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.StandardClaims
	Jti          string       `json:"username"`
	UserID       string       `json:"user_id"`
	HumanID      string       `json:"human_id"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	FullName     string       `json:"full_name"`
	Email        string       `json:"email"`
	IsActive     bool         `json:"is_active"`
	IsStaff      bool         `json:"is_staff"`
	IsSuperuser  bool         `json:"is_superuser"`
	PhoneNumber  string       `json:"phone_number"`
	Groups       []string     `json:"groups"`
	Organization Organization `json:"organization"`
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(jw *config.ConfigJWTEnv, logger *zap.Logger) *JWTManager {
	secretKey := jw.GetSecretVal()
	tokenDurationX := jw.GetDurationRefreshVal()
	return &JWTManager{
		secretKey,
		tokenDurationX,
		logger,
	}
}

// Generate generates and signs a new token for a user
func (manager *JWTManager) Generate(user *User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		UserID:  user.UserID,
		HumanID: user.HumanID,
		Email:   user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				manager.Logger.Info("unexpected token signing method")
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		manager.Logger.Info(err.Error())
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		manager.Logger.Info("invalid token claims")
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
