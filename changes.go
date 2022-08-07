package tmdb

type ChangesOptions struct {
	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// Filter the results with a start date.
	// format: YYYY-MM-DD
	StartDate string `url:"start_date,omitempty" json:"start_date,omitempty"`

	// Filter the results with an end date.
	// format: YYYY-MM-DD
	EndDate string `url:"end_date,omitempty" json:"end_date,omitempty"`
}

type changeItem struct {
	Id            string      `json:"id"`
	Action        string      `json:"action"`
	Time          string      `json:"time"`
	ISO6391       string      `json:"iso_639_1"`
	ISO31661      string      `json:"iso_3166_1"`
	OriginalValue interface{} `json:"original_value"`
	Value         interface{} `json:"value"`
}

type Change struct {
	Key   string       `json:"key"`
	Items []changeItem `json:"items"`
}

type Changes struct {
	Changes []Change `json:"changes"`
}

type mediaChange struct {
	Adult *bool `json:"adult"`
	Id    int   `json:"id"`
}

type MediaChanges struct {
	pagination
	Changes []mediaChange `json:"results"`
}
