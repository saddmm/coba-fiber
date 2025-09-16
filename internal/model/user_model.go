package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null, type:varchar(20)"`
	Email    string `json:"email" gorm:"unique, type:varchar(20)"`
	Password string `json:"password" gorm:"not null, type:varchar(50)"`
	Posts    []Post `json:"posts" gorm:"foreignKey:UserID"`
}
