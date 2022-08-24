package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// CompaniesResource handles company-related requests of TMDb API.
type CompaniesResource struct {
	client *Client
}

// Company represents a company in TMDb.
type Company struct {
	Name     string  `json:"name"`
	ID       int     `json:"id"`
	LogoPath *string `json:"logo_path"`
}

// CompanyDetails represents company details in TMDb.
type CompanyDetails struct {
	Description   string   `json:"description"`
	Headquarters  string   `json:"headquarters"`
	Homepage      string   `json:"homepage"`
	ID            int      `json:"id"`
	LogoPath      *string  `json:"logo_path"`
	Name          string   `json:"name"`
	OriginCountry string   `json:"origin_country"`
	ParentCompany *Company `json:"parent_company"`
}

// GetCompany retrieves company details by id.
func (cr *CompaniesResource) GetCompany(id int) (*CompanyDetails, *http.Response, error) {
	path := fmt.Sprintf("/company/%d", id)
	var company CompanyDetails
	resp, err := cr.client.get(path, &company)
	return &company, resp, errors.Wrap(err, "failed to get company")
}

// CompanyAlternativeName represents a company alternative name in TMDb.
type CompanyAlternativeName struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// CompanyAlternativeNames represents company alternative names in TMDb.
type CompanyAlternativeNames struct {
	ID    int                      `json:"id"`
	Names []CompanyAlternativeName `json:"results"`
}

// GetAlternativeNames retrieves the alternative names of a company.
func (cr *CompaniesResource) GetAlternativeNames(id int) (*CompanyAlternativeNames, *http.Response, error) {
	path := fmt.Sprintf("/company/%d/alternative_names", id)
	var names CompanyAlternativeNames
	resp, err := cr.client.get(path, &names)
	return &names, resp, errors.Wrap(err, "failed to get company alternative names")
}

// CompanyLogo represents a company logo in TMDb.
type CompanyLogo struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	ID          string  `json:"id"`
	FileType    string  `json:"file_type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

// CompanyImages represents company images in TMDb.
type CompanyImages struct {
	ID    int           `json:"id"`
	Logos []CompanyLogo `json:"logos"`
}

// GetImages retrieves company logos by id.
// There are two image formats that are supported for companies, PNG's and SVG's.
func (cr *CompaniesResource) GetImages(id int) (*CompanyImages, *http.Response, error) {
	path := fmt.Sprintf("/company/%d/images", id)
	var images CompanyImages
	resp, err := cr.client.get(path, &images)
	return &images, resp, errors.Wrap(err, "failed to get company images")
}
