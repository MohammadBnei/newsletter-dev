package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (r Router) Mount(router chi.Router) {
	router.Use(corsMiddleware)
	router.Get("/", r.Base)
	router.Get("/article", r.ShowArticles)
	router.Delete("/article/{id}", r.DeleteArticle)
}

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (r Router) handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	log.Println(err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)

	return true
}
