package models

import "time"

type BaseModel struct {
	CreationDate     time.Time
	ModificationDate time.Time
	Status           int
}

func NewBaseModel() BaseModel {
	return BaseModel{
		Status:           1,
		ModificationDate: time.Now(),
		CreationDate:     time.Now(),
	}
}


