package models


type Permission struct {
	BaseModel
	Id       int    `json:"id"`
	Name     string `json:"name"`
	MetaData string `json:"metaData"`
}

