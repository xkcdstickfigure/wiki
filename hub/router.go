package hub

import (
	"alles/wiki/store"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db store.Store) chi.Router {
	// create router
	r := chi.NewRouter()
	h := handlers{db}

	// discord
	r.Get("/discord", h.discordAuth)
	r.Get("/discord/callback", h.discordCallback)

	// google
	r.Get("/auth", h.googleAuth)
	r.Get("/auth/callback", h.googleCallback)

	// gitea
	r.Get("/gitea", h.giteaAuth)
	r.Post("/gitea/token", h.giteaToken)
	r.Get("/gitea/profile", h.giteaProfile)

	return r
}

type handlers struct {
	db store.Store
}
