package processor

import (
	"fmt"
	"net"

	"github.com/miekg/dns"

	d "github.com/aerialls/meereen/pkg/data"
	p "github.com/aerialls/meereen/pkg/processor"
)

// DNS processor
// Check that the domain can be resolved with a specified resolver
type DNS struct {
	domain   string
	resolver string
}

// NewDNS validates and returns a new DNS processor
func NewDNS(data map[string]string) (p.Processor, error) {
	domain, err := d.GetRequiredParameter(data, "domain")
	if err != nil {
		return nil, err
	}

	resolver := d.GetParameter(data, "resolver", "")

	return &DNS{
		domain:   domain,
		resolver: resolver,
	}, nil
}

// Process the HTTP processor
func (d *DNS) Process() (p.State, string) {
	if d.resolver == "" {
		// Resolver will be the one configured on the server
		_, err := net.LookupIP(d.domain)
		if err != nil {
			return p.Error, fmt.Sprintf("Unable to resolve %s (%s)", d.domain, err)
		}

		return p.Ok, ""
	}

	// Resolution with a dedicated resolver
	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(d.domain+".", dns.TypeA)
	r, _, err := c.Exchange(&m, d.resolver+":53")
	if err != nil {
		return p.Error, fmt.Sprintf("Unable to resolve %s (%s)", d.domain, err)
	}

	if len(r.Answer) == 0 {
		return p.Error, fmt.Sprintf("No results for resolving %s", d.domain)
	}

	return p.Ok, ""
}

func init() {
	p.RegisterProcessor("dns", NewDNS)
}
