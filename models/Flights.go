package models

import "gorm.io/gorm"

type Flight struct {
	gorm.Model

	Id          uint    `gorm:"primaryKey"`
	Origin      string  `gorm:"size:255 not null"`
	Destination string  `gorm:"size:255 not null"`
	Duration    int     `gorm:"not null"`
	Distance    float32 `gorm:"not null"`
	Price       float32 `gorm:"not null"`
}
