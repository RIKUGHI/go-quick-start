package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rikughi/go-quick-start/internal/delivery/http/controller"
)

type Router struct {
	HelloController *controller.HelloController
}

func (r *Router) Handler() http.Handler {
	c := chi.NewRouter()

	c.Use(middleware.RequestID)
	c.Use(middleware.RealIP)
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Use(middleware.Timeout(60 * time.Second))

	c.Get("/", r.HelloController.Hello)

	return c
}
