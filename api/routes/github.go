package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mick-roper/pr-leaderboard/api/types"
)

const (
	eventGithubPullRequest              = "pull_request"
	eventGithubPullRequestReview        = "pull_request_review"
	eventGithubPullRequestReviewComment = "pull_request_review_comment"
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
	githubEvent := req.Header.Get("X-GitHub-Event")
	// githubSecret := req.Header.Get("X-Hub-Signature")

	defer req.Body.Close()

	if req.Method != http.MethodPost {
		res.WriteHeader(405)
		return
	}

	if !(githubEvent == eventGithubPullRequest || githubEvent == eventGithubPullRequestReview || githubEvent == eventGithubPullRequestReviewComment) {
		// short circuit the code - we don't care about this event
		log.Printf("Unsupported event received: '%v'", githubEvent)
		res.WriteHeader(204)
		return
	}

	event := types.GithubWebhookEvent{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&event); err != nil {
		log.Print(err)
		res.WriteHeader(500)
		return
	}

	switch event.Action {
	case "opened":
		{
			h.store.IncrementPullRequestOpened(event.Sender.Login)
		}
	case "closed":
		{
			h.store.IncrementPullRequestComment(event.Sender.Login)
		}
	case "submitted":
		{
			h.store.IncrementPullRequestApproved(event.Sender.Login)
		}
	case "pull_request_closed":
		{
			h.store.IncrementPullRequestClosed(event.Sender.Login)
		}
	default:
		{
			log.Printf("unsupported event recieved: '%v'", event.Action)
		}
	}

	res.WriteHeader(204)
}
