package impl

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/review/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type reviewRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewReviewRepository(db *database.TransactionManagerImpl) repository.ReviewRepository {
	return &reviewRepositoryImpl{
		TransactionManagerImpl: db,
	}
}

func (r *reviewRepositoryImpl) Create(ctx context.Context, review *models.Review) error {
	return r.GetTx(ctx).Create(review).Error
}

func (r *reviewRepositoryImpl) GetByProductID(ctx context.Context, productID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.GetTx(ctx).
		Preload("User").
		Where("product_id = ?", productID).
		Order("created_at DESC").
		Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepositoryImpl) GetAverageRating(ctx context.Context, productID uint) (float64, error) {
	var avg float64
	err := r.GetTx(ctx).Model(&models.Review{}).
		Where("product_id = ?", productID).
		Select("COALESCE(AVG(rating), 0)").
		Row().Scan(&avg)
	return avg, err
}
