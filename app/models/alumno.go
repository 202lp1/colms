package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Alumno struct {
	Id         int `gorm:"primary_key;"`
	Nombres    string
	Codigo     string      //`gorm:"column:my_ciudad"`
	Matriculas []Matricula // `gorm:"foreignkey:ManagerID"`
}

func (tab Alumno) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", tab.Id, tab.Nombres)
}

func (alumno Alumno) FindAll(conn *gorm.DB) ([]Alumno, error) {
	var alumnos []Alumno
	conn.Preload("Matriculas").Find(&alumnos)
	return alumnos, nil
}

//"github.com/twinj/uuid"

//func (tab *Alumno) BeforeCreate(*gorm.DB) error {
//	tab.Id = uuid.NewV4().String()
//	return nil
//}
