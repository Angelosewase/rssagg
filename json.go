package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responWithError(w http.ResponseWriter,code int , msg string){
	if code > 499{
		log.Printf("responding with a 5XX error: %v",msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJson(w,code,errResponse{
		Error: msg,
	})

}


func respondWithJson(w http.ResponseWriter,code int, payload interface{}){
	data,err := json.Marshal(payload)

	if err != nil{
	    log.Printf("Failed to marshall json : %v ",err)	
		w.WriteHeader(500)
		return
	}

	  w.Header().Add("Content-Type:","application/json")
	  w.WriteHeader(code)
	  w.Write(data)
}