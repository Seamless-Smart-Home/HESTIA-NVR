package models

import "gorm.io/gorm"

type Areas struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Cameras []Cameras
}
