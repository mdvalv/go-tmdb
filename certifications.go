package tmdb

import (
	"net/http"

	"github.com/pkg/errors"
)

// CertificationsResource handles certification-related requests of TMDb API.
type CertificationsResource struct {
	client *Client
}

// Certification represents a certification in TMDb.
type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int    `json:"order"`
}

// MovieCertifications represents movie certifications in TMDb.
type MovieCertifications struct {
	AU []Certification `json:"AU"`
	BR []Certification `json:"BR"`
	CA []Certification `json:"CA"`
	DE []Certification `json:"DE"`
	FR []Certification `json:"FR"`
	GB []Certification `json:"GB"`
	IN []Certification `json:"IN"`
	NZ []Certification `json:"NZ"`
	US []Certification `json:"US"`
}

// TVCertifications represents tv certifications in TMDb.
type TVCertifications struct {
	AU []Certification `json:"AU"`
	BR []Certification `json:"BR"`
	CA []Certification `json:"CA"`
	DE []Certification `json:"DE"`
	FR []Certification `json:"FR"`
	GB []Certification `json:"GB"`
	KR []Certification `json:"KR"`
	RU []Certification `json:"RU"`
	TH []Certification `json:"TH"`
	US []Certification `json:"US"`
}

// MovieCertificationsResponse represents the response for getting movie certifications from TMDb.
type MovieCertificationsResponse struct {
	Certifications MovieCertifications `json:"certifications"`
}

// TVCertificationsResponse represents the response for getting tv certifications from TMDb.
type TVCertificationsResponse struct {
	Certifications TVCertifications `json:"certifications"`
}

// GetMovieCertifications gets an up to date list of the officially supported movie certifications on TMDB.
func (cr *CertificationsResource) GetMovieCertifications() (*MovieCertificationsResponse, *http.Response, error) {
	path := "/certification/movie/list"
	var certifications MovieCertificationsResponse
	resp, err := cr.client.get(path, &certifications)
	return &certifications, resp, errors.Wrap(err, "failed to get movie certifications")
}

// GetTVCertifications gets an up to date list of the officially supported TV show certifications on TMDB.
func (cr *CertificationsResource) GetTVCertifications() (*TVCertificationsResponse, *http.Response, error) {
	path := "/certification/tv/list"
	var certifications TVCertificationsResponse
	resp, err := cr.client.get(path, &certifications)
	return &certifications, resp, errors.Wrap(err, "failed to get tv certifications")
}
