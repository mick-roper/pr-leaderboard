package db

import (
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mick-roper/pr-leaderboard/api/types"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string) (*RedisStore, error) {
	if addr == "" {
		return nil, errors.New("address is empty")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)

	if err != nil {
		return nil, err
	}

	return &RedisStore{client}, nil
}

func (s *RedisStore) GetReviewers(from, to time.Time) ([]types.PullRequestReviewer, error) {
	return nil, errors.New("not implemented")
}

func (s *RedisStore) IncrementPullRequestOpened(author string) error {
	return errors.New("not implemented")
}

func (s *RedisStore) IncrementPullRequestComment(author string) error {
	return errors.New("not implemented")
}

func (s *RedisStore) IncrementPullRequestClosed(author string) error {
	return errors.New("not implemented")
}

func (s *RedisStore) IncrementPullRequestApproved(author string) error {
	return errors.New("not implemented")
}
