package api

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"snapp-featureflag/internal/app/featureflag/feature"
	"strconv"
)

type Handler struct {
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

func NewApiHandler(commandHandler feature.CommandHandler, queryHandler feature.QueryHandler) *Handler {
	return &Handler{commandHandler, queryHandler}
}

func (a *Handler) CreateFeature(w http.ResponseWriter, r *http.Request) {
	log.Println("Call to CreateFeature")
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

	if err = a.commandHandler.Create(context.Background(), feature.CreateFeatureCommand{
		Name:       request.Name,
		MinVersion: request.MinVersion,
		Coverage:   request.Coverage,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *Handler) UpdateFeature(w http.ResponseWriter, r *http.Request) {
	log.Println("Call to UpdateFeature")
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

	if err = a.commandHandler.Update(context.Background(), feature.UpdateFeatureCommand{
		Name:       request.Name,
		MinVersion: request.MinVersion,
		Coverage:   request.Coverage,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *Handler) DeleteFeature(w http.ResponseWriter, r *http.Request) {
	log.Println("Call to DeleteFeature")
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

func (a *Handler) GetFeature(w http.ResponseWriter, r *http.Request) {
	log.Println("Call to GetFeature")
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

func (a *Handler) GetUserActiveFeatures(w http.ResponseWriter, r *http.Request) {
	log.Println("Call to GetUserActiveFeatures")
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
