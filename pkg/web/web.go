package web

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"syscall"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	router      *mux.Router
	shutdown    chan os.Signal
	middlewares []Middleware
}

func NewApp(router *mux.Router, shutdown chan os.Signal, middlewares ...Middleware) *App {
	return &App{
		shutdown:    shutdown,
		router:      router,
		middlewares: middlewares,
	}
}

func (a *App) Shutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *App) Handle(method string, group string, path string, handler Handler, middlewares ...Middleware) {
	handler = wrapMiddleware(handler, middlewares)
	handler = wrapMiddleware(handler, a.middlewares)

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			log.Println(err.Error())
			a.Shutdown()
			return
		}
	}

	if group != "" {
		path = "/" + group + path
	}

	a.router.HandleFunc(path, h).Methods(method)
}
