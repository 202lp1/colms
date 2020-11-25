package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

func (tab *Alumno) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

type Alumno struct {
	Id      string `gorm:"primary_key;"`
	Nombres string
	Codigo  string //`gorm:"column:my_ciudad"`
}
