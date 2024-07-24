package auth

import (
	"errors"
	"net/http"
	"strings"
)



func GetApiKey(header http.Header) (string, error) {
   val := header.Get("Authorisation")

   if val == "" {
       return "", errors.New("no authorization info found")
   }

   vals := strings.Split(val, " ")

   if len(vals) != 2 {
       return "", errors.New("invalid format of the authorization header")
   }

   if vals[0] != "ApiKey" {
       return "", errors.New("invalid format of the authorization header")
   }

   return vals[1], nil
}