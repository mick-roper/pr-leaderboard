package routes

import (
	"encoding/json"
	"errors"
	"log"
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
	// githubEvent := req.Header.Get("X-GitHub-Event")
	// githubSecret := req.Header.Get("X-Hub-Signature")

	switch req.Method {
	case http.MethodPost:
		{
			event := types.GithubWebhookEvent{}
			decoder := json.NewDecoder(req.Body)
			if err := decoder.Decode(&event); err != nil {
				log.Print(err)
				res.WriteHeader(500)
				return
			}

			switch event.Action {
			case "pull_request_created":
				{
					h.store.IncrementPullRequestOpened(event.Sender.Login)
				}
			case "pull_request_comment":
				{
					h.store.IncrementPullRequestComment(event.Sender.Login)
				}
			case "pull_request_approved":
				{
					h.store.IncrementPullRequestApproved(event.Sender.Login)
				}
			case "pull_request_closed":
				{
					h.store.IncrementPullRequestClosed(event.Sender.Login)
				}
			}

			res.WriteHeader(204)
		}
	default:
		{
			res.WriteHeader(405)
		}
	}
}
