package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mick-roper/pr-leaderboard/api/auth"
	"github.com/mick-roper/pr-leaderboard/api/types"
)

type (
	apiHandler struct {
		store    types.Store
		keyStore auth.APIKeyStore
	}
)

// ConfigureAPIRoutes for the server
func ConfigureAPIRoutes(mux *http.ServeMux, dataStore types.Store, apiKeyStore auth.APIKeyStore) error {
	if mux == nil {
		errors.New("mux is nil")
	}

	if dataStore == nil {
		errors.New("dataStore is nil")
	}

	handler := apiHandler{
		store:    dataStore,
		keyStore: apiKeyStore,
	}
	mux.Handle("/api", &handler)

	return nil
}

func (h *apiHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if !h.authorise(req) {
		res.WriteHeader(403)
		return
	}

	switch req.Method {
	case http.MethodGet:
		{
			from := time.Now()
			to := time.Now()
			items, err := h.store.GetReviewers(from, to)

			if err != nil {
				log.Print(err)
				res.WriteHeader(500)
				return
			}

			if err = json.NewEncoder(res).Encode(items); err != nil {
				log.Print(err)
			}
		}
	default:
		{
			res.WriteHeader(405)
		}
	}
}

func (h *apiHandler) authorise(req *http.Request) bool {
	requestAPIKey := req.Header.Get("x-api-key")
	apiKey := h.keyStore.GetAPIKey()

	return apiKey == requestAPIKey
}
