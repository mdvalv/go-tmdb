package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ReviewsResource handles review-related requests of TMDb API.
type ReviewsResource struct {
	client *Client
}

// AuthorDetails represents author details in TMDb.
type AuthorDetails struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	AvatarPath string  `json:"avatar_path"`
	Rating     float64 `json:"rating"`
}

// Review represents review in TMDb.
type Review struct {
	ID            string        `json:"id"`
	Author        string        `json:"author"`
	AuthorDetails AuthorDetails `json:"author_details"`
	Content       string        `json:"content"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	URL           string        `json:"url"`
}

// ReviewDetails represents review details in TMDb.
type ReviewDetails struct {
	Author        string        `json:"author"`
	AuthorDetails AuthorDetails `json:"author_details"`
	Content       string        `json:"content"`
	CreatedAt     string        `json:"created_at"`
	ID            string        `json:"id"`
	ISO6391       string        `json:"iso_639_1"`
	MediaID       int           `json:"media_id"`
	MediaTitle    string        `json:"media_title"`
	MediaType     string        `json:"media_type"`
	UpdatedAt     string        `json:"updated_at"`
	URL           string        `json:"url"`
}

// GetReview retrieves the details of a movie or TV show review.
func (rr *ReviewsResource) GetReview(id string) (*ReviewDetails, *http.Response, error) {
	path := fmt.Sprintf("/review/%s", id)
	var review ReviewDetails
	resp, err := rr.client.get(path, &review)
	return &review, resp, errors.Wrap(err, "failed to get review")
}
