package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null;unique_index" json:"title"`
	Description string `json:"description"`
	TypeTask    string `gorm:"not null" json:"type_task"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `json:"user_id"`
}
