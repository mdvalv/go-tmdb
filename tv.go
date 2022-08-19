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

type tvShowAppendToResponse struct {
	AggregateCredits     *aggregateCredits        `json:"aggregate_credits"`
	AlternativeTitles    *alternativeTVShowTitles `json:"alternative_titles"`
	Changes              *Changes                 `json:"changes"`
	ContentRatings       *contentRatings          `json:"content_ratings"`
	Credits              *tvShowCredits           `json:"credits"`
	EpisodeGroups        *episodeGroups           `json:"episode_groups"`
	ExternalIds          *tvShowExternalIds       `json:"external_ids"`
	Images               *images                  `json:"images"`
	Keywords             *tvShowKeywords          `json:"keywords"`
	Recommendations      *RecommendedTVShows      `json:"recommendations"`
	Reviews              *tvShowReviews           `json:"reviews"`
	ScreenedTheatrically *screenedTheatrically    `json:"screened_theatrically"`
	Similar              *SimilarTVShows          `json:"similar"`
	Translations         *tvShowTranslations      `json:"translations"`
	Videos               *videos                  `json:"videos"`
}

type TVShowDetails struct {
	tvShowAppendToResponse
	tvShowInfo
}

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

// Get the TV season details by id.
func (tr *TVResource) GetTVShow(tvId int, opt *TVShowDetailsOptions) (*TVShowDetails, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d", tvId)
	var tvShow TVShowDetails
	resp, err := tr.client.get(path, &tvShow, WithQueryParams(opt))
	return &tvShow, resp, errors.Wrap(err, "failed to get tv show")
}

type Season struct {
	season
	MediaType string `json:"media_type"`
}

type season struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
}

type TVShow struct {
	tv
	MediaType string `json:"media_type"`
}

type Episode struct {
	episode
	MediaType string `json:"media_type"`
}

type tv struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
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

type episode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowId         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

type paginatedTVShows struct {
	pagination
	TVShows []tv `json:"results"`
}

// Grab the following account states for a session:
// - Movie rating
// - If it belongs to the watchlist
// - If it belongs to the favorite list
func (tr *TVResource) GetAccountStates(tvId int, sessionId string) (*AccountStates, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/account_states", tvId)
	var states AccountStates
	resp, err := tr.client.get(path, &states, WithQueryParam("session_id", sessionId))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

type AggregateCreditsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

type Role struct {
	CreditId     string `json:"credit_id"`
	Character    string `json:"character"`
	EpisodeCount int    `json:"episode_count"`
}

type Job struct {
	CreditId     string `json:"credit_id"`
	Job          string `json:"job"`
	EpisodeCount int    `json:"episode_count"`
}

type AggregateCreditsCast struct {
	Adult              bool    `json:"adult"`
	Gender             *int    `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	Order              int     `json:"order"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	Roles              []Role  `json:"roles"`
	TotalEpisodeCount  int     `json:"total_episode_count"`
}

type AggregateCreditsCrew struct {
	Adult              bool    `json:"adult"`
	Department         string  `json:"department"`
	Gender             *int    `json:"gender"`
	Id                 int     `json:"id"`
	Jobs               []Job   `json:"jobs"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	TotalEpisodeCount  int     `json:"total_episode_count"`
}

type aggregateCredits struct {
	Cast []AggregateCreditsCast `json:"cast"`
	Crew []AggregateCreditsCrew `json:"crew"`
}

type AggregateCredits struct {
	Id int `json:"id"`
	aggregateCredits
}

// Get the aggregate credits (cast and crew) that have been added to a TV show.
// This call differs from the main `credits` call in that it does not return the newest season but rather,
// is a view of all the entire cast & crew for all episodes belonging to a TV show.
func (tr *TVResource) GetAggregateCredits(tvId int, opt *AggregateCreditsOptions) (*AggregateCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/aggregate_credits", tvId)
	var credits AggregateCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get aggregate credits")
}

type TVShowAlternativeTitlesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

type alternativeTVShowTitles struct {
	Titles []Title `json:"results"`
}

type AlternativeTVShowTitles struct {
	alternativeTVShowTitles
	Id int `json:"id"`
}

// Get all of the alternative titles for a tv show.
func (tr *TVResource) GetAlternativeTitles(tvId int, opt *TVShowAlternativeTitlesOptions) (*AlternativeTVShowTitles, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/alternative_titles", tvId)
	var titles AlternativeTVShowTitles
	resp, err := tr.client.get(path, &titles, WithQueryParams(opt))
	return &titles, resp, errors.Wrap(err, "failed to get alternative titles")
}

// Get the changes for a TV show. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the start_date and end_date query parameters.
// TV show changes are different than movie changes in that there are some edits on seasons and episodes
// that will create a change entry at the show level.
// These can be found under the season and episode keys.
// These keys will contain a series_id and episode_id.
func (tr *TVResource) GetChanges(tvId int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/changes", tvId)
	var changes Changes
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get changes")
}

type ContentRatingsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

type Rating struct {
	ISO31661 string `json:"iso_3166_1"`
	Rating   string `json:"rating"`
}

type contentRatings struct {
	Ratings []Rating `json:"results"`
}

type ContentRatings struct {
	contentRatings
	Id int `json:"id"`
}

// Get the list of content ratings (certifications) that have been added to a TV show.
func (tr *TVResource) GetContentRatings(tvId int, opt *ContentRatingsOptions) (*ContentRatings, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/content_ratings", tvId)
	var ratings ContentRatings
	resp, err := tr.client.get(path, &ratings, WithQueryParams(opt))
	return &ratings, resp, errors.Wrap(err, "failed to get content ratings")
}

type TVShowCast struct {
	Adult              bool    `json:"adult"`
	Character          string  `json:"character"`
	CreditId           string  `json:"credit_id"`
	Gender             *int    `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	Order              int     `json:"order"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

type TVShowCrew struct {
	Adult              bool    `json:"adult"`
	CreditId           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Gender             *int    `json:"gender"`
	Id                 int     `json:"id"`
	Job                string  `json:"job"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

type tvShowCredits struct {
	Cast []TVShowCast `json:"cast"`
	Crew []TVShowCrew `json:"crew"`
}

type TVShowCredits struct {
	Id int `json:"id"`
	tvShowCredits
}

// Get the credits (cast and crew) that have been added to a TV show.
func (tr *TVResource) GetCredits(tvId int, opt *CreditsOptions) (*TVShowCredits, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/credits", tvId)
	var credits TVShowCredits
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get credits")
}

type TVShowNetwork struct {
	Id            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type TVShowEpisodeGroup struct {
	Description  string        `json:"description"`
	EpisodeCount int           `json:"episode_count"`
	GroupCount   int           `json:"group_count"`
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Network      TVShowNetwork `json:"network"`
	Type         int           `json:"type"`
}

type episodeGroups struct {
	Groups []TVShowEpisodeGroup `json:"results"`
}

type EpisodeGroups struct {
	episodeGroups
	Id int `json:"id"`
}

type EpisodeGroupsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// Get all of the episode groups that have been created for a TV show.
func (tr *TVResource) GetEpisodeGroups(tvId int, opt *EpisodeGroupsOptions) (*EpisodeGroups, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/episode_groups", tvId)
	var credits EpisodeGroups
	resp, err := tr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get episode groups")
}

type TVShowExternalIds struct {
	Id int `json:"id"`
	tvShowExternalIds
}

type tvShowExternalIds struct {
	FacebookId  *string `json:"facebook_id"`
	FreebaseId  *string `json:"freebase_id"`
	FreebaseMId *string `json:"freebase_mid"`
	IMDbId      *string `json:"imdb_id"`
	InstagramId *string `json:"instagram_id"`
	TVDbId      *int    `json:"tvdb_id"`
	TVRageId    *int    `json:"tvrage_id"`
	TwitterId   *string `json:"twitter_id"`
}

type ExternalIdsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// Get the external ids for a TV show.
func (tr *TVResource) GetExternalIds(tvId int, opt *ExternalIdsOptions) (*TVShowExternalIds, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/external_ids", tvId)
	var ids TVShowExternalIds
	resp, err := tr.client.get(path, &ids, WithQueryParams(opt))
	return &ids, resp, errors.Wrap(err, "failed to get external ids")
}

// Get the images that belong to a TV show.
// Querying images with a language parameter will filter the results.
// To include a fallback language (especially useful for backdrops), use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (tr *TVResource) GetImages(tvId int, opt *ImagesOptions) (*Images, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/images", tvId)
	var images Images
	resp, err := tr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

type tvShowKeywords struct {
	Keywords []Keyword `json:"results"`
}

type TVShowKeywords struct {
	Id int `json:"id"`
	tvShowKeywords
}

// Get the keywords that have been added to a TV show.
func (tr *TVResource) GetKeywords(tvId int) (*TVShowKeywords, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/keywords", tvId)
	var keywords TVShowKeywords
	resp, err := tr.client.get(path, &keywords)
	return &keywords, resp, errors.Wrap(err, "failed to get keywords")
}

type RecommendationsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

type RecommendedTVShows struct {
	pagination
	TVShows []TVShow `json:"results"`
}

// Get the list of TV show recommendations for this item.
func (tr *TVResource) GetRecommendations(tvId int, opt *RecommendationsOptions) (*RecommendedTVShows, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/recommendations", tvId)
	var tvShows RecommendedTVShows
	resp, err := tr.client.get(path, &tvShows, WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to get tv shows recommendations")
}

type ReviewsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

type tvShowReviews struct {
	pagination
	Reviews []review `json:"results"`
}

type TVShowReviews struct {
	Id int `json:"id"`
	tvShowReviews
}

// Get the reviews for a TV show.
func (tr *TVResource) GetReviews(tvId int, opt *ReviewsOptions) (*TVShowReviews, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/reviews", tvId)
	var reviews TVShowReviews
	resp, err := tr.client.get(path, &reviews, WithQueryParams(opt))
	return &reviews, resp, errors.Wrap(err, "failed to get reviews")
}

type theatricalScreens struct {
	Id            int `json:"id"`
	EpisodeNumber int `json:"episode_number"`
	SeasonNumber  int `json:"season_number"`
}

type screenedTheatrically struct {
	Results []theatricalScreens `json:"results"`
}

type ScreenedTheatrically struct {
	screenedTheatrically
	Id int `json:"id"`
}

// Get a list of seasons or episodes that have been screened in a film festival or theatre.
func (tr *TVResource) GetScreenedTheatrically(tvId int) (*ScreenedTheatrically, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/screened_theatrically", tvId)
	var screenedTheatrically ScreenedTheatrically
	resp, err := tr.client.get(path, &screenedTheatrically)
	return &screenedTheatrically, resp, errors.Wrap(err, "failed to get screened theatrically info")
}

type SimilarTVShows paginatedTVShows

type SimilarTVShowsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Get a list of similar TV shows. These items are assembled by looking at keywords and genres.
func (tr *TVResource) GetSimilar(tvId int, opt *SimilarTVShowsOptions) (*SimilarTVShows, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/similar", tvId)
	var similar SimilarTVShows
	resp, err := tr.client.get(path, &similar, WithQueryParams(opt))
	return &similar, resp, errors.Wrap(err, "failed to get similar tv shows")
}

type TVData struct {
	Homepage string `json:"homepage"`
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Tagline  string `json:"tagline"`
}

type TVShowTranslation struct {
	ISO31661    string `json:"iso_3166_1"`
	ISO6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
	Data        TVData `json:"data"`
}

type tvShowTranslations struct {
	Translations []TVShowTranslation `json:"translations"`
}

type TVShowTranslations struct {
	Id int `json:"id"`
	tvShowTranslations
}

// Get a list of the translations that exist for a TV show.
func (tr *TVResource) GetTranslations(tvId int) (*TVShowTranslations, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/translations", tvId)
	var translations TVShowTranslations
	resp, err := tr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

// Get the videos that have been added to a TV show.
func (tr *TVResource) GetVideos(tvId int, opt *VideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/videos", tvId)
	var videos Videos
	resp, err := tr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get tv show videos")
}

// Powered by the partnership with JustWatch, use this method to get a list of the availabilities per country by provider.
// This is not going to return full deep links, but rather, it's just enough information to display what's available where.
// Link to the provided TMDB URL to help support TMDB and provide the actual deep links to the content.
// Please note: In order to use this data you MUST attribute the source of the data as JustWatch.
// If any usage is found not complying with these terms the access to the API will be revoked.
func (tr *TVResource) GetWatchProviders(tvId int) (*WatchProviders, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/watch/providers", tvId)
	var providers WatchProviders
	resp, err := tr.client.get(path, &providers)
	return &providers, resp, errors.Wrap(err, "failed to get tv show watch providers")
}

type TVCreatedBy struct {
	Type       int     `json:"type"`
	CreditId   string  `json:"credit_id"`
	Name       string  `json:"name"`
	Gender     int     `json:"gender"`
	PosterPath *string `json:"poster_path"`
}

type LatestTVShow tvShowInfo

type tvShowInfo struct {
	Adult               bool                `json:"adult"`
	BackdropPath        *string             `json:"backdrop_path"`
	CreatedBy           []TVCreatedBy       `json:"created_by"`
	EpisodeRunTime      []int               `json:"episode_run_time"`
	FirstAirDate        string              `json:"first_air_date"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	Id                  int                 `json:"id"`
	InProduction        bool                `json:"in_production"`
	Languages           []string            `json:"languages"`
	LastAirDate         string              `json:"last_air_date"`
	LastEpisodeToAir    *episode            `json:"last_episode_to_air"`
	Name                string              `json:"name"`
	Networks            []TVShowNetwork     `json:"networks"`
	NextEpisodeToAir    *episode            `json:"next_episode_to_air"`
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
	Seasons             []season            `json:"seasons"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
}

// Get the most newly created TV show. This is a live response and will continuously change.
func (tr *TVResource) GetLatest(opt *LatestOptions) (*LatestTVShow, *http.Response, error) {
	path := "/tv/latest"
	var latest LatestTVShow
	resp, err := tr.client.get(path, &latest, WithQueryParams(opt))
	return &latest, resp, errors.Wrap(err, "failed to get latest tv show")
}

type TVShowAiring struct {
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
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

type TVShowsAiring struct {
	pagination
	TVShows []TVShowAiring `json:"results"`
}

type TVShowsAiringOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Get a list of TV shows that are airing today.
// This query is purely day based as TMDb currently doesn't support airing times.
func (tr *TVResource) GetAiringToday(opt *TVShowsAiringOptions) (*TVShowsAiring, *http.Response, error) {
	path := "/tv/airing_today"
	var tvShows TVShowsAiring
	resp, err := tr.client.get(path, &tvShows, WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to get airing today")
}

// Get a list of all of the movie ids that have been changed in the past 24 hours.
// Query it for up to 14 days worth of changed IDs at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
func (tr *TVResource) GetTVShowsChanges(opt *ChangesOptions) (*MediaChanges, *http.Response, error) {
	path := "/tv/changes"
	var changes MediaChanges
	resp, err := tr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get tv shows changes")
}

// Get a list of shows that are currently on the air.
// This query looks for any TV show that has an episode with an air date in the next 7 days.
func (tr *TVResource) GetOnTheAir(opt *TVShowsAiringOptions) (*TVShowsAiring, *http.Response, error) {
	path := "/tv/on_the_air"
	var tvShows TVShowsAiring
	resp, err := tr.client.get(path, &tvShows, WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to get on the air")
}

type PopularTVShow struct {
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
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

type PopularTVShows struct {
	pagination
	TVShows []PopularTVShow `json:"results"`
}

type PopularTVShowsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Get a list of the current popular TV shows on TMDB. This list updates daily.
func (tr *TVResource) GetPopular(opt *PopularTVShowsOptions) (*PopularTVShows, *http.Response, error) {
	path := "/tv/popular"
	var popular PopularTVShows
	resp, err := tr.client.get(path, &popular, WithQueryParams(opt))
	return &popular, resp, errors.Wrap(err, "failed to get popular tv shows")
}

type TopRatedTVShow struct {
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
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

type TopRatedTVShows struct {
	pagination
	TVShows []TopRatedTVShow `json:"results"`
}

type TopRatedTVShowOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Get a list of the top rated TV shows on TMDB.
func (tr *TVResource) GetTopRated(opt *TopRatedTVShowOptions) (*TopRatedTVShows, *http.Response, error) {
	path := "/tv/popular"
	var topRated TopRatedTVShows
	resp, err := tr.client.get(path, &topRated, WithQueryParams(opt))
	return &topRated, resp, errors.Wrap(err, "failed to get top rated tv shows")
}

type episodeGroup struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Order          int     `json:"order"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowId         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

type Group struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Order    int            `json:"order"`
	Episodes []episodeGroup `json:"episodes"`
	Locked   bool           `json:"locked"`
}

type EpisodeGroup struct {
	Description  string        `json:"description"`
	EpisodeCount int           `json:"episode_count"`
	GroupCount   int           `json:"group_count"`
	Groups       []Group       `json:"groups"`
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Network      TVShowNetwork `json:"network"`
	Type         int           `json:"type"`
}

type EpisodeGroupOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// Get the details of a TV episode group. Groups support 7 different types which are enumerated as the following:
// 1. Original air date
// 2. Absolute
// 3. DVD
// 4. Digital
// 5. Story arc
// 6. Production
// 7. TV
func (tr *TVResource) GetEpisodeGroup(groupId string, opt *EpisodeGroupOptions) (*EpisodeGroup, *http.Response, error) {
	path := fmt.Sprintf("/tv/episode_group/%s", groupId)
	var groups EpisodeGroup
	resp, err := tr.client.get(path, &groups, WithQueryParams(opt))
	return &groups, resp, errors.Wrap(err, "failed to get episode groups")
}

// Rate a TV show.
// A valid session or guest session ID is required.
func (tr *TVResource) Rate(tvId int, rating float64, sessionId Auth) (*RateResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/rating", tvId)
	var response RateResponse
	resp, err := tr.client.post(path, &response, WithBody(map[string]float64{"value": rating}), WithQueryParams(sessionId))
	return &response, resp, errors.Wrap(err, "failed to rate tv show")
}

// Remove a rating for a TV show.
// A valid session or guest session ID is required.
func (tr *TVResource) DeleteRating(movieId int, sessionId Auth) (*DeleteRatingResponse, *http.Response, error) {
	path := fmt.Sprintf("/tv/%d/rating", movieId)
	var response DeleteRatingResponse
	resp, err := tr.client.delete(path, &response, WithQueryParams(sessionId))
	return &response, resp, errors.Wrap(err, "failed to delete tv show rating")
}
