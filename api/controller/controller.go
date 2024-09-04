package controller

import (
	"DJMIL/config"
	"strconv"
)

type BaseController struct {
	*config.ApiConfig
}

func (controller *BaseController) parseSectionID(sectionIdStr string) (uint64, error) {
	return strconv.ParseUint(sectionIdStr, 10, 64)
}
