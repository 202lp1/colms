package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Alumno struct {
	Id         string `gorm:"primaryKey;"`
	Nombres    string
	Codigo     string
	Matriculas []Matricula
}

func (tab Alumno) ToString() string {
	return tab.Nombres
}

func (tab *Alumno) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (alumno Alumno) FindAll(conn *gorm.DB) ([]Alumno, error) {
	var alumnos []Alumno
	conn.Preload("Matriculas").Find(&alumnos)
	return alumnos, nil
}

func (alumno Alumno) GetAll(conn *gorm.DB) ([]Alumno, error) {
	var alumnos []Alumno
	conn.Find(&alumnos)
	return alumnos, nil
}
