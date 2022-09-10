package featureflag

import (
	"net/http"
)

func NewHttpServeMux(apiHandler *ApiHandler) (*http.ServeMux, error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(apiHandler.Home))

	return mux, nil
}
