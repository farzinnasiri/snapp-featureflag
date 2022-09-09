package featureflag

import (
	"net/http"
)

func NewHttpServeMux() (*http.ServeMux, error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(Home))

	return mux, nil
}
