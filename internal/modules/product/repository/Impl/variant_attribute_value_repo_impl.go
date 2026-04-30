package impl

import (
	"context"
	"errors"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/product/repository"
	"gorm.io/gorm"
)

type variantAttributeValueRepo struct {
	DB *gorm.DB
}

func NewVariantAttributeValueRepository(db *gorm.DB) repository.VariantAttributeValueRepository {
	return &variantAttributeValueRepo{DB: db}
}

func (r *variantAttributeValueRepo) Create(ctx context.Context, v *models.VariantAttributeValue) error {
	return r.DB.WithContext(ctx).Create(v).Error
}

func (r *variantAttributeValueRepo) GetByID(ctx context.Context, id int) (*models.VariantAttributeValue, error) {
	var v models.VariantAttributeValue
	err := r.DB.WithContext(ctx).First(&v, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("variant attribute value not found")
		}
		return nil, err
	}
	return &v, nil
}

func (r *variantAttributeValueRepo) GetByVariantID(ctx context.Context, variantID int) ([]models.VariantAttributeValue, error) {
	var results []models.VariantAttributeValue
	err := r.DB.WithContext(ctx).Where("variant_id = ?", variantID).Find(&results).Error
	return results, err
}

func (r *variantAttributeValueRepo) Update(ctx context.Context, v *models.VariantAttributeValue) error {

	return r.DB.WithContext(ctx).Model(v).Updates(v).Error
}

func (r *variantAttributeValueRepo) Delete(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&models.VariantAttributeValue{}, id).Error
}
