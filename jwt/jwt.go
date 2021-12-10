package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	SignKey = "test.me"

	ValidationErrorSignatureInvalid = errors.New("couldn't handle this token")
	ValidationErrorMalformed        = errors.New("that's not even a token")
	ValidationErrorExpired          = errors.New("token is expired")
	ValidationErrorNotValidYet      = errors.New("token not active yet")
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

func New() *JWT {
	return &JWT{
		[]byte(SignKey),
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ValidationErrorMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ValidationErrorExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ValidationErrorNotValidYet
			} else {
				return nil, ValidationErrorSignatureInvalid
			}
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ValidationErrorSignatureInvalid
}

func (j *JWT) RefreshToken(tokenString string, expiresAt time.Time) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = expiresAt.Unix()
		return j.CreateToken(*claims)
	}
	return "", ValidationErrorSignatureInvalid
}
