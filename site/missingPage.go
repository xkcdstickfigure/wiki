package site

import (
	"net/http"

	"alles/wiki/env"
	"alles/wiki/store"
)

func (h handlers) sendMissingPage(w http.ResponseWriter, r *http.Request, site store.Site) {
	w.WriteHeader(http.StatusNotFound)
	h.templates.ExecuteTemplate(w, "missing.html", struct {
		Site          string
		SiteName      string
		Origin        string
		StorageOrigin string
	}{
		Site:          site.Name,
		SiteName:      site.DisplayName,
		Origin:        env.Origin,
		StorageOrigin: env.StorageOrigin,
	})
}
