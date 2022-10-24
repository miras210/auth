package service

import (
	"auth/internal/core/auth"
	"auth/internal/core/user"
	"auth/internal/sys"
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	userRepo  user.Repository
	tokenRepo auth.TokenRepository
	logger    *zap.SugaredLogger
	config    sys.TokenConfig
}

func NewAuthService(userRepo user.Repository, tokenRepo auth.TokenRepository, logger *zap.SugaredLogger, config sys.TokenConfig) auth.Service {
	return AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		logger:    logger,
		config:    config,
	}
}

func (a AuthService) Auth(ctx context.Context, creds *user.SignIn) (*auth.TokenPair, error) {
	login := GetLogin(creds.Email, creds.Username)
	retrievedUser, err := a.userRepo.GetByEmailOrUsername(ctx, login)
	if err != nil {
		a.logger.Errorf("error occurred while getting user: %s", err.Error())
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(retrievedUser.PassHash), []byte(creds.Password)); err != nil {
		a.logger.Errorf("invalid password")
		//TODO change error
		return nil, err
	}

	return a.GenerateTokens(ctx, retrievedUser)
}

func (a AuthService) Refresh(ctx context.Context, refreshString string) (*auth.TokenPair, error) {
	token, err := a.tokenRepo.GetByToken(ctx, refreshString)
	if err != nil {
		return nil, err
	}

	err = a.tokenRepo.Delete(ctx, token.ID)
	if err != nil {
		return nil, err
	}

	if time.Now().Unix() > token.ExpiresAt {
		return nil, ErrTokenExpired
	}

	u := user.User{ID: token.UserID}

	tokens, err := a.GenerateTokens(ctx, u)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a AuthService) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) ValidateAccess(ctx context.Context, accessString string) (jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(accessString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.PubKey), nil
	})
	if err != nil {
		return jwt.StandardClaims{}, err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return jwt.StandardClaims{}, err
	}

	return *claims, nil
}

func (a AuthService) GenerateTokens(ctx context.Context, user user.User) (*auth.TokenPair, error) {
	accessToken, err := createAccessToken(user, a.config.PubKey, a.config.AccessExpiration)
	if err != nil {
		a.logger.Errorf("error occurred while creating access token: %s", err.Error())
		return nil, err
	}

	refreshToken, err := createRefreshToken(user, a.config.PrivKey, a.config.RefreshExpiration)
	if err != nil {
		a.logger.Errorf("error occurred while creating access token: %s", err.Error())
		return nil, err
	}

	err = a.tokenRepo.Create(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	tokenPair := &auth.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, nil
}

func createAccessToken(user user.User, secret string, ttl time.Duration) (*auth.Token, error) {
	iat := time.Now().Unix()
	//TODO: change to env var
	exp := time.Now().Add(ttl)
	atClaims := jwt.MapClaims{}
	atClaims["sub"] = user.ID
	atClaims["iat"] = iat
	atClaims["exp"] = exp.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, err := at.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	token := &auth.Token{
		UserID:     user.ID,
		TokenValue: tokenString,
		ExpiresAt:  exp.Unix(),
	}

	return token, nil
}

func createRefreshToken(user user.User, secret string, ttl time.Duration) (*auth.Token, error) {
	iat := time.Now().Unix()
	//TODO: change to env var
	exp := time.Now().Add(ttl)
	atClaims := jwt.MapClaims{}
	atClaims["sub"] = user.ID
	atClaims["iat"] = iat
	atClaims["exp"] = exp.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, err := at.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	token := &auth.Token{
		UserID:     user.ID,
		TokenValue: tokenString,
		ExpiresAt:  exp.Unix(),
	}

	return token, nil
}
