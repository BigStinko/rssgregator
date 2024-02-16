package main

import (
	"net/http"
	"time"

	"github.com/BigStinko/rssgregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) postFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn'tdecode parameters")
		return
	}

	feed, err := cfg.DB.CreatFeed(r.Context(), database.CreatFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserId: user.ID,
		Name: params.Name,
		Url: params.URL,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) getFeedHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feeds))
}
