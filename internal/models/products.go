package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey"`
	ShopID      uint
	CategoryID  uint
	Name        string  `gorm:"size:255;not null"`
	Slug        string  `gorm:"size:255;uniqueIndex;not null"`
	Description *string `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(15,2);not null"`
	Weight      int     `gorm:"default:100"`
	Unit        string
	RatingAvg   float64
	RatingCount int
	IsActive    bool
	Thumbnail   string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Shop     Shop
	Category Category
	Variants []ProductVariant
	Images   []ProductImage
	Reviews  []Review
}

type ProductImage struct {
	ID uint `gorm:"primaryKey"`

	ProductID uint
	ImageURL  string
	IsPrimary bool

	CreatedAt time.Time

	Product Product
}

type ProductVariant struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint
	PriceExtra float64 `gorm:"type:decimal(15,2)"`

	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Product Product

	AttributeValues []VariantAttributeValue `gorm:"foreignKey:VariantID;constraint:OnDelete:CASCADE"`
	Inventory       *Inventory              `gorm:"foreignKey:VariantID"`
}

func (Product) TableName() string {
	return "products"
}

func (ProductVariant) TableName() string {
	return "product_variants"
}
