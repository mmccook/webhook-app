package service

import (
	"DJMIL/entity"
	"context"
	"net/http"
)

type WebhookService struct {
	BaseService
}

func (service *WebhookService) CreateWebhook(sectionId string, originUrl string, header http.Header, body map[string]interface{}) (entity.Webhook, error) {
	webhook := entity.NewWebhook(sectionId, originUrl, header, body)
	redisCtx := context.Background()

	rcmd := service.DB.JSONSet(redisCtx, webhook.RedisKey(), "$", webhook)

	_, err := rcmd.Result()
	if err != nil {
		service.Log.Log().Timestamp().Msg(err.Error())
		return entity.Webhook{}, err
	}

	return webhook, nil
}

func (service *WebhookService) CreateIndex() error {
	redisCtx := context.Background()
	if !service.IndexExists("webhookIdx") {
		rcmd := service.DB.Do(redisCtx, "FT.CREATE", "webhookIdx", "ON", "JSON", "PREFIX", "1", "webhook:", "SCHEMA", "$.id as id", "NUMERIC", "$.createdAt as createdAt", "NUMERIC", "$.originUrl as originUrl", "TEXT", "$.sectionId as sectionId", "TEXT")
		_, err := rcmd.Result()
		if err != nil {
			service.Log.Log().Timestamp().Err(err).Msg(err.Error())
			return err
		}
		return err
	}
	return nil
}
