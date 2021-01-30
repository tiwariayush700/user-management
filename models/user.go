package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName" validate:"required" binding:"required" gorm:"type:text;not null"`
	LastName  string `json:"lastName" validate:"required" binding:"required" gorm:"type:text;not null"`
	Email     string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Role      string `json:"role" binding:"required" validate:"required,oneof=ADMIN USER" gorm:"type:text;check:role = 'ADMIN' or role = 'USER';not null"`
	Phone     string `json:"phone" binding:"required" gorm:"type:text"`
	Password  string `json:"password" binding:"required" validate:"required" gorm:"type:text"`
}
