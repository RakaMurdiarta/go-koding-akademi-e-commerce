package models

import "time"

type Inventory struct {
	ID            uint `gorm:"primaryKey"`
	VariantID     uint `gorm:"uniqueIndex"`
	TotalStock    int  `gorm:"default:0"`
	ReservedStock int  `gorm:"default:0"`
	Version       int

	CreatedAt time.Time
	UpdatedAt time.Time

	Variant ProductVariant `gorm:"foreignKey:VariantID"`
}

func (Inventory) TableName() string {
	return "inventory"
}
