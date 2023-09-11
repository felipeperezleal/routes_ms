package models

import "gorm.io/gorm"

type RouteGraph struct {
	gorm.Model
	ID       int `gorm:"primaryKey" json:"id" db:"id"`
	NumNodes int `json:"numNodes" db:"numNodes"`
}
