package utils

import (
	"fmt"
	"golang_restaurant_backend/db"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(), /* Access token with 1 day validity */
		})

	tokenString, err := token.SignedString(db.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(email string) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(), /* Refresh token with 7 days validity  */
		})

	refreshTokenString, err := refreshToken.SignedString(db.RefreshSecretKey)
	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}

/* Validating Signing method and expiry date */
func ValidateToken(token string) (*jwt.Token, error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid token")
		}
		return db.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		expireTime := int64(claims["exp"].(float64))
		if expireTime < time.Now().Unix() {
			return nil, fmt.Errorf("token has expired")
		}
		return tk, nil
	}

	return nil, fmt.Errorf("invalid token")
}
