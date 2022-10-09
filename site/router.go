package site

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"alles/wiki/markup"
	"alles/wiki/render"
	"alles/wiki/store"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db store.Store) chi.Router {
	// html templates
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v\n", err)
	}

	// create router
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

		// get site
		site, err := db.SiteGetByName(r.Context(), subdomain)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// get article
		article, err := db.ArticleGetBySlug(r.Context(), site.Id, slug)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// parse article
		articleData, err := markup.ParseArticle(article.Source)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// render article html
		articleHtml, err := render.RenderArticle(articleData, render.PageContext{
			Title:         article.Title,
			Site:          site.Name,
			Domain:        os.Getenv("DOMAIN"),
			PageSlug:      article.Slug,
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
			Site:          site.Name,
			SiteName:      site.DisplayName,
			StorageOrigin: os.Getenv("STORAGE_ORIGIN"),
		})

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Write(html.Bytes())
	})

	return r
}

func getSubdomain(r *http.Request) string {
	domain := os.Getenv("DOMAIN")
	if strings.HasSuffix(r.Host, "."+domain) {
		return strings.TrimSuffix(r.Host, "."+domain)
	} else {
		return ""
	}
}
