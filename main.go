package main

import (
	"alles/wiki/markup"
	"alles/wiki/render"
	"bytes"
	"context"
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

//go:embed cmd/parser/eye_of_cthulhu.txt
var source string

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

		// parse article
		article, err := markup.ParseArticle(source)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// render article html
		articleHtml, err := render.RenderArticle(article, render.PageContext{
			Title:         "Eye of Cthulhu",
			Site:          subdomain,
			Domain:        os.Getenv("DOMAIN"),
			PageSlug:      slug,
			StorageOrigin: os.Getenv("STORAGE_ORIGIN"),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// render page
		html := new(bytes.Buffer)
		err = tmpl.ExecuteTemplate(html, "article.html", struct {
			Content       template.HTML
			Site          string
			SiteName      string
			StorageOrigin string
		}{
			Content:       template.HTML(articleHtml),
			Site:          "terraria",
			SiteName:      "Terraria",
			StorageOrigin: os.Getenv("STORAGE_ORIGIN"),
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
