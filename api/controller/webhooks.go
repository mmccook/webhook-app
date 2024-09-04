package controller

import (
	"DJMIL/config"
	"DJMIL/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type WebhooksController struct {
	BaseController
	webhookService service.WebhookService
}

func (controller *WebhooksController) Route() {
	var group = controller.App.Group("/sections")
	group.POST("/:sectionId/webhooks", controller.postWebhookHandler)
}

func (controller *WebhooksController) postWebhookHandler(c echo.Context) error {
	var sectionIdStr = c.Param("sectionId")

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Error(err)
		log.Error("empty json body")
	}

	webhook, err := controller.webhookService.CreateWebhook(sectionIdStr, c.Request().RemoteAddr, c.Request().Header, jsonBody)

	if err != nil {
		controller.Log.Log().Timestamp().Msg(err.Error())
		return err
	}

	return c.JSON(200, webhook)
}

func InitWebhooksRouter(config *config.ApiConfig) WebhooksController {
	webhooks := WebhooksController{
		BaseController: BaseController{
			config,
		},
		webhookService: service.WebhookService{
			BaseService: service.BaseService{
				DB:  config.DB,
				Log: &config.Log,
			},
		},
	}
	_ = webhooks.webhookService.CreateIndex()

	webhooks.Route()
	return webhooks
}
