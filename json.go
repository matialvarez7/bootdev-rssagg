package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Lo que hace la función es formatear el mensaje "string" que nos llega en un objeto json
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", msg)
	}

	type errReponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errReponse{Error: msg})
}

// Lo que hace esta función es transformar la informacion que llega como JSON
// y convertirlo a bytes
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	// En caso de error imprimimos en un log lo que falló y enviamos el código
	// en el encabezado
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}
