package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type ReviewRepository interface {
	Create(ctx context.Context, review *models.Review) error
	GetByProductID(ctx context.Context, productID uint) ([]models.Review, error)
	GetAverageRating(ctx context.Context, productID uint) (float64, error)

	database.TransactionManager
}
