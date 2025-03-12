package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/vallesda/goserver/internal/database"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

func databaseChirpToChirp(dbChirp database.Chirp) Chirp {
	return Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		UserID:    dbChirp.UserID,
		Body:      dbChirp.Body,
	}
}

func databaseChirpToChirps(dbChirps []database.Chirp) []Chirp {
	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, databaseChirpToChirp(dbChirp))
	}
	return chirps
}
