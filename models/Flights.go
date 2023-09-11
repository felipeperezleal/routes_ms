package models

import "gorm.io/gorm"

type Flight struct {
	gorm.Model

	Origin      string  `json:"origin" db:"origin"`
	Destination string  `json:"destination" db:"destination"`
	Duration    int     `json:"duration" db:"duration"`
	Distance    float64 `json:"distance" db:"distance"`
	Price       float64 `json:"price" db:"price"`
}
