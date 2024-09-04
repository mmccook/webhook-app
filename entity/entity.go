package entity

import (
	"github.com/godruoyi/go-snowflake"
	"time"
)

type BaseEntity struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt"`
}

func NewBaseEntity() BaseEntity {
	return BaseEntity{
		Id:        snowflake.ID(),
		CreatedAt: time.Now().Unix(),
	}
}
