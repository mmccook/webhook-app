package entity

import (
	"fmt"
	"strconv"
)

type Section struct {
	Id         uint64 `json:"id" validate:"required"`
	LastActive string `json:"lastActive" validate:"required,datetime"`
}

func (entity *Section) RedisKey() string {
	idStr := strconv.FormatUint(entity.Id, 10)
	return fmt.Sprintf("section:%s", idStr)
}
