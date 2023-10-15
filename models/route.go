package models

import "gorm.io/gorm"

type Routes struct {
	gorm.Model

	Origin   string `json:"origin" db:"origin"`
	Destiny  string `json:"destiny" db:"destiny"`
	NumNodes int    `json:"numNodes" db:"numNodes"`
	Ordering string `json:"ordering" db:"ordering"`
}
