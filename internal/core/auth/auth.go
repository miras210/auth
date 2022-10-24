package auth

import (
	"time"
)

const (
	AccessTokenExpiration  = time.Minute * 15
	RefreshTokenExpiration = time.Hour * 24 * 7
)

// Token - structure for holding all token related data
type Token struct {
	ID         string `json:"-"`
	UserID     string `json:"user_id"`
	TokenValue string `json:"token_value"`
	ExpiresAt  int64  `json:"expires_at"`
}

// TokenPair - structure for holding access and refresh token
type TokenPair struct {
	AccessToken  *Token `json:"access_token"`
	RefreshToken *Token `json:"refresh_token"`
}
