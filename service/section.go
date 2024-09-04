package service

import (
	"DJMIL/entity"
	"context"
)

type SectionService struct {
	BaseService
}

func (service *SectionService) CreateSection() (entity.Section, error) {
	section := entity.NewSection()
	redisCtx := context.Background()

	rcmd := service.DB.JSONSet(redisCtx, section.RedisKey(), "$", section)

	_, err := rcmd.Result()
	if err != nil {
		service.Log.Log().Timestamp().Msg(err.Error())
		return entity.Section{}, err
	}

	return section, nil
}

var IndexExists string = "Index already exists"

func (service *SectionService) CreateIndex() error {
	redisCtx := context.Background()

	if !service.IndexExists("sectionIdx") {
		rcmd := service.DB.Do(redisCtx, "FT.CREATE", "sectionIdx", "ON", "JSON", "PREFIX", "1", "section:", "SCHEMA", "$.id as id", "NUMERIC", "$.createdAt as createdAt", "NUMERIC", "$.lastActive as lastActive", "NUMERIC")
		_, err := rcmd.Result()
		if err != nil {
			service.Log.Log().Timestamp().Err(err).Msg(err.Error())
			return err
		}
		return err
	}
	return nil
}
