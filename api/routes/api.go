package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
func ConfigureAPIRoutes(router *mux.Router, dataStore types.Store, apiKeyStore auth.APIKeyStore) error {
	if router == nil {
		errors.New("mux is nil")
	}

	if dataStore == nil {
		errors.New("dataStore is nil")
	}

	handler := apiHandler{
		store:    dataStore,
		keyStore: apiKeyStore,
	}
	router.Handle("/api", &handler)

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
			items, err := h.store.GetReviewers()

			if err != nil {
				log.Print(err)
				res.WriteHeader(500)
				return
			}

			res.Header().Set("content-type", "application/json")

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
