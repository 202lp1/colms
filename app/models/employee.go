package models

import (
	"gorm.io/gorm"
)

type Empleado struct {
	gorm.Model
	//Id   int
	Name string
	City string
}
