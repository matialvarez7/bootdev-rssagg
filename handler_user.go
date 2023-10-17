package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/matialvarez7/bootdev-rssagg/internal/database"
)

// En este handler lo convertimos en un método para que pueda además tener como información adicional la apiConfig que creamos
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Estructura para contener lo que nos llega en el body del JSON
	type parameters struct {
		Name string `json:"name"`
	}

	//Transformamos la información que nos llega por request para usarlo en la estructura
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	//Decodificamos en una instancia de la estructura creada
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't creat user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUsertoUser(user))
}
