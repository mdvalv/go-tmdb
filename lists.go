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

// List represents a list in TMDb.
type List struct {
	CreatedBy     string      `json:"created_by"`
	Description   string      `json:"description"`
	FavoriteCount int         `json:"favorite_count"`
	ID            string      `json:"id"`
	ISO6391       string      `json:"iso_639_1"`
	ItemCount     int         `json:"item_count"`
	Items         []MovieOrTV `json:"items"`
	Name          string      `json:"name"`
	PosterPath    *string     `json:"poster_path"`
}

// ListOptions represents the available options for the request.
type ListOptions languageOptions

// GetList retrieves the details of a list.
func (lr *ListsResource) GetList(listID string, opt *ListOptions) (*List, *http.Response, error) {
	path := fmt.Sprintf("/list/%s", listID)
	var list List
	resp, err := lr.client.get(path, &list, WithQueryParams(opt))
	return &list, resp, errors.Wrap(err, "failed to get list")
}

// ItemStatus represents an item status in TMDb.
type ItemStatus struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}

// GetItemStatus checks if a movie has already been added to the list.
func (lr *ListsResource) GetItemStatus(listID string, movieID int) (*ItemStatus, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/item_status", listID)
	var status ItemStatus
	resp, err := lr.client.get(path, &status, WithQueryParam("movie_id", fmt.Sprint(movieID)))
	return &status, resp, errors.Wrap(err, "failed to get item status")
}

// CreateList represents list to be created in TMDb.
type CreateList struct {
	Name        string `url:"name,omitempty" json:"name,omitempty"`
	Description string `url:"description,omitempty" json:"description,omitempty"`
	Language    string `url:"language,omitempty" json:"language,omitempty"`
}

// CreateListResponse represents response for creating a list in TMDb.
type CreateListResponse struct {
	ListID        int    `json:"list_id"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// CreateList creates a list.
func (lr *ListsResource) CreateList(sessionID string, list CreateList) (*CreateListResponse, *http.Response, error) {
	path := "/list"
	var response CreateListResponse
	resp, err := lr.client.post(path, &response, WithBody(list), WithSessionID(sessionID))
	return &response, resp, errors.Wrap(err, "failed to get item status")
}

// AddItemResponse represents the response for adding an item to a list.
type AddItemResponse statusResponse

// AddMovie adds a movie to a list.
func (lr *ListsResource) AddMovie(sessionID, listID string, itemID int) (*AddItemResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/add_item", listID)
	var response AddItemResponse
	resp, err := lr.client.post(path, &response, WithQueryParam("media_id", fmt.Sprint(itemID)), WithSessionID(sessionID))
	return &response, resp, errors.Wrap(err, "failed to add movie")
}

// RemoveItemResponse represents the response for removing an item from a list.
type RemoveItemResponse statusResponse

// RemoveMovie removes a movie from a list.
func (lr *ListsResource) RemoveMovie(sessionID, listID string, itemID int) (*RemoveItemResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/remove_item", listID)
	var response RemoveItemResponse
	resp, err := lr.client.post(path, &response, WithQueryParam("media_id", fmt.Sprint(itemID)), WithSessionID(sessionID))
	return &response, resp, errors.Wrap(err, "failed to remove movie")
}

// ClearListResponse represents the response for clearing all items from a list.
type ClearListResponse statusResponse

// Clear clears all of the items from a list.
func (lr *ListsResource) Clear(sessionID, listID string) (*ClearListResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s/clear", listID)
	var response ClearListResponse
	resp, err := lr.client.post(path, &response, WithQueryParam("confirm", "true"), WithSessionID(sessionID))
	return &response, resp, errors.Wrap(err, "failed to clear list")
}

// DeleteListResponse represents the response for deleting a list.
type DeleteListResponse statusResponse

// Delete deletes a list.
func (lr *ListsResource) Delete(sessionID, listID string) (*DeleteListResponse, *http.Response, error) {
	path := fmt.Sprintf("/list/%s", listID)
	var response DeleteListResponse
	resp, err := lr.client.delete(path, &response, WithSessionID(sessionID))
	return &response, resp, errors.Wrap(err, "failed to delete list")
}
