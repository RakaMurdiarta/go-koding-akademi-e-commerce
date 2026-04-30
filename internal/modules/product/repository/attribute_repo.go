package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type AttributeRepository interface {
	CreateAttribute(ctx context.Context, attribute *models.Attribute) error
	CreateAttributeValues(ctx context.Context, attrValues []models.AttributeValue) error
	GetByID(ctx context.Context, id uint) (*models.Attribute, error)
	GetByName(ctx context.Context, name string) (*models.Attribute, error)
	GetAll(ctx context.Context) ([]models.Attribute, error)
	CheckValueExists(ctx context.Context, attrID uint, value string) (bool, error)
	UpdateValue(ctx context.Context, valueID uint, newValue string) error
	Update(ctx context.Context, attribute *models.Attribute) error
	Delete(ctx context.Context, id uint) error

	database.TransactionManager
}

type VariantAttributeValueRepository interface {
	Create(ctx context.Context, v *models.VariantAttributeValue) error
	GetByID(ctx context.Context, id int) (*models.VariantAttributeValue, error)
	GetByVariantID(ctx context.Context, variantID int) ([]models.VariantAttributeValue, error)
	Update(ctx context.Context, v *models.VariantAttributeValue) error
	Delete(ctx context.Context, id int) error
}
