package controller

import (
	"DJMIL/api/template"
	"DJMIL/config"
	"DJMIL/entity"
	"DJMIL/utils"
	"context"
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type Sections struct {
	BaseController
}

func (sectionController *Sections) Route() {
	var group = sectionController.App.Group("/sections")
	sectionController.App.GET("/", sectionController.getHome)
	sectionController.App.POST("/", sectionController.postSectionInitialize)
	group.GET("/:sectionId", sectionController.getSectionDash)
}

func (sectionController *Sections) getHome(c echo.Context) error {
	_, err := utils.CreateSession(c, sectionController.Config.GetString("SESSION_NAME"), utils.Session{
		LastActive: time.DateTime,
	})
	if err != nil {
		return err
	}
	return template.AssertRender(c, http.StatusOK, template.Index())
}

func (sectionController *Sections) postSectionInitialize(c echo.Context) error {
	id := snowflake.ID()

	section := entity.Section{
		Id:         id,
		LastActive: time.DateTime,
	}
	sectionIdStr := strconv.FormatUint(id, 10)

	webhook := entity.Webhook{
		SectionId: sectionIdStr,
	}

	redisCtx := context.Background()

	sectionController.DB.JSONSet(redisCtx, section.RedisKey(), "$", section)
	sectionController.DB.JSONSet(redisCtx, webhook.RedisKey(), "$", "[]")

	c.Response().Header().Set("Hx-Redirect", fmt.Sprintf("/sections/%d", id))
	return c.NoContent(201)
}

func (sectionController *Sections) getSectionDash(c echo.Context) error {
	var sectionIdStr = c.Param("sectionId")
	return template.AssertRender(c, http.StatusOK, template.SectionDash(sectionIdStr))
}

func InitSectionRouter(config *config.ApiConfig) Sections {
	sections := Sections{
		BaseController{
			config,
		},
	}
	sections.Route()
	return sections
}
