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

// Movie represents a movie in TMDb.
type Movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
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

// MovieResult represents a movie in TMDb.
type MovieResult struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
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

// paginatedMovies represents movies in TMDb.
type paginatedMovies struct {
	pagination
	Movies []MovieResult `json:"results"`
}

// MovieDetails represents movie details in TMDb.
type MovieDetails struct {
	Adult               bool                 `json:"adult"`
	BackdropPath        *string              `json:"backdrop_path"`
	BelongsToCollection *BelongsToCollection `json:"belongs_to_collection"`
	Budget              int                  `json:"budget"`
	GenreIDs            []int                `json:"genre_ids"`
	Genres              []Genre              `json:"genres"`
	Homepage            string               `json:"homepage"`
	ID                  int                  `json:"id"`
	IMDbID              string               `json:"imdb_id"`
	OriginalLanguage    string               `json:"original_language"`
	OriginalTitle       string               `json:"original_title"`
	Overview            string               `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          *string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany  `json:"production_companies"`
	ProductionCountries []ProductionCountry  `json:"production_countries"`
	ReleaseDate         string               `json:"release_date"`
	Revenue             int                  `json:"revenue"`
	Runtime             int                  `json:"runtime"`
	SpokenLanguages     []SpokenLanguage     `json:"spoken_languages"`
	Status              string               `json:"status"`
	Tagline             string               `json:"tagline"`
	Title               string               `json:"title"`
	Video               bool                 `json:"video"`
	VoteAverage         float64              `json:"vote_average"`
	VoteCount           int                  `json:"vote_count"`

	// append to response
	AlternativeTitles *AlternativeMovieTitles `json:"alternative_titles"`
	Changes           *Changes                `json:"changes"`
	Credits           *MovieCredits           `json:"credits"`
	ExternalIDs       *MovieExternalIDs       `json:"external_ids"`
	Images            *Images                 `json:"images"`
	Keywords          *MovieKeywords          `json:"keywords"`
	Lists             *MovieLists             `json:"lists"`
	Recommendations   *RecommendedMovies      `json:"recommendations"`
	ReleaseDates      *MovieReleaseDates      `json:"release_dates"`
	Reviews           *MovieReviews           `json:"reviews"`
	Similar           *SimilarMovies          `json:"similar"`
	Translations      *MovieTranslations      `json:"translations"`
	Videos            *Videos                 `json:"videos"`
}

// MovieDetailsOptions represents the available options for the request.
type MovieDetailsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	AppendToResponse string `url:"append_to_response,omitempty" json:"append_to_response,omitempty"`
}

// GetMovie retrieves the primary information about a movie.
func (mr *MoviesResource) GetMovie(movieID int, opt *MovieDetailsOptions) (*MovieDetails, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d", movieID)
	var movie MovieDetails
	resp, err := mr.client.get(path, &movie, WithQueryParams(opt))
	return &movie, resp, errors.Wrap(err, "failed to get movie")
}

// GetMoviesChanges retrieves a list of all of the movie ids that have been changed in the past 24 hours.
// Query it for up to 14 days worth of changed IDs at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
func (mr *MoviesResource) GetMoviesChanges(opt *ChangesOptions) (*MediaChanges, *http.Response, error) {
	path := "/movie/changes"
	var changes MediaChanges
	resp, err := mr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get movies changes")
}

// ProductionCompany represents a production company in TMDb.
type ProductionCompany struct {
	Name          string  `json:"name"`
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	OriginCountry string  `json:"origin_country"`
}

// ProductionCountry represents a production country in TMDb.
type ProductionCountry struct {
	Name     string `json:"name"`
	ISO31661 string `json:"iso_3166_1"`
}

// SpokenLanguage represents a spoken language in TMDb.
type SpokenLanguage struct {
	ISO6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
}

// BelongsToCollection represents a belongs to collection object in TMDb.
type BelongsToCollection struct {
	BackdropPath *string `json:"backdrop_path"`
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	PosterPath   *string `json:"poster_path"`
}

// LatestMovie represents the latest movie in TMDb.
type LatestMovie struct {
	Adult               bool                 `json:"adult"`
	BackdropPath        *string              `json:"backdrop_path"`
	BelongsToCollection *BelongsToCollection `json:"belongs_to_collection"`
	Budget              int                  `json:"budget"`
	GenreIDs            []int                `json:"genre_ids"`
	Genres              []Genre              `json:"genres"`
	Homepage            string               `json:"homepage"`
	ID                  int                  `json:"id"`
	IMDbID              string               `json:"imdb_id"`
	OriginalLanguage    string               `json:"original_language"`
	OriginalTitle       string               `json:"original_title"`
	Overview            string               `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          *string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany  `json:"production_companies"`
	ProductionCountries []ProductionCountry  `json:"production_countries"`
	ReleaseDate         string               `json:"release_date"`
	Revenue             int                  `json:"revenue"`
	Runtime             int                  `json:"runtime"`
	SpokenLanguages     []SpokenLanguage     `json:"spoken_languages"`
	Status              string               `json:"status"`
	Tagline             string               `json:"tagline"`
	Title               string               `json:"title"`
	Video               bool                 `json:"video"`
	VoteAverage         float64              `json:"vote_average"`
	VoteCount           int                  `json:"vote_count"`
}

// LatestOptions represents the available options for the request.
type LatestOptions languageOptions

// GetLatest retrieves the most newly created movie. This is a live response and will continuously change.
func (mr *MoviesResource) GetLatest(opt *LatestOptions) (*LatestMovie, *http.Response, error) {
	path := "/movie/latest"
	var latest LatestMovie
	resp, err := mr.client.get(path, &latest, WithQueryParams(opt))
	return &latest, resp, errors.Wrap(err, "failed to get latest movie")
}

// NowPlayingMoviesOptions represents the available options for the request.
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

// DateRange represents a date range in TMDb.
type DateRange struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}

// NowPlayingMovies represents the now playing movies in TMDb.
type NowPlayingMovies struct {
	pagination
	Movies []MovieResult `json:"results"`
	Dates  DateRange     `json:"dates"`
}

// GetNowPlaying retrieves a list of movies in theatres.
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

// PopularMoviesOptions represents the available options for the request.
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

// PopularMovies represents popular movies in TMDb.
type PopularMovies paginatedMovies

// GetPopular retrieves a list of the current popular movies on TMDB. This list updates daily.
func (mr *MoviesResource) GetPopular(opt *PopularMoviesOptions) (*PopularMovies, *http.Response, error) {
	path := "/movie/popular"
	var movies PopularMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get popular movies")
}

// TopRatedMoviesOptions represents the available options for the request.
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

// TopRatedMovies represents the top rated movies in TMDb.
type TopRatedMovies paginatedMovies

// GetTopRated retrieves the top rated movies on TMDB.
func (mr *MoviesResource) GetTopRated(opt *TopRatedMoviesOptions) (*TopRatedMovies, *http.Response, error) {
	path := "/movie/top_rated"
	var movies TopRatedMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get top rated movies")
}

// UpcomingMoviesOptions represents the available options for the request.
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

// UpcomingMovies represents the upcoming movies in TMDb.
type UpcomingMovies struct {
	pagination
	Movies []MovieResult `json:"results"`
	Dates  DateRange     `json:"dates"`
}

// GetUpcoming retrieves a list of upcoming movies in theatres.
// This is a release type query that looks for all movies that have a release type of 2 or 3 within the specified date range.
// Optionally specify a region parameter which will narrow the search to only look for theatrical release dates
// within the specified country.
func (mr *MoviesResource) GetUpcoming(opt *UpcomingMoviesOptions) (*UpcomingMovies, *http.Response, error) {
	path := "/movie/upcoming"
	var movies UpcomingMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get upcoming movies")
}

// AccountStates represents account states in TMDb.
type AccountStates struct {
	ID        int         `json:"id"`
	Favorite  bool        `json:"favorite"`
	Rated     interface{} `json:"rated"`
	Watchlist bool        `json:"watchlist"`
}

// GetAccountStates retrieves the following account states for a session:
// - Movie rating
// - If it belongs to the watchlist
// - If it belongs to the favorite list
func (mr *MoviesResource) GetAccountStates(movieID int, sessionID string) (*AccountStates, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/account_states", movieID)
	var states AccountStates
	resp, err := mr.client.get(path, &states, WithSessionID(sessionID))
	return &states, resp, errors.Wrap(err, "failed to get account states")
}

// Auth represents auth in TMDb.
type Auth struct {
	SessionID      string `url:"session_id,omitempty" json:"session_id,omitempty"`
	GuestSessionID string `url:"guest_session_id,omitempty" json:"guest_session_id,omitempty"`
}

// RateResponse represents a rate response in TMDb.
type RateResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// Rate rates a movie.
// A valid session or guest session ID is required.
func (mr *MoviesResource) Rate(movieID int, rating float64, sessionID Auth) (*RateResponse, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/rating", movieID)
	var response RateResponse
	resp, err := mr.client.post(path, &response, WithBody(map[string]float64{"value": rating}), WithQueryParams(sessionID))
	return &response, resp, errors.Wrap(err, "failed to rate movie")
}

// DeleteRatingResponse represents the delete rating response in TMDb.
type DeleteRatingResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// DeleteRating removes a rating for a movie.
// A valid session or guest session ID is required.
func (mr *MoviesResource) DeleteRating(movieID int, sessionID Auth) (*DeleteRatingResponse, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/rating", movieID)
	var response DeleteRatingResponse
	resp, err := mr.client.delete(path, &response, WithQueryParams(sessionID))
	return &response, resp, errors.Wrap(err, "failed to delete movie rating")
}

// Title represents title in TMDb.
type Title struct {
	ISO31661 string `json:"iso_3166_1"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

// AlternativeMovieTitles represents alternative movie titles in TMDb.
type AlternativeMovieTitles struct {
	ID     *int    `json:"id"`
	Titles []Title `json:"titles"`
}

// MovieAlternativeTitlesOptions represents the available options for the request.
type MovieAlternativeTitlesOptions struct {
	Country string `url:"country,omitempty" json:"country,omitempty"`
}

// GetAlternativeTitles retrieves all of the alternative titles for a movie.
func (mr *MoviesResource) GetAlternativeTitles(movieID int, opt *MovieAlternativeTitlesOptions) (*AlternativeMovieTitles, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/alternative_titles", movieID)
	var titles AlternativeMovieTitles
	resp, err := mr.client.get(path, &titles, WithQueryParams(opt))
	return &titles, resp, errors.Wrap(err, "failed to get alternative titles")
}

// GetChanges retrieves the changes for a movie. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the `start_date` and `end_date` query parameters.
func (mr *MoviesResource) GetChanges(movieID int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/changes", movieID)
	var changes Changes
	resp, err := mr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get movie changes")
}

// MovieCast represents a movie cast in TMDb.
type MovieCast struct {
	Adult              bool    `json:"adult"`
	CastID             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	Order              int     `json:"order"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

// MovieCrew represents a movie crew in TMDb.
type MovieCrew struct {
	Adult              bool    `json:"adult"`
	CreditID           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	Job                string  `json:"job"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

// MovieCredits represents movie credits in TMDb.
type MovieCredits struct {
	ID   *int        `json:"id"`
	Cast []MovieCast `json:"cast"`
	Crew []MovieCrew `json:"crew"`
}

// GetCredits retrieves the cast and crew for a movie.
func (mr *MoviesResource) GetCredits(movieID int, opt *CreditsOptions) (*MovieCredits, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/credits", movieID)
	var credits MovieCredits
	resp, err := mr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get movie credits")
}

// MovieExternalIDs represents movie external ids in TMDb.
type MovieExternalIDs struct {
	ID          *int    `json:"id"`
	IMDbID      *string `json:"imdb_id"`
	FacebookID  *string `json:"facebook_id"`
	InstagramID *string `json:"instagram_id"`
	TwitterID   *string `json:"twitter_id"`
}

// GetExternalIDs retrieves the external ids for a movie.
func (mr *MoviesResource) GetExternalIDs(movieID int) (*MovieExternalIDs, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/external_ids", movieID)
	var ids MovieExternalIDs
	resp, err := mr.client.get(path, &ids)
	return &ids, resp, errors.Wrap(err, "failed to get movie external ids")
}

// Logo represents a logo in TMDb.
type Logo Image

// Images represents images in TMDb.
type Images struct {
	ID        *int       `json:"id"`
	Backdrops []Backdrop `json:"backdrops"`
	Posters   []Poster   `json:"posters"`
	Logos     []Logo     `json:"logos"`
}

// ImagesOptions represents the available options for the request.
type ImagesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	IncludeImageLanguage string `url:"include_image_language,omitempty" json:"include_image_language,omitempty"`
}

// GetImages retrieves the images that belong to a movie.
// Querying images with a language parameter will filter the results.
// To include a fallback language (especially useful for backdrops), use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
func (mr *MoviesResource) GetImages(movieID int, opt *ImagesOptions) (*Images, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/images", movieID)
	var images Images
	resp, err := mr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get movie images")
}

// MovieKeywords represents movie keywords in TMDb.
type MovieKeywords struct {
	ID       *int      `json:"id"`
	Keywords []Keyword `json:"keywords"`
}

// GetKeywords retrieves the keywords that have been added to a movie.
func (mr *MoviesResource) GetKeywords(movieID int) (*MovieKeywords, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/keywords", movieID)
	var keywords MovieKeywords
	resp, err := mr.client.get(path, &keywords)
	return &keywords, resp, errors.Wrap(err, "failed to get movie keywords")
}

// MovieList represents a movie list in TMDb.
type MovieList struct {
	Description   string  `json:"description"`
	FavoriteCount int     `json:"favorite_count"`
	ID            int     `json:"id"`
	ISO6391       string  `json:"iso_639_1"`
	ItemCount     int     `json:"item_count"`
	ListType      string  `json:"list_type"`
	Name          string  `json:"name"`
	PosterPath    *string `json:"poster_path"`
}

// MovieLists represents movie lists in TMDb.
type MovieLists struct {
	pagination
	Lists []MovieList `json:"results"`
}

// MoviesOptions represents the available options for the request.
type MoviesOptions languagePageOptions

// GetLists retrieves 
func (mr *MoviesResource) GetLists(movieID int, opt *MoviesOptions) (*MovieLists, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/lists", movieID)
	var lists MovieLists
	resp, err := mr.client.get(path, &lists, WithQueryParams(opt))
	return &lists, resp, errors.Wrap(err, "failed to get movie lists")
}

// RecommendedMovies represents recommended movies in TMDb.
type RecommendedMovies struct {
	pagination
	Movies []Movie `json:"results"`
}

// GetRecommendations retrieves a list of recommended movies for a movie.
func (mr *MoviesResource) GetRecommendations(movieID int, opt *MoviesOptions) (*RecommendedMovies, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/recommendations", movieID)
	var movies RecommendedMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get movie recommendations")
}

// MovieReleaseDate represents movie release date in TMDb.
type MovieReleaseDate struct {
	Certification string  `json:"certification"`
	ISO6391       *string `json:"iso_639_1"`
	ReleaseDate   string  `json:"release_date"`
	Type          int     `json:"type"`
	Note          string  `json:"note"`
}

// MovieRelease represents a movie release in TMDb.
type MovieRelease struct {
	ISO31661     string             `json:"iso_3166_1"`
	ReleaseDates []MovieReleaseDate `json:"release_dates"`
}

// MovieReleaseDates represents movie release dates in TMDb.
type MovieReleaseDates struct {
	ID       *int           `json:"id"`
	Releases []MovieRelease `json:"results"`
}

// GetReleaseDates retrieves the release date along with the certification for a movie.
// Release dates support different types:
// 1. Premiere
// 2. Theatrical (limited)
// 3. Theatrical
// 4. Digital
// 5. Physical
// 6. TV
func (mr *MoviesResource) GetReleaseDates(movieID int) (*MovieReleaseDates, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/release_dates", movieID)
	var dates MovieReleaseDates
	resp, err := mr.client.get(path, &dates)
	return &dates, resp, errors.Wrap(err, "failed to get movie release dates")
}

// MovieReviews represents movie reviews in TMDb.
type MovieReviews struct {
	pagination
	ID      *int     `json:"id"`
	Reviews []Review `json:"results"`
}

// GetReviews retrieves the user reviews for a movie.
func (mr *MoviesResource) GetReviews(movieID int, opt *MoviesOptions) (*MovieReviews, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/reviews", movieID)
	var reviews MovieReviews
	resp, err := mr.client.get(path, &reviews, WithQueryParams(opt))
	return &reviews, resp, errors.Wrap(err, "failed to get movie reviews")
}

// SimilarMovies represents similar movies in TMDb.
type SimilarMovies paginatedMovies

// GetSimilar retrieves a list of similar movies.
// This is not the same as the "Recommendation" system on the website.
// These items are assembled by looking at keywords and genres.
func (mr *MoviesResource) GetSimilar(movieID int, opt *MoviesOptions) (*SimilarMovies, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/similar", movieID)
	var movies SimilarMovies
	resp, err := mr.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get similar movies")
}

// MovieTranslations represents movie translations in TMDb.
type MovieTranslations struct {
	ID           *int               `json:"id"`
	Translations []MovieTranslation `json:"translations"`
}

// MovieData represents movie data in TMDb.
type MovieData struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
	Tagline  string `json:"tagline"`
	Runtime  int    `json:"runtime"`
}

// MovieTranslation represents a movie translation in TMDb.
type MovieTranslation struct {
	ISO31661    string    `json:"iso_3166_1"`
	ISO6391     string    `json:"iso_639_1"`
	Name        string    `json:"name"`
	EnglishName string    `json:"english_name"`
	Data        MovieData `json:"data"`
}

// GetTranslations retrieves a list of translations that have been created for a movie.
func (mr *MoviesResource) GetTranslations(movieID int) (*MovieTranslations, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/translations", movieID)
	var translations MovieTranslations
	resp, err := mr.client.get(path, &translations)
	return &translations, resp, errors.Wrap(err, "failed to get movie translations")
}

// VideosOptions represents the available options for the request.
type VideosOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	IncludeVideoLanguage string `url:"include_video_language,omitempty" json:"include_video_language,omitempty"`
}

// Video represents a video in TMDb.
type Video struct {
	ID          string `json:"id"`
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

// Videos represents videos in TMDb.
type Videos struct {
	ID     *int    `json:"id"`
	Videos []Video `json:"results"`
}

// GetVideos retrieves the videos that have been added to a movie.
func (mr *MoviesResource) GetVideos(movieID int, opt *VideosOptions) (*Videos, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/videos", movieID)
	var videos Videos
	resp, err := mr.client.get(path, &videos, WithQueryParams(opt))
	return &videos, resp, errors.Wrap(err, "failed to get movie videos")
}

// Providers represents providers in TMDb.
type Providers map[string]interface{}

// WatchProviders represents watch providers in TMDb.
type WatchProviders struct {
	ID        int       `json:"id"`
	Providers Providers `json:"results"`
}

// GetWatchProviders retrieves watch providers for a movie.
// Powered by the partnership with JustWatch, use this method to get a list of the availabilities per country by provider.
// This is not going to return full deep links, but rather, it's just enough information to display what's available where.
// Link to the provided TMDB URL to help support TMDB and provide the actual deep links to the content.
// Please note: In order to use this data it's REQUIRED to attribute the source of the data as JustWatch.
// If any usage is found not complying with these terms the access to the API will be revoked.
func (mr *MoviesResource) GetWatchProviders(movieID int) (*WatchProviders, *http.Response, error) {
	path := fmt.Sprintf("/movie/%d/watch/providers", movieID)
	var providers WatchProviders
	resp, err := mr.client.get(path, &providers)
	return &providers, resp, errors.Wrap(err, "failed to get movie watch providers")
}
