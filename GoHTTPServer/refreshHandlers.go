package main

import (
	"net/http"
	"time"

	"github.com/ha36ad/BootsDevProjects/GoHTTPServer/internal/auth"
)

func (cfg *apiConfig) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}
	tokenStr, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or malformed token", err)
		return
	}

	refreshToken, err := cfg.DB.GetRefreshToken(r.Context(), tokenStr)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		respondWithError(w, http.StatusUnauthorized, "Token expired", nil)
		return
	}

	if refreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Token revoked", nil)
		return
	}

	jwtStr, err := auth.MakeJWT(refreshToken.UserID, cfg.JWTSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{Token: jwtStr})
}

func (cfg *apiConfig) revokeHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or malformed token", err)
		return
	}

	_, err = cfg.DB.RevokeRefreshToken(r.Context(), tokenStr)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Token not found or already revoked", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
