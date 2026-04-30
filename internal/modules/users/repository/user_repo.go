package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	IsSeller(ctx context.Context, userID uint) (bool, error)
	FindByProvider(ctx context.Context, provider common.Provider, providerID string) (*models.User, error)
	FindBySession(ctx context.Context, token string) (*models.User, error)

	database.TransactionManager
}

type UserAddressRepository interface {
	CreateAddress(ctx context.Context, address *models.UserAddress) error
	GetAddressesByUserID(ctx context.Context, userID uint) ([]models.UserAddress, error)
	GetAddressByID(ctx context.Context, id uint) (*models.UserAddress, error)
	UpdateAddress(ctx context.Context, address *models.UserAddress) error
	SetDefaultAddress(ctx context.Context, userID uint, addressID uint) error
	ResetDefaultAddress(ctx context.Context, userID uint) error
	UpdateAddressDefault(ctx context.Context, id uint, status bool) error
	DeleteAddress(ctx context.Context, addressID uint) error

	database.TransactionManager
}
