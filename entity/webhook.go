package entity

import "encoding/json"

type Webhook struct {
	SectionId string       `json:"sectionId"`
	OriginUrl string       `json:"originUrl" validate:"required, url"`
	Headers   []HttpHeader `json:"headers"`
}

func (m *Webhook) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *Webhook) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

type HttpHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
