package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	BaseEntity
	SectionId string                 `json:"sectionId"`
	OriginUrl string                 `json:"originUrl" validate:"required, url"`
	Headers   []HttpHeader           `json:"headers"`
	Body      map[string]interface{} `json:"body"`
}

func NewWebhook(sectionId string, originUrl string, header http.Header, body map[string]interface{}) Webhook {
	var headers []HttpHeader
	for name, values := range header {
		for _, value := range values {
			headers = append(headers, NewHttpHeader(name, value))
		}
	}
	baseEntity := NewBaseEntity()
	return Webhook{
		BaseEntity: baseEntity,
		SectionId:  sectionId,
		OriginUrl:  originUrl,
		Headers:    headers,
		Body:       body,
	}
}

func (entity *Webhook) RedisKey() string {
	return fmt.Sprintf("webhook:%s", entity.Id)
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

func NewHttpHeader(name string, value string) HttpHeader {
	return HttpHeader{
		Name:  name,
		Value: value,
	}
}
