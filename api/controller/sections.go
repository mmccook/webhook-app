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
	Config *config.ApiConfig
}

func (sectionController *Sections) Route() {
	var group = sectionController.Config.App.Group("/sections")
	sectionController.Config.App.GET("/", getHome(sectionController.Config))
	sectionController.Config.App.POST("/", postSectionInitialize(sectionController.Config))
	group.GET("/:sectionId", getSectionDash(sectionController.Config))
}

func getHome(config *config.ApiConfig) func(c echo.Context) error {
	return func(c echo.Context) error {
		utils.CreateSession(c, config.Config.GetString("SESSION_NAME"), utils.Session{
			LastActive: time.DateTime,
		})
		return template.AssertRender(c, http.StatusOK, template.Index())

	}
}
func postSectionInitialize(config *config.ApiConfig) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := snowflake.ID()
		c.Response().Header().Set("Hx-Redirect", fmt.Sprintf("/sections/%d", id))
		section := entity.Section{
			Id:         id,
			LastActive: time.DateTime,
		}
		sectionIdStr := strconv.FormatUint(id, 10)
		redisKey := fmt.Sprintf("%s::webhooks", sectionIdStr)
		redisCtx := context.Background()
		config.DB.JSONSet(redisCtx, sectionIdStr, "$", section)
		config.DB.JSONSet(redisCtx, redisKey, "$", "[]")
		return c.NoContent(201)
	}
}

func getSectionDash(config *config.ApiConfig) func(c echo.Context) error {
	return func(c echo.Context) error {
		var sectionId, err = strconv.ParseUint(c.Param("sectionId"), 10, 64)
		if err != nil {
			return err
		}
		return template.AssertRender(c, http.StatusOK, template.SectionDash(strconv.FormatUint(sectionId, 10)))
	}
}

func InitSectionRouter(config *config.ApiConfig) Sections {
	sections := Sections{
		Config: config,
	}
	sections.Route()
	return sections
}
