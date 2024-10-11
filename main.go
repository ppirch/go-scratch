package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func responseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with error:", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errorResponse{
		Error: message,
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server listening on port %s", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
