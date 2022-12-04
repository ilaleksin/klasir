package handlers

import (
	"encoding/json"
	"fmt"
	"kvasir/pkg/models"
	"kvasir/pkg/repository"
	"net/http"
	"strconv"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

type Manager struct {
	repo repository.Repository
}

func NewManager(repo repository.Repository) *Manager {
	return &Manager{
		repo: repo,
	}
}

func (api *Manager) AddWord(w http.ResponseWriter, r *http.Request) {
	var row models.DictionaryRow
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&row)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = api.repo.CreateDictonaryRow(ctx, row)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *Manager) GetDictionary(w http.ResponseWriter, r *http.Request) {
	var (
		limit  = defaultLimit
		offset = defaultOffset
		err    error
	)

	ctx := r.Context()
	if limitStr := r.FormValue("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			err := fmt.Errorf("limit argument should be integer, received %s", limitStr)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if offsetStr := r.FormValue("offset"); offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			err := fmt.Errorf("limit argument should be integer, received %s", offsetStr)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	records, err := api.repo.GetDictionary(ctx, limit, offset)
	if err != nil {
		err := fmt.Errorf("failed to get dictionary from %d to %d rows", limit, offset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(records)
	if err != nil {
		err := fmt.Errorf("failed to convert response in json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *Manager) MakeReview(w http.ResponseWriter, r *http.Request) {
	var filter models.ReviewOptions
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		err = fmt.Errorf("failed to parse review options")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
