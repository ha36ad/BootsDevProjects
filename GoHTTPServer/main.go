package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ha36ad/BootsDevProjects/GoHTTPServer/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	jwtSecret := os.Getenv("JWT_SECRET")
	polkaKey := os.Getenv("POLKA_KEY")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Initialize SQLC Queries
	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"
	apiCfg := apiConfig{
		DB:        dbQueries,
		Platform:  platform,
		JWTSecret: jwtSecret,
		PolkaKey:  polkaKey,
	}

	mux := http.NewServeMux()

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("GET  /api/healthz", apiCfg.healthHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetHandler)

	mux.HandleFunc("POST  /api/chirps", apiCfg.createChirpHandler)
	mux.HandleFunc("GET  /api/chirps", apiCfg.getChirpsHandler)
	mux.HandleFunc("GET  /api/chirps/{chirpID}", apiCfg.getChirpHandler)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.deleteChirpHandler)

	mux.HandleFunc("POST /api/users", apiCfg.createUserHandler)
	mux.HandleFunc("PUT /api/users", apiCfg.updateHandler)

	mux.HandleFunc("POST /api/login", apiCfg.loginHandler)
	mux.HandleFunc("POST  /api/refresh", apiCfg.RefreshHandler)
	mux.HandleFunc("POST  /api/revoke", apiCfg.revokeHandler)

	mux.HandleFunc("POST /api/polka/webhooks", apiCfg.polkaHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
