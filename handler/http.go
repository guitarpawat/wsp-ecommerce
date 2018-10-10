package handler

import (
	"net"
	"net/http"
	"os"
)

func RedirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		host = r.Host
	}
	port := ":443"
	if os.Getenv("SolidTesting") == "true" {
		port = ":4433"
	}
	target := "https://" + host + port + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}

	http.Redirect(w, r, target, http.StatusPermanentRedirect)
}
