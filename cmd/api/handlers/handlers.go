package handlers

import (
	"auth/cmd/api/handlers/v1/authgrp"
	"auth/cmd/api/handlers/v1/usrgrp"
	"auth/internal/repository/postgres"
	"auth/internal/service"
	"auth/internal/sys"
	"auth/internal/web/v1/middleware"
	"auth/pkg/web"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
	"os"
)

const version = "v1"

func API(shutdown chan os.Signal, log *zap.SugaredLogger, db *pgxpool.Pool, tokenConf sys.TokenConfig) (http.Handler, error) {
	app := web.NewApp(mux.NewRouter(), shutdown, middleware.Errors(log))

	{
		tokenRepo := postgres.NewTokenRepository(db, log)
		userRepo := postgres.NewUserRepository(db, log)
		authService := service.NewAuthService(userRepo, tokenRepo, log, tokenConf)
		authHandler := authgrp.NewHandler(authService)
		app.Handle(http.MethodPost, version, "/sign-in", authHandler.Signin)
		app.Handle(http.MethodPost, version, "/refresh", authHandler.Refresh)
	}
	{
		tokenRepo := postgres.NewTokenRepository(db, log)
		userRepo := postgres.NewUserRepository(db, log)
		userService := service.NewUserService(userRepo, log)
		authService := service.NewAuthService(userRepo, tokenRepo, log, tokenConf)
		userHandler := usrgrp.NewHandler(userService)
		app.Handle(http.MethodPost, version, "/sign-up", userHandler.Signup)
		app.Handle(http.MethodGet, version, "/test", userHandler.Test, middleware.Authorize(authService))
	}

	return app, nil
}
