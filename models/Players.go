package models

import "gorm.io/gorm"

type Players struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name"`
	Group_ID   uint8  `gorm:"not null" json:"group_id"`
	Account_Id int16  `gorm:"not null" json:"account_id"`
	Level      int64  `gorm:"not null" json:"level"`
	Town_ID    uint8  `gorm:"not null" json:"town_id"`
	PosX       uint64 `gorm:"not null" json:"posx"`
	PosY       uint64 `gorm:"not null" json:"posy"`
	PosZ       int8   `gorm:"not null" json:"posz"`
}
