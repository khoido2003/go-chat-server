package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/khoido2003/go-chat-server/router"
)

func main() {

	godotenv.Load()

	portStr := os.Getenv("PORT")

	if portStr == "" {
		log.Fatal("PORT IS NOT FOUND!")
	}
	fmt.Println("PORT: ", portStr)

	// dbURL := os.Getenv("DB_URL")

	// if dbURL == "" {
	// 	log.Fatal("DB_URL IS NOT FOUND!")
	// }

	//////////////////////////////////////////////

	// Setup router
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	r.Mount("/v1/api", router.V1Router())

	///////////////////////////////////////////////////

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + portStr,
	}

	log.Printf("Server starting on port %v", portStr)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
