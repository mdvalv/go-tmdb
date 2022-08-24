package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// PeopleResource handles person-related requests of TMDb API.
type PeopleResource struct {
	client *Client
}

// Person represents a person in TMDb.
type Person struct {
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	ID                 int         `json:"id"`
	KnownForDepartment string      `json:"known_for_department"`
	MediaType          string      `json:"media_type"`
	Name               string      `json:"name"`
	OriginalName       string      `json:"original_name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        string      `json:"profile_path"`
	KnownFor           []MovieOrTV `json:"known_for"`
}

// PersonDetails represents person details in TMDb.
type PersonDetails struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          *string  `json:"biography"`
	Birthday           *string  `json:"birthday"`
	Deathday           *string  `json:"deathday"`
	Gender             int      `json:"gender"`
	Homepage           *string  `json:"homepage"`
	ID                 int      `json:"id"`
	ImdbID             *string  `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       *string  `json:"place_of_birth"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`

	// append to response
	Changes         *Changes             `json:"changes"`
	CombinedCredits *CombinedCredits     `json:"combined_credits"`
	ExternalIDs     *PersonExternalIDs   `json:"external_ids"`
	Images          *PersonImages        `json:"images"`
	MovieCredits    *PersonMovieCredits  `json:"movie_credits"`
	TaggedImages    *TaggedImages        `json:"tagged_images"`
	Translations    *PersonTranslations  `json:"translations"`
	TVShowCredits   *PersonTVShowCredits `json:"tv_credits"`
}

// PersonDetailsOptions represents the available options for the request.
type PersonDetailsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	AppendToResponse string `url:"append_to_response,omitempty" json:"append_to_response,omitempty"`
}

// GetPerson retrieves the primary person details by id.
func (pr *PeopleResource) GetPerson(personID int, opt *PersonDetailsOptions) (*PersonDetails, *http.Response, error) {
	path := fmt.Sprintf("/person/%d", personID)
	var person PersonDetails
	resp, err := pr.client.get(path, &person, WithQueryParams(opt))
	return &person, resp, errors.Wrap(err, "failed to get person")
}

// Profile represents a profile in TMDb.
type Profile struct {
	FilePath *string `json:"file_path"`
}

// GetChanges retrieves the changes for a person. By default only the last 24 hours are returned.
// Query up to 14 days in a single query by using the start_date and end_date query parameters.
func (pr *PeopleResource) GetChanges(personID int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/changes", personID)
	var changes Changes
	resp, err := pr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get changes")
}

// MovieCastPerson represents a person in a movie cast in TMDb.
type MovieCastPerson struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	Character        string  `json:"character"`
	CreditID         string  `json:"credit_id"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	Order            int     `json:"order"`
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

// MovieCrewPerson represents a person in a movie crew in TMDb.
type MovieCrewPerson struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	CreditID         string  `json:"credit_id"`
	Department       string  `json:"department"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	Job              string  `json:"job"`
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

// PersonMovieCredits represents a person movie credits in TMDb.
type PersonMovieCredits struct {
	ID   *int              `json:"id"`
	Cast []MovieCastPerson `json:"cast"`
	Crew []MovieCrewPerson `json:"crew"`
}

// GetMovieCredits retrieves the movie credits for a person.
func (pr *PeopleResource) GetMovieCredits(personID int, opt *CreditsOptions) (*PersonMovieCredits, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/movie_credits", personID)
	var credits PersonMovieCredits
	resp, err := pr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get movie credits")
}

// TVShowCastPerson represents a person in a tv show cast in TMDb.
type TVShowCastPerson struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	Character        string   `json:"character"`
	CreditID         string   `json:"credit_id"`
	EpisodeCount     int      `json:"episode_count"`
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

// TVShowCrewPerson represents a person in a tv show crew in TMDb.
type TVShowCrewPerson struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	CreditID         string   `json:"credit_id"`
	Department       string   `json:"department"`
	EpisodeCount     int      `json:"episode_count"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Job              string   `json:"job"`
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

// PersonTVShowCredits represents a person tv show credits in TMDb.
type PersonTVShowCredits struct {
	ID   *int               `json:"id"`
	Cast []TVShowCastPerson `json:"cast"`
	Crew []TVShowCrewPerson `json:"crew"`
}

// GetTVCredits retrieves the tv show credits for a person.
func (pr *PeopleResource) GetTVCredits(personID int, opt *CreditsOptions) (*PersonTVShowCredits, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/tv_credits", personID)
	var credits PersonTVShowCredits
	resp, err := pr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get tv show credits")
}

// CombinedCreditsCast represents a cast in TMDb.
type CombinedCreditsCast map[string]interface{}

// CombinedCreditsMovieCast represents a movie cast in TMDb.
type CombinedCreditsMovieCast struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	Character        string  `json:"character"`
	CreditID         string  `json:"credit_id"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	MediaType        string  `json:"media_type"`
	Order            int     `json:"order"`
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

// CombinedCreditsTVShowCast represents a tv show cast in TMDb.
type CombinedCreditsTVShowCast struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	Character        string   `json:"character"`
	CreditID         string   `json:"credit_id"`
	EpisodeCount     int      `json:"episode_count"`
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

// CombinedCreditsCrew represents a crew in TMDb.
type CombinedCreditsCrew map[string]interface{}

// CombinedCreditsMovieCrew represents a movie crew in TMDb.
type CombinedCreditsMovieCrew struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	CreditID         string  `json:"credit_id"`
	Department       string  `json:"department"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	Job              string  `json:"job"`
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

// CombinedCreditsTVShowCrew represents a tv show crew in TMDb.
type CombinedCreditsTVShowCrew struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	CreditID         string   `json:"credit_id"`
	Department       string   `json:"department"`
	EpisodeCount     int      `json:"episode_count"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Job              string   `json:"job"`
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

// CombinedCredits represents combined credits in TMDb.
type CombinedCredits struct {
	ID   *int                  `json:"id"`
	Cast []CombinedCreditsCast `json:"cast"`
	Crew []CombinedCreditsCrew `json:"crew"`
}

// GetMediaType retrieves the media type from a combined credits cast.
func (cc CombinedCreditsCast) GetMediaType() string {
	return cc["media_type"].(string)
}

// ToMovieCast converts the data to a movie cast.
func (cc CombinedCreditsCast) ToMovieCast() (*CombinedCreditsMovieCast, error) {
	var credits CombinedCreditsMovieCast
	err := convert("movie", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to movie cast")
}

// ToTVShowCast converts the data to a tv show cast.
func (cc CombinedCreditsCast) ToTVShowCast() (*CombinedCreditsTVShowCast, error) {
	var credits CombinedCreditsTVShowCast
	err := convert("tv", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to tv cast")
}

// GetMediaType retrieves the media type from a combined credits crew.
func (cc CombinedCreditsCrew) GetMediaType() string {
	return cc["media_type"].(string)
}

// ToMovieCrew converts the data to a movie crew.
func (cc CombinedCreditsCrew) ToMovieCrew() (*CombinedCreditsMovieCrew, error) {
	var credits CombinedCreditsMovieCrew
	err := convert("movie", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to movie crew")
}

// ToTVShowCrew converts the data to a tv show crew.
func (cc CombinedCreditsCrew) ToTVShowCrew() (*CombinedCreditsTVShowCrew, error) {
	var credits CombinedCreditsTVShowCrew
	err := convert("tv", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to tv crew")
}

// GetCombinedCredits retrieves the movie and TV credits together in a single response.
func (pr *PeopleResource) GetCombinedCredits(personID int, opt *CreditsOptions) (*CombinedCredits, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/combined_credits", personID)
	var credits CombinedCredits
	resp, err := pr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get combined credits")
}

// PersonExternalIDs represents a person external ids in TMDb.
type PersonExternalIDs struct {
	ID          *int    `json:"id"`
	FacebookID  *string `json:"facebook_id"`
	FreebaseID  *string `json:"freebase_id"`
	FreebaseMID *string `json:"freebase_mid"`
	IMDbID      *string `json:"imdb_id"`
	InstagramID *string `json:"instagram_id"`
	TVRageID    *int    `json:"tvrage_id"`
	TwitterID   *string `json:"twitter_id"`
}

// ExternalIDOptions represents the available options for the request.
type ExternalIDOptions languageOptions

// GetExternalIDs retrieves the external ids for a person.
// Currently supported external sources:
// IMDB ID, Facebook, Freebase MID, Freebase ID, Instagram, TVRage ID, Twitter
func (pr *PeopleResource) GetExternalIDs(personID int, opt *ExternalIDOptions) (*PersonExternalIDs, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/external_ids", personID)
	var externalIDs PersonExternalIDs
	resp, err := pr.client.get(path, &externalIDs, WithQueryParams(opt))
	return &externalIDs, resp, errors.Wrap(err, "failed to get external ids")
}

// PersonImages represents person images in TMDb.
type PersonImages struct {
	ID       *int    `json:"id"`
	Profiles []Image `json:"profiles"`
}

// GetImages retrieves the images for a person.
func (pr *PeopleResource) GetImages(personID int) (*PersonImages, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/images", personID)
	var images PersonImages
	resp, err := pr.client.get(path, &images)
	return &images, resp, errors.Wrap(err, "failed to get images")
}

// TaggedImagesOptions represents the available options for the request.
type TaggedImagesOptions languagePageOptions

// TaggedImage represents a tagged image in TMDb.
type TaggedImage struct {
	AspectRatio float64   `json:"aspect_ratio"`
	FilePath    string    `json:"file_path"`
	Height      int       `json:"height"`
	ID          string    `json:"id"`
	ImageType   string    `json:"image_type"`
	ISO6391     *string   `json:"iso_639_1"`
	Media       MovieOrTV `json:"media"`
	MediaType   string    `json:"media_type"`
	VoteAverage float64   `json:"vote_average"`
	VoteCount   int       `json:"vote_count"`
	Width       int       `json:"width"`
}

// TaggedImages represents tagged images in TMDb.
type TaggedImages struct {
	pagination
	ID     *int          `json:"id"`
	Images []TaggedImage `json:"results"`
}

// GetTaggedImages retrieves the images that this person has been tagged in.
func (pr *PeopleResource) GetTaggedImages(personID int, opt *TaggedImagesOptions) (*TaggedImages, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/tagged_images", personID)
	var images TaggedImages
	resp, err := pr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

// PersonTranslationsOptions represents the available options for the request.
type PersonTranslationsOptions languageOptions

// PersonTranslations represents person translations in TMDb.
type PersonTranslations struct {
	ID           *int                `json:"id"`
	Translations []PersonTranslation `json:"translations"`
}

// PersonData represents person data in TMDb.
type PersonData struct {
	Biography string `json:"biography"`
}

// PersonTranslation represents a person translation in TMDb.
type PersonTranslation struct {
	ISO31661    string     `json:"iso_3166_1"`
	ISO6391     string     `json:"iso_639_1"`
	Name        string     `json:"name"`
	EnglishName string     `json:"english_name"`
	Data        PersonData `json:"data"`
}

// GetTranslations retrieves a list of translations that have been created for a person.
func (pr *PeopleResource) GetTranslations(personID int, opt *PersonTranslationsOptions) (*PersonTranslations, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/translations", personID)
	var translations PersonTranslations
	resp, err := pr.client.get(path, &translations, WithQueryParams(opt))
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

// LatestPersonOptions represents the available options for the request.
type LatestPersonOptions languageOptions

// LatestPerson represents the latest person in TMDb.
type LatestPerson struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          *string  `json:"biography"`
	Birthday           *string  `json:"birthday"`
	Deathday           *string  `json:"deathday"`
	Gender             int      `json:"gender"`
	Homepage           *string  `json:"homepage"`
	ID                 int      `json:"id"`
	ImdbID             *string  `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       *string  `json:"place_of_birth"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
}

// GetLatest retrieves the most newly created person. This is a live response and will continuously change.
func (pr *PeopleResource) GetLatest(opt *LatestPersonOptions) (*LatestPerson, *http.Response, error) {
	path := "/person/latest"
	var latest LatestPerson
	resp, err := pr.client.get(path, &latest, WithQueryParams(opt))
	return &latest, resp, errors.Wrap(err, "failed to get latest person")
}

// PopularPeopleOptions represents the available options for the request.
type PopularPeopleOptions languagePageOptions

// PopularPerson represents a popular person in TMDb.
type PopularPerson struct {
	KnownFor           []MovieOrTV `json:"known_for"`
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	ID                 int         `json:"id"`
	KnownForDepartment string      `json:"known_for_department"`
	Name               string      `json:"name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        string      `json:"profile_path"`
}

// PopularPeople represents popular people in TMDb.
type PopularPeople struct {
	pagination
	People []PopularPerson `json:"results"`
}

// GetPopular retrieves the list of popular people on TMDB. This list updates daily.
func (pr *PeopleResource) GetPopular(opt *PopularPeopleOptions) (*PopularPeople, *http.Response, error) {
	path := "/person/popular"
	var popular PopularPeople
	resp, err := pr.client.get(path, &popular, WithQueryParams(opt))
	return &popular, resp, errors.Wrap(err, "failed to get popular people")
}

// GetPeopleChanges retrieves a list of all of the person ids that have been changed in the past 24 hours.
// Query it for up to 14 days worth of changed IDs at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
func (pr *PeopleResource) GetPeopleChanges(opt *ChangesOptions) (*MediaChanges, *http.Response, error) {
	path := "/person/changes"
	var changes MediaChanges
	resp, err := pr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get people changes")
}
