package jwt

import (
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

func TestNewJWT(t *testing.T) {
	t.Parallel()
	t.Log(New())
}

func TestJWT_CreateToken(t *testing.T) {
	t.Parallel()
	token, err := New().CreateToken(CustomClaims{
		User: "username",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 3600,
			Issuer:    SignKey},
	})
	t.Log(token, err)
}

func TestJWT_ParseToken(t *testing.T) {
	t.Parallel()
	t.Log(New().ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidXNlcm5hbWUiLCJleHAiOjE1OTQzMzU1NTMsImlzcyI6InRlc3QubWUiLCJuYmYiOjE1OTQyOTgzMTR9.aNwKepeUDpNnwdidXToA58_ONiRdd-JP1D1HvXZ0Ee8"))
}

func TestJWT_RefreshToken(t *testing.T) {
	t.Parallel()
	t.Log(New().RefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoidXNlcm5hbWUiLCJleHAiOjE1OTQzMDI5MTQsImlzcyI6InRlc3QubWUiLCJuYmYiOjE1OTQyOTgzMTR9.2mjRxNe1b_rt8V7Jj6Lav7S0qVdM_Z8W2c3e1qgqfLA", time.Now().Add(-10*time.Hour)))
}
