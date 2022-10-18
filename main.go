package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"alles/wiki/env"
	"alles/wiki/hub"
	"alles/wiki/site"
	"alles/wiki/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// connect to database
	conn, err := pgxpool.New(context.Background(), env.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	db := store.Store{Conn: conn}

	// router
	hubRouter := hub.NewRouter(db)
	siteRouter := site.NewRouter(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname := strings.Split(r.Host, ":")[0]
		if hostname == env.Domain {
			// domain
			hubRouter.ServeHTTP(w, r)
		} else if strings.HasSuffix(hostname, "."+env.Domain) {
			// subdomain
			siteRouter.ServeHTTP(w, r)
		}
	})

	fmt.Println("starting http server on :3000")
	http.ListenAndServe(":3000", nil)
}
