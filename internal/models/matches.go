package models

import (
	"gorm.io/gorm"
)

type Matches struct {
	gorm.Model
	Img     string    `gorm:"not null"`
	People  []*People `gorm:"many2many:people_matches;"`
	AreasID uint
	Areas   Areas
}
