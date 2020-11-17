package types

type (
	// GithubWebhookEvent that comes from Github
	GithubWebhookEvent struct {
		Action      string            `json:"action"`
		Repository  GithubRepository  `json:"repository"`
		PullRequest GithubPullRequest `json:"pull_request"`
		Sender      GithubUser        `json:"sender"`
	}

	// GithubUser represents the a user
	GithubUser struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
		Type      string `json:"type"`
	}

	// GithubRepository represent repo info
	GithubRepository struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	// GithubPullRequest represents a pull request
	GithubPullRequest struct {
		User  GithubUser `json:"user"`
		URL   string     `json:"url"`
		Title string     `json:"title"`
	}
)
