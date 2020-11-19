package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/mick-roper/pr-leaderboard/api/types"
)

type RedisStore struct {
	client *redis.Client
	logger *log.Logger
	ctx    context.Context
}

const (
	opened    = "opened"
	closed    = "closed"
	comment   = "comment"
	reviewed  = "reviewed"
	separator = "/"
)

func NewRedisStore(addr, password string) (*RedisStore, error) {
	if addr == "" {
		return nil, errors.New("address is empty")
	}

	logger := log.New(log.Writer(), "REDIS ", log.LUTC|log.Lmsgprefix)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	logger.Println(pong, err)

	if err != nil {
		return nil, err
	}

	return &RedisStore{
		client: client,
		ctx:    context.Background(),
		logger: logger,
	}, nil
}

func (s *RedisStore) GetReviewers() ([]types.PullRequestReviewer, error) {
	keys, err := s.getAllKeys()

	if err != nil {
		return nil, err
	}

	type aggregate struct {
		opened   int
		closed   int
		comments int
		reviewed int
	}

	m := map[string]*aggregate{}

	for _, key := range keys {
		val, err := s.client.Get(key).Result()
		if err != nil {
			return nil, err
		}

		intVal, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		splits := strings.Split(key, separator)
		prEventType := splits[0]
		author := splits[1]

		var item *aggregate
		var exists bool

		if item, exists = m[author]; !exists {
			item = &aggregate{}
			m[author] = item
		}

		switch prEventType {
		case opened:
			{
				item.opened = intVal
			}
		case closed:
			{
				item.closed = intVal
			}
		case comment:
			{
				item.comments = intVal
			}
		case reviewed:
			{
				item.reviewed = intVal
			}
		}
	}

	i := 0
	items := make([]types.PullRequestReviewer, len(m))
	for key, value := range m {
		items[i].AuthorName = key
		items[i].PullRequestsOpened = value.opened
		items[i].PullRequestsCommented = value.comments
		items[i].PullRequestsClosed = value.closed
		items[i].PullRequestsReviewed = value.reviewed
		i++
	}

	return items, nil
}

func (s *RedisStore) IncrementPullRequestOpened(author string) error {
	key := fmt.Sprint(opened, separator, author)
	return s.increment(key)
}

func (s *RedisStore) IncrementPullRequestComment(author string) error {
	key := fmt.Sprint(comment, separator, author)
	return s.increment(key)
}

func (s *RedisStore) IncrementPullRequestClosed(author string) error {
	key := fmt.Sprint(closed, separator, author)
	return s.increment(key)
}

func (s *RedisStore) IncrementPullRequestApproved(author string) error {
	key := fmt.Sprint(reviewed, separator, author)
	return s.increment(key)
}

func (s *RedisStore) ResetCounters() error {
	keys, err := s.getAllKeys()

	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := s.client.Set(key, 0, 0).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (s *RedisStore) increment(key string) error {
	result, err := s.client.Incr(key).Result()
	s.logger.Println("INCREMENT", result)

	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStore) getAllKeys() ([]string, error) {
	keys := []string{}
	wildcards := []string{opened + "*", closed + "*", comment + "*", reviewed + "*"}
	for _, wildcard := range wildcards {
		if k, err := s.client.Keys(wildcard).Result(); err == nil {
			keys = append(keys, k...)
		} else {
			return nil, err
		}
	}
	return keys, nil
}
