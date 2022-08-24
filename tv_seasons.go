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

// SeasonEpisode represents a season episode in TMDb.
type SeasonEpisode struct {
	AirDate        string       `json:"air_date"`
	Crew           []TVShowCrew `json:"crew"`
	EpisodeNumber  int          `json:"episode_number"`
	GuestStars     []TVShowCast `json:"guest_stars"`
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Overview       string       `json:"overview"`
	ProductionCode *string      `json:"production_code"`
	Runtime        int          `json:"runtime"`
	SeasonNumber   int          `json:"season_number"`
	StillPath      *string      `json:"still_path"`
	VoteAverage    float64      `json:"vote_average"`
	VoteCount      int          `json:"vote_count"`
}

// TVSeasonDetails represents season details in TMDb.
type TVSeasonDetails struct {
	ID           int             `json:"id"`
	AirDate      string          `json:"air_date"`
	Episodes     []SeasonEpisode `json:"episodes"`
	Name         string          `json:"name"`
	Overview     string          `json:"overview"`
	PosterPath   *string         `json:"poster_path"`
	SeasonNumber int             `json:"season_number"`

	// append to response
	AggregateCredits *AggregateCredits     `json:"aggregate_credits"`
	Credits          *TVShowCredits        `json:"credits"`
	ExternalIDs      *TVSeasonExternalIDs  `json:"external_ids"`
	Images           *TVSeasonImages       `json:"images"`
	Translations     *TVSeasonTranslations `json:"translations"`
	Videos           *Videos               `json:"videos"`
}

// TVSeasonDetailsOptions represents the available options for the request.
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

// GetSeason retrieves the TV season details by id.
func (tr *TVSeasonsResource) GetSeason(tvID, seasonNumber int, opt *TVSeasonDetailsOptions) (*TVSeasonDetails, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d", tvID, seasonNumber)
	var season TVSeasonDetails
	resp, err := tr.client.get(path, &season, WithQueryParams(opt))
	return &season, resp, errors.Wrap(err, "failed to get season")
}

// AccountStateSeason represents account state for a season in TMDb.
type AccountStateSeason struct {
	ID            int         `json:"id"`
	EpisodeNumber int         `json:"episode_number"`
	Rated         interface{} `json:"rated"`
}

// AccountStatesSeason represents account states for a season in TMDb.
type AccountStatesSeason struct {
	ID      int                  `json:"id"`
	Results []AccountStateSeason `json:"results"`
}

// GetAccountStates returns all of the user ratings for the season's episodes.
func (tr *TVSeasonsResource) GetAccountStates(tvID, seasonNumber int, sessionID string) (*AccountStatesSeason, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/account_states", tvID, seasonNumber)
	var states AccountStatesSeason
	resp, err := tr.client.get(path, &states, WithSessionID(sessionID))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

// GetAggregateCredits retrieves the aggregate credits for TV season.
// This call differs from the main credits call in that it does not only return the season credits,
// but rather is a view of all the cast & crew for all of the episodes belonging to a season.
func (tr *TVSeasonsResource) GetAggregateCredits(tvID, seasonNumber int, opt *AggregateCreditsOptions) (*AggregateCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/aggregate_credits", tvID, seasonNumber)
	var credits AggregateCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get aggregate credits")
}

// GetChanges retrieves the changes for a TV season. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the start_date and end_date query parameters.
func (tr *TVSeasonsResource) GetChanges(seasonID int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/tv/season/%d/changes", seasonID)
	var changes Changes
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get season changes")
}

// GetCredits retrieves the credits for TV season.
func (tr *TVSeasonsResource) GetCredits(tvID, seasonNumber int, opt *CreditsOptions) (*TVShowCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/credits", tvID, seasonNumber)
	var credits TVShowCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get credits")
}

// TVSeasonExternalIDs represents season external ids in TMDb.
type TVSeasonExternalIDs struct {
	ID          *int    `json:"id"`
	FreebaseID  *string `json:"freebase_id"`
	FreebaseMID *string `json:"freebase_mid"`
	TVDbID      *int    `json:"tvdb_id"`
	TVRageID    *int    `json:"tvrage_id"`
}

// GetExternalIDs retrieves the external ids for a TV season.
func (tr *TVSeasonsResource) GetExternalIDs(tvID, seasonNumber int, opt *ExternalIDsOptions) (*TVSeasonExternalIDs, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/external_ids", tvID, seasonNumber)
	var ids TVSeasonExternalIDs
	resp, err := tr.client.get(path, &ids, WithQueryParams(opt))
	return &ids, resp, errors.Wrap(err, "failed to get external ids")
}

// TVSeasonImages represents season images in TMDb.
type TVSeasonImages struct {
	ID      *int     `json:"id"`
	Posters []Poster `json:"posters"`
}

// GetImages retrieves the images that belong to a TV season.
// Querying images with a language parameter will filter the results.
// To include a fallback language (especially useful for backdrops), use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (tr *TVSeasonsResource) GetImages(tvID, seasonNumber int, opt *ImagesOptions) (*TVSeasonImages, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/images", tvID, seasonNumber)
	var images TVSeasonImages
	resp, err := tr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

// TVSeasonData represents season data in TMDb.
type TVSeasonData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

// TVSeasonTranslation represents a season translation in TMDb.
type TVSeasonTranslation struct {
	ISO31661    string       `json:"iso_3166_1"`
	ISO6391     string       `json:"iso_639_1"`
	Name        string       `json:"name"`
	EnglishName string       `json:"english_name"`
	Data        TVSeasonData `json:"data"`
}

// TVSeasonTranslations represents season translations in TMDb.
type TVSeasonTranslations struct {
	ID           int                   `json:"id"`
	Translations []TVSeasonTranslation `json:"translations"`
}

// GetTranslations retrieves a list of the translations that exist for a TV show.
func (tr *TVSeasonsResource) GetTranslations(tvID, seasonNumber int) (*TVSeasonTranslations, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/translations", tvID, seasonNumber)
	var translations TVSeasonTranslations
	resp, err := tr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

// GetVideos retrieves the videos that have been added to a TV season.
func (tr *TVSeasonsResource) GetVideos(tvID, seasonNumber int, opt *VideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/videos", tvID, seasonNumber)
	var videos Videos
	resp, err := tr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get tv show videos")
}
