package impl

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/product/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type productVariantReposiotry struct {
	*database.TransactionManagerImpl
}

func NewProductVariantRepository(db *database.TransactionManagerImpl) repository.ProductVariantRepository {
	return &productVariantReposiotry{
		TransactionManagerImpl: db,
	}
}

func (r *productVariantReposiotry) Create(ctx context.Context, v *models.ProductVariant) error {

	return r.GetTx(ctx).
		Create(v).
		Error
}

func (r *productVariantReposiotry) GetByID(ctx context.Context, id uint) (*models.ProductVariant, error) {

	var variant models.ProductVariant

	err := r.GetTx(ctx).Preload("Inventory").
		First(&variant, id).
		Error

	if err != nil {
		return nil, err
	}

	return &variant, nil
}

func (r *productVariantReposiotry) GetByProductID(ctx context.Context, productID uint) ([]models.ProductVariant, error) {

	var variants []models.ProductVariant

	err := r.GetTx(ctx).
		Where("product_id = ?", productID).
		Order("id ASC").
		Find(&variants).
		Error

	return variants, err
}

func (r *productVariantReposiotry) Delete(ctx context.Context, id uint) error {

	return r.GetTx(ctx).
		Delete(&models.ProductVariant{}, id).
		Error
}

func (r *productVariantReposiotry) ExistsBySKU(ctx context.Context, sku string) (bool, error) {

	var count int64

	err := r.GetTx(ctx).
		Model(&models.ProductVariant{}).
		Where("sku = ?", sku).
		Count(&count).
		Error

	return count > 0, err
}
