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

// TVEpisodeDetails represents tv episode details in TMDb.
type TVEpisodeDetails struct {
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

	// append to response
	Credits      *TVEpisodeCredits      `json:"credits"`
	ExternalIDs  *TVEpisodeExternalIDs  `json:"external_ids"`
	Images       *TVEpisodeImages       `json:"images"`
	Translations *TVEpisodeTranslations `json:"translations"`
	Videos       *Videos                `json:"videos"`
}

// TVEpisodeDetailsOptions represents the available options for the request.
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

// GetEpisode retrieves the TV episode details by id.
func (tr *TVEpisodesResource) GetEpisode(tvID, seasonNumber, episodeNumber int, opt *TVEpisodeDetailsOptions) (*TVEpisodeDetails, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d", tvID, seasonNumber, episodeNumber)
	var episode TVEpisodeDetails
	resp, err := tr.client.get(path, &episode, WithQueryParams(opt))
	return &episode, resp, errors.Wrap(err, "failed to get episode")
}

// AccountStatesEpisode represents account states for a episode in TMDb.
type AccountStatesEpisode struct {
	ID    int         `json:"id"`
	Rated interface{} `json:"rated"`
}

// GetAccountStates returns all of the user ratings for the season's episodes.
func (tr *TVEpisodesResource) GetAccountStates(tvID, seasonNumber, episodeNumber int, sessionID string) (*AccountStatesEpisode, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/account_states", tvID, seasonNumber, episodeNumber)
	var states AccountStatesEpisode
	resp, err := tr.client.get(path, &states, WithSessionID(sessionID))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

// GetChanges retrieves the changes for a TV episode. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the start_date and end_date query parameters.
func (tr *TVEpisodesResource) GetChanges(episodeID int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/tv/episode/%d/changes", episodeID)
	var changes Changes
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get episode changes")
}

// TVEpisodeCredits represents tv episode credits in TMDb.
type TVEpisodeCredits struct {
	Cast       []TVShowCast `json:"cast"`
	Crew       []TVShowCrew `json:"crew"`
	GuestStars []TVShowCast `json:"guest_stars"`
	ID         *int         `json:"id"`
}

// GetCredits retrieves the credits (cast, crew and guest stars) for a TV episode.
func (tr *TVEpisodesResource) GetCredits(tvID, seasonNumber, episodeNumber int, opt *CreditsOptions) (*TVEpisodeCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/credits", tvID, seasonNumber, episodeNumber)
	var credits TVEpisodeCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get credits")
}

// TVEpisodeExternalIDs represents tv episode external ids in TMDb.
type TVEpisodeExternalIDs struct {
	ID          *int    `json:"id"`
	FreebaseID  *string `json:"freebase_id"`
	FreebaseMID *string `json:"freebase_mid"`
	IMDbID      *string `json:"imdb_id"`
	TVDbID      *int    `json:"tvdb_id"`
	TVRageID    *int    `json:"tvrage_id"`
}

// GetExternalIDs retrieves the external ids for a TV season.
func (tr *TVEpisodesResource) GetExternalIDs(tvID, seasonNumber, episodeNumber int, opt *ExternalIDsOptions) (*TVEpisodeExternalIDs, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/external_ids", tvID, seasonNumber, episodeNumber)
	var ids TVEpisodeExternalIDs
	resp, err := tr.client.get(path, &ids, WithQueryParams(opt))
	return &ids, resp, errors.Wrap(err, "failed to get external ids")
}

// Still represents a still in TMDb.
type Still Image

// TVEpisodeImages represents tv episode images in TMDb.
type TVEpisodeImages struct {
	ID     *int    `json:"id"`
	Stills []Still `json:"stills"`
}

// GetImages retrieves the images that belong to a TV episode.
// Querying images with a language parameter will filter the results.
// To include a fallback language (especially useful for backdrops), use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (tr *TVEpisodesResource) GetImages(tvID, seasonNumber, episodeNumber int, opt *ImagesOptions) (*TVEpisodeImages, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/images", tvID, seasonNumber, episodeNumber)
	var images TVEpisodeImages
	resp, err := tr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

// TVEpisodeData represents tv episode data in TMDb.
type TVEpisodeData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

// TVEpisodeTranslation represents a tv episode translation in TMDb.
type TVEpisodeTranslation struct {
	ISO31661    string        `json:"iso_3166_1"`
	ISO6391     string        `json:"iso_639_1"`
	Name        string        `json:"name"`
	EnglishName string        `json:"english_name"`
	Data        TVEpisodeData `json:"data"`
}

// TVEpisodeTranslations represents tv episode translations in TMDb.
type TVEpisodeTranslations struct {
	ID           *int                   `json:"id"`
	Translations []TVEpisodeTranslation `json:"translations"`
}

// GetTranslations retrieves a list of the translations that exist for a TV show.
func (tr *TVEpisodesResource) GetTranslations(tvID, seasonNumber, episodeNumber int) (*TVEpisodeTranslations, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/translations", tvID, seasonNumber, episodeNumber)
	var translations TVEpisodeTranslations
	resp, err := tr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

// GetVideos retrieves the videos that have been added to a TV season.
func (tr *TVEpisodesResource) GetVideos(tvID, seasonNumber, episodeNumber int, opt *VideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/videos", tvID, seasonNumber, episodeNumber)
	var videos Videos
	resp, err := tr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get tv show videos")
}

// Rate rates a TV episode.
// A valid session or guest session ID is required.
func (tr *TVEpisodesResource) Rate(tvID, seasonNumber, episodeNumber int, rating float64, sessionID Auth) (*RateResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/rating", tvID, seasonNumber, episodeNumber)
	var response RateResponse
	resp, err := tr.client.post(path, &response, WithBody(map[string]float64{"value": rating}), WithQueryParams(sessionID))
	return &response, resp, errors.Wrap(err, "failed to rate tv show episode")
}

// DeleteRating removes a rating for a TV episode.
// A valid session or guest session ID is required.
func (tr *TVEpisodesResource) DeleteRating(tvID, seasonNumber, episodeNumber int, sessionID Auth) (*DeleteRatingResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/season/%d/episode/%d/rating", tvID, seasonNumber, episodeNumber)
	var response DeleteRatingResponse
	resp, err := tr.client.delete(path, &response, WithQueryParams(sessionID))
	return &response, resp, errors.Wrap(err, "failed to delete tv show episode rating")
}
