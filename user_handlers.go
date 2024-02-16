package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/BigStinko/rssgregator/internal/database"
	"github.com/google/uuid"
)

var (
	ErrCouldntDecode = errors.New("couldn't decode parameters")
	ErrCouldntCreateUser = errors.New("couldn't create user")
)

func (cfg *apiConfig) postUsersHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrCouldntDecode.Error())
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrCouldntCreateUser.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
