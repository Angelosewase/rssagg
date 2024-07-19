package main

import "net/http"


func handleReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJson(w,200,struct{}{})
}