package migration

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Payment Payment
}
