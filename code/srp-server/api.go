package main

// ErrorResponse stores the error returned to the client
type ErrorResponse struct {
	Message string
}

// AuthChallengeResponse stores the response to an /auth/challenge request
type AuthChallengeResponse struct {
	Salt []byte
	B    []byte
}

// AuthAuthenticateRequest stores the request to an /auth/authenticate request
type AuthAuthenticateRequest struct {
	Username string
	A        []byte
	M1       []byte
}

// AuthAuthenticateResponse stores the response to an /auth/authenticate request
type AuthAuthenticateResponse struct {
	M2 []byte
}
