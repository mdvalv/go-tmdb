package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// CreditsResource handles credit-related requests of TMDb API.
type CreditsResource struct {
	client *Client
}

type Season struct {
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}

type Media struct {
	Adult            bool      `json:"adult"`
	BackdropPath     string    `json:"backdrop_path"`
	Character        string    `json:"character"`
	Episodes         []episode `json:"episodes"`
	FirstAirDate     string    `json:"first_air_date"`
	GenreIds         []int     `json:"genre_ids"`
	Id               int       `json:"id"`
	MediaType        string    `json:"media_type"`
	Name             string    `json:"name"`
	OriginalLanguage string    `json:"original_language"`
	OriginalName     string    `json:"original_name"`
	OriginCountry    []string  `json:"origin_country"`
	Overview         string    `json:"overview"`
	Popularity       float64   `json:"popularity"`
	PosterPath       string    `json:"poster_path"`
	Seasons          []Season  `json:"seasons"`
	VoteAverage      float64   `json:"vote_average"`
	VoteCount        int       `json:"vote_count"`
}

type Credit struct {
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Id         string `json:"id"`
	Job        string `json:"job"`
	Media      Media  `json:"media"`
	MediaType  string `json:"media_type"`
	Person     person `json:"person"`
}

// Get a movie or TV credit details by id.
func (cr *CreditsResource) GetCredit(id string) (*Credit, *http.Response, error) {
	path := fmt.Sprintf("/credit/%s", id)
	var credit Credit
	resp, err := cr.client.get(path, &credit)
	return &credit, resp, errors.Wrap(err, "failed to get credit")
}
