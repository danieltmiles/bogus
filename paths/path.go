package paths

import (
	"net/http"
	"strings"
)

// Path represents an endpoint added to a bogus server and how it should respond
type Path struct {
	Hits    int
	payload []byte
	status  int
	methods []string
}

// New returns a newly instantiated path object with everything initialized as
// needed.
func New() *Path {
	return &Path{}
}

// SetPayload sets the response payload for the path and returns the path for
// additional configuration
func (p *Path) SetPayload(payload []byte) *Path {
	p.payload = payload
	return p
}

// SetStatus sets the http status for the path and returns the path for
// additional configuration
func (p *Path) SetStatus(status int) *Path {
	p.status = status
	return p
}

// SetMethods accepts a list of methods the path should respond to
func (p *Path) SetMethods(methods ...string) *Path {
	for i, m := range methods {
		methods[i] = strings.ToUpper(m)
	}

	p.methods = methods
	return p
}

func (p *Path) HandleRequest(w http.ResponseWriter, r *http.Request) {
	payload := []byte("")
	status := http.StatusForbidden

	if p.hasMethod(r.Method) {
		p.Hits++
		w.WriteHeader(p.status)
		w.Write(p.payload)
		return
	}

	w.WriteHeader(status)
	w.Write(payload)
	return
}

func (p *Path) hasMethod(method string) bool {
	method = strings.ToUpper(method)

	// Check for configured methods first
	if len(p.methods) != 0 {
		for _, m := range p.methods {
			if m == method {
				return true
			}
		}

		return false
	}

	// If no methods configured, handle GET by default for convenience
	if method == "GET" {
		return true
	}

	return false
}
