package site

import (
	"bytes"
	"net/http"

	"alles/wiki/env"
	"alles/wiki/store"
)

func (h handlers) sendMissingPage(w http.ResponseWriter, r *http.Request, site store.Site) {
	html := new(bytes.Buffer)
	h.templates.ExecuteTemplate(html, "missing.html", struct {
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

	w.WriteHeader(http.StatusNotFound)
	w.Write(html.Bytes())
}
