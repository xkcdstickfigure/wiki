package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"alles/wiki/env"
	"alles/wiki/site"
	"alles/wiki/store"

	"github.com/jackc/pgx/v5"
)

func main() {
	// connect to database
	conn, err := pgx.Connect(context.Background(), env.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	db := store.Store{Conn: conn}

	// site
	siteRouter := site.NewRouter(db)

	// router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Host == env.Domain {
			// domain
			w.Write([]byte("glaffle"))
		} else if strings.HasSuffix(r.Host, "."+env.Domain) {
			// subdomain
			siteRouter.ServeHTTP(w, r)
		}
	})

	http.ListenAndServe(":3000", nil)
}
