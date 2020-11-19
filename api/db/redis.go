package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/mick-roper/pr-leaderboard/api/types"
)

type RedisStore struct {
	client *redis.Client
}

const (
	opened   = "opened/"
	closed   = "closed/"
	comment  = "comment/"
	approved = "approved/"
)

var ctx = context.Background()

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

func (s *RedisStore) GetReviewers() ([]types.PullRequestReviewer, error) {
	keys, err := s.getAllKeys()

	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		log.Println(key)
	}

	return nil, errors.New("not implemented")
}

func (s *RedisStore) IncrementPullRequestOpened(author string) error {
	key := fmt.Sprint(opened, author)
	return s.increment(key)
}

func (s *RedisStore) IncrementPullRequestComment(author string) error {
	key := fmt.Sprint(comment, author)
	return s.increment(key)
}

func (s *RedisStore) IncrementPullRequestClosed(author string) error {
	key := fmt.Sprint(closed, author)
	return s.increment(key)
}

func (s *RedisStore) IncrementPullRequestApproved(author string) error {
	key := fmt.Sprint(approved, author)
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
	log.Println(result)

	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStore) getAllKeys() ([]string, error) {
	keys := []string{}
	wildcards := []string{opened + "*", closed + "*", comment + "*", approved + "*"}
	for _, wildcard := range wildcards {
		if k, err := s.client.Keys(wildcard).Result(); err == nil {
			keys = append(keys, k...)
		} else {
			return nil, err
		}
	}
	return keys, nil
}
