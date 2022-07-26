package tmdb

import (
	"net/http"

	"github.com/pkg/errors"
)

// AuthenticationResource handles authentication-related requests of TMDb API.
type AuthenticationResource struct {
	client *Client
}

// AuthToken represents the authentication object in TMDb.
type AuthToken struct {
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
	Success      bool   `json:"success"`
}

// CreateRequestToken creates a temporary request token that can be used to validate a TMDB user login.
func (ar *AuthenticationResource) CreateRequestToken() (*AuthToken, *http.Response, error) {
	path := "/authentication/token/new"
	var response AuthToken
	resp, err := ar.client.get(path, &response)
	return &response, resp, errors.Wrap(err, "failed to get request token")
}

// GuestSession represents a guest session object in TMDb.
type GuestSession struct {
	ExpiresAt      string `json:"expires_at"`
	GuestSessionID string `json:"guest_session_id"`
	Success        bool   `json:"success"`
}

// CreateGuestSession creates a new guest session.
func (ar *AuthenticationResource) CreateGuestSession() (*GuestSession, *http.Response, error) {
	path := "/authentication/guest_session/new"
	var session GuestSession
	resp, err := ar.client.get(path, &session)
	return &session, resp, errors.Wrap(err, "failed to get guest session")
}

// Session represents the session object in TMDb.
type Session struct {
	SessionID string `json:"session_id"`
	Success   bool   `json:"success"`
}

// CreateSession creates a fully valid session ID once a user has validated the request token.
func (ar *AuthenticationResource) CreateSession(requestToken string) (*Session, *http.Response, error) {
	path := "/authentication/session/new"
	opt := map[string]string{
		"request_token": requestToken,
	}
	var session Session
	resp, err := ar.client.post(path, &session, WithBody(opt))
	return &session, resp, errors.Wrap(err, "failed to get session")
}

// ValidateRequestToken allows an application to validate a request token by entering a username and password.
func (ar *AuthenticationResource) ValidateRequestToken(username, password, requestToken string) (*AuthToken, *http.Response, error) {
	path := "/authentication/token/validate_with_login"
	opt := map[string]string{
		"request_token": requestToken,
		"username":      username,
		"password":      password,
	}
	var session AuthToken
	resp, err := ar.client.post(path, &session, WithBody(opt))
	return &session, resp, errors.Wrap(err, "failed to get session")
}

// CreateSessionWithV4Token creates a v3 session ID from a valid v4 access token.
// The v4 token needs to be authenticated by the user.
// The standard "read token" will not validate to create a session ID.
func (ar *AuthenticationResource) CreateSessionWithV4Token(accessToken string) (*Session, *http.Response, error) {
	path := "/authentication/session/convert/4"
	opt := map[string]string{
		"access_token": accessToken,
	}
	var session Session
	resp, err := ar.client.post(path, &session, WithBody(opt))
	return &session, resp, errors.Wrap(err, "failed to get session")
}

// DeleteSessionResponse represents the response for deleting a session id.
type DeleteSessionResponse struct {
	Success bool `json:"success"`
}

// DeleteSession deletes (or "logout") from a session.
func (ar *AuthenticationResource) DeleteSession(sessionID string) (*DeleteSessionResponse, *http.Response, error) {
	path := "/authentication/session"
	opt := map[string]string{
		"session_id": sessionID,
	}
	var deleteResponse DeleteSessionResponse
	resp, err := ar.client.delete(path, &deleteResponse, WithBody(opt))
	return &deleteResponse, resp, errors.Wrap(err, "failed to delete session")
}
