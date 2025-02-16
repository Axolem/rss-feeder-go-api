package main

import (
	"net/http"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("'PORT' is not founf in the enviroment")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AlloweMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handleRediness)
	v1Router.Get("/error", handleError)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("Server listerning on PORT: %v", PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
