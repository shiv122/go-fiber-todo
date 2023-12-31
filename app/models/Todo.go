package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint       `gorm:"primary_key;autoIncrement"`
	UserID      uint       `gorm:"column:user_id"`
	Name        string     `gorm:"column:name;type:varchar(500);not null"`
	Description string     `grom:"column:description"`
	Status      string     `gorm:"column:status;type:varchar(500);not null;default:pending"`
	CompletedAt *time.Time `gorm:"column:completed_at"`
	User        *User
	gorm.Model
}

/*
|======================================================================
| Scopes
|======================================================================
*/

func CompletedTodo(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", "completed")
}
