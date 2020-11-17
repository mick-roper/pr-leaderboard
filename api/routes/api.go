package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mick-roper/pr-leaderboard/api/types"
)

type (
	apiHandler struct {
		store types.Store
	}
)

// ConfigureAPIRoutes for the server
func ConfigureAPIRoutes(mux *http.ServeMux, store types.Store) error {
	if mux == nil {
		errors.New("mux is nil")
	}

	if store == nil {
		errors.New("store is nil")
	}

	handler := apiHandler{store}
	mux.Handle("/api", &handler)

	return nil
}

func (h *apiHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		{
			from := time.Now()
			to := time.Now()
			items, err := h.store.GetReviewers(from, to)

			if err != nil {
				log.Print(err)
				log.Print(err)
				return
			}

			if err = json.NewEncoder(res).Encode(items); err != nil {
				res.WriteHeader(204)
			} else {
				log.Print(err)
				res.WriteHeader(500)
			}
		}
	default:
		{
			res.WriteHeader(405)
		}
	}
}
