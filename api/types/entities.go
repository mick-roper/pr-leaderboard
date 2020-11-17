package types

type (
	PullRequestReviewer struct {
		AuthorName            string
		AuthorImageURL        string
		PullRequestsOpened    int
		PullRequestsCommented int
		PullRequestsReviewed  int
	}
)
