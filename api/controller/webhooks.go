package controller

import (
	"DJMIL/config"
	"DJMIL/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Webhooks struct {
	Config *config.ApiConfig
}

func (controller *Webhooks) Route(sectionGroup *echo.Group) {
	sectionGroup.GET("/:sectionId/webhooks", getWebhookListHandler(controller.Config))
	sectionGroup.POST("/:sectionId/webhooks", postWebhookHandler(controller.Config))
}

func getWebhookListHandler(config *config.ApiConfig) func(c echo.Context) error {
	return func(c echo.Context) error {
		var sectionId, err = strconv.ParseUint(c.Param("sectionId"), 10, 64)
		if err != nil {
			return err
		}

		sectionIdStr := strconv.FormatUint(sectionId, 10)

		redisKey := fmt.Sprintf("%s::webhooks", sectionIdStr)

		redisCtx := context.Background()

		rcmd := config.DB.JSONGet(redisCtx, redisKey, "$")

		result, err := rcmd.Result()
		if err != nil {
			config.Log.Log().Timestamp().Msg(err.Error())
			return err
		}

		var data []entity.Webhook
		err = json.Unmarshal([]byte(result), &data)
		if err != nil {
			config.Log.Log().Timestamp().Msg(err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(200, data)
	}
}
func postWebhookHandler(config *config.ApiConfig) func(c echo.Context) error {
	return func(c echo.Context) error {
		var sectionId, err = strconv.ParseUint(c.Param("sectionId"), 10, 64)
		if err != nil {
			return err
		}
		sectionIdStr := strconv.FormatUint(sectionId, 10)

		redisKey := fmt.Sprintf("%s::webhooks", sectionIdStr)

		var headers []entity.HttpHeader
		for name, values := range c.Request().Header {
			// Loop over all values for the name.
			for _, value := range values {
				headers = append(headers, entity.HttpHeader{
					Name:  name,
					Value: value,
				})
			}
		}

		webhook := entity.Webhook{
			SectionId: sectionIdStr,
			OriginUrl: c.Request().RemoteAddr,
			Headers:   headers,
		}
		redisCtx := context.Background()

		result, err := webhook.MarshalBinary()
		if err != nil {
			config.Log.Log().Timestamp().Msg(err.Error())
			return err
		}
		rcmd := config.DB.JSONArrAppend(redisCtx, redisKey, "$", result)
		err = rcmd.Err()
		if err != nil {
			config.Log.Log().Timestamp().Msg(err.Error())
			return err
		}
		return c.JSON(200, webhook)

	}
}

func InitWebhooksRouter(config *config.ApiConfig, group *echo.Group) Webhooks {
	webhooks := Webhooks{
		Config: config,
	}
	webhooks.Route(group)
	return webhooks
}
