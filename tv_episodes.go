package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TVEpisodesResource handles tv-related requests of TMDb API.
type TVEpisodesResource struct {
	client *Client
}

type TVEpisodeDetails struct {
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

	// append to response
	Credits      *TVEpisodeCredits      `json:"credits"`
	ExternalIds  *TVEpisodeExternalIds  `json:"external_ids"`
	Images       *TVEpisodeImages       `json:"images"`
	Translations *TVEpisodeTranslations `json:"translations"`
	Videos       *Videos                `json:"videos"`
}

type TVEpisodeDetailsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// supported values:
	// - credits
	// - external_ids
	// - images
	// - translations
	// - videos
	// provide them separated by commas, example: images,videos
	AppendToResponse string `url:"append_to_response,omitempty" json:"append_to_response,omitempty"`
}

// Get the TV episode details by id.
func (tr *TVEpisodesResource) GetEpisode(tvId, seasonNumber, episodeNumber int, opt *TVEpisodeDetailsOptions) (*TVEpisodeDetails, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d", tvId, seasonNumber, episodeNumber)
	var episode TVEpisodeDetails
	resp, err := tr.client.get(path, &episode, WithQueryParams(opt))
	return &episode, resp, errors.Wrap(err, "failed to get episode")
}

type AccountStatesEpisode struct {
	Id    int         `json:"id"`
	Rated interface{} `json:"rated"`
}

// Returns all of the user ratings for the season's episodes.
func (tr *TVEpisodesResource) GetAccountStates(tvId, seasonNumber, episodeNumber int, sessionId string) (*AccountStatesEpisode, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/account_states", tvId, seasonNumber, episodeNumber)
	var states AccountStatesEpisode
	resp, err := tr.client.get(path, &states, WithQueryParam("session_id", sessionId))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

// Get the changes for a TV episode. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the start_date and end_date query parameters.
func (tr *TVEpisodesResource) GetChanges(episodeId int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/tv/episode/%d/changes", episodeId)
	var changes Changes
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get episode changes")
}

type TVEpisodeCredits struct {
	Cast       []TVShowCast `json:"cast"`
	Crew       []TVShowCrew `json:"crew"`
	GuestStars []TVShowCast `json:"guest_stars"`
	Id         *int         `json:"id"`
}

// Get the credits (cast, crew and guest stars) for a TV episode.
func (tr *TVEpisodesResource) GetCredits(tvId, seasonNumber, episodeNumber int, opt *CreditsOptions) (*TVEpisodeCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/credits", tvId, seasonNumber, episodeNumber)
	var credits TVEpisodeCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get credits")
}

type TVEpisodeExternalIds struct {
	Id          *int    `json:"id"`
	FreebaseId  *string `json:"freebase_id"`
	FreebaseMId *string `json:"freebase_mid"`
	IMDbId      *string `json:"imdb_id"`
	TVDbId      *int    `json:"tvdb_id"`
	TVRageId    *int    `json:"tvrage_id"`
}

// Get the external ids for a TV season.
func (tr *TVEpisodesResource) GetExternalIds(tvId, seasonNumber, episodeNumber int, opt *ExternalIdsOptions) (*TVEpisodeExternalIds, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/external_ids", tvId, seasonNumber, episodeNumber)
	var ids TVEpisodeExternalIds
	resp, err := tr.client.get(path, &ids, WithQueryParams(opt))
	return &ids, resp, errors.Wrap(err, "failed to get external ids")
}

type Still Image

type TVEpisodeImages struct {
	Id     *int    `json:"id"`
	Stills []Still `json:"stills"`
}

// Get the images that belong to a TV episode.
// Querying images with a language parameter will filter the results.
// To include a fallback language (especially useful for backdrops), use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (tr *TVEpisodesResource) GetImages(tvId, seasonNumber, episodeNumber int, opt *ImagesOptions) (*TVEpisodeImages, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/images", tvId, seasonNumber, episodeNumber)
	var images TVEpisodeImages
	resp, err := tr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

type TVEpisodeData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

type TVEpisodeTranslation struct {
	ISO31661    string        `json:"iso_3166_1"`
	ISO6391     string        `json:"iso_639_1"`
	Name        string        `json:"name"`
	EnglishName string        `json:"english_name"`
	Data        TVEpisodeData `json:"data"`
}

type TVEpisodeTranslations struct {
	Id           *int                   `json:"id"`
	Translations []TVEpisodeTranslation `json:"translations"`
}

// Get a list of the translations that exist for a TV show.
func (tr *TVEpisodesResource) GetTranslations(tvId, seasonNumber, episodeNumber int) (*TVEpisodeTranslations, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/translations", tvId, seasonNumber, episodeNumber)
	var translations TVEpisodeTranslations
	resp, err := tr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

// Get the videos that have been added to a TV season.
func (tr *TVEpisodesResource) GetVideos(tvId, seasonNumber, episodeNumber int, opt *VideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/videos", tvId, seasonNumber, episodeNumber)
	var videos Videos
	resp, err := tr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get tv show videos")
}

// Rate a TV episode.
// A valid session or guest session ID is required.
func (tr *TVEpisodesResource) Rate(tvId, seasonNumber, episodeNumber int, rating float64, sessionId Auth) (*RateResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/rating", tvId, seasonNumber, episodeNumber)
	var response RateResponse
	resp, err := tr.client.post(path, &response, WithBody(map[string]float64{"value": rating}), WithQueryParams(sessionId))
	return &response, resp, errors.Wrap(err, "failed to rate tv show episode")
}

// Remove a rating for a TV episode.
// A valid session or guest session ID is required.
func (tr *TVEpisodesResource) DeleteRating(tvId, seasonNumber, episodeNumber int, sessionId Auth) (*DeleteRatingResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/rating", tvId, seasonNumber, episodeNumber)
	var response DeleteRatingResponse
	resp, err := tr.client.delete(path, &response, WithQueryParams(sessionId))
	return &response, resp, errors.Wrap(err, "failed to delete tv show episode rating")
}
