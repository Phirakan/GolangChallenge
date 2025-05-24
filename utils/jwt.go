package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		
		secret = "secret_key_develop" 
	}
	return []byte(secret)
}

var jwtSecret = getJWTSecret()

type JWTClaim struct {
	UserID string `json:"user_id"` 
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(userID string, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	
	claims := &JWTClaim{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}


func ValidateToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		},
	)
	
	if err != nil {
		return nil, err
	}
	

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	
	return claims, nil
}

func GetJWTSecret() []byte {
	return jwtSecret
}