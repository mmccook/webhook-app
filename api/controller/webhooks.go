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
	BaseController
}

func (controller *Webhooks) Route() {
	var group = controller.App.Group("/sections")
	group.GET("/:sectionId/webhooks", controller.getWebhookListHandler)
	group.POST("/:sectionId/webhooks", controller.postWebhookHandler)
}

func (controller *Webhooks) getWebhookListHandler(c echo.Context) error {
	var sectionId, err = controller.parseSectionID(c.Param("sectionId"))
	if err != nil {
		return err
	}

	sectionIdStr := strconv.FormatUint(sectionId, 10)

	redisKey := fmt.Sprintf("%s::webhooks", sectionIdStr)

	redisCtx := context.Background()

	rcmd := controller.DB.JSONGet(redisCtx, redisKey, "$")

	result, err := rcmd.Result()
	if err != nil {
		controller.Log.Log().Timestamp().Msg(err.Error())
		return err
	}

	var data []entity.Webhook
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		controller.Log.Log().Timestamp().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, data)

}
func (controller *Webhooks) postWebhookHandler(c echo.Context) error {

	var sectionId, err = controller.parseSectionID(c.Param("sectionId"))
	if err != nil {
		return err
	}
	sectionIdStr := strconv.FormatUint(sectionId, 10)

	redisKey := fmt.Sprintf("section:%s:webhooks", sectionIdStr)

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
		controller.Log.Log().Timestamp().Msg(err.Error())
		return err
	}
	rcmd := controller.DB.JSONArrAppend(redisCtx, redisKey, "$", result)
	err = rcmd.Err()
	if err != nil {
		controller.Log.Log().Timestamp().Msg(err.Error())
		return err
	}
	return c.JSON(200, webhook)

}

func InitWebhooksRouter(config *config.ApiConfig) Webhooks {
	webhooks := Webhooks{
		BaseController{
			config,
		},
	}
	webhooks.Route()
	return webhooks
}
