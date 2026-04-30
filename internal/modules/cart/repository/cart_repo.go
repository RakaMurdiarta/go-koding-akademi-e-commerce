package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type CartRepository interface {
	AddToCart(ctx context.Context, cart *models.Cart) error
	GetCartByUserAndVariant(ctx context.Context, userID uint, variantID uint) (*models.Cart, error)
	GetCartByUserID(ctx context.Context, userID uint) ([]models.Cart, error)
	UpdateQuantity(ctx context.Context, cartID uint, qty int) error
	DeleteCartItem(ctx context.Context, userID uint, cartID uint) error
	ClearCart(ctx context.Context, userID uint) error
	GetCartByIDs(ctx context.Context, userID uint, ids []uint) ([]models.Cart, error)
	DeleteMultiple(ctx context.Context, userID uint, ids []uint) error

	database.TransactionManager
}
