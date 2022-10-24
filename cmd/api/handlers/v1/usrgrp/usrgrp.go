package usrgrp

import (
	"auth/internal/core/user"
	v1 "auth/internal/web/v1"
	"auth/pkg/web"
	"context"
	"net/http"
)

type Handler struct {
	UserService user.Service
}

func NewHandler(userService user.Service) *Handler {
	return &Handler{UserService: userService}
}

func (h Handler) Signup(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var u user.NewUser
	if err := web.Decode(r, &u); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	err := h.UserService.Create(ctx, &u)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	return web.Respond(ctx, w, nil, http.StatusCreated)
}

func (h Handler) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, "Successful", http.StatusOK)
}
