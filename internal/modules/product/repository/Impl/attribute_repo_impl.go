package impl

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/product/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type attributeRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewAttributeRepository(db *database.TransactionManagerImpl) repository.AttributeRepository {
	return &attributeRepositoryImpl{
		TransactionManagerImpl: db,
	}
}

func (r *attributeRepositoryImpl) CreateAttribute(ctx context.Context, attribute *models.Attribute) error {
	return r.GetTx(ctx).
		Create(attribute).
		Error
}

func (r *attributeRepositoryImpl) CreateAttributeValues(ctx context.Context, attrValues []models.AttributeValue) error {
	return r.GetTx(ctx).
		Create(&attrValues).
		Error
}

func (r *attributeRepositoryImpl) GetByName(ctx context.Context, name string) (*models.Attribute, error) {
	var attr models.Attribute
	err := r.GetTx(ctx).Where("name = ?", name).First(&attr).Error
	return &attr, err
}

func (r *attributeRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.Attribute, error) {

	var attribute models.Attribute

	err := r.GetTx(ctx).
		First(&attribute, id).
		Error

	if err != nil {
		return nil, err
	}

	return &attribute, nil
}

func (r *attributeRepositoryImpl) GetAll(ctx context.Context) ([]models.Attribute, error) {

	var attributes []models.Attribute

	err := r.GetTx(ctx).
		Preload("Values").
		Order("id ASC").
		Find(&attributes).
		Error

	return attributes, err
}

func (r *attributeRepositoryImpl) Update(ctx context.Context, attribute *models.Attribute) error {
	return r.GetTx(ctx).
		Save(attribute).
		Error
}

func (r *attributeRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.GetTx(ctx).
		Delete(&models.Attribute{}, id).
		Error
}

func (r *attributeRepositoryImpl) CheckValueExists(ctx context.Context, attrID uint, value string) (bool, error) {
	var count int64
	err := r.GetTx(ctx).Model(&models.AttributeValue{}).
		Where("attribute_id = ? AND value = ?", attrID, value).
		Count(&count).Error
	return count > 0, err
}

func (r *attributeRepositoryImpl) UpdateValue(ctx context.Context, valueID uint, newValue string) error {
	return r.GetTx(ctx).Model(&models.AttributeValue{}).
		Where("id = ?", valueID).
		Update("value", newValue).Error
}
