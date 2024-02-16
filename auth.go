package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/BigStinko/rssgregator/internal/database"
)

var (
	ErrNoAuthHeader = errors.New("no authorization header included")
	ErrBadAuthHeader = errors.New("malformed authorization header")
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)'
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}
		handler(w, r, user)
	}
}

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", ErrBadAuthHeader
	}
	return splitAuth[1], nil
}
