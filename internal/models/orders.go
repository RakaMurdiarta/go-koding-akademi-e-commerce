package models

import (
	"time"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
)

type Order struct {
	ID             uint `gorm:"primaryKey"`
	TransactionID  uint
	ShopID         uint
	TotalPrice     float64            `gorm:"type:decimal(15,2);not null"`
	ShippingCost   float64            `gorm:"type:decimal(15,2);default:0"`
	Status         common.OrderStatus `gorm:"size:20;default:waiting_payment"`
	TrackingNumber *string            `gorm:"size:100"`

	ReceiverName    string
	ReceiverPhone   string
	ShippingAddress string

	CreatedAt time.Time

	Transaction Transaction
	Shop        Shop
	OrderItems  []OrderItem
}

type OrderItem struct {
	ID         uint `gorm:"primaryKey"`
	OrderID    uint
	VariantID  uint
	Qty        int
	PriceAtBuy float64 `gorm:"type:decimal(15,2);not null"`

	Order   Order
	Variant ProductVariant
	Review  *Review
}
