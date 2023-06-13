package models

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	Name           string `gorm:"not null" json:"name,omitempty"`
	Password       string `gorm:"not null" json:"password,omitempty"`
	Email          string `gorm:"not null;unique_index" json:"email,omitempty"`
	Type           string `gorm:"not null" json:"type,omitempty"`
	Premium_Points string `gorm:"not null;unique_index" json:"premium_points,omitempty"`
}
