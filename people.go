package tmdb

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// PeopleResource handles person-related requests of TMDb API.
type PeopleResource struct {
	client *Client
}

type personKnownFor struct {
	person
	KnownFor []KnownFor `json:"known_for"`
}

type person struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	MediaType          string  `json:"media_type"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
}

type KnownFor map[string]interface{}

func (kf KnownFor) GetMediaType() string {
	return kf["media_type"].(string)
}

func (kf KnownFor) ToMovie() (*Movie, error) {
	if kf.GetMediaType() != "movie" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to movie", kf.GetMediaType()))
	}
	return convertToMovie(kf)
}

func (kf KnownFor) ToTVShow() (*TVShow, error) {
	if kf.GetMediaType() != "tv" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to tv", kf.GetMediaType()))
	}
	return convertToTVShow(kf)
}

func convertToMovie(obj interface{}) (*Movie, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var movie Movie
	err = json.Unmarshal(result, &movie)
	return &movie, err
}

func convertToTVShow(obj interface{}) (*TVShow, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var tvShow TVShow
	err = json.Unmarshal(result, &tvShow)
	return &tvShow, err
}
