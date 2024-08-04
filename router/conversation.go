package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func conversation(router *chi.Mux) {

	router.Get("/c", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world! Khoi"))
	})

}
