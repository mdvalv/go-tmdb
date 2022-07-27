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

type AuthorDetails struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	AvatarPath string  `json:"avatar_path"`
	Rating     float64 `json:"rating"`
}

type Review struct {
	Id            string        `json:"id"`
	Author        string        `json:"author"`
	AuthorDetails AuthorDetails `json:"author_details"`
	Content       string        `json:"content"`
	CreatedAt     string        `json:"created_at"`
	ISO6391       string        `json:"iso_639_1"`
	MediaId       int           `json:"media_id"`
	MediaTitle    string        `json:"media_title"`
	MediaType     string        `json:"media_type"`
	UpdatedAt     string        `json:"updated_at"`
	Url           string        `json:"url"`
}

// Retrieve the details of a movie or TV show review.
func (rr *ReviewsResource) GetReview(id string) (*Review, *http.Response, error) {
	path := fmt.Sprintf("/review/%s", id)
	var review Review
	resp, err := rr.client.getResource(path, nil, &review)
	return &review, resp, errors.Wrap(err, "failed to get review")
}
