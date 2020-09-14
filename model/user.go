package model

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Job  string `gorm:"not null" json:"job"`
	Age  int    `gorm:"not null" json:"Age"`
}
