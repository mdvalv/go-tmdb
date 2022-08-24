package tmdb

// ChangesOptions represents the available options for the request.
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

// ChangeItem represents a change item in TMDb.
type ChangeItem struct {
	ID            string      `json:"id"`
	Action        string      `json:"action"`
	Time          string      `json:"time"`
	ISO6391       string      `json:"iso_639_1"`
	ISO31661      string      `json:"iso_3166_1"`
	OriginalValue interface{} `json:"original_value"`
	Value         interface{} `json:"value"`
}

// Change represents a change in TMDb.
type Change struct {
	Key   string       `json:"key"`
	Items []ChangeItem `json:"items"`
}

// Changes represents changes in TMDb.
type Changes struct {
	Changes []Change `json:"changes"`
}

// MediaChange represents a media change in TMDb.
type MediaChange struct {
	Adult *bool `json:"adult"`
	ID    int   `json:"id"`
}

// MediaChanges represents media changes in TMDb.
type MediaChanges struct {
	pagination
	Changes []MediaChange `json:"results"`
}
