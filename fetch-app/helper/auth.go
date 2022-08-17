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

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Username string
	jwt.StandardClaims
}

// GenerateToken generates a jwt token
func (j *JwtWrapper) GenerateToken(username string) (signedToken string, err error) {
	claims := &JwtClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
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
