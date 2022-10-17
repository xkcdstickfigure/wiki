package site

import (
	"html/template"
	"net/http"
	"strings"

	"alles/wiki/env"
	"alles/wiki/markup"
	"alles/wiki/render"
	"alles/wiki/sessionAuth"
	"alles/wiki/store"

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
		h.templates.ExecuteTemplate(w, "missing.html", struct {
			Site          store.Site
			Origin        string
			StorageOrigin string
		}{
			Site:          site,
			Origin:        env.Origin,
			StorageOrigin: env.StorageOrigin,
		})
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
		Domain:        env.Domain,
		Slug:          article.Slug,
		StorageOrigin: env.StorageOrigin,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get session
	session, err := sessionAuth.UseSession(h.db, w, r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// add to history
	_, err = h.db.ArticleViewCreate(r.Context(), store.ArticleView{
		SessionId: session.Id,
		ArticleId: article.Id,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// render page
	h.templates.ExecuteTemplate(w, "article.html", struct {
		Site          store.Site
		Origin        string
		StorageOrigin string
		Title         string
		Content       template.HTML
	}{
		Site:          site,
		Origin:        env.Origin,
		StorageOrigin: env.StorageOrigin,
		Title:         article.Title,
		Content:       template.HTML(articleHtml),
	})
}
