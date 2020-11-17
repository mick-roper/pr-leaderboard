package types

type (
	// GithubWebhookEvent that comes from Github
	GithubWebhookEvent struct {
		Action     string           `json:"action"`
		Sender     GithubSender     `json:"sender"`
		Repository GithubRepository `json:"repository"`
	}

	// GithubSender represents the sender of an event
	GithubSender struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
		Type      string `json:"type"`
	}

	// GithubRepository represent repo info
	GithubRepository struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
)
