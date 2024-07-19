package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responWithError(w, 400, "something went long")
}