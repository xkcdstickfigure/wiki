package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect to database
	_, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// http router
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	r.Get("/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.ToLower(chi.URLParam(r, "slug"))
		subdomain := getSubdomain(r)
		w.Write([]byte("this page is " + slug + " on " + subdomain))
	})

	http.ListenAndServe(":3000", r)
}

func getSubdomain(r *http.Request) string {
	domain := os.Getenv("DOMAIN")
	if strings.HasSuffix(r.Host, "."+domain) {
		return strings.TrimSuffix(r.Host, "."+domain)
	} else {
		return ""
	}
}
