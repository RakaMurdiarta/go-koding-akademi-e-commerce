package models

import "time"

type Review struct {
	ID          uint `gorm:"primaryKey"`
	OrderItemID uint `gorm:"uniqueIndex"`
	UserID      uint
	ProductID   uint
	Rating      int `gorm:"check:rating >= 1 AND rating <= 5"`
	Comment     *string

	CreatedAt time.Time

	User      User
	Product   Product
	OrderItem OrderItem
}
