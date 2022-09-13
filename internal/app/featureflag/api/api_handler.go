package api

import (
	"context"
	"encoding/json"
	"net/http"
	"snapp-featureflag/internal/app/featureflag/feature"
	"strconv"
)

type ApiHandler struct {
	commandHandler feature.CommandHandler
	queryHandler   feature.QueryHandler
}

type CreateFeatureRequest struct {
	Name       string  `json:"name"`
	MinVersion string  `json:"min_version"`
	Coverage   float32 `json:"coverage"`
}

type UpdateFeatureRequest struct {
	Name       string  `json:"name"`
	MinVersion string  `json:"min_version"`
	Coverage   float32 `json:"coverage"`
}

type DeleteFeatureRequest struct {
	Name string `json:"name"`
}

type GetFeatureRequest struct {
	Name string `json:"name"`
}

type GetFeatureResponse struct {
	Name       string  `json:"name"`
	MinVersion string  `json:"min_version,omitempty"`
	Coverage   float32 `json:"coverage,omitempty"`
}

func NewApiHandler(commandHandler feature.CommandHandler, queryHandler feature.QueryHandler) *ApiHandler {
	return &ApiHandler{commandHandler, queryHandler}
}

func (a *ApiHandler) CreateFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	request := &CreateFeatureRequest{}
	err := decoder.Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	if err = a.commandHandler.Upsert(context.Background(), feature.UpsertFeatureCommand{
		Name:       request.Name,
		MinVersion: request.MinVersion,
		Coverage:   request.Coverage,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *ApiHandler) UpdateFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	request := &UpdateFeatureRequest{}
	err := decoder.Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	if err = a.commandHandler.Upsert(context.Background(), feature.UpsertFeatureCommand{
		Name:       request.Name,
		MinVersion: request.MinVersion,
		Coverage:   request.Coverage,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *ApiHandler) DeleteFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	request := &DeleteFeatureRequest{}
	err := decoder.Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	if err = a.commandHandler.Delete(context.Background(), feature.DeleteFeatureCommand{
		Name: request.Name,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *ApiHandler) GetFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	featureName := r.URL.Query().Get("name")

	queryResp, err := a.queryHandler.Get(context.Background(), feature.GetFeatureQuery{
		Name: featureName,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := GetFeatureResponse{
		Name:       queryResp.Name,
		MinVersion: queryResp.MinVersion,
		Coverage:   queryResp.Coverage,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

func (a *ApiHandler) GetUserActiveFeatures(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userId, err := strconv.ParseUint(
		r.URL.Query().Get("user_id"), 0, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	version := r.URL.Query().Get("version")

	queryResp, err := a.queryHandler.GetActiveFeaturesByUserId(context.Background(),
		feature.GetActiveFeaturesByUserIdQuery{
			UserId:  uint(userId),
			Version: version,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(queryResp.FeaturesNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}