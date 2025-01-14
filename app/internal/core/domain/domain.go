package domain

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Id    string `json:"id,omitempty" gorm:"primaryKey"`
	Name  string `json:"name,omitempty" gorm:"size:256"`
	Price int    `json:"price,omitempty"` // TODO Could not be float  64?

	// GORM Variables
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
