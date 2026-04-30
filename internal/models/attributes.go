package models

type Attribute struct {
	ID     uint             `gorm:"primaryKey"`
	Name   string           `gorm:"type:varchar(255);not null;index:idx_category_attr,unique"`
	Values []AttributeValue `gorm:"foreignKey:AttributeID"`
}

func (Attribute) TableName() string {
	return "attributes"
}

type AttributeValue struct {
	ID          uint   `gorm:"primaryKey"`
	AttributeID uint   `gorm:"index"`
	Value       string `gorm:"type:varchar(255);not null"`

	Attribute Attribute `gorm:"foreignKey:AttributeID"`

	VariantLinks []VariantAttributeValue `gorm:"foreignKey:AttributeValueID"`
}

func (AttributeValue) TableName() string {
	return "attribute_values"
}
