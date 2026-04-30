package impl

import (
	"context"
	"errors"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/users/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewUserRepository(db *database.TransactionManagerImpl) repository.UserRepository {
	return &userRepositoryImpl{
		TransactionManagerImpl: db,
	}
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	return r.GetTx(ctx).Create(user).Error
}

func (r *userRepositoryImpl) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.GetTx(ctx).Preload("Addresses").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.GetTx(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepositoryImpl) FindByProvider(ctx context.Context, provider common.Provider, providerID string) (*models.User, error) {
	var user models.User
	err := r.GetTx(ctx).Where("provider = ? AND provider_id = ?", provider, providerID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) UpdateUser(ctx context.Context, user *models.User) error {
	return r.GetTx(ctx).Model(user).Updates(user).Error
}

func (r *userRepositoryImpl) IsSeller(ctx context.Context, userID uint) (bool, error) {
	var result struct {
		ID uint
	}

	err := r.GetTx(ctx).
		Model(&models.User{}).
		Select("id").
		Where("id = ? AND role = ?", userID, "seller").
		Take(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *userRepositoryImpl) FindBySession(ctx context.Context, token string) (*models.User, error) {
	var user models.User
	err := r.GetTx(ctx).Where("session = ?", token).First(&user).Error
	return &user, err
}
