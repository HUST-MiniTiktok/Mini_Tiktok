package redis

import (
	redis "github.com/go-redis/redis/v7"
)

type RDClient struct {
	Client *redis.Client
}