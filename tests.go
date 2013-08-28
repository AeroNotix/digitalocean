package digitalocean

import (
	"net/http"
	"net/http/httptest"
)

func httpTestsSetUp(f http.HandlerFunc) {
	Endpoint = ts.URL + "/droplets/%s?client_id=%s&api_key=%s"
	if f != nil {
		functionalTest = f
	}
}

var th = testHandler{}
var ts = httptest.NewServer(th)
var functionalTest http.HandlerFunc

type testHandler struct {
	Status         bool
	Message        string
	FunctionalTest http.HandlerFunc
}

func (th testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	functionalTest(w, req)
}
