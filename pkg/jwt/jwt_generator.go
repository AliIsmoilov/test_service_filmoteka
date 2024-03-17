package jwt

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"test_service_filmoteka/pkg/constatnts"
	"time"

	"github.com/golang-jwt/jwt"
)

type Tokens struct {
	Access  string
	Refresh string
}

// GenerateNewCustomerTokens func for generate a new Access & Refresh tokens.
// Credentials intentionally added func method signature in case of needed
func GenerateNewUserTokens(id string, credentials map[string]string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewUserAccessToken(id, credentials)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewUserAccessToken(id string, credentials map[string]string) (string, error) {

	// conf := config.Config()
	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["name"] = credentials["name"]
	claims["phone_number"] = credentials["phone_number"]
	claims["role"] = credentials["role"]
	claims["user_name"] = credentials["user_name"]
	claims["password"] = credentials["password"]

	const day = time.Hour * 24

	// in acces 10 days
	claims["expires"] = time.Now().Add(day * 10).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(constatnts.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	sha256 := sha256.New()

	// Create a new now date and time string with salt.
	refresh := constatnts.JWTSecretKey + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := sha256.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(constatnts.JWTRefreshKeyExpireHours)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(sha256.Sum(nil)) + "." + expireTime

	return t, nil
}
