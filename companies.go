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

type company struct {
	Name     string  `json:"name"`
	Id       int     `json:"id"`
	LogoPath *string `json:"logo_path"`
}

type Company struct {
	Description   string   `json:"description"`
	Headquarters  string   `json:"headquarters"`
	Homepage      string   `json:"homepage"`
	Id            int      `json:"id"`
	LogoPath      *string  `json:"logo_path"`
	Name          string   `json:"name"`
	OriginCountry string   `json:"origin_country"`
	ParentCompany *company `json:"parent_company"`
}

// Get company details by id.
func (cr *CompaniesResource) GetCompany(id int) (*Company, *http.Response, error) {
	path := fmt.Sprintf("/company/%d", id)
	var company Company
	resp, err := cr.client.get(path, &company)
	return &company, resp, errors.Wrap(err, "failed to get company")
}

type CompanyAlternativeName struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CompanyAlternativeNames struct {
	Id    int                      `json:"id"`
	Names []CompanyAlternativeName `json:"results"`
}

// Get the alternative names of a company.
func (cr *CompaniesResource) GetAlternativeNames(id int) (*CompanyAlternativeNames, *http.Response, error) {
	path := fmt.Sprintf("/company/%d/alternative_names", id)
	var names CompanyAlternativeNames
	resp, err := cr.client.get(path, &names)
	return &names, resp, errors.Wrap(err, "failed to get company alternative names")
}

type CompanyLogo struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	Id          string  `json:"id"`
	FileType    string  `json:"file_type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

type CompanyImages struct {
	Id    int           `json:"id"`
	Logos []CompanyLogo `json:"logos"`
}

// Get company logos by id.
// There are two image formats that are supported for companies, PNG's and SVG's.
func (cr *CompaniesResource) GetImages(id int) (*CompanyImages, *http.Response, error) {
	path := fmt.Sprintf("/company/%d/images", id)
	var images CompanyImages
	resp, err := cr.client.get(path, &images)
	return &images, resp, errors.Wrap(err, "failed to get company images")
}
