package mounts

import "github.com/go-redis/redis"

type ILimiter interface {
	Allow(method string, mobile string ,) bool
	Remain() int
}

type redisLimiter struct {
	source *redis.Client
}


