package impl

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/shop/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type shopRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewShopRepository(db *database.TransactionManagerImpl) repository.ShopRepository {
	return &shopRepositoryImpl{
		TransactionManagerImpl: db,
	}
}

func (r *shopRepositoryImpl) Create(ctx context.Context, shop *models.Shop) error {
	return r.GetTx(ctx).
		Create(shop).
		Error
}

func (r *shopRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.Shop, error) {

	var shop models.Shop

	err := r.GetTx(ctx).
		First(&shop, id).
		Error

	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func (r *shopRepositoryImpl) GetByUserID(ctx context.Context, userID uint) (*models.Shop, error) {

	var shop models.Shop

	err := r.GetTx(ctx).
		Where("user_id = ?", userID).
		First(&shop).
		Error

	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func (r *shopRepositoryImpl) GetAll(ctx context.Context) ([]models.Shop, error) {

	var shops []models.Shop

	err := r.GetTx(ctx).
		Order("id ASC").
		Find(&shops).
		Error

	return shops, err
}

func (r *shopRepositoryImpl) Update(ctx context.Context, shop *models.Shop) error {
	return r.GetTx(ctx).
		Save(shop).
		Error
}

func (r *shopRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.GetTx(ctx).
		Delete(&models.Shop{}, id).
		Error
}

func (r *shopRepositoryImpl) ExistsByName(ctx context.Context, name string) (bool, error) {

	var count int64

	err := r.GetTx(ctx).
		Model(&models.Shop{}).
		Where("name = ?", name).
		Count(&count).
		Error

	return count > 0, err
}
