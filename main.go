package main

import (
	"database/sql"
	"github.com/axolem/rss-feeder/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("'PORT' is not founf in the enviroment")
	}

	GOOSE_DBSTRING := os.Getenv("GOOSE_DBSTRING")

	if GOOSE_DBSTRING == "" {
		log.Fatal("'GOOSE_DBSTRING' is not founf in the enviroment")
	}

	conn, err := sql.Open("postgres", GOOSE_DBSTRING)

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handleHealth)
	v1Router.Get("/error", handleError)

	v1UserRouter := chi.NewRouter()
	v1UserRouter.Post("/create", apiCfg.handleCreateUser)
	v1UserRouter.Get("/", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1FeedRouter := chi.NewRouter()
	v1FeedRouter.Post("/create", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))

	v1Router.Mount("/users", v1UserRouter)
	v1Router.Mount("/feed", v1FeedRouter)
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
