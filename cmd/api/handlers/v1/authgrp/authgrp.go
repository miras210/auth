package authgrp

import (
	"auth/internal/core/auth"
	"auth/internal/core/user"
	v1 "auth/internal/web/v1"
	"auth/pkg/web"
	"context"
	"net/http"
)

type Handler struct {
	AuthService auth.Service
}

func NewHandler(authService auth.Service) *Handler {
	return &Handler{
		AuthService: authService,
	}
}

func (h Handler) Signin(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var u user.SignIn
	if err := web.Decode(r, &u); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	tokens, err := h.AuthService.Auth(ctx, &u)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	return web.Respond(ctx, w, tokens, http.StatusOK)
}

func (h Handler) Refresh(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	t := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	if err := web.Decode(r, &t); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	tokens, err := h.AuthService.Refresh(ctx, t.RefreshToken)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	return web.Respond(ctx, w, tokens, http.StatusOK)
}
