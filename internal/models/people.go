package models

import "gorm.io/gorm"

type People struct {
	gorm.Model
	Name             string     `gorm:"unique;not null"`
	Resident         bool       `gorm:"default:false"`
	PrivilegedAccess bool       `gorm:"default:false"`
	Matches          []*Matches `gorm:"many2many:people_matches;"`
}
