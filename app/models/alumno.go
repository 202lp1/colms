package models

import (
	"fmt"

	"github.com/202lp1/colms/cfig"
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

func (alumno Alumno) FindAll() ([]Alumno, error) {
	var alumnos []Alumno
	cfig.DB.Preload("Matriculas").Find(&alumnos)
	return alumnos, nil
}

//"github.com/twinj/uuid"

//func (tab *Alumno) BeforeCreate(*gorm.DB) error {
//	tab.Id = uuid.NewV4().String()
//	return nil
//}
