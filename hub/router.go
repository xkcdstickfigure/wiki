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

	return r
}

type handlers struct {
	db store.Store
}
