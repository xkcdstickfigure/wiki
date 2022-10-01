package main

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect to database
	_, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// html templates
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v\n", err)
	}

	// http router
	r := chi.NewRouter()

	// assets
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// homepage
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	// article
	r.Get("/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.ToLower(chi.URLParam(r, "slug"))
		subdomain := getSubdomain(r)

		html := new(bytes.Buffer)
		err := tmpl.ExecuteTemplate(html, "article.html", struct {
			Site    string
			Domain  string
			Slug    string
			Title   string
			Content template.HTML
		}{
			Site:    subdomain,
			Domain:  os.Getenv("DOMAIN"),
			Slug:    slug,
			Title:   "Eye of Cthulhu",
			Content: "The Eye of Cthulu is a <a href='/pre-hardmode'>pre-Hardmode</a> <a href='/boss'>boss</a>. It is one of the first bosses the player may encounter, as it spawns automatically when a relatively early level of game advancement is achieved.",
		})

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Write(html.Bytes())
	})

	http.ListenAndServe(":3000", r)
}

func getSubdomain(r *http.Request) string {
	domain := os.Getenv("DOMAIN")
	if strings.HasSuffix(r.Host, "."+domain) {
		return strings.TrimSuffix(r.Host, "."+domain)
	} else {
		return ""
	}
}
