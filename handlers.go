package main

import "net/http"

func readinessGetHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func errGetHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError,
		"Internal Server Error")
}
