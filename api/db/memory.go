package db

import (
	"time"

	"github.com/mick-roper/pr-leaderboard/api/types"
)

type (
	// MemoryStore stores data in memory
	MemoryStore struct {
		entries map[string]aggregate
	}

	aggregate struct {
		opened   int
		comments int
		closed   int
		reviewed int
	}
)

// NewMemoryStore creates a new memory store instance
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		entries: make(map[string]aggregate),
	}
}

func (s *MemoryStore) GetReviewers(from, to time.Time) ([]types.PullRequestReviewer, error) {
	// ignore times
	array := make([]types.PullRequestReviewer, len(s.entries))
	i := 0

	for key, entry := range s.entries {
		array[i].AuthorName = key
		array[i].PullRequestsOpened = entry.opened
		array[i].PullRequestsCommented = entry.comments
		array[i].PullRequestsClosed = entry.closed
		array[i].PullRequestsReviewed = entry.reviewed

		i++
	}

	return array, nil
}

func (s *MemoryStore) IncrementPullRequestOpened(author string) error {
	if value, exists := s.entries[author]; exists {
		value.opened++
	} else {
		s.entries[author] = aggregate{
			opened: 1,
		}
	}

	return nil
}

func (s *MemoryStore) IncrementPullRequestComment(author string) error {
	if value, exists := s.entries[author]; exists {
		value.comments++
	} else {
		s.entries[author] = aggregate{
			comments: 1,
		}
	}

	return nil
}

func (s *MemoryStore) IncrementPullRequestClosed(author string) error {
	if value, exists := s.entries[author]; exists {
		value.closed++
	} else {
		s.entries[author] = aggregate{
			closed: 1,
		}
	}

	return nil
}

func (s *MemoryStore) IncrementPullRequestReviewed(author string) error {
	if value, exists := s.entries[author]; exists {
		value.reviewed++
	} else {
		s.entries[author] = aggregate{
			reviewed: 1,
		}
	}

	return nil
}
