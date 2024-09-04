package controller

import (
	"DJMIL/api/template"
	"DJMIL/config"
	"DJMIL/service"
	"DJMIL/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type SectionController struct {
	BaseController
	sectionService service.SectionService
}

func (sectionController *SectionController) Route() {
	var group = sectionController.App.Group("/sections")
	sectionController.App.GET("/", sectionController.getHome)
	sectionController.App.POST("/", sectionController.postSectionInitialize)
	group.GET("/:sectionId", sectionController.getSectionDash)
}

func (sectionController *SectionController) getHome(c echo.Context) error {
	_, err := utils.CreateSession(c, sectionController.Config.GetString("SESSION_NAME"), utils.Session{
		LastActive: time.DateTime,
	})
	if err != nil {
		return err
	}
	return template.AssertRender(c, http.StatusOK, template.Index())
}

func (sectionController *SectionController) postSectionInitialize(c echo.Context) error {
	section, err := sectionController.sectionService.CreateSection()
	if err != nil {
		return err
	}
	c.Response().Header().Set("Hx-Redirect", fmt.Sprintf("/sections/%d", section.Id))
	return c.NoContent(201)
}

func (sectionController *SectionController) getSectionDash(c echo.Context) error {
	var sectionIdStr = c.Param("sectionId")
	return template.AssertRender(c, http.StatusOK, template.SectionDash(sectionIdStr))
}

func InitSectionRouter(config *config.ApiConfig) SectionController {
	sections := SectionController{
		BaseController: BaseController{
			config,
		},
		sectionService: service.SectionService{
			BaseService: service.BaseService{
				DB:  config.DB,
				Log: &config.Log,
			},
		},
	}
	_ = sections.sectionService.CreateIndex()

	sections.Route()
	return sections
}
