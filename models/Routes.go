package models

import "gorm.io/gorm"

type Routes struct {
	gorm.Model

	NumNodes int    `json:"numNodes" db:"numNodes"`
	Ordering string `json:"ordering" db:"ordering"`
}
