package routes

import (
	// "log"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mdalasini/rssagg/internal/handler"
)

func NewRounter() *chi.Mux{
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handler.HandleReadiness)
	v1Router.Get("/err", handler.HandleErr)

	mux.Mount("/v1", v1Router)

	return mux
}
