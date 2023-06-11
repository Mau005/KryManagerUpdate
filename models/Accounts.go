package models

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	Name           string `gorm:"not null" json:"name"`
	Password       string `gorm:"not null" json:"-"`
	Email          string `gorm:"not null;unique_index" json:"email"`
	Type           string `gorm:"not null" json:"type"`
	Premium_Points string `gorm:"not null;unique_index" json:"premium_points"`
}
