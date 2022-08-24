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

// GetMediaType retrieves the media type from a media credit.
func (mc MediaCredit) GetMediaType() string {
	return mc["media_type"].(string)
}

// ToMovieCredit converts the data to a movie credit.
func (mc MediaCredit) ToMovieCredit() (*MovieCredit, error) {
	var credit MovieCredit
	err := convert("movie", mc, &credit)
	return &credit, errors.Wrap(err, "failed to convert object to movie credit")
}

// ToTVShowCredit converts the data to a tv show credit.
func (mc MediaCredit) ToTVShowCredit() (*TVShowCredit, error) {
	var credit TVShowCredit
	err := convert("tv", mc, &credit)
	return &credit, errors.Wrap(err, "failed to convert object to tv credit")
}

// MovieCredit represents a movie credit in TMDb.
type MovieCredit struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	Character        string  `json:"character"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	MediaType        string  `json:"media_type"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       *string `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

// TVShowCredit represents a tv show credit in TMDb.
type TVShowCredit struct {
	Adult            bool        `json:"adult"`
	BackdropPath     *string     `json:"backdrop_path"`
	Episodes         []TVEpisode `json:"episodes"`
	FirstAirDate     string      `json:"first_air_date"`
	GenreIDs         []int       `json:"genre_ids"`
	ID               int         `json:"id"`
	MediaType        string      `json:"media_type"`
	Name             string      `json:"name"`
	OriginalLanguage string      `json:"original_language"`
	OriginalName     string      `json:"original_name"`
	OriginCountry    []string    `json:"origin_country"`
	Overview         string      `json:"overview"`
	Popularity       float64     `json:"popularity"`
	PosterPath       *string     `json:"poster_path"`
	Seasons          []Season    `json:"seasons"`
	VoteAverage      float64     `json:"vote_average"`
	VoteCount        int         `json:"vote_count"`
}

// MediaCredit represents a media credit from TMSb.
type MediaCredit map[string]interface{}

// PersonCredit represents a person credit in TMDb.
type PersonCredit struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	MediaType          string  `json:"media_type"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
}

// Credit represents a credit in TMDb.
type Credit struct {
	CreditType string       `json:"credit_type"`
	Department string       `json:"department"`
	ID         string       `json:"id"`
	Job        string       `json:"job"`
	Media      MediaCredit  `json:"media"`
	MediaType  string       `json:"media_type"`
	Person     PersonCredit `json:"person"`
}

// GetCredit retrieves a movie or TV credit details by id.
func (cr *CreditsResource) GetCredit(id string) (*Credit, *http.Response, error) {
	path := fmt.Sprintf("/credit/%s", id)
	var credit Credit
	resp, err := cr.client.get(path, &credit)
	return &credit, resp, errors.Wrap(err, "failed to get credit")
}
