package models

//"github.com/twinj/uuid"

//func (tab *Alumno) BeforeCreate(*gorm.DB) error {
//	tab.Id = uuid.NewV4().String()
//	return nil
//}

type Alumno struct {
	Id      int `gorm:"primary_key;"`
	Nombres string
	Codigo  string //`gorm:"column:my_ciudad"`
	//Matricula []Matricula // `gorm:"foreignkey:ManagerID"`
}
