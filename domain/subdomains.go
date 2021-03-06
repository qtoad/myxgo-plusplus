package domain

import (
	"net/http"
	"strings"
)

type Subdomains map[string]http.Handler

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")

	if mux := subdomains[domainParts[0]]; mux != nil {
		mux.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not found", 404)
	}
}
