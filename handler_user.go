package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Angelosewase/rssagg/internal/auth"
	"github.com/Angelosewase/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
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
		CreatedAt: time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Name:      Params.Name,
		ApiKey: fmt.Sprintf("%v", uuid.New()),
	})

	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Couldn't create user:%v", err))
	}

	insertedUserID, err := user.LastInsertId()
	if err != nil {
		log.Fatalf("Error retriving the last inserted user id %v",err)
		return
	}
	respondWithJson(w, 200, struct{ id any }{id: insertedUserID})

}

func (apiCfg *apiConfig) handleGetuser(w http.ResponseWriter, r *http.Request) {
	key,err:=auth.GetApiKey(r.Header)
	if err != nil{
		responWithError(w,403,fmt.Sprintf("error parsing the api key:%v",err))
		return 
	}
    user,err:=apiCfg.DB.GetUserByApiKey(r.Context(),key)
	if err !=nil{
		responWithError(w,403,fmt.Sprintf("error fetching the user %v",err))
	}

	respondWithJson(w,200,user)
}