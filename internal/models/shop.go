package models

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"uniqueIndex"`
	Name        string  `gorm:"size:100;uniqueIndex;not null"`
	Description *string `gorm:"type:text"`
	LogoURL     *string `gorm:"type:text"`
	IsActive    bool    `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User     User `gorm:"foreignKey:UserID"`
	Products []Product
	Orders   []Order
}

func (Shop) TableName() string {
	return "shops"
}
