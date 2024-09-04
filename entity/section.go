package entity

import (
	"fmt"
	"strconv"
	"time"
)

type Section struct {
	BaseEntity
	LastActive int64 `json:"lastActive" validate:"required,datetime"`
}

func (entity *Section) RedisKey() string {
	idStr := strconv.FormatUint(entity.Id, 10)

	return fmt.Sprintf("section:%s", idStr)
}

func NewSection() Section {
	baseEntity := NewBaseEntity()
	return Section{
		BaseEntity: baseEntity,
		LastActive: time.Now().Unix(),
	}
}
