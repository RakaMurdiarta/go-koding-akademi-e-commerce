package impl

import (
	"context"
	"errors"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/users/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
	"gorm.io/gorm"
)

type userAddressRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewUserAddressRepository(db *database.TransactionManagerImpl) repository.UserAddressRepository {
	return &userAddressRepositoryImpl{TransactionManagerImpl: db}
}

func (r *userAddressRepositoryImpl) CreateAddress(ctx context.Context, address *models.UserAddress) error {
	return r.GetTx(ctx).Create(address).Error
}

func (r *userAddressRepositoryImpl) GetAddressesByUserID(ctx context.Context, userID uint) ([]models.UserAddress, error) {
	var addresses []models.UserAddress
	err := r.GetTx(ctx).Where("user_id = ?", userID).Order("created_at desc").Find(&addresses).Error
	return addresses, err
}

func (r *userAddressRepositoryImpl) GetAddressByID(ctx context.Context, id uint) (*models.UserAddress, error) {
	var address models.UserAddress
	err := r.GetTx(ctx).First(&address, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &address, nil
}

func (r *userAddressRepositoryImpl) UpdateAddress(ctx context.Context, address *models.UserAddress) error {
	return r.GetTx(ctx).Model(address).Updates(address).Error
}

func (r *userAddressRepositoryImpl) SetDefaultAddress(ctx context.Context, userID uint, addressID uint) error {
	return r.GetTx(ctx).Transaction(func(tx *gorm.DB) error {

		err := tx.Model(&models.UserAddress{}).
			Where("user_id = ?", userID).
			Update("is_default", false).Error
		if err != nil {
			return err
		}

		err = tx.Model(&models.UserAddress{}).
			Where("id = ? AND user_id = ?", addressID, userID).
			Update("is_default", true).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *userAddressRepositoryImpl) ResetDefaultAddress(ctx context.Context, userID uint) error {
	return r.GetTx(ctx).Model(&models.UserAddress{}).
		Where("user_id = ?", userID).
		Update("is_default", false).Error
}

func (r *userAddressRepositoryImpl) UpdateAddressDefault(ctx context.Context, id uint, status bool) error {
	return r.GetTx(ctx).Model(&models.UserAddress{}).
		Where("id = ?", id).
		Update("is_default", status).Error
}

func (r *userAddressRepositoryImpl) DeleteAddress(ctx context.Context, addressID uint) error {
	return r.GetTx(ctx).Delete(&models.UserAddress{}, addressID).Error
}
