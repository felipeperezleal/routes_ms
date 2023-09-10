package models

import "gorm.io/gorm"

type RouteGraph struct {
	gorm.Model
	Id       uint `gorm:"primaryKey"`
	NumNodes uint `gorm:"not null"`
}
