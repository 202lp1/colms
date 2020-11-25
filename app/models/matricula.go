package models

//func (tab *Matricula) BeforeCreate(*gorm.DB) error {
//	tab.Id = uuid.NewV4().String()
//	return nil
//}

type Matricula struct {
	Id int `gorm:"primary_key;"`
	//Fecha  string
	Semestre string //`gorm:"column:my_ciudad"`
	AlumnoId int    // `gorm:"column:alumno_id;type:varchar(191);"`
	Alumno   Alumno //`gorm:"foreignkey:AlumnoId"`
}
