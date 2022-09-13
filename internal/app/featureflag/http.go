package featureflag

import (
	"net/http"
)

func NewHttpServeMux(apiHandler *ApiHandler) (*http.ServeMux, error) {
	mux := http.NewServeMux()
	mux.Handle("/feature/create",
		http.HandlerFunc(apiHandler.CreateFeature))
	mux.Handle("/feature/update",
		http.HandlerFunc(apiHandler.UpdateFeature))
	mux.Handle("/feature/delete",
		http.HandlerFunc(apiHandler.DeleteFeature))
	mux.Handle("/feature/get",
		http.HandlerFunc(apiHandler.GetFeature))
	mux.Handle("/feature/get-active-features",
		http.HandlerFunc(apiHandler.GetUserActiveFeatures))

	return mux, nil
}
