package models

type Note struct {
	Id      uint32 `json:"-"`
	Message string
	Tag     string `json:"tag,omitempty"`
}
