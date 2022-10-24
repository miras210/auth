package middleware

import (
	"auth/internal/core/auth"
	v1 "auth/internal/web/v1"
	"auth/pkg/web"
	"context"
	"errors"
	"net/http"
	"strings"
)

func Authorize(a auth.Service) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			header := r.Header.Get("Authorization")
			bearer := strings.Split(header, " ")
			if len(bearer) != 2 {
				return v1.NewRequestError(errors.New("invalid bearer token"), http.StatusForbidden)
			}

			payload, err := a.ValidateAccess(ctx, bearer[1])
			if err != nil {
				return v1.NewRequestError(err, http.StatusForbidden)
			}

			ctx = context.WithValue(ctx, "payload", payload)

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
