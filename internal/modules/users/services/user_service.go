package services

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/users/delivery"
)

type UserService interface {
	AddAddress(ctx context.Context, userID uint, req *delivery.AddressRequest) error
	SetDefaultAddress(ctx context.Context, userID uint, addressID uint) error
	DeleteAddress(ctx context.Context, userID uint, addressID uint) error
	GetUserAddresses(ctx context.Context, userID uint) ([]delivery.AddressResponse, error)
	UpdateAddress(ctx context.Context, userID uint, addressID uint, req delivery.AddressRequest) error
}
