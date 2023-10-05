package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key;autoIncrement;column:id`
	FirstName string `gorm:"column:first_name;type:varchar(255);not null"`
	LastName  string `gorm:"column:last_name;type:varchar(255);not null"`
	Email     string `gorm:"unique;column:email;not null;type:varchar(500)"`
	Phone     string `gorm:"column:phone;type:varchar(255)"`
	Password  string `gorm:"not null;type:varchar(255)" json:"-"`
	Todos     []Todo
	gorm.Model
}
