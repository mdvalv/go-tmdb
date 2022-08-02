package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ListsResource handles list-related requests of TMDb API.
type ListsResource struct {
	client *Client
}

type List struct {
	CreatedBy     string      `json:"created_by"`
	Description   string      `json:"description"`
	FavoriteCount int         `json:"favorite_count"`
	Id            string      `json:"id"`
	ISO6391       string      `json:"iso_639_1"`
	ItemCount     int         `json:"item_count"`
	Items         []MovieOrTV `json:"items"`
	Name          string      `json:"name"`
	PosterPath    *string     `json:"poster_path"`
}

type ListOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// Get the details of a list.
func (lr *ListsResource) GetList(listId string, opt *ListOptions) (*List, *http.Response, error) {
	path := fmt.Sprintf("/list/%s", listId)
	var list List
	resp, err := lr.client.get(path, &list, WithQueryParams(opt))
	return &list, resp, errors.Wrap(err, "failed to get list")
}

type ItemStatus struct {
	Id          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}

// Check if a movie has already been added to the list.
func (lr *ListsResource) GetItemStatus(listId string, movieId int) (*ItemStatus, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/item_status", listId)
	var status ItemStatus
	resp, err := lr.client.get(path, &status, WithQueryParam("movie_id", fmt.Sprint(movieId)))
	return &status, resp, errors.Wrap(err, "failed to get item status")
}

type CreateList struct {
	Name        string `url:"name,omitempty" json:"name,omitempty"`
	Description string `url:"description,omitempty" json:"description,omitempty"`
	Language    string `url:"language,omitempty" json:"language,omitempty"`
}

type CreateListResponse struct {
	ListId        int    `json:"list_id"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// Create a list.
func (lr *ListsResource) CreateList(sessionId string, list CreateList) (*CreateListResponse, *http.Response, error) {
	path := "/list"
	var response CreateListResponse
	resp, err := lr.client.post(path, &response, WithBody(list), WithQueryParam("session_id", sessionId))
	return &response, resp, errors.Wrap(err, "failed to get item status")
}

type AddItemResponse statusResponse

// Add a movie to a list.
func (lr *ListsResource) AddMovie(sessionId, listId string, itemId int) (*AddItemResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/add_item", listId)
	var response AddItemResponse
	resp, err := lr.client.post(path, &response, WithQueryParam("media_id", fmt.Sprint(itemId)), WithQueryParam("session_id", sessionId))
	return &response, resp, errors.Wrap(err, "failed to add movie")
}

type RemoveItemResponse statusResponse

// Remove a movie from a list.
func (lr *ListsResource) RemoveMovie(sessionId, listId string, itemId int) (*RemoveItemResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/remove_item", listId)
	var response RemoveItemResponse
	resp, err := lr.client.post(path, &response, WithQueryParam("media_id", fmt.Sprint(itemId)), WithQueryParam("session_id", sessionId))
	return &response, resp, errors.Wrap(err, "failed to remove movie")
}

type ClearListResponse statusResponse

// Clear all of the items from a list.
func (lr *ListsResource) Clear(sessionId, listId string) (*ClearListResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/clear", listId)
	var response ClearListResponse
	resp, err := lr.client.post(path, &response, WithQueryParam("confirm", "true"), WithQueryParam("session_id", sessionId))
	return &response, resp, errors.Wrap(err, "failed to clear list")
}

type DeleteListResponse statusResponse

// Delete a list.
func (lr *ListsResource) Delete(sessionId, listId string) (*DeleteListResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s", listId)
	var response DeleteListResponse
	resp, err := lr.client.delete(path, &response, WithQueryParam("session_id", sessionId))
	return &response, resp, errors.Wrap(err, "failed to delete list")
}
