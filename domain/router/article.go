package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"handler/function/domain/model"
	"handler/function/domain/template"
)

func (router Router) ShowArticles(w http.ResponseWriter, r *http.Request) {
	articlePtrs, err := router.service.GetArticles(r.Context())
	if router.handleError(w, err) {
		return
	}

	articles := make([]model.Article, len(articlePtrs))
	for i, articlePtr := range articlePtrs {
		articles[i] = *articlePtr
	}

	template.Layout(template.ArticleList(articles)).Render(r.Context(), w)
}

func (router Router) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		router.handleError(w, fmt.Errorf("missing article id"))
		return
	}

	articleID, err := primitive.ObjectIDFromHex(id)
	if router.handleError(w, err) {
		return
	}

	article := &model.Article{
		ID: articleID,
	}

	err = router.service.DeleteArticle(r.Context(), article)
	if router.handleError(w, err) {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
