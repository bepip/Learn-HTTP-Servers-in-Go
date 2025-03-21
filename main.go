package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/bepip/Learn-HTTP-Servers-in-Go/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
}

func main() {
	const port = "8080"
	const filepathRoot = "."

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
	}

	mux := http.NewServeMux()

	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("POST /api/validate_chirp", validateChripHandler)
	mux.HandleFunc("GET /api/healthz", readinessHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetMetricHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Server starting on :%s\n", port)
	log.Fatal(server.ListenAndServe())
}
