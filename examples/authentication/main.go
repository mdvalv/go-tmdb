package main

import (
	"fmt"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) CreateGuestSession() {
	session, _, err := e.client.Authentication.CreateGuestSession()
	if err != nil {
		panic(errors.Wrap(err, "failed to create guest session"))
	}
	examples.PrettyPrint(*session)
}

func (e example) CreateRequestToken() {
	token, _, err := e.client.Authentication.CreateRequestToken()
	if err != nil {
		panic(errors.Wrap(err, "failed to create request token"))
	}
	examples.PrettyPrint(*token)
}

func (e example) ValidateRequestToken() {
	token, _, err := e.client.Authentication.ValidateRequestToken("username", "password", "request_token")
	if err != nil {
		panic(errors.Wrap(err, "failed to validate request token"))
	}
	examples.PrettyPrint(*token)
}

func (e example) CreateSession() {
	session, _, err := e.client.Authentication.CreateSession("request_token")
	if err != nil {
		panic(errors.Wrap(err, "failed to create session"))
	}
	examples.PrettyPrint(*session)
}

func (e example) CreateSessionWithV4Token() {
	session, _, err := e.client.Authentication.CreateSessionWithV4Token("v4_access_token")
	if err != nil {
		panic(errors.Wrap(err, "failed to create session"))
	}
	examples.PrettyPrint(*session)
}

func (e example) DeleteSession() {
	success, _, err := e.client.Authentication.DeleteSession("session_id")
	if err != nil {
		panic(errors.Wrap(err, "failed to delete session"))
	}
	fmt.Println("Deleted successfully?", success)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.CreateGuestSession,
		example.CreateRequestToken,
		example.ValidateRequestToken,
		example.CreateSession,
		example.CreateSessionWithV4Token,
		example.DeleteSession,
	)
}
