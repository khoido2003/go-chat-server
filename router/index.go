package router

import (
	"github.com/go-chi/chi/v5"
)

func V1Router() *chi.Mux {
	v1Router := chi.NewRouter()

	conversation(v1Router)

	return v1Router
}
