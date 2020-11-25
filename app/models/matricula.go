package models

import (
	"fmt"
)

type Matricula struct {
	Id int `gorm:"primary_key;"`
	//Fecha  string
	Semestre string //`gorm:"column:my_ciudad"`
	AlumnoId int    // `gorm:"column:alumno_id;type:varchar(191);"`
	Alumno   Alumno //para crear el FK `gorm:"foreignkey:AlumnoId"`
}

func (tab Matricula) ToString() string {
	return fmt.Sprintf("id: %d\nSemestre: %s", tab.Id, tab.Semestre)
}

//func (tab *Matricula) BeforeCreate(*gorm.DB) error {
//	tab.Id = uuid.NewV4().String()
//	return nil
//}
