package entity

import (
	"encoding/json"
	"fmt"
)

type Webhook struct {
	SectionId string       `json:"sectionId"`
	OriginUrl string       `json:"originUrl" validate:"required, url"`
	Headers   []HttpHeader `json:"headers"`
}

func (entity *Webhook) RedisKey() string {
	return fmt.Sprintf("section:%s:webhooks", entity.SectionId)
}

func (entity *Webhook) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, entity)
}

func (entity *Webhook) MarshalBinary() (data []byte, err error) {
	return json.Marshal(entity)
}

type HttpHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
