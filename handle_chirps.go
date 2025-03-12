package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vallesda/goserver/internal/auth"
	"github.com/vallesda/goserver/internal/database"
)

func (cfg *apiConfig) handleCreateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type Chirp struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UserID    uuid.UUID `json:"user_id"`
		Body      string    `json:"body"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	cleaned_body := badWordReplacement(params.Body)
	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   cleaned_body,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func (cfg *apiConfig) handleGetChirp(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.PathValue("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse to uuid", err)
	}
	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
	}

	respondWithJSON(w, http.StatusOK, databaseChirpToChirp(chirp))
}

func (cfg *apiConfig) handleGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
	}

	respondWithJSON(w, http.StatusOK, databaseChirpToChirps(chirps))
}
