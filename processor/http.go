package processor

import (
	"fmt"
	"net/http"

	d "github.com/aerialls/meereen/pkg/data"
	p "github.com/aerialls/meereen/pkg/processor"
)

// HTTP processor
// Check the status code for an HTTP page (2xx or 3xx)
type HTTP struct {
	URL string
}

// NewHTTP validates and returns a new HTTP processor
func NewHTTP(data map[string]string) (p.Processor, error) {
	url, err := d.GetRequiredParameter(data, "url")
	if err != nil {
		return nil, err
	}

	return &HTTP{
		URL: url,
	}, nil
}

// Process the HTTP processor
func (h *HTTP) Process() (p.State, string) {
	res, err := http.Get(h.URL)
	if err != nil {
		return p.Error, fmt.Sprintf("Unable to fetch the remote URL (%s)", err)
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return p.Ok, ""
	}

	return p.Error, fmt.Sprintf("Unexpected status code (%d)", res.StatusCode)
}

func init() {
	p.RegisterProcessor("http", NewHTTP)
}
