package models

import (
	"time"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
)

type Transaction struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	PromotionID      *uint
	GrandTotal       float64              `gorm:"type:decimal(15,2);not null"`
	PaymentMethod    string               `gorm:"size:50"`
	PaymentStatus    common.PaymentStatus `gorm:"size:20;default:pending"`
	SnapToken        *string              `gorm:"type:text"`
	PaymentExpiredAt *time.Time

	CreatedAt time.Time

	User      User
	Promotion *Promotion
	Orders    []Order
}

type PaymentLog struct {
	ID uint `gorm:"primaryKey"`

	TransactionID uint

	Payload string `gorm:"type:jsonb"`

	CreatedAt time.Time

	Transaction Transaction
}
