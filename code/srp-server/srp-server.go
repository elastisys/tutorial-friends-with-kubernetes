package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kong/go-srp"
)

var authSessionCache = map[string](*srp.SRPServer){}

func handleError(w http.ResponseWriter, statusCode int, message string) {
	log.Printf("Error: %q\n", message)
	errorResponse := ErrorResponse{message}
	handleResponse(w, statusCode, errorResponse)
}

func handleResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleChallenge(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		handleError(w, http.StatusBadRequest, "You need to supply a username")
		return
	}

	authValue, ok := authDatabase[username]
	if !ok {
		handleError(w, http.StatusBadRequest, "Username not found")
		return
	}

	verifier, err := base64.StdEncoding.DecodeString(authValue.verifier)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "")
		return
	}

	salt, err := base64.StdEncoding.DecodeString(authValue.salt)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "")
		return
	}

	params := srp.GetParams(4096)
	secret2 := srp.GenKey()
	server := srp.NewServer(params, verifier, secret2)
	srpB := server.ComputeB()

	// We assume one login flow per user
	authSessionCache[username] = server

	handleResponse(w, http.StatusOK, AuthChallengeResponse{salt, srpB})
	log.Printf("Challenge sent for %q\n", username)
}

func handleAuthentication(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var areq AuthAuthenticateRequest
	err := decoder.Decode(&areq)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	server, ok := authSessionCache[areq.Username]
	defer delete(authSessionCache, areq.Username)
	if !ok {
		handleError(w, http.StatusBadRequest, "No authentication session found")
		return
	}

	server.SetA(areq.A)
	srpM2, err := server.CheckM1(areq.M1)
	if err != nil {
		handleError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	handleResponse(w, http.StatusOK, AuthAuthenticateResponse{srpM2})
	log.Printf("%q authenticated successfully\n", areq.Username)
}

func main() {
	log.Println("Starting up ...")

	http.HandleFunc("/auth/challenge", handleChallenge)
	http.HandleFunc("/auth/authenticate", handleAuthentication)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
