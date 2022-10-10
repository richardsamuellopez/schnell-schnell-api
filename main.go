package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"
)

type authEvent struct {
	Email		string `json:"Email"`
	Password	string `json:"Password`
	Token 		string `json:"Token`
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	var newAuthEvent authEvent
	err := json.NewDecoder(r.Body).Decode(&newAuthEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newAuthEvent.Email == "" || newAuthEvent.Password == "" || newAuthEvent.Token == "" {
		log.Printf("Invalid request body")
		fmt.Fprintf(w, "Invalid request body")
	}

	var isValidCredentials = validateCredentials(newAuthEvent.Email, newAuthEvent.Password)
	var isValidToken = validateToken(newAuthEvent.Token)

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	if isValidCredentials && isValidToken {
		log.Printf("Success")
		// fmt.Fprintf(w, "Valid auth")
		// w.WriteHeader(http.StatusOK)
		resp["message"] = "Status OK"
		
	} else {
		log.Printf("Failure")
		// fmt.Fprintf(w, "INVALID AUTH")
		// w.WriteHeader(http.StatusBadRequest)
		resp["message"] = "Status UNAUTHORIZED"
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func validateCredentials(e string, p string) bool {
	// TODO implement non hard coded credential validation
	if e == "c137@onecause.com" && p == "#th@nH@rm#y#r!$100%D0p#" {
		return true
	}
	return false
}

func validateToken(t string) bool {
	now := time.Now()
	hour := fmt.Sprintf("%02d", now.Hour()) // Hour is 24 format
	minute := fmt.Sprintf("%02d", now.Minute())

	if t == hour + minute {
		return true
	}
	return false
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", authenticate).
	HeadersRegexp("Content-Type", "application/json").
	Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
