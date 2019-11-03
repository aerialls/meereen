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
	url string
}

// NewHTTP validates and returns a new HTTP processor
func NewHTTP(data map[string]string) (p.Processor, error) {
	url, err := d.GetRequiredParameter(data, "url")
	if err != nil {
		return nil, err
	}

	return &HTTP{
		url: url,
	}, nil
}

// Process the HTTP processor
func (h *HTTP) Process() (p.State, string) {
	res, err := http.Get(h.url)

	if err != nil {
		return p.Error, fmt.Sprintf("Unable to fetch the remote URL (%s)", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return p.Ok, ""
	}

	return p.Error, fmt.Sprintf("Unexpected status code (%d)", res.StatusCode)
}

func init() {
	p.RegisterProcessor("http", NewHTTP)
}
