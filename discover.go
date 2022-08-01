package tmdb

import (
	"net/http"

	"github.com/pkg/errors"
)

// DiscoverResource handles discover-related requests of TMDb API.
type DiscoverResource struct {
	client *Client
}

type DiscoverMovies paginatedMovies
type DiscoverTVShows paginatedTVShows

type DiscoverMoviesOptions struct {
	// Choose from one of the many available sort options:
	//    popularity.asc / popularity.desc
	//    release_date.asc / release_date.desc
	//    revenue.asc / revenue.desc
	//    primary_release_date.asc / primary_release_date.desc
	//    original_title.asc / original_title.desc
	//    vote_average.asc / vote_average.desc
	//    vote_count.asc / vote_count.des
	// default: popularity.desc
	SortBy string `url:"sort_by,omitempty" json:"sort_by,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`

	// Used in conjunction with the certification filter, use this to specify a country with a valid certification.
	CertificationCountry string `url:"certification_country,omitempty" json:"certification_country,omitempty"`

	// Filter results with a valid certification from the 'certification_country' field.
	Certification string `url:"certification,omitempty" json:"certification,omitempty"`

	// Filter and only include movies that have a certification that is less than or equal to the specified value.
	CertificationLte string `url:"certification.lte,omitempty" json:"certification.lte,omitempty"`

	// Filter and only include movies that have a certification that is greater than or equal to the specified value.
	CertificationGte string `url:"certification.gte,omitempty" json:"certification.gte,omitempty"`

	// A filter and include or exclude adult movies.
	// default: false
	IncludeAdult bool `url:"include_adult,omitempty" json:"include_adult,omitempty"`

	// A filter to include or exclude videos.
	// default: false
	IncludeVideo bool `url:"include_video,omitempty" json:"include_video,omitempty"`

	// Specify a language to query translatable fields with.
	// default: en-US
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify the page of results to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// A filter to limit the results to a specific primary release year.
	PrimaryReleaseYear *int `url:"primary_release_year,omitempty" json:"primary_release_year,omitempty"`

	// Filter and only include movies that have a primary release date that is greater or equal to the specified value.
	// format: YYYY-MM-DD
	PrimaryReleaseDateGte string `url:"primary_release_date.gte,omitempty" json:"primary_release_date.gte,omitempty"`

	// Filter and only include movies that have a primary release date that is less than or equal to the specified value.
	// format: YYYY-MM-DD
	PrimaryReleaseDateLte string `url:"primary_release_date.lte,omitempty" json:"primary_release_date.lte,omitempty"`

	// Filter and only include movies that have a release date (looking at all release dates) that is greater or equal to the specified value.
	// format: YYYY-MM-DD
	ReleaseDateGte string `url:"release_date.gte,omitempty" json:"release_date.gte,omitempty"`

	// Filter and only include movies that have a release date (looking at all release dates) that is less than or equal to the specified value.
	// format: YYYY-MM-DD
	ReleaseDateLte string `url:"release_date.lte,omitempty" json:"release_date.lte,omitempty"`

	// Specify a comma (,) for AND or pipe (|) for OR separated values to filter release types by.
	// These release types map to the same values found on the movie release date method.
	WithReleaseType string `url:"with_release_type,omitempty" json:"with_release_type,omitempty"`

	// A filter to limit the results to a specific year (looking at all release dates).
	Year *int `url:"year,omitempty" json:"year,omitempty"`

	// Filter and only include movies that have a vote count that is greater or equal to the specified value.
	VoteCountGte *int `url:"vote_count.gte,omitempty" json:"vote_count.gte,omitempty"`

	// Filter and only include movies that have a vote count that is less than or equal to the specified value.
	VoteCountLte *int `url:"vote_count.lte,omitempty" json:"vote_count.lte,omitempty"`

	// Filter and only include movies that have a rating that is greater or equal to the specified value.
	VoteAverageGte *float64 `url:"vote_average.gte,omitempty" json:"vote_average.gte,omitempty"`

	// Filter and only include movies that have a rating that is less than or equal to the specified value.
	VoteAverageLte *float64 `url:"vote_average.lte,omitempty" json:"vote_average.lte,omitempty"`

	// A comma separated list of person ID's. Only include movies that have one of the ID's added as an actor.
	WithCast string `url:"with_cast,omitempty" json:"with_cast,omitempty"`

	// A comma separated list of person ID's. Only include movies that have one of the ID's added as a crew member.
	WithCrew string `url:"with_crew,omitempty" json:"with_crew,omitempty"`

	// A comma separated list of person ID's. Only include movies that have one of the ID's added as a either a actor or a crew member.
	WithPeople string `url:"with_people,omitempty" json:"with_people,omitempty"`

	// A comma separated list of production company ID's. Only include movies that have one of the ID's added as a production company.
	WithCompanies string `url:"with_companies,omitempty" json:"with_companies,omitempty"`

	// Comma separated value of genre ids that you want to include in the results.
	WithGenres string `url:"with_genres,omitempty" json:"with_genres,omitempty"`

	// Comma separated value of genre ids that you want to exclude from the results.
	WithoutGenres string `url:"without_genres,omitempty" json:"without_genres,omitempty"`

	// A comma separated list of keyword ID's. Only includes movies that have one of the ID's added as a keyword.
	WithKeywords string `url:"with_keywords,omitempty" json:"with_keywords,omitempty"`

	// Exclude items with certain keywords. You can comma and pipe separate these values to create an 'AND' or 'OR' logic.
	WithoutKeywords string `url:"without_keywords,omitempty" json:"without_keywords,omitempty"`

	// Filter and only include movies that have a runtime that is greater or equal to a value.
	WithRuntimeGte *int `url:"with_runtime.gte,omitempty" json:"with_runtime.gte,omitempty"`

	// Filter and only include movies that have a runtime that is less than or equal to a value.
	WithRuntimeLte *int `url:"with_runtime.lte,omitempty" json:"with_runtime.lte,omitempty"`

	// Specify an ISO 639-1 string to filter results by their original language value.
	WithOriginalLanguage string `url:"with_original_language,omitempty" json:"with_original_language,omitempty"`

	// A comma or pipe separated list of watch provider ID's.
	// Combine this filter with `watch_region` in order to filter your results by a specific watch provider in a specific region.
	WithWatchProviders string `url:"with_watch_providers,omitempty" json:"with_watch_providers,omitempty"`

	// An ISO 3166-1 code.
	// Combine this filter with `with_watch_providers` in order to filter your results by a specific watch provider in a specific region.
	WatchRegion string `url:"watch_region,omitempty" json:"watch_region,omitempty"`

	// In combination with `watch_region`, you can filter by monetization type.
	// types: flatrate, free, ads, rent, buy
	WithWatchMonetizationTypes string `url:"with_watch_monetization_types,omitempty" json:"with_watch_monetization_types,omitempty"`

	// Filter the results to exclude the specific production companies you specify here. `AND` / `OR` filters are supported.
	WithoutCompanies string `url:"without_companies,omitempty" json:"without_companies,omitempty"`
}

// Discover movies by different types of data like average rating, number of votes, genres and certifications.
//
// Please note, when using certification/certification.lte you must also specify certification_country.
// These two parameters work together in order to filter the results.
// You can only filter results with the countries we have added to our certifications list.
//
// If you specify the region parameter, the regional release date will be used instead of the primary release date.
// The date returned will be the first date based on your query (ie. if a with_release_type is specified).
// It's important to note the order of the release types that are used.
// Specifying "2|3" would return the limited theatrical release date
// as opposed to "3|2" which would return the theatrical date.
//
// Also note that a number of filters support being comma (,) or pipe (|) separated.
// Comma's are treated like an AND and query while pipe's are an OR.
//
// Some examples can be found here: https://www.themoviedb.org/documentation/api/discover
func (dr *DiscoverResource) DiscoverMovies(opt *DiscoverMoviesOptions) (*DiscoverMovies, *http.Response, error) {
	path := "/discover/movie"
	var discover DiscoverMovies
	resp, err := dr.client.get(path, &discover, WithQueryParams(opt))
	return &discover, resp, errors.Wrap(err, "failed to discover movies")
}

// path := "/discover/tv"

type DiscoverTVShowsOptions struct {
	// Choose from one of the many available sort options:
	//    vote_average.desc / vote_average.asc
	//    first_air_date.desc / first_air_date.asc
	//    popularity.desc / popularity.asc
	// default: popularity.desc
	SortBy string `url:"sort_by,omitempty" json:"sort_by,omitempty"`

	// Filter and only include TV shows that have a air date (by looking at all episodes) that is greater or equal to the specified value.
	// format: YYYY-MM-DD
	AirDateGte string `url:"air_date.gte,omitempty" json:"air_date.gte,omitempty"`

	// Filter and only include TV shows that have a air date (by looking at all episodes) that is less than or equal to the specified value.
	// format: YYYY-MM-DD
	AirDateLte string `url:"air_date.lte,omitempty" json:"air_date.lte,omitempty"`

	// Filter and only include TV shows that have a original air date that is greater or equal to the specified value.
	// Can be used in conjunction with the "include_null_first_air_dates" filter if you want to include items with no air date.
	// format: YYYY-MM-DD
	FirstAirDateGte string `url:"first_air_date.gte,omitempty" json:"first_air_date.gte,omitempty"`

	// Filter and only include TV shows that have a original air date that is less than or equal to the specified value.
	// Can be used in conjunction with the "include_null_first_air_dates" filter if you want to include items with no air date.
	FirstAirDateLte string `url:"first_air_date.lte,omitempty" json:"first_air_date.lte,omitempty"`

	// Filter and only include TV shows that have a original air date year that equal to the specified value.
	// Can be used in conjunction with the "include_null_first_air_dates" filter if you want to include items with no air date.
	FirstAirDateYear *int `url:"first_air_date_year,omitempty" json:"first_air_date_year,omitempty"`

	// Specify a language to query translatable fields with.
	// default: en-US
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify the page of results to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// Used in conjunction with the air_date.gte/lte filter to calculate the proper UTC offset.
	// default: America/New_York
	Timezone string `url:"timezone,omitempty" json:"timezone,omitempty"`

	// Filter and only include movies that have a rating that is greater or equal to the specified value.
	VoteAverageGte *float64 `url:"vote_average.gte,omitempty" json:"vote_average.gte,omitempty"`

	// Filter and only include movies that have a rating that is less or equal to the specified value.
	VoteAverageLte *float64 `url:"vote_average.lte,omitempty" json:"vote_average.lte,omitempty"`

	// Filter and only include movies that have a vote count that is greater or equal to the specified value.
	VoteCountGte *int `url:"vote_count.gte,omitempty" json:"vote_count.gte,omitempty"`

	// Filter and only include movies that have a vote count that is less or equal to the specified value.
	VoteCountLte *int `url:"vote_count.lte,omitempty" json:"vote_count.lte,omitempty"`

	// Comma separated value of genre ids that you want to include in the results.
	WithGenres string `url:"with_genres,omitempty" json:"with_genres,omitempty"`

	// Comma separated value of network ids that you want to include in the results.
	WithNetworks string `url:"with_networks,omitempty" json:"with_networks,omitempty"`

	// Comma separated value of genre ids that you want to exclude from the results.
	WithoutGenres string `url:"without_genres,omitempty" json:"without_genres,omitempty"`

	// Filter and only include TV shows with an episode runtime that is greater than or equal to a value.
	WithRuntimeGte *int `url:"with_runtime.gte,omitempty" json:"with_runtime.gte,omitempty"`

	// Filter and only include TV shows with an episode runtime that is less than or equal to a value.
	WithRuntimeLte *int `url:"with_runtime.lte,omitempty" json:"with_runtime.lte,omitempty"`

	// Use this filter to include TV shows that don't have an air date while using any of the "first_air_date" filters.
	// default: false
	IncludeNullFirstAirDates bool `url:"include_null_first_air_dates,omitempty" json:"include_null_first_air_dates,omitempty"`

	// Specify an ISO 639-1 string to filter results by their original language value.
	WithOriginalLanguage string `url:"with_original_language,omitempty" json:"with_original_language,omitempty"`

	// Exclude items with certain keywords. You can comma and pipe separate these values to create an 'AND' or 'OR' logic.
	WithoutKeywords string `url:"without_keywords,omitempty" json:"without_keywords,omitempty"`

	// Filter results to include items that have been screened theatrically.
	ScreenedTheatrically *bool `url:"screened_theatrically,omitempty" json:"screened_theatrically,omitempty"`

	// A comma separated list of production company ID's. Only include movies that have one of the ID's added as a production company.
	WithCompanies string `url:"with_companies,omitempty" json:"with_companies,omitempty"`

	// A comma separated list of keyword ID's. Only includes TV shows that have one of the ID's added as a keyword.
	WithKeywords string `url:"with_keywords,omitempty" json:"with_keywords,omitempty"`

	// A comma or pipe separated list of watch provider ID's.
	// Combine this filter with `watch_region` in order to filter your results by a specific watch provider in a specific region.
	WithWatchProviders string `url:"with_watch_providers,omitempty" json:"with_watch_providers,omitempty"`

	// An ISO 3166-1 code.
	// Combine this filter with `with_watch_providers` in order to filter your results by a specific watch provider in a specific region.
	WatchRegion string `url:"watch_region,omitempty" json:"watch_region,omitempty"`

	// In combination with `watch_region`, you can filter by monetization type.
	// types: flatrate, free, ads, rent, buy
	WithWatchMonetizationTypes string `url:"with_watch_monetization_types,omitempty" json:"with_watch_monetization_types,omitempty"`

	// Filter TV shows by their status.
	// Returning Series: 0
	// Planned: 1
	// In Production: 2
	// Ended: 3
	// Cancelled: 4
	// Pilot: 5
	WithStatus string `url:"with_status,omitempty" json:"with_status,omitempty"`

	// Filter TV shows by their type.
	// Documentary: 0
	// News: 1
	// Miniseries: 2
	// Reality: 3
	// Scripted: 4
	// Talk Show: 5
	// Video: 6
	WithType string `url:"with_type,omitempty" json:"with_type,omitempty"`

	// Filter the results to exclude the specific production companies you specify here. `AND` / `OR` filters are supported.
	WithoutCompanies string `url:"without_companies,omitempty" json:"without_companies,omitempty"`
}

// Discover TV shows by different types of data like average rating, number of votes,
// genres, the network they aired on and air dates.
//
// Also note that a number of filters support being comma (,) or pipe (|) separated.
// Comma's are treated like an AND and query while pipe's are an OR.
//
// Some examples can be found here: https://www.themoviedb.org/documentation/api/discover
func (dr *DiscoverResource) DiscoverTVShows(opt *DiscoverTVShowsOptions) (*DiscoverTVShows, *http.Response, error) {
	path := "/discover/tv"
	var discover DiscoverTVShows
	resp, err := dr.client.get(path, &discover, WithQueryParams(opt))
	return &discover, resp, errors.Wrap(err, "failed to discover tv shows")
}
