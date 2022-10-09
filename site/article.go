package site

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"strings"

	"alles/wiki/markup"
	"alles/wiki/render"

	"github.com/go-chi/chi/v5"
)

func (h handlers) article(w http.ResponseWriter, r *http.Request) {
	slug := strings.ToLower(chi.URLParam(r, "slug"))
	subdomain := getSubdomain(r)

	// get site
	site, err := h.db.SiteGetByName(r.Context(), subdomain)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get article
	article, err := h.db.ArticleGetBySlug(r.Context(), site.Id, slug)
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
	err = h.templates.ExecuteTemplate(html, "article.html", struct {
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
}
