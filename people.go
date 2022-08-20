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

type Person struct {
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	Id                 int         `json:"id"`
	KnownForDepartment string      `json:"known_for_department"`
	MediaType          string      `json:"media_type"`
	Name               string      `json:"name"`
	OriginalName       string      `json:"original_name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        string      `json:"profile_path"`
	KnownFor           []MovieOrTV `json:"known_for"`
}

type PersonDetails struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          *string  `json:"biography"`
	Birthday           *string  `json:"birthday"`
	Deathday           *string  `json:"deathday"`
	Gender             int      `json:"gender"`
	Homepage           *string  `json:"homepage"`
	Id                 int      `json:"id"`
	ImdbId             *string  `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       *string  `json:"place_of_birth"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`

	// append to response
	Changes         *Changes             `json:"changes"`
	CombinedCredits *CombinedCredits     `json:"combined_credits"`
	ExternalIds     *PersonExternalIds   `json:"external_ids"`
	Images          *PersonImages        `json:"images"`
	MovieCredits    *PersonMovieCredits  `json:"movie_credits"`
	TaggedImages    *TaggedImages        `json:"tagged_images"`
	Translations    *PersonTranslations  `json:"translations"`
	TVShowCredits   *PersonTVShowCredits `json:"tv_credits"`
}

type PersonDetailsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	AppendToResponse string `url:"append_to_response,omitempty" json:"append_to_response,omitempty"`
}

func (pr *PeopleResource) GetPerson(personId int, opt *PersonDetailsOptions) (*PersonDetails, *http.Response, error) {
	path := fmt.Sprintf("/person/%d", personId)
	var person PersonDetails
	resp, err := pr.client.get(path, &person, WithQueryParams(opt))
	return &person, resp, errors.Wrap(err, "failed to get person")
}

type Profile struct {
	FilePath *string `json:"file_path"`
}

func (pr *PeopleResource) GetChanges(personId int, opt *ChangesOptions) (*Changes, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/changes", personId)
	var changes Changes
	resp, err := pr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get changes")
}

type MovieCastPerson struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	Character        string  `json:"character"`
	CreditId         string  `json:"credit_id"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
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

type MovieCrewPerson struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	CreditId         string  `json:"credit_id"`
	Department       string  `json:"department"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
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

type PersonMovieCredits struct {
	Id   *int              `json:"id"`
	Cast []MovieCastPerson `json:"cast"`
	Crew []MovieCrewPerson `json:"crew"`
}

// Get the movie credits for a person
func (pr *PeopleResource) GetMovieCredits(personId int, opt *CreditsOptions) (*PersonMovieCredits, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/movie_credits", personId)
	var credits PersonMovieCredits
	resp, err := pr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get movie credits")
}

type TVShowCastPerson struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	Character        string   `json:"character"`
	CreditId         string   `json:"credit_id"`
	EpisodeCount     int      `json:"episode_count"`
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

type TVShowCrewPerson struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	CreditId         string   `json:"credit_id"`
	Department       string   `json:"department"`
	EpisodeCount     int      `json:"episode_count"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
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

type PersonTVShowCredits struct {
	Id   *int               `json:"id"`
	Cast []TVShowCastPerson `json:"cast"`
	Crew []TVShowCrewPerson `json:"crew"`
}

// Get the TV show credits for a person.
func (pr *PeopleResource) GetTVCredits(personId int, opt *CreditsOptions) (*PersonTVShowCredits, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/tv_credits", personId)
	var credits PersonTVShowCredits
	resp, err := pr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get tv show credits")
}

type CombinedCreditsCast map[string]interface{}

type CombinedCreditsMovieCast struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	Character        string  `json:"character"`
	CreditId         string  `json:"credit_id"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
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

type CombinedCreditsTVShowCast struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	Character        string   `json:"character"`
	CreditId         string   `json:"credit_id"`
	EpisodeCount     int      `json:"episode_count"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
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

type CombinedCreditsCrew map[string]interface{}

type CombinedCreditsMovieCrew struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	CreditId         string  `json:"credit_id"`
	Department       string  `json:"department"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
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

type CombinedCreditsTVShowCrew struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	CreditId         string   `json:"credit_id"`
	Department       string   `json:"department"`
	EpisodeCount     int      `json:"episode_count"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
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

type CombinedCredits struct {
	Id   *int                  `json:"id"`
	Cast []CombinedCreditsCast `json:"cast"`
	Crew []CombinedCreditsCrew `json:"crew"`
}

func (cc CombinedCreditsCast) GetMediaType() string {
	return cc["media_type"].(string)
}

func (cc CombinedCreditsCast) ToMovieCast() (*CombinedCreditsMovieCast, error) {
	var credits CombinedCreditsMovieCast
	err := convert("movie", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to movie cast")
}

func (cc CombinedCreditsCast) ToTVShowCast() (*CombinedCreditsTVShowCast, error) {
	var credits CombinedCreditsTVShowCast
	err := convert("tv", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to tv cast")
}

func (cc CombinedCreditsCrew) GetMediaType() string {
	return cc["media_type"].(string)
}

func (cc CombinedCreditsCrew) ToMovieCrew() (*CombinedCreditsMovieCrew, error) {
	var credits CombinedCreditsMovieCrew
	err := convert("movie", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to movie crew")
}

func (cc CombinedCreditsCrew) ToTVShowCrew() (*CombinedCreditsTVShowCrew, error) {
	var credits CombinedCreditsTVShowCrew
	err := convert("tv", cc, &credits)
	return &credits, errors.Wrap(err, "failed to convert object to tv crew")
}

// Get the movie and TV credits together in a single response.
func (pr *PeopleResource) GetCombinedCredits(personId int, opt *CreditsOptions) (*CombinedCredits, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/combined_credits", personId)
	var credits CombinedCredits
	resp, err := pr.client.get(path, &credits, WithQueryParams(opt))
	return &credits, resp, errors.Wrap(err, "failed to get combined credits")
}

type PersonExternalIds struct {
	Id          *int    `json:"id"`
	FacebookId  *string `json:"facebook_id"`
	FreebaseId  *string `json:"freebase_id"`
	FreebaseMId *string `json:"freebase_mid"`
	IMDbId      *string `json:"imdb_id"`
	InstagramId *string `json:"instagram_id"`
	TVRageId    *int    `json:"tvrage_id"`
	TwitterId   *string `json:"twitter_id"`
}

type ExternalIdOptions languageOptions

// Get the external ids for a person.
// Currently supported external sources:
// IMDB ID, Facebook, Freebase MID, Freebase ID, Instagram, TVRage ID, Twitter
func (pr *PeopleResource) GetExternalIDs(personId int, opt *ExternalIdOptions) (*PersonExternalIds, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/external_ids", personId)
	var externalIds PersonExternalIds
	resp, err := pr.client.get(path, &externalIds, WithQueryParams(opt))
	return &externalIds, resp, errors.Wrap(err, "failed to get external ids")
}

type PersonImages struct {
	Id       *int    `json:"id"`
	Profiles []Image `json:"profiles"`
}

// Get the images for a person.
func (pr *PeopleResource) GetImages(personId int) (*PersonImages, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/images", personId)
	var images PersonImages
	resp, err := pr.client.get(path, &images)
	return &images, resp, errors.Wrap(err, "failed to get images")
}

type TaggedImagesOptions languagePageOptions

type TaggedImage struct {
	AspectRatio float64   `json:"aspect_ratio"`
	FilePath    string    `json:"file_path"`
	Height      int       `json:"height"`
	Id          string    `json:"id"`
	ImageType   string    `json:"image_type"`
	ISO6391     *string   `json:"iso_639_1"`
	Media       MovieOrTV `json:"media"`
	MediaType   string    `json:"media_type"`
	VoteAverage float64   `json:"vote_average"`
	VoteCount   int       `json:"vote_count"`
	Width       int       `json:"width"`
}

type TaggedImages struct {
	pagination
	Id     *int          `json:"id"`
	Images []TaggedImage `json:"results"`
}

// Get the images that this person has been tagged in.
func (pr *PeopleResource) GetTaggedImages(personId int, opt *TaggedImagesOptions) (*TaggedImages, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/tagged_images", personId)
	var images TaggedImages
	resp, err := pr.client.get(path, &images, WithQueryParams(opt))
	return &images, resp, errors.Wrap(err, "failed to get images")
}

type PersonTranslationsOptions languageOptions

type PersonTranslations struct {
	Id           *int                `json:"id"`
	Translations []PersonTranslation `json:"translations"`
}

type PersonData struct {
	Biography string `json:"biography"`
}

type PersonTranslation struct {
	ISO31661    string     `json:"iso_3166_1"`
	ISO6391     string     `json:"iso_639_1"`
	Name        string     `json:"name"`
	EnglishName string     `json:"english_name"`
	Data        PersonData `json:"data"`
}

// Get a list of translations that have been created for a person.
func (pr *PeopleResource) GetTranslations(personId int, opt *PersonTranslationsOptions) (*PersonTranslations, *http.Response, error) {
	path := fmt.Sprintf("/person/%d/translations", personId)
	var translations PersonTranslations
	resp, err := pr.client.get(path, &translations, WithQueryParams(opt))
	return &translations, resp, errors.Wrap(err, "failed to get translations")
}

type LatestPersonOptions languageOptions

type LatestPerson struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          *string  `json:"biography"`
	Birthday           *string  `json:"birthday"`
	Deathday           *string  `json:"deathday"`
	Gender             int      `json:"gender"`
	Homepage           *string  `json:"homepage"`
	Id                 int      `json:"id"`
	ImdbId             *string  `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       *string  `json:"place_of_birth"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        *string  `json:"profile_path"`
}

// Get the most newly created person. This is a live response and will continuously change.
func (pr *PeopleResource) GetLatest(opt *LatestPersonOptions) (*LatestPerson, *http.Response, error) {
	path := "/person/latest"
	var latest LatestPerson
	resp, err := pr.client.get(path, &latest, WithQueryParams(opt))
	return &latest, resp, errors.Wrap(err, "failed to get latest person")
}

type PopularPeopleOptions languagePageOptions

type PopularPerson struct {
	KnownFor           []MovieOrTV `json:"known_for"`
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	Id                 int         `json:"id"`
	KnownForDepartment string      `json:"known_for_department"`
	Name               string      `json:"name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        string      `json:"profile_path"`
}

type PopularPeople struct {
	pagination
	People []PopularPerson `json:"results"`
}

// Get the list of popular people on TMDB. This list updates daily.
func (pr *PeopleResource) GetPopular(opt *PopularPeopleOptions) (*PopularPeople, *http.Response, error) {
	path := "/person/popular"
	var popular PopularPeople
	resp, err := pr.client.get(path, &popular, WithQueryParams(opt))
	return &popular, resp, errors.Wrap(err, "failed to get popular people")
}

// Get a list of all of the person ids that have been changed in the past 24 hours.
// Query it for up to 14 days worth of changed IDs at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
func (pr *PeopleResource) GetPeopleChanges(opt *ChangesOptions) (*MediaChanges, *http.Response, error) {
	path := "/person/changes"
	var changes MediaChanges
	resp, err := pr.client.get(path, &changes, WithQueryParams(opt))
	return &changes, resp, errors.Wrap(err, "failed to get people changes")
}
