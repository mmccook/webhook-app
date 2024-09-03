package entity

type Section struct {
	Id         uint64 `json:"id" validate:"required"`
	LastActive string `json:"lastActive" validate:"required,datetime"`
}
