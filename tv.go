package tmdb

// TVResource handles tv-related requests of TMDb API.
type TVResource struct {
	client *Client
}

type tv struct {
	PosterPath       *string  `json:"poster_path"`
	Popularity       float64  `json:"popularity"`
	Id               int      `json:"id"`
	BackdropPath     *string  `json:"backdrop_path"`
	VoteAverage      float64  `json:"vote_average"`
	Overview         string   `json:"overview"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	GenreIds         []int    `json:"genre_ids"`
	OriginalLanguage string   `json:"original_language"`
	VoteCount        int      `json:"vote_count"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
	Adult            bool     `json:"adult"`
}
