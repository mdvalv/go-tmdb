package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// CollectionsResource handles collection-related requests of TMDb API.
type CollectionsResource struct {
	client *Client
}

type Part struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	PosterPath       *string `json:"poster_path"`
	Popularity       float64 `json:"popularity"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type Collection struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	BackdropPath *string `json:"backdrop_path"`
	Parts        []Part  `json:"parts"`
}

type CollectionsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// Get collection details by id.
func (cr *CollectionsResource) GetCollection(id int, opt *CollectionsOptions) (*Collection, *http.Response, error) {
	path := fmt.Sprintf("/collection/%d", id)
	var collection Collection
	resp, err := cr.client.get(path, &collection, WithQueryParams(opt))
	return &collection, resp, errors.Wrap(err, "failed to get collection")
}

type image struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	ISO6391     *string `json:"iso_639_1"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

type Backdrop image

type Poster image

type CollectionImages struct {
	Id        int        `json:"id"`
	Backdrops []Backdrop `json:"backdrops"`
	Posters   []Poster   `json:"posters"`
}

// Get the images for a collection by id.
func (cr *CollectionsResource) GetImages(id int, opt *CollectionsOptions) (*CollectionImages, *http.Response, error) {
	path := fmt.Sprintf("/collection/%d/images", id)
	var images CollectionImages
	resp, err := cr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get collection images")
}

type CollectionData struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
}

type CollectionTranslation struct {
	ISO31661    string         `json:"iso_3166_1"`
	ISO6391     string         `json:"iso_639_1"`
	Name        string         `json:"name"`
	EnglishName string         `json:"english_name"`
	Data        CollectionData `json:"data"`
}

type CollectionTranslations struct {
	Id           int                     `json:"id"`
	Translations []CollectionTranslation `json:"translations"`
}

// Get the list translations for a collection by id.
func (cr *CollectionsResource) GetTranslations(id int, opt *CollectionsOptions) (*CollectionTranslations, *http.Response, error) {
	path := fmt.Sprintf("/collection/%d/translations", id)
	var translations CollectionTranslations
	resp, err := cr.client.get(path, &translations, WithQueryParams(opt))
	return &translations, resp, errors.Wrap(err, "failed to get collection translations")
}
