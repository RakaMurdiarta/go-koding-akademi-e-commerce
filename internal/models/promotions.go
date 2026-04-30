package models

import (
	"time"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
)

type Promotion struct {
	ID            uint                `gorm:"primaryKey"`
	Code          string              `gorm:"size:20;uniqueIndex;not null"`
	DiscountType  common.DiscountType `gorm:"size:10"`
	DiscountValue float64             `gorm:"type:decimal(15,2);not null"`
	MinSpend      float64             `gorm:"type:decimal(15,2);default:0"`
	MaxDiscount   *float64            `gorm:"type:decimal(15,2)"`

	StartDate *time.Time
	EndDate   *time.Time

	IsActive bool `gorm:"default:true"`
}

func (Promotion) TableName() string {
	return "promotions"
}
