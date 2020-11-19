package types

type (
	// Store holds info about reviewers
	Store interface {
		GetReviewers() ([]PullRequestReviewer, error)
		IncrementPullRequestOpened(author string) error
		IncrementPullRequestComment(author string) error
		IncrementPullRequestClosed(author string) error
		IncrementPullRequestApproved(author string) error
		ResetCounters() error
	}
)
