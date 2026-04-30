package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type ShopRepository interface {
	Create(ctx context.Context, shop *models.Shop) error
	GetByID(ctx context.Context, id uint) (*models.Shop, error)
	GetByUserID(ctx context.Context, userID uint) (*models.Shop, error)
	GetAll(ctx context.Context) ([]models.Shop, error)
	Update(ctx context.Context, shop *models.Shop) error
	Delete(ctx context.Context, id uint) error
	ExistsByName(ctx context.Context, name string) (bool, error)

	//Impl Transaction
	database.TransactionManager
}
