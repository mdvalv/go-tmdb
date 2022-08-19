package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// MoviesResource handles movie-related requests of TMDb API.
type MoviesResource struct {
	client *Client
}

type Movie struct {
	movie
	MediaType string `json:"media_type"`
}

type movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
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

type paginatedMovies struct {
	pagination
	Movies []movie `json:"results"`
}

type movieAppendToResponse struct {
	AlternativeTitles *alternativeMovieTitles `json:"alternative_titles"`
	Changes           *Changes                `json:"changes"`
	Credits           *movieCredits           `json:"credits"`
	ExternalIds       *movieExternalIds       `json:"external_ids"`
	Images            *images                 `json:"images"`
	Keywords          *movieKeywords          `json:"keywords"`
	Lists             *MovieLists             `json:"lists"`
	Recommendations   *RecommendedMovies      `json:"recommendations"`
	ReleaseDates      *movieReleaseDates      `json:"release_dates"`
	Reviews           *movieReviews           `json:"reviews"`
	Similar           *SimilarMovies          `json:"similar"`
	Translations      *movieTranslations      `json:"translations"`
	Videos            *videos                 `json:"videos"`
}

type MovieDetails struct {
	movieInfo
	movieAppendToResponse
}

type MovieDetailsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	AppendToResponse string `url:"append_to_response,omitempty" json:"append_to_response,omitempty"`
}

// Get the primary information about a movie.
func (mr *MoviesResource) GetMovie(movieId int, opt *MovieDetailsOptions) (*MovieDetails, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d", movieId)
	var movie MovieDetails
	resp, err := mr.client.get(path, &movie, WithQueryParams(opt))
	return &movie, resp, errors.Wrap(err, "failed to get movie")
}

// Get a list of all of the movie ids that have been changed in the past 24 hours.
// You can query it for up to 14 days worth of changed IDs at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
func (mr *MoviesResource) GetMoviesChanges(opt *ChangesOptions) (*MediaChanges, *http.Response, error) {
	path := "/movie/changes"
	var changes MediaChanges
	resp, err := mr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get movies changes")
}

type ProductionCompany struct {
	Name          string  `json:"name"`
	Id            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	OriginCountry string  `json:"origin_country"`
}

type ProductionCountry struct {
	Name     string `json:"name"`
	ISO31661 string `json:"iso_3166_1"`
}

type SpokenLanguage struct {
	ISO6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
}

type BelongsToCollection struct {
	BackdropPath *string `json:"backdrop_path"`
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	PosterPath   *string `json:"poster_path"`
}

type LatestMovie movieInfo

type movieInfo struct {
	movie
	BelongsToCollection *BelongsToCollection `json:"belongs_to_collection"`
	Budget              int                  `json:"budget"`
	Genres              []Genre              `json:"genres"`
	Homepage            string               `json:"homepage"`
	IMDbId              string               `json:"imdb_id"`
	ProductionCompanies []ProductionCompany  `json:"production_companies"`
	ProductionCountries []ProductionCountry  `json:"production_countries"`
	Revenue             int                  `json:"revenue"`
	Runtime             int                  `json:"runtime"`
	SpokenLanguages     []SpokenLanguage     `json:"spoken_languages"`
	Status              string               `json:"status"`
	Tagline             string               `json:"tagline"`
}

type LatestMovieOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// Get the most newly created movie. This is a live response and will continuously change.
func (mr *MoviesResource) GetLatest(opt *LatestMovieOptions) (*LatestMovie, *http.Response, error) {
	path := "/movie/latest"
	var latest LatestMovie
	resp, err := mr.client.get(path, &latest, WithQueryParams(opt))
	return &latest, resp, errors.Wrap(err, "failed to get latest movie")
}

type NowPlayingMoviesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`
}

type DateRange struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}

type NowPlayingMovies struct {
	paginatedMovies
	Dates DateRange `json:"dates"`
}

// Get a list of movies in theatres.
// This is a release type query that looks for all movies that have a release type of
// 2 or 3 within the specified date range.
// Optionally specify a region parameter which will narrow the search to only look for
// theatrical release dates within the specified country.
func (mr *MoviesResource) GetNowPlaying(opt *NowPlayingMoviesOptions) (*NowPlayingMovies, *http.Response, error) {
	path := "/movie/now_playing"
	var movies NowPlayingMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get movies playing now")
}

type PopularMoviesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`
}

type PopularMovies paginatedMovies

// Get a list of the current popular movies on TMDB. This list updates daily.
func (mr *MoviesResource) GetPopular(opt *PopularMoviesOptions) (*PopularMovies, *http.Response, error) {
	path := "/movie/popular"
	var movies PopularMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get popular movies")
}

type TopRatedMoviesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`
}

type TopRatedMovies paginatedMovies

// Get the top rated movies on TMDB.
func (mr *MoviesResource) GetTopRated(opt *TopRatedMoviesOptions) (*TopRatedMovies, *http.Response, error) {
	path := "/movie/top_rated"
	var movies TopRatedMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get top rated movies")
}

type UpcomingMoviesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`
}

type UpcomingMovies struct {
	paginatedMovies
	Dates DateRange `json:"dates"`
}

// Get a list of upcoming movies in theatres.
// This is a release type query that looks for all movies that have a release type of 2 or 3 within the specified date range.
// You can optionally specify a region parameter which will narrow the search to only look for theatrical release dates
// within the specified country.
func (mr *MoviesResource) GetUpcoming(opt *UpcomingMoviesOptions) (*UpcomingMovies, *http.Response, error) {
	path := "/movie/upcoming"
	var movies UpcomingMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get upcoming movies")
}

type AccountStates struct {
	Id        int         `json:"id"`
	Favorite  bool        `json:"favorite"`
	Rated     interface{} `json:"rated"`
	Watchlist bool        `json:"watchlist"`
}

// Grab the following account states for a session:
// - Movie rating
// - If it belongs to the watchlist
// - If it belongs to the favorite list
func (mr *MoviesResource) GetAccountStates(movieId int, sessionId string) (*AccountStates, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/account_states", movieId)
	var states AccountStates
	resp, err := mr.client.get(path, &states, WithQueryParam("session_id", sessionId))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

type Auth struct {
	SessionId      string `url:"session_id,omitempty" json:"session_id,omitempty"`
	GuestSessionId string `url:"guest_session_id,omitempty" json:"guest_session_id,omitempty"`
}

type RateMovieResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// Rate a movie.
// A valid session or guest session ID is required.
func (mr *MoviesResource) Rate(movieId int, rating float64, sessionId Auth) (*RateResponse, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/rating", movieId)
	var response RateResponse
	resp, err := mr.client.post(path, &response, WithBody(map[string]float64{"value": rating}), WithQueryParams(sessionId))
	return &response, resp, errors.Wrap(err, "failed to rate movie")
}

type DeleteRatingResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// Remove rating for a movie.
// A valid session or guest session ID is required.
func (mr *MoviesResource) DeleteRating(movieId int, sessionId Auth) (*DeleteRatingResponse, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/rating", movieId)
	var response DeleteRatingResponse
	resp, err := mr.client.delete(path, &response, WithQueryParams(sessionId))
	return &response, resp, errors.Wrap(err, "failed to delete movie rating")
}

type Title struct {
	ISO31661 string `json:"iso_3166_1"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

type alternativeMovieTitles struct {
	Titles []Title `json:"titles"`
}

type AlternativeMovieTitles struct {
	alternativeMovieTitles
	Id int `json:"id"`
}

type MovieAlternativeTitlesOptions struct {
	Country string `url:"country,omitempty" json:"country,omitempty"`
}

// Get all of the alternative titles for a movie.
func (mr *MoviesResource) GetAlternativeTitles(movieId int, opt *MovieAlternativeTitlesOptions) (*AlternativeMovieTitles, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/alternative_titles", movieId)
	var titles AlternativeMovieTitles
	resp, err := mr.client.get(path, &titles, WithQueryParams(opt))
	return &titles, resp, errors.Wrap(err, "failed to get alternative titles")
}

// Get the changes for a movie. By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using the `start_date` and `end_date` query parameters.
func (mr *MoviesResource) GetChanges(movieId int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/changes", movieId)
	var changes Changes
	resp, err := mr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get movie changes")
}

type movieCast struct {
	Adult              bool    `json:"adult"`
	CastId             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditId           string  `json:"credit_id"`
	Gender             int     `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	Order              int     `json:"order"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

type movieCrew struct {
	Adult              bool    `json:"adult"`
	CreditId           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Gender             int     `json:"gender"`
	Id                 int     `json:"id"`
	Job                string  `json:"job"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

type movieCredits struct {
	Cast []movieCast `json:"cast"`
	Crew []movieCrew `json:"crew"`
}

type MovieCredits struct {
	Id int `json:"id"`
	movieCredits
}

// Get the cast and crew for a movie.
func (mr *MoviesResource) GetCredits(movieId int, opt *CreditsOptions) (*MovieCredits, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/credits", movieId)
	var credits MovieCredits
	resp, err := mr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get movie credits")
}

type movieExternalIds struct {
	IMDbId      *string `json:"imdb_id"`
	FacebookId  *string `json:"facebook_id"`
	InstagramId *string `json:"instagram_id"`
	TwitterId   *string `json:"twitter_id"`
}

type MovieExternalIds struct {
	Id int `json:"id"`
	movieExternalIds
}

// Get the external ids for a movie.
func (mr *MoviesResource) GetExternalIds(movieId int) (*MovieExternalIds, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/external_ids", movieId)
	var ids MovieExternalIds
	resp, err := mr.client.get(path, &ids)
	return &ids, resp, errors.Wrap(err, "failed to get movie external ids")
}

type Logo image

type images struct {
	Backdrops []Backdrop `json:"backdrops"`
	Posters   []Poster   `json:"posters"`
	Logos     []Logo     `json:"logos"`
}

type Images struct {
	Id int `json:"id"`
	images
}

type ImagesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	IncludeImageLanguage string `url:"include_image_language,omitempty" json:"include_image_language,omitempty"`
}

// Get the images that belong to a movie.
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (mr *MoviesResource) GetImages(movieId int, opt *ImagesOptions) (*Images, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/images", movieId)
	var images Images
	resp, err := mr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get movie images")
}

type movieKeywords struct {
	Keywords []Keyword `json:"keywords"`
}

type MovieKeywords struct {
	Id int `json:"id"`
	movieKeywords
}

// Get the keywords that have been added to a movie.
func (mr *MoviesResource) GetKeywords(movieId int) (*MovieKeywords, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/keywords", movieId)
	var keywords MovieKeywords
	resp, err := mr.client.get(path, &keywords)
	return &keywords, resp, errors.Wrap(err, "failed to get movie keywords")
}

type MovieLists struct {
	pagination
	Lists []list `json:"results"`
}

type MoviesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

func (mr *MoviesResource) GetLists(movieId int, opt *MoviesOptions) (*MovieLists, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/lists", movieId)
	var lists MovieLists
	resp, err := mr.client.get(path, &lists, WithQueryParams(opt))
	return &lists, resp, errors.Wrap(err, "failed to get movie lists")
}

type RecommendedMovies struct {
	pagination
	Movies []Movie `json:"results"`
}

// Get a list of recommended movies for a movie.
func (mr *MoviesResource) GetRecommendations(movieId int, opt *MoviesOptions) (*RecommendedMovies, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/recommendations", movieId)
	var movies RecommendedMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get movie recommendations")
}

type MovieReleaseDate struct {
	Certification string  `json:"certification"`
	ISO6391       *string `json:"iso_639_1"`
	ReleaseDate   string  `json:"release_date"`
	Type          int     `json:"type"`
	Note          string  `json:"note"`
}

type MovieRelease struct {
	Iso31661     string             `json:"iso_3166_1"`
	ReleaseDates []MovieReleaseDate `json:"release_dates"`
}

type movieReleaseDates struct {
	Releases []MovieRelease `json:"results"`
}

type MovieReleaseDates struct {
	Id int `json:"id"`
	movieReleaseDates
}

// Get the release date along with the certification for a movie.
// Release dates support different types:
// 1. Premiere
// 2. Theatrical (limited)
// 3. Theatrical
// 4. Digital
// 5. Physical
// 6. TV
func (mr *MoviesResource) GetReleaseDates(movieId int) (*MovieReleaseDates, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/release_dates", movieId)
	var dates MovieReleaseDates
	resp, err := mr.client.get(path, &dates)
	return &dates, resp, errors.Wrap(err, "failed to get movie release dates")
}

type movieReviews struct {
	pagination
	Reviews []review `json:"results"`
}

type MovieReviews struct {
	Id int `json:"id"`
	movieReviews
}

// Get the user reviews for a movie.
func (mr *MoviesResource) GetReviews(movieId int, opt *MoviesOptions) (*MovieReviews, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/reviews", movieId)
	var reviews MovieReviews
	resp, err := mr.client.get(path, &reviews, WithQueryParams(opt))
	return &reviews, resp, errors.Wrap(err, "failed to get movie reviews")
}

type SimilarMovies paginatedMovies

func (mr *MoviesResource) GetSimilar(movieId int, opt *MoviesOptions) (*SimilarMovies, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/similar", movieId)
	var movies SimilarMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get similar movies")
}

type movieTranslations struct {
	Translations []MovieTranslation `json:"translations"`
}

type MovieTranslations struct {
	Id int `json:"id"`
	movieTranslations
}

type MovieData struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
	Tagline  string `json:"tagline"`
	Runtime  int    `json:"runtime"`
}

type MovieTranslation struct {
	ISO31661    string    `json:"iso_3166_1"`
	ISO6391     string    `json:"iso_639_1"`
	Name        string    `json:"name"`
	EnglishName string    `json:"english_name"`
	Data        MovieData `json:"data"`
}

// Get a list of translations that have been created for a movie.
func (mr *MoviesResource) GetTranslations(movieId int) (*MovieTranslations, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/translations", movieId)
	var translations MovieTranslations
	resp, err := mr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get movie translations")
}

type MovieVideosOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	IncludeVideoLanguage string `url:"include_video_language,omitempty" json:"include_video_language,omitempty"`
}

type Video struct {
	Id          string `json:"id"`
	ISO31661    string `json:"iso_3166_1"`
	ISO6391     string `json:"iso_639_1"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Official    bool   `json:"official"`
	PublishedAt string `json:"published_at"`
	Site        string `json:"site"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
}

type videos struct {
	Videos []Video `json:"results"`
}

type Videos struct {
	Id int `json:"id"`
	videos
}

// Get the videos that have been added to a movie.
func (mr *MoviesResource) GetVideos(movieId int, opt *MovieVideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/videos", movieId)
	var videos Videos
	resp, err := mr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get movie videos")
}

type MovieProviders map[string]interface{}

type MovieWatchProviders struct {
	Id        int            `json:"id"`
	Providers MovieProviders `json:"results"`
}

// Powered by the partnership with JustWatch, use this method to get a list of the availabilities per country by provider.
// This is not going to return full deep links, but rather, it's just enough information to display what's available where.
// Link to the provided TMDB URL to help support TMDB and provide the actual deep links to the content.
// Please note: In order to use this data you MUST attribute the source of the data as JustWatch.
// If any usage is found not complying with these terms the access to the API will be revoked.
func (mr *MoviesResource) GetWatchProviders(movieId int) (*MovieWatchProviders, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/watch/providers", movieId)
	var providers MovieWatchProviders
	resp, err := mr.client.get(path, &providers)
	return &providers, resp, errors.Wrap(err, "failed to get movie watch providers")
}
