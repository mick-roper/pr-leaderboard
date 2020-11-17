package routes

import (
	"errors"
	"net/http"

	"github.com/mick-roper/pr-leaderboard/api/types"
)

type (
	githubHandler struct {
		store types.Store
	}
)

// ConfigureGithubRoutes for the server
func ConfigureGithubRoutes(mux *http.ServeMux, store types.Store) error {
	if mux == nil {
		errors.New("mux is nil")
	}

	if store == nil {
		errors.New("store is nil")
	}

	handler := githubHandler{store}
	mux.Handle("/github", &handler)

	return nil
}

func (h *githubHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		{
			res.WriteHeader(501)
			res.Write([]byte("Not Implemented"))
		}
	default:
		{
			res.WriteHeader(405)
		}
	}
}
