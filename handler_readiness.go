package main

import "net/http"

func handlerReadinees(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
