package types

import (
	"time"
)

type (
	// Store holds info about reviewers
	Store interface {
		GetReviewers(from, to time.Time) ([]PullRequestReviewer, error)
		IncrementPullRequestOpened(author string) error
		IncrementPullRequestComment(author string) error
		IncrementPullRequestClosed(author string) error
		IncrementPullRequestReviewed(author string) error
	}
)
