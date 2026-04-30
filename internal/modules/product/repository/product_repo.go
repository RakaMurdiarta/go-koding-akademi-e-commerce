package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, p *models.Product) error
	CreateImages(ctx context.Context, imgs []models.ProductImage) error
	CreateVariants(ctx context.Context, vs []models.ProductVariant) error
	CreateVariant(ctx context.Context, variant *models.ProductVariant) error
	CreateInventories(ctx context.Context, invs []models.Inventory) error
	CreateInventory(ctx context.Context, inv *models.Inventory) error
	CreateAttributeValues(ctx context.Context, avs []models.VariantAttributeValue) error
	CreateVariantAttributes(ctx context.Context, attrLinks []models.VariantAttributeValue) error
	FindVariantByID(ctx context.Context, id uint) (*models.ProductVariant, error)
	DeleteVariantAttributesByVariantID(ctx context.Context, variantID uint) error

	GetProductByID(ctx context.Context, id uint) (*models.Product, error)
	FindBySlug(ctx context.Context, slug string, preloads ...string) (*models.Product, error)
	UpdateProduct(ctx context.Context, p *models.Product) error
	DeleteImagesByProductID(ctx context.Context, productID uint) error
	UpdateVariant(ctx context.Context, v *models.ProductVariant) error
	UpdateInventoryStock(ctx context.Context, variantID uint, stock int) error
	DeleteAttributeValuesByVariantID(ctx context.Context, variantID uint) error
	DeleteImagesByVariantID(ctx context.Context, vID uint) error

	FindByID(ctx context.Context, id uint, preloads ...string) (*models.Product, error)
	FindAll(ctx context.Context, limit, offset int, search string, preloads ...string) ([]models.Product, int64, error)
	FindByUserID(ctx context.Context, userID uint, limit int, offset int) ([]models.Product, int64, error)
	Delete(ctx context.Context, id uint, userID uint) error
	GetImagesByProductID(ctx context.Context, productID uint) ([]models.ProductImage, error)

	database.TransactionManager
}

type ProductVariantRepository interface {
	Create(ctx context.Context, variant *models.ProductVariant) error

	GetByID(ctx context.Context, id uint) (*models.ProductVariant, error)

	GetByProductID(ctx context.Context, productID uint) ([]models.ProductVariant, error)

	Delete(ctx context.Context, id uint) error

	ExistsBySKU(ctx context.Context, sku string) (bool, error)
}
