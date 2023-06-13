package models

import "gorm.io/gorm"

type Reports struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name"`
	Position   string `gorm:"not null" json:"position"`
	Status     bool   `gorm:"defult:false" json:"status"`
	ID_account int64  `json:"id_acocunt"`
}
