package models

import (
	"github.com/twinj/uuid" //jwt-best-practices
	"gorm.io/gorm"
)

//https://gorm.io/docs/conventions.html
//type Tabler interface {
//TableName() string
//}

// TableName overrides the table name used by Empleado to `employee`
func (Empleado) TableName() string {
	return "employee2"
}

// BeforeCreate will set a UUID rather than numeric ID. https://gorm.io/docs/create.html

func (tab *Empleado) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

//https://gorm.io/docs/models.html
type Empleado struct {
	//gorm.Model

	//ID uint `gorm:"primaryKey"`
	//ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	//ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Id string `gorm:"primary_key;column:id"` //;default:UUID()
	//UUID   string `gorm:"primaryKey"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`

	Name string
	City string `gorm:"column:my_ciudad"`
}

/*
type User struct {
  gorm.Model
  Name string
}
// equals
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
*/
