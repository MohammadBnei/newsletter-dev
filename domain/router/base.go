package router

import (
	"net/http"

	"handler/function/domain/template"
)

func (router Router) Base(w http.ResponseWriter, r *http.Request) {
	template.Page("base").Render(r.Context(), w)
}
