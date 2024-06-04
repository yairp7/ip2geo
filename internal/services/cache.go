package services

import (
	"context"
	"errors"
)

type InMemCache struct {
	cache map[string]any
}

func NewInMemCache() *InMemCache {
	return &InMemCache{
		cache: make(map[string]any),
	}
}

func (c *InMemCache) Set(ctx context.Context, key string, value any) error {
	c.cache[key] = value
	return nil
}

func (c *InMemCache) Get(ctx context.Context, key string) (any, error) {
	if v, ok := c.cache[key]; ok {
		return v, nil
	}
	return nil, errors.New("no such key")
}
