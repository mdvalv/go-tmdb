package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// NetworksResource handles network-related requests of TMDb API.
type NetworksResource struct {
	client *Client
}

// Network represents a network in TMDb.
type Network struct {
	Headquarters  string  `json:"headquarters"`
	Homepage      string  `json:"homepage"`
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

// GetNetwork retrieves network details by id.
func (nr *NetworksResource) GetNetwork(id int) (*Network, *http.Response, error) {
	path := fmt.Sprintf("/network/%d", id)
	var network Network
	resp, err := nr.client.get(path, &network)
	return &network, resp, errors.Wrap(err, "failed to get network")
}

// NetworkAlternativeName represents a network alternative name in TMDb.
type NetworkAlternativeName struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// NetworkAlternativeNames represents network alternative names in TMDb.
type NetworkAlternativeNames struct {
	ID    int                      `json:"id"`
	Names []NetworkAlternativeName `json:"results"`
}

// GetAlternativeNames retrieves the alternative names of a network.
func (nr *NetworksResource) GetAlternativeNames(id int) (*NetworkAlternativeNames, *http.Response, error) {
	path := fmt.Sprintf("/network/%d/alternative_names", id)
	var names NetworkAlternativeNames
	resp, err := nr.client.get(path, &names)
	return &names, resp, errors.Wrap(err, "failed to get network alternative names")
}

// NetworkImages represents network images in TMDb.
type NetworkImages struct {
	ID    int           `json:"id"`
	Logos []CompanyLogo `json:"logos"`
}

// GetImages retrieves network logos by id.
// There are two image formats that are supported for networks, PNG's and SVG's.
func (nr *NetworksResource) GetImages(id int) (*NetworkImages, *http.Response, error) {
	path := fmt.Sprintf("/network/%d/images", id)
	var images NetworkImages
	resp, err := nr.client.get(path, &images)
	return &images, resp, errors.Wrap(err, "failed to get network images")
}
