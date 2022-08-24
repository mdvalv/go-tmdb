package main

import (
	"os"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

const (
	username      = "username"
	password      = "password"
	requestToken  = "request_token"
	accessTokenV4 = "v4_access_token"
)

var sessionID = os.Getenv("SESSIONID")

func (e example) CreateGuestSession() {
	session, _, err := e.client.Authentication.CreateGuestSession()
	examples.PanicOnError(err)
	examples.PrettyPrint(*session)
}

func (e example) CreateRequestToken() {
	token, _, err := e.client.Authentication.CreateRequestToken()
	examples.PanicOnError(err)
	examples.PrettyPrint(*token)
}

func (e example) ValidateRequestToken() {
	token, _, err := e.client.Authentication.ValidateRequestToken(username, password, requestToken)
	examples.PanicOnError(err)
	examples.PrettyPrint(*token)
}

func (e example) CreateSession() {
	session, _, err := e.client.Authentication.CreateSession(requestToken)
	examples.PanicOnError(err)
	examples.PrettyPrint(*session)
}

func (e example) CreateSessionWithV4Token() {
	session, _, err := e.client.Authentication.CreateSessionWithV4Token(accessTokenV4)
	examples.PanicOnError(err)
	examples.PrettyPrint(*session)
}

func (e example) DeleteSession() {
	success, _, err := e.client.Authentication.DeleteSession(sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*success)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.CreateGuestSession,       // 1
		example.CreateRequestToken,       // 2
		example.ValidateRequestToken,     // 3
		example.CreateSession,            // 4
		example.CreateSessionWithV4Token, // 5
		example.DeleteSession,            // 6
	)
}
