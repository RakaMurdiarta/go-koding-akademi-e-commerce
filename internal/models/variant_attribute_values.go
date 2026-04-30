package models

type VariantAttributeValue struct {
	VariantID        uint `gorm:"primaryKey"`
	AttributeValueID uint `gorm:"primaryKey"`

	Variant        *ProductVariant `gorm:"foreignKey:VariantID"`
	AttributeValue *AttributeValue `gorm:"foreignKey:AttributeValueID"`
}

func (VariantAttributeValue) TableName() string {
	return "variant_attribute_values"
}
