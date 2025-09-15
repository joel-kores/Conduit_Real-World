package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claim struct {
	UserID string `json:"user_id"`
	Email string	`json:"email"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager{
	return &JWTManager{
		secretKey: secretKey,
		tokenDuration: tokenDuration,
	}
}

func (m *JWTManager) Generate(userID, email string) (string, error) {
	claims := Claim{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	 return token.SignedString([]byte(m.secretKey))
}
