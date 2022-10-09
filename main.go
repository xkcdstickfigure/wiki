package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"alles/wiki/site"
	"alles/wiki/store"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect to database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	db := store.Store{Conn: conn}

	// site
	siteRouter := site.NewRouter(db)

	// router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		domain := os.Getenv("DOMAIN")
		if r.Host == domain {
			// domain
			w.Write([]byte("glaffle"))
		} else if strings.HasSuffix(r.Host, "."+domain) {
			// subdomain
			siteRouter.ServeHTTP(w, r)
		}
	})

	http.ListenAndServe(":3000", nil)
}
