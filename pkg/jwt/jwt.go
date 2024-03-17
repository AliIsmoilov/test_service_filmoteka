package jwt

import (
	"errors"
	"log"
	"strings"
	"test_service_filmoteka/pkg/constatnts"

	jwtgo "github.com/golang-jwt/jwt"
)

// ExtractClaims extracts roles from the claims of JWT token
func ExtractClaims(tokenString string, signingKey []byte) (jwtgo.MapClaims, error) {
	claims := jwtgo.MapClaims{}
	if tokenString == "" {
		claims["role"] = constatnts.UnAuthorized
		return claims, nil
	}
	if strings.Contains(tokenString, "Basic") {
		claims["role"] = constatnts.UnAuthorized
		return claims, nil
	}
	token, err := jwtgo.ParseWithClaims(tokenString, claims, func(token *jwtgo.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, errors.Join(err, constatnts.InvalidToken)
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !(ok && token.Valid) {
		return nil, constatnts.InvalidToken
	}

	return claims, nil
}

func GetUserMetadata(signingKey, accessToken string) (map[string]any, error) {

	claims, err := ExtractClaims(accessToken, []byte(signingKey))
	if err != nil {
		claims, err = ExtractClaims(accessToken, []byte(constatnts.JWTKEY))
		if err != nil {
			log.Println("could not extract claims:", err)
			return nil, err
		}

	}
	return claims, nil
}
