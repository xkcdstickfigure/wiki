package site

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"alles/wiki/env"
	"alles/wiki/store"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db store.Store) chi.Router {
	// html templates
	tmpl, err := template.ParseGlob("site/templates/*.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v\n", err)
	}

	// create router
	r := chi.NewRouter()
	h := handlers{
		db:        db,
		templates: tmpl,
	}

	// assets
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("site/assets"))))

	// article
	r.Get("/{slug}", h.article)

	return r
}

func getSubdomain(r *http.Request) string {
	hostname := strings.Split(r.Host, ":")[0]
	if strings.HasSuffix(hostname, "."+env.Domain) {
		return strings.TrimSuffix(hostname, "."+env.Domain)
	} else {
		return ""
	}
}

type handlers struct {
	db        store.Store
	templates *template.Template
}
