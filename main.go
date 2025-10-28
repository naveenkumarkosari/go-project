package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/naveenkumarkosari/go-project.git/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = string("8080")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == " " {
		log.Fatal("DB URL is not provided")
	}
	conn, err := sql.Open("postgres", dbURL)

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	if err != nil {
		log.Fatal("error connecting to Database")
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/", handleHealthRequest)
	v1Router.Get("/err", handleErrRequest)
	v1Router.Post("/user", apiCfg.handleCreateUser)
	v1Router.Get("/user", apiCfg.GetUser)
	v1Router.Get("/list", apiCfg.GetAllUser)

	// feeds route ====
	v1Router.Post("/feed", apiCfg.CreateFeed)
	v1Router.Get("/myfeeds", apiCfg.GetUserFeeds)
	r.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}
	fmt.Println("sever started at port", port)
	srv.ListenAndServe()
	fmt.Println("server started at port", port)
}
