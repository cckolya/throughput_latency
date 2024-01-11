package main

import (
	"devhands/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	cli *redis.Client
}

func NewRedis(cfg config.Redis) (*Redis, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + cfg.Port,
	})
	return &Redis{cli}, nil
}
