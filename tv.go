package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TVResource handles tv-related requests of TMDb API.
type TVResource struct {
	client *Client
}

// TVShowDetails represents tv show details in TMDb.
type TVShowDetails struct {
	Adult               bool                `json:"adult"`
	BackdropPath        *string             `json:"backdrop_path"`
	CreatedBy           []TVCreatedBy       `json:"created_by"`
	EpisodeRunTime      []int               `json:"episode_run_time"`
	FirstAirDate        string              `json:"first_air_date"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int                 `json:"id"`
	InProduction        bool                `json:"in_production"`
	Languages           []string            `json:"languages"`
	LastAirDate         string              `json:"last_air_date"`
	LastEpisodeToAir    *TVEpisode          `json:"last_episode_to_air"`
	Name                string              `json:"name"`
	Networks            []TVShowNetwork     `json:"networks"`
	NextEpisodeToAir    *TVEpisode          `json:"next_episode_to_air"`
	NumberOfEpisodes    int                 `json:"number_of_episodes"`
	NumberOfSeasons     int                 `json:"number_of_seasons"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalName        string              `json:"original_name"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          *string             `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	Seasons             []Season            `json:"seasons"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`

	// append to response
	AggregateCredits     *AggregateCredits        `json:"aggregate_credits"`
	AlternativeTitles    *TVShowAlternativeTitles `json:"alternative_titles"`
	Changes              *Changes                 `json:"changes"`
	ContentRatings       *ContentRatings          `json:"content_ratings"`
	Credits              *TVShowCredits           `json:"credits"`
	EpisodeGroups        *EpisodeGroups           `json:"episode_groups"`
	ExternalIDs          *TVShowExternalIDs       `json:"external_ids"`
	Images               *Images                  `json:"images"`
	Keywords             *TVShowKeywords          `json:"keywords"`
	Recommendations      *RecommendedTVShows      `json:"recommendations"`
	Reviews              *TVShowReviews           `json:"reviews"`
	ScreenedTheatrically *ScreenedTheatrically    `json:"screened_theatrically"`
	Similar              *SimilarTVShows          `json:"similar"`
	Translations         *TVShowTranslations      `json:"translations"`
	Videos               *Videos                  `json:"videos"`
}

// TVShowDetailsOptions represents the available options for the request.
type TVShowDetailsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// supported values:
	// - aggregate_credits
	// - alternative_titles
	// - changes
	// - content_ratings
	// - credits
	// - episode_groups
	// - external_ids
	// - images
	// - keywords
	// - recommendations
	// - reviews
	// - screened_theatrically
	// - similar
	// - translations
	// - videos
	// provide them separated by commas, example: images,videos
	AppendToResponse string `url:"append_to_response,omitempty" json:"append_to_response,omitempty"`
}

// GetTVShow retrieves the primary TV show details by id.
func (tr *TVResource) GetTVShow(tvID int, opt *TVShowDetailsOptions) (*TVShowDetails, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d", tvID)
	var tvShow TVShowDetails
	resp, err := tr.client.get(path, &tvShow, WithQueryParams(opt))
	return &tvShow, resp, errors.Wrap(err, "failed to get tv show")
}

// Season represents a tv season in TMDb.
type Season struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
}

// TVShow represents a tv show in TMDb.
type TVShow struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	MediaType        string   `json:"media_type"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// TVEpisode represents a tv episode in TMDb.
type TVEpisode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

// TVShowResult represents a tv show result in TMDb.
type TVShowResult struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

type paginatedTVShows struct {
	pagination
	TVShows []TVShowResult `json:"results"`
}

// GetAccountStates retrieves the following account states for a session:
// - Movie rating
// - If it belongs to the watchlist
// - If it belongs to the favorite list
func (tr *TVResource) GetAccountStates(tvID int, sessionID string) (*AccountStates, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/account_states", tvID)
	var states AccountStates
	resp, err := tr.client.get(path, &states, WithSessionID(sessionID))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

// AggregateCreditsOptions represents the available options for the request.
type AggregateCreditsOptions languageOptions

// Role represents a role in TMDb.
type Role struct {
	CreditID     string `json:"credit_id"`
	Character    string `json:"character"`
	EpisodeCount int    `json:"episode_count"`
}

// Job represents a job in TMDb.
type Job struct {
	CreditID     string `json:"credit_id"`
	Job          string `json:"job"`
	EpisodeCount int    `json:"episode_count"`
}

// AggregateCreditsCast represents aggregate credits cast in TMDb.
type AggregateCreditsCast struct {
	Adult              bool    `json:"adult"`
	Gender             *int    `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	Order              int     `json:"order"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	Roles              []Role  `json:"roles"`
	TotalEpisodeCount  int     `json:"total_episode_count"`
}

// AggregateCreditsCrew represents aggregate credits crew in TMDb.
type AggregateCreditsCrew struct {
	Adult              bool    `json:"adult"`
	Department         string  `json:"department"`
	Gender             *int    `json:"gender"`
	ID                 int     `json:"id"`
	Jobs               []Job   `json:"jobs"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	TotalEpisodeCount  int     `json:"total_episode_count"`
}

// AggregateCredits represents aggregate credits in TMDb.
type AggregateCredits struct {
	ID   *int                   `json:"id"`
	Cast []AggregateCreditsCast `json:"cast"`
	Crew []AggregateCreditsCrew `json:"crew"`
}

// GetAggregateCredits retrieves the aggregate credits (cast and crew) that have been added to a TV show.
// This call differs from the main `credits` call in that it does not return the newest season but rather,
// is a view of all the entire cast & crew for all episodes belonging to a TV show.
func (tr *TVResource) GetAggregateCredits(tvID int, opt *AggregateCreditsOptions) (*AggregateCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/aggregate_credits", tvID)
	var credits AggregateCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get aggregate credits")
}

// TVShowAlternativeTitlesOptions represents the available options for the request.
type TVShowAlternativeTitlesOptions languageOptions

// TVShowAlternativeTitles represents tv show alternative titles in TMDb.
type TVShowAlternativeTitles struct {
	ID     *int    `json:"id"`
	Titles []Title `json:"results"`
}

// GetAlternativeTitles retrieves all of the alternative titles for a tv show.
func (tr *TVResource) GetAlternativeTitles(tvID int, opt *TVShowAlternativeTitlesOptions) (*TVShowAlternativeTitles, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/alternative_titles", tvID)
	var titles TVShowAlternativeTitles
	resp, err := tr.client.get(path, &titles, WithQueryParams(opt))
	return &titles, resp, errors.Wrap(err, "failed to get alternative titles")
}

// GetChanges retrieves the changes for a TV show. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the start_date and end_date query parameters.
// TV show changes are different than movie changes in that there are some edits on seasons and episodes
// that will create a change entry at the show level.
// These can be found under the season and episode keys.
// These keys will contain a series_id and episode_id.
func (tr *TVResource) GetChanges(tvID int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/changes", tvID)
	var changes Changes
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get changes")
}

// ContentRatingsOptions represents the available options for the request.
type ContentRatingsOptions languageOptions

// Rating represents rating in TMDb.
type Rating struct {
	ISO31661 string `json:"iso_3166_1"`
	Rating   string `json:"rating"`
}

// ContentRatings represents content ratings in TMDb.
type ContentRatings struct {
	ID      *int     `json:"id"`
	Ratings []Rating `json:"results"`
}

// GetContentRatings retrieves the list of content ratings (certifications) that have been added to a TV show.
func (tr *TVResource) GetContentRatings(tvID int, opt *ContentRatingsOptions) (*ContentRatings, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/content_ratings", tvID)
	var ratings ContentRatings
	resp, err := tr.client.get(path, &ratings, WithQueryParams(opt))
	return &ratings, resp, errors.Wrap(err, "failed to get content ratings")
}

// TVShowCast represents a tv show cast in TMDb.
type TVShowCast struct {
	Adult              bool    `json:"adult"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Gender             *int    `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	Order              int     `json:"order"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

// TVShowCrew represents a tv show crew in TMDb.
type TVShowCrew struct {
	Adult              bool    `json:"adult"`
	CreditID           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Gender             *int    `json:"gender"`
	ID                 int     `json:"id"`
	Job                string  `json:"job"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

// TVShowCredits represents tv show credits in TMDb.
type TVShowCredits struct {
	ID   *int         `json:"id"`
	Cast []TVShowCast `json:"cast"`
	Crew []TVShowCrew `json:"crew"`
}

// GetCredits retrieves the credits (cast and crew) that have been added to a TV show.
func (tr *TVResource) GetCredits(tvID int, opt *CreditsOptions) (*TVShowCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/credits", tvID)
	var credits TVShowCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get credits")
}

// TVShowNetwork represents tv show network in TMDb.
type TVShowNetwork struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

// TVShowEpisodeGroup represents tv show episode group in TMDb.
type TVShowEpisodeGroup struct {
	Description  string        `json:"description"`
	EpisodeCount int           `json:"episode_count"`
	GroupCount   int           `json:"group_count"`
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Network      TVShowNetwork `json:"network"`
	Type         int           `json:"type"`
}

// EpisodeGroups represents episode groups in TMDb.
type EpisodeGroups struct {
	ID     *int                 `json:"id"`
	Groups []TVShowEpisodeGroup `json:"results"`
}

// EpisodeGroupsOptions represents the available options for the request.
type EpisodeGroupsOptions languageOptions

// GetEpisodeGroups retrieves all of the episode groups that have been created for a TV show.
func (tr *TVResource) GetEpisodeGroups(tvID int, opt *EpisodeGroupsOptions) (*EpisodeGroups, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/episode_groups", tvID)
	var credits EpisodeGroups
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get episode groups")
}

// TVShowExternalIDs represents tv show external ids in TMDb.
type TVShowExternalIDs struct {
	ID          *int    `json:"id"`
	FacebookID  *string `json:"facebook_id"`
	FreebaseID  *string `json:"freebase_id"`
	FreebaseMID *string `json:"freebase_mid"`
	IMDbID      *string `json:"imdb_id"`
	InstagramID *string `json:"instagram_id"`
	TVDbID      *int    `json:"tvdb_id"`
	TVRageID    *int    `json:"tvrage_id"`
	TwitterID   *string `json:"twitter_id"`
}

// ExternalIDsOptions represents the available options for the request.
type ExternalIDsOptions languageOptions

// GetExternalIDs retrieves the external ids for a TV show.
func (tr *TVResource) GetExternalIDs(tvID int, opt *ExternalIDsOptions) (*TVShowExternalIDs, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/external_ids", tvID)
	var ids TVShowExternalIDs
	resp, err := tr.client.get(path, &ids, WithQueryParams(opt))
	return &ids, resp, errors.Wrap(err, "failed to get external ids")
}

// GetImages retrieves the images that belong to a TV show.
// Querying images with a language parameter will filter the results.
// To include a fallback language (especially useful for backdrops), use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (tr *TVResource) GetImages(tvID int, opt *ImagesOptions) (*Images, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/images", tvID)
	var images Images
	resp, err := tr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

// TVShowKeywords represents tv show keywords in TMDb.
type TVShowKeywords struct {
	ID       *int      `json:"id"`
	Keywords []Keyword `json:"results"`
}

// GetKeywords retrieves the keywords that have been added to a TV show.
func (tr *TVResource) GetKeywords(tvID int) (*TVShowKeywords, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/keywords", tvID)
	var keywords TVShowKeywords
	resp, err := tr.client.get(path, &keywords)
	return &keywords, resp, errors.Wrap(err, "failed to get keywords")
}

// RecommendationsOptions represents the available options for the request.
type RecommendationsOptions languagePageOptions

// RecommendedTVShows represents recommended tv shows in TMDb.
type RecommendedTVShows struct {
	pagination
	TVShows []TVShow `json:"results"`
}

// GetRecommendations retrieves the list of TV show recommendations for this item.
func (tr *TVResource) GetRecommendations(tvID int, opt *RecommendationsOptions) (*RecommendedTVShows, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/recommendations", tvID)
	var tvShows RecommendedTVShows
	resp, err := tr.client.get(path, &tvShows, WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to get tv shows recommendations")
}

// ReviewsOptions represents the available options for the request.
type ReviewsOptions languagePageOptions

// TVShowReviews represents tv show reviews in TMDb.
type TVShowReviews struct {
	pagination
	ID      *int     `json:"id"`
	Reviews []Review `json:"results"`
}

// GetReviews retrieves the reviews for a TV show.
func (tr *TVResource) GetReviews(tvID int, opt *ReviewsOptions) (*TVShowReviews, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/reviews", tvID)
	var reviews TVShowReviews
	resp, err := tr.client.get(path, &reviews, WithQueryParams(opt))
	return &reviews, resp, errors.Wrap(err, "failed to get reviews")
}

// TheatricalScreens represents theatrical screens in TMDb.
type TheatricalScreens struct {
	ID            int `json:"id"`
	EpisodeNumber int `json:"episode_number"`
	SeasonNumber  int `json:"season_number"`
}

// ScreenedTheatrically represents screened theatrically in TMDb.
type ScreenedTheatrically struct {
	ID      *int                `json:"id"`
	Results []TheatricalScreens `json:"results"`
}

// GetScreenedTheatrically retrieves a list of seasons or episodes that have been screened in a film festival or theatre.
func (tr *TVResource) GetScreenedTheatrically(tvID int) (*ScreenedTheatrically, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/screened_theatrically", tvID)
	var screenedTheatrically ScreenedTheatrically
	resp, err := tr.client.get(path, &screenedTheatrically)
	return &screenedTheatrically, resp, errors.Wrap(err, "failed to get screened theatrically info")
}

// SimilarTVShows represents similar tv shows in TMDb.
type SimilarTVShows paginatedTVShows

// SimilarTVShowsOptions represents the available options for the request.
type SimilarTVShowsOptions languagePageOptions

// GetSimilar retrieves a list of similar TV shows. These items are assembled by looking at keywords and genres.
func (tr *TVResource) GetSimilar(tvID int, opt *SimilarTVShowsOptions) (*SimilarTVShows, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/similar", tvID)
	var similar SimilarTVShows
	resp, err := tr.client.get(path, &similar, WithQueryParams(opt))
	return &similar, resp, errors.Wrap(err, "failed to get similar tv shows")
}

// TVData represents tv data in TMDb.
type TVData struct {
	Homepage string `json:"homepage"`
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Tagline  string `json:"tagline"`
}

// TVShowTranslation represents a tv show translation in TMDb.
type TVShowTranslation struct {
	ISO31661    string `json:"iso_3166_1"`
	ISO6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
	Data        TVData `json:"data"`
}

// TVShowTranslations represents tv show translations in TMDb.
type TVShowTranslations struct {
	ID           *int                `json:"id"`
	Translations []TVShowTranslation `json:"translations"`
}

// GetTranslations retrieves a list of the translations that exist for a TV show.
func (tr *TVResource) GetTranslations(tvID int) (*TVShowTranslations, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/translations", tvID)
	var translations TVShowTranslations
	resp, err := tr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

// GetVideos retrieves the videos that have been added to a TV show.
func (tr *TVResource) GetVideos(tvID int, opt *VideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/videos", tvID)
	var videos Videos
	resp, err := tr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get tv show videos")
}

// GetWatchProviders retrieves watch providers for a tv show.
// Powered by the partnership with JustWatch, use this method to get a list of the availabilities per country by provider.
// This is not going to return full deep links, but rather, it's just enough information to display what's available where.
// Link to the provided TMDB URL to help support TMDB and provide the actual deep links to the content.
// Please note: In order to use this data it's REQUIRED to attribute the source of the data as JustWatch.
// If any usage is found not complying with these terms the access to the API will be revoked.
func (tr *TVResource) GetWatchProviders(tvID int) (*WatchProviders, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/watch/providers", tvID)
	var providers WatchProviders
	resp, err := tr.client.get(path, &providers)
	return &providers, resp, errors.Wrap(err, "failed to get tv show watch providers")
}

// TVCreatedBy represents created by content in TMDb.
type TVCreatedBy struct {
	Type       int     `json:"type"`
	CreditID   string  `json:"credit_id"`
	Name       string  `json:"name"`
	Gender     int     `json:"gender"`
	PosterPath *string `json:"poster_path"`
}

// LatestTVShow represents the latest tv show in TMDb.
type LatestTVShow struct {
	Adult               bool                `json:"adult"`
	BackdropPath        *string             `json:"backdrop_path"`
	CreatedBy           []TVCreatedBy       `json:"created_by"`
	EpisodeRunTime      []int               `json:"episode_run_time"`
	FirstAirDate        string              `json:"first_air_date"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int                 `json:"id"`
	InProduction        bool                `json:"in_production"`
	Languages           []string            `json:"languages"`
	LastAirDate         string              `json:"last_air_date"`
	LastEpisodeToAir    *TVEpisode          `json:"last_episode_to_air"`
	Name                string              `json:"name"`
	Networks            []TVShowNetwork     `json:"networks"`
	NextEpisodeToAir    *TVEpisode          `json:"next_episode_to_air"`
	NumberOfEpisodes    int                 `json:"number_of_episodes"`
	NumberOfSeasons     int                 `json:"number_of_seasons"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalName        string              `json:"original_name"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          *string             `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	Seasons             []Season            `json:"seasons"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
}

// GetLatest retrieves the most newly created TV show. This is a live response and will continuously change.
func (tr *TVResource) GetLatest(opt *LatestOptions) (*LatestTVShow, *http.Response, error) {
	path := "/tv/latest"
	var latest LatestTVShow
	resp, err := tr.client.get(path, &latest, WithQueryParams(opt))
	return &latest, resp, errors.Wrap(err, "failed to get latest tv show")
}

// TVShowAiring represents a airing tv show in TMDb.
type TVShowAiring struct {
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// TVShowsAiring represents airing tv shows in TMDb.
type TVShowsAiring struct {
	pagination
	TVShows []TVShowAiring `json:"results"`
}

// TVShowsAiringOptions represents the available options for the request.
type TVShowsAiringOptions languagePageOptions

// GetAiringToday retrieves a list of TV shows that are airing today.
// This query is purely day based as TMDb currently doesn't support airing times.
func (tr *TVResource) GetAiringToday(opt *TVShowsAiringOptions) (*TVShowsAiring, *http.Response, error) {
	path := "/tv/airing_today"
	var tvShows TVShowsAiring
	resp, err := tr.client.get(path, &tvShows, WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to get airing today")
}

// GetTVShowsChanges retrieves a list of all of the movie ids that have been changed in the past 24 hours.
// Query it for up to 14 days worth of changed IDs at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
func (tr *TVResource) GetTVShowsChanges(opt *ChangesOptions) (*MediaChanges, *http.Response, error) {
	path := "/tv/changes"
	var changes MediaChanges
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get tv shows changes")
}

// GetOnTheAir retrieves a list of shows that are currently on the air.
// This query looks for any TV show that has an episode with an air date in the next 7 days.
func (tr *TVResource) GetOnTheAir(opt *TVShowsAiringOptions) (*TVShowsAiring, *http.Response, error) {
	path := "/tv/on_the_air"
	var tvShows TVShowsAiring
	resp, err := tr.client.get(path, &tvShows, WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to get on the air")
}

// PopularTVShow represents a popular tv show in TMDb.
type PopularTVShow struct {
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// PopularTVShows represents popular tv shows in TMDb.
type PopularTVShows struct {
	pagination
	TVShows []PopularTVShow `json:"results"`
}

// PopularTVShowsOptions represents the available options for the request.
type PopularTVShowsOptions languagePageOptions

// GetPopular retrieves a list of the current popular TV shows on TMDB. This list updates daily.
func (tr *TVResource) GetPopular(opt *PopularTVShowsOptions) (*PopularTVShows, *http.Response, error) {
	path := "/tv/popular"
	var popular PopularTVShows
	resp, err := tr.client.get(path, &popular, WithQueryParams(opt))
	return &popular, resp, errors.Wrap(err, "failed to get popular tv shows")
}

// TopRatedTVShow represents a top rated tv show in TMDb.
type TopRatedTVShow struct {
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// TopRatedTVShows represents top rated tv shows in TMDb.
type TopRatedTVShows struct {
	pagination
	TVShows []TopRatedTVShow `json:"results"`
}

// TopRatedTVShowOptions represents the available options for the request.
type TopRatedTVShowOptions languagePageOptions

// GetTopRated retrieves a list of the top rated TV shows on TMDB.
func (tr *TVResource) GetTopRated(opt *TopRatedTVShowOptions) (*TopRatedTVShows, *http.Response, error) {
	path := "/tv/popular"
	var topRated TopRatedTVShows
	resp, err := tr.client.get(path, &topRated, WithQueryParams(opt))
	return &topRated, resp, errors.Wrap(err, "failed to get top rated tv shows")
}

// GroupEpisode represents a group episode in TMDb.
type GroupEpisode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Order          int     `json:"order"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

// Group represents a group in TMDb.
type Group struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Order    int            `json:"order"`
	Episodes []GroupEpisode `json:"episodes"`
	Locked   bool           `json:"locked"`
}

// EpisodeGroup represents an episode group in TMDb.
type EpisodeGroup struct {
	Description  string        `json:"description"`
	EpisodeCount int           `json:"episode_count"`
	GroupCount   int           `json:"group_count"`
	Groups       []Group       `json:"groups"`
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Network      TVShowNetwork `json:"network"`
	Type         int           `json:"type"`
}

// EpisodeGroupOptions represents the available options for the request.
type EpisodeGroupOptions languageOptions

// GetEpisodeGroup retrieves the details of a TV episode group.
// Groups support 7 different types which are enumerated as the following:
// 1. Original air date
// 2. Absolute
// 3. DVD
// 4. Digital
// 5. Story arc
// 6. Production
// 7. TV
func (tr *TVResource) GetEpisodeGroup(groupID string, opt *EpisodeGroupOptions) (*EpisodeGroup, *http.Response, error) {
	path := fmt.Sprintf("/tv/episode_group/%s", groupID)
	var groups EpisodeGroup
	resp, err := tr.client.get(path, &groups, WithQueryParams(opt))
	return &groups, resp, errors.Wrap(err, "failed to get episode groups")
}

// Rate rates a TV show.
// A valid session or guest session ID is required.
func (tr *TVResource) Rate(tvID int, rating float64, sessionID Auth) (*RateResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/rating", tvID)
	var response RateResponse
	resp, err := tr.client.post(path, &response, WithBody(map[string]float64{"value": rating}), WithQueryParams(sessionID))
	return &response, resp, errors.Wrap(err, "failed to rate tv show")
}

// DeleteRating removes a rating for a TV show.
// A valid session or guest session ID is required.
func (tr *TVResource) DeleteRating(movieID int, sessionID Auth) (*DeleteRatingResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/rating", movieID)
	var response DeleteRatingResponse
	resp, err := tr.client.delete(path, &response, WithQueryParams(sessionID))
	return &response, resp, errors.Wrap(err, "failed to delete tv show rating")
}
