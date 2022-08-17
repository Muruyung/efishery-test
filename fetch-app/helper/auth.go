package helper

import (
	"MyAPI/config"
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(
		signedToken,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if int64(claims["exp"].(float64)) < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return
}

// BearerTokenAuth validate the jwt token from bearer token
func BearerTokenAuth(bearerToken string) (jwt.MapClaims, error) {
	bearer := strings.Split(bearerToken, " ")
	token := bearer[1]

	jwtWrapper := JwtWrapper{
		SecretKey:       config.JWTSECRETKEY(),
		Issuer:          "AuthService",
		ExpirationHours: int64(config.JWTEXPIRATIONHOUR()),
	}

	payload, err := jwtWrapper.ValidateToken(token)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	return payload, nil
}
