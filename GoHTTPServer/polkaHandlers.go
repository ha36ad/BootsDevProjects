package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ha36ad/BootsDevProjects/GoHTTPServer/internal/auth"
)

func (cfg *apiConfig) polkaHandler(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key", err)
		return
	}

	if apiKey != cfg.PolkaKey {
		respondWithError(w, http.StatusUnauthorized, "API key is invalid", err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if request.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, struct{}{})
		return
	}

	userID, err := uuid.Parse(request.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID format", err)
		return
	}

	_, err = cfg.DB.UpgradeToChirpyRed(r.Context(), userID)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, struct{}{})

}
