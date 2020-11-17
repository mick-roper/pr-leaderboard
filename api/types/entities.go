package types

type (
	PullRequestReviewer struct {
		AuthorName            string
		AuthorImageURL        string
		PullRequestsOpened    int
		PullRequestsCommented int
		PullRequestsClosed    int
		PullRequestsReviewed  int
	}
)
