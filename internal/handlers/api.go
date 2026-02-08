package handlers

import (
	"github.com/AN01KU/money-manager/internal/middleware"
	"github.com/AN01KU/money-manager/internal/tools"
	"github.com/go-chi/chi"
)

func (h *Handlers) RegisterRoutes(r *chi.Mux) {
	//GLOBAL MIDDLEWARES
	r.Use(middleware.Logger)
	r.Use(middleware.Recovery)

	r.Route("/auth", func(router chi.Router) {
		router.Post("/signup", h.signup)
		router.Post("/login", h.login)
	})

	r.Group(func(router chi.Router) {
		router.Use(middleware.Auth(h.DB))

		router.Post("/groups", h.createGroups)
	})
}

type Handlers struct {
	DB tools.DatabaseInterface
}

func NewHandler(db tools.DatabaseInterface) *Handlers {
	return &Handlers{DB: db}
}
