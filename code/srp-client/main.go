package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Kong/go-srp"
)

func post(baseURL *url.URL, relURLString string, body []byte) (*http.Response, error) {
	relURL, err := url.Parse(relURLString)
	if err != nil {
		return nil, err
	}
	resolvedURL := baseURL.ResolveReference(relURL)
	return http.Post(resolvedURL.String(), "application/json", bytes.NewBuffer(body))
}

func doAuth(baseURL *url.URL) {
	fmt.Printf("%v: Getting challenge ... ", time.Now().UTC().Format(time.RFC3339))
	resp, err := post(baseURL, "/auth/challenge?username=test@example.com", nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	var challengeResponse AuthChallengeResponse
	err = json.NewDecoder(resp.Body).Decode(&challengeResponse)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	params := srp.GetParams(4096)
	username := "test@example.com"
	password := "testPassword"

	secret1 := srp.GenKey()
	client := srp.NewClient(params, challengeResponse.Salt, []byte(username), []byte(password), secret1)
	srpA := client.ComputeA()
	client.SetB(challengeResponse.B)
	srpM1 := client.ComputeM1()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Authentication ... ")
	var authenticateReq AuthAuthenticateRequest
	authenticateReq.Username = username
	authenticateReq.A = srpA
	authenticateReq.M1 = srpM1
	body, err := json.Marshal(&authenticateReq)
	resp, err = post(baseURL, "/auth/authenticate", body)

	if err != nil {
		fmt.Println("LOW-LEVEL FAIL ❌")
	} else if resp.StatusCode == 200 {
		fmt.Println("PASS ✅")
	} else {
		fmt.Println("FAIL ❌")
	}
}

func main() {
	flag.Parse()

	baseURL, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatalf("Cannot parse base URL: %v", err)
	}

	for {
		doAuth(baseURL)
	}
}
