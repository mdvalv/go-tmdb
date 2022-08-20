package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TVSeasonsResource handles tv-related requests of TMDb API.
type TVSeasonsResource struct {
	client *Client
}

type SeasonEpisode struct {
	AirDate        string       `json:"air_date"`
	Crew           []TVShowCrew `json:"crew"`
	EpisodeNumber  int          `json:"episode_number"`
	GuestStars     []TVShowCast `json:"guest_stars"`
	Id             int          `json:"id"`
	Name           string       `json:"name"`
	Overview       string       `json:"overview"`
	ProductionCode *string      `json:"production_code"`
	Runtime        int          `json:"runtime"`
	SeasonNumber   int          `json:"season_number"`
	StillPath      *string      `json:"still_path"`
	VoteAverage    float64      `json:"vote_average"`
	VoteCount      int          `json:"vote_count"`
}

type TVSeasonDetails struct {
	Id           int             `json:"id"`
	AirDate      string          `json:"air_date"`
	Episodes     []SeasonEpisode `json:"episodes"`
	Name         string          `json:"name"`
	Overview     string          `json:"overview"`
	PosterPath   *string         `json:"poster_path"`
	SeasonNumber int             `json:"season_number"`

	// append to response
	AggregateCredits *AggregateCredits     `json:"aggregate_credits"`
	Credits          *TVShowCredits        `json:"credits"`
	ExternalIds      *TVSeasonExternalIds  `json:"external_ids"`
	Images           *TVSeasonImages       `json:"images"`
	Translations     *TVSeasonTranslations `json:"translations"`
	Videos           *Videos               `json:"videos"`
}

type TVSeasonDetailsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// supported values:
	// - aggregate_credits
	// - credits
	// - external_ids
	// - images
	// - translations
	// - videos
	// provide them separated by commas, example: images,videos
	AppendToResponse string `url:"append_to_response,omitempty" json:"append_to_response,omitempty"`
}

// Get the TV season details by id.
func (tr *TVSeasonsResource) GetSeason(tvId, seasonNumber int, opt *TVSeasonDetailsOptions) (*TVSeasonDetails, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d", tvId, seasonNumber)
	var season TVSeasonDetails
	resp, err := tr.client.get(path, &season, WithQueryParams(opt))
	return &season, resp, errors.Wrap(err, "failed to get season")
}

type AccountStateSeason struct {
	Id            int         `json:"id"`
	EpisodeNumber int         `json:"episode_number"`
	Rated         interface{} `json:"rated"`
}

type AccountStatesSeason struct {
	Id      int                  `json:"id"`
	Results []AccountStateSeason `json:"results"`
}

// Returns all of the user ratings for the season's episodes.
func (tr *TVSeasonsResource) GetAccountStates(tvId, seasonNumber int, sessionId string) (*AccountStatesSeason, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/account_states", tvId, seasonNumber)
	var states AccountStatesSeason
	resp, err := tr.client.get(path, &states, WithQueryParam("session_id", sessionId))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

// Get the aggregate credits for TV season.
// This call differs from the main credits call in that it does not only return the season credits,
// but rather is a view of all the cast & crew for all of the episodes belonging to a season.
func (tr *TVSeasonsResource) GetAggregateCredits(tvId, seasonNumber int, opt *AggregateCreditsOptions) (*AggregateCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/aggregate_credits", tvId, seasonNumber)
	var credits AggregateCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get aggregate credits")
}

// Get the changes for a TV season. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the start_date and end_date query parameters.
func (tr *TVSeasonsResource) GetChanges(seasonId int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/tv/season/%d/changes", seasonId)
	var changes Changes
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get season changes")
}

// Get the credits for TV season.
func (tr *TVSeasonsResource) GetCredits(tvId, seasonNumber int, opt *CreditsOptions) (*TVShowCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/credits", tvId, seasonNumber)
	var credits TVShowCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get credits")
}

type TVSeasonExternalIds struct {
	Id          *int    `json:"id"`
	FreebaseId  *string `json:"freebase_id"`
	FreebaseMId *string `json:"freebase_mid"`
	TVDbId      *int    `json:"tvdb_id"`
	TVRageId    *int    `json:"tvrage_id"`
}

// Get the external ids for a TV season.
func (tr *TVSeasonsResource) GetExternalIds(tvId, seasonNumber int, opt *ExternalIdsOptions) (*TVSeasonExternalIds, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/external_ids", tvId, seasonNumber)
	var ids TVSeasonExternalIds
	resp, err := tr.client.get(path, &ids, WithQueryParams(opt))
	return &ids, resp, errors.Wrap(err, "failed to get external ids")
}

type TVSeasonImages struct {
	Id      *int     `json:"id"`
	Posters []Poster `json:"posters"`
}

// Get the images that belong to a TV season.
// Querying images with a language parameter will filter the results.
// To include a fallback language (especially useful for backdrops), use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (tr *TVSeasonsResource) GetImages(tvId, seasonNumber int, opt *ImagesOptions) (*TVSeasonImages, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/images", tvId, seasonNumber)
	var images TVSeasonImages
	resp, err := tr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

type TVSeasonData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

type TVSeasonTranslation struct {
	ISO31661    string       `json:"iso_3166_1"`
	ISO6391     string       `json:"iso_639_1"`
	Name        string       `json:"name"`
	EnglishName string       `json:"english_name"`
	Data        TVSeasonData `json:"data"`
}

type TVSeasonTranslations struct {
	Id           int                   `json:"id"`
	Translations []TVSeasonTranslation `json:"translations"`
}

// Get a list of the translations that exist for a TV show.
func (tr *TVSeasonsResource) GetTranslations(tvId, seasonNumber int) (*TVSeasonTranslations, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/translations", tvId, seasonNumber)
	var translations TVSeasonTranslations
	resp, err := tr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

// Get the videos that have been added to a TV season.
func (tr *TVSeasonsResource) GetVideos(tvId, seasonNumber int, opt *VideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/videos", tvId, seasonNumber)
	var videos Videos
	resp, err := tr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get tv show videos")
}
