package response

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenResponse struct {
	Username    string `json:"username"`
	TokenAccess string `json:"token_access"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
}

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
