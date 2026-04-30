package models

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	VariantID uint
	Quantity  int `gorm:"check:quantity > 0"`

	CreatedAt time.Time

	User    User
	Variant ProductVariant
}
