package models

import "gorm.io/gorm"

type Cameras struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	FrigateName string `gorm:"unique"`
	AreasID     uint
}
