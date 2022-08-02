package tmdb

// PeopleResource handles person-related requests of TMDb API.
type PeopleResource struct {
	client *Client
}

type personKnownFor struct {
	person
	KnownFor []MovieOrTV `json:"known_for"`
}

type person struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	MediaType          string  `json:"media_type"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
}
