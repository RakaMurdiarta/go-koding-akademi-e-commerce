package impl

import (
	"context"
	"errors"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/product/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewProductRepository(tm *database.TransactionManagerImpl) repository.ProductRepository {
	return &productRepositoryImpl{
		TransactionManagerImpl: tm,
	}
}

func (r *productRepositoryImpl) CreateProduct(ctx context.Context, p *models.Product) error {
	return r.GetTx(ctx).Create(p).Error
}

func (r *productRepositoryImpl) CreateImages(ctx context.Context, imgs []models.ProductImage) error {
	return r.GetTx(ctx).Create(&imgs).Error
}

func (r *productRepositoryImpl) CreateVariants(ctx context.Context, vs []models.ProductVariant) error {
	return r.GetTx(ctx).Create(&vs).Error
}

func (r *productRepositoryImpl) CreateVariant(ctx context.Context, variant *models.ProductVariant) error {
	return r.GetTx(ctx).Create(variant).Error
}

func (r *productRepositoryImpl) CreateInventories(ctx context.Context, invs []models.Inventory) error {
	return r.GetTx(ctx).Create(&invs).Error
}

func (r *productRepositoryImpl) CreateInventory(ctx context.Context, inv *models.Inventory) error {
	return r.GetTx(ctx).Create(inv).Error
}

func (r *productRepositoryImpl) CreateAttributeValues(ctx context.Context, avs []models.VariantAttributeValue) error {
	return r.GetTx(ctx).Create(&avs).Error
}

func (r *productRepositoryImpl) GetProductByID(ctx context.Context, id uint) (*models.Product, error) {
	var p models.Product
	err := r.GetTx(ctx).Preload("Variants").First(&p, id).Error
	return &p, err
}

func (r *productRepositoryImpl) FindBySlug(ctx context.Context, slug string, preloads ...string) (*models.Product, error) {
	var product models.Product
	db := r.GetTx(ctx)

	for _, p := range preloads {
		db = db.Preload(p)
	}

	err := db.Where("slug = ?", slug).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepositoryImpl) UpdateProduct(ctx context.Context, p *models.Product) error {
	return r.GetTx(ctx).Save(p).Error
}

func (r *productRepositoryImpl) DeleteImagesByProductID(ctx context.Context, pID uint) error {
	return r.GetTx(ctx).Where("product_id = ? AND variant_id IS NULL", pID).Delete(&models.ProductImage{}).Error
}

func (r *productRepositoryImpl) DeleteImagesByVariantID(ctx context.Context, vID uint) error {
	return r.GetTx(ctx).Where("variant_id = ?", vID).Delete(&models.ProductImage{}).Error
}

func (r *productRepositoryImpl) UpdateVariant(ctx context.Context, v *models.ProductVariant) error {
	return r.GetTx(ctx).Save(v).Error
}

func (r *productRepositoryImpl) UpdateInventoryStock(ctx context.Context, variantID uint, newStock int) error {
	return r.GetTx(ctx).Model(&models.Inventory{}).
		Where("variant_id = ? AND ? >= reserved_stock", variantID, newStock).
		Updates(map[string]interface{}{
			"total_stock": newStock,
			"version":     gorm.Expr("version + 1"),
		}).Error
}

func (r *productRepositoryImpl) DeleteAttributeValuesByVariantID(ctx context.Context, variantID uint) error {
	return r.GetTx(ctx).Where("variant_id = ?", variantID).Delete(&models.VariantAttributeValue{}).Error
}

func (r *productRepositoryImpl) FindByID(ctx context.Context, id uint, preloads ...string) (*models.Product, error) {
	var product models.Product
	db := r.GetTx(ctx)

	// Dynamic Preloading base on Service needed!
	for _, p := range preloads {
		db = db.Preload(p)
	}

	err := db.First(&product, id).Error
	return &product, err
}

func (r *productRepositoryImpl) FindAll(ctx context.Context, limit, offset int, search string, preloads ...string) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.GetTx(ctx).Model(&models.Product{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	for _, p := range preloads {
		query = query.Preload(p)
	}

	err := query.Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func (r *productRepositoryImpl) FindByUserID(ctx context.Context, userID uint, limit int, offset int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	db := r.GetTx(ctx).Model(&models.Product{}).
		Joins("JOIN shops ON shops.id = products.shop_id").
		Where("shops.user_id = ?", userID)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Limit(limit).Offset(offset).
		Preload("Shop").
		Preload("Images", "is_primary = ?", true).
		Find(&products).Error

	return products, total, err
}

func (r *productRepositoryImpl) Delete(ctx context.Context, id uint, userID uint) error {

	res := r.GetTx(ctx).
		Where("id = ?", id).
		Where("shop_id IN (?)", r.GetTx(ctx).Model(&models.Shop{}).Select("id").Where("user_id = ?", userID)).
		Delete(&models.Product{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("product not found or you are not the owner")
	}

	return nil
}

func (r *productRepositoryImpl) CreateVariantAttributes(ctx context.Context, attrLinks []models.VariantAttributeValue) error {
	return r.GetTx(ctx).Create(&attrLinks).Error
}

func (r *productRepositoryImpl) FindVariantByID(ctx context.Context, id uint) (*models.ProductVariant, error) {
	var variant models.ProductVariant
	err := r.GetTx(ctx).First(&variant, id).Error
	return &variant, err
}

func (r *productRepositoryImpl) DeleteVariantAttributesByVariantID(ctx context.Context, variantID uint) error {

	return r.GetTx(ctx).
		Where("variant_id = ?", variantID).
		Delete(&models.VariantAttributeValue{}).Error
}

func (r *productRepositoryImpl) GetImagesByProductID(ctx context.Context, productID uint) ([]models.ProductImage, error) {
	var images []models.ProductImage

	err := r.GetTx(ctx).
		Where("product_id = ?", productID).
		Order("is_primary DESC, created_at ASC").
		Find(&images).
		Error

	if err != nil {
		return nil, err
	}

	return images, nil
}
