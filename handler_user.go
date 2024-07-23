package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Angelosewase/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})

	type parmaters struct {
		Name string `json:"name"`
	}
	Params := parmaters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&Params)

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Error parsing json %v:", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        fmt.Sprintf("%v", uuid.New()),
		Createdat: time.Now().UTC(),
		Updateat:  time.Now().UTC(),
		Name:      Params.Name,
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't create user:%v", err))
	}

	insertedUserID, err := user.LastInsertId()
	if err != nil {
		return
	}

	respondWithJson(w, 200, struct{ id any }{id: insertedUserID})

}
