package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

/**
	A reverse proxy is a type of server that sits between client requests and backend services, f
	orwarding client requests to the appropriate backend service and then returning the service's 
	response to the client.
*/


func NewReverseProxy(target string) http.Handler {
	url, err := url.Parse(target) //url.parse("http://localhost:4000")
	if err != nil {
		panic("Invalid target URL: " + target)
	}
	//This line uses Go's built-in httputil package to create a reverse proxy 
	// that forwards incoming HTTP requests to the target backend.
	return httputil.NewSingleHostReverseProxy(url)
}
