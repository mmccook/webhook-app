package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type BaseService struct {
	DB  *redis.Client
	Log *zerolog.Logger
}

type IBaseService interface {
	IndexExists() bool
	CreateIndex() error
}

func (service *BaseService) IndexExists(indexName string) bool {
	redisCtx := context.Background()
	rcmd := service.DB.Do(redisCtx, "FT.INFO", indexName)
	_, err := rcmd.Result()
	if err != nil {
		return false
	}
	return true
}
