package gorbit

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestNewJWT(t *testing.T) {
	t.Parallel()
	t.Log(NewJWT())
}

func TestJWT_CreateToken(t *testing.T) {
	t.Parallel()
	token, err := NewJWT().CreateToken(CustomClaims{
		User: "username",
		JSC: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    SignKey},
	})
	t.Log(token, err)
}

func TestJWT_ParseToken(t *testing.T) {
	t.Parallel()
	t.Log(NewJWT().ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidXNlcm5hbWUiLCJKU0MiOnsiZXhwIjoxNTk0MDE3NDE1LCJpc3MiOiJ0ZXN0Lm1lIiwibmJmIjoxNTk0MDEyODE1fX0.UUGVDk9ruWTHIQ9im5q4bLyyq4O933WEhuTdq_ybE8I"))
}

func TestJWT_RefreshToken(t *testing.T) {
	t.Parallel()
	t.Log(NewJWT().RefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidXNlcm5hbWUiLCJKU0MiOnsiZXhwIjoxNTk0MDE3NDE1LCJpc3MiOiJ0ZXN0Lm1lIiwibmJmIjoxNTk0MDEyODE1fX0.UUGVDk9ruWTHIQ9im5q4bLyyq4O933WEhuTdq_ybE8I"))
}
