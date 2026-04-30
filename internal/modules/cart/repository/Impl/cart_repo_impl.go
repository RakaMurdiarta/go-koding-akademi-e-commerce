package impl

import (
	"context"
	"errors"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/cart/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
	"gorm.io/gorm"
)

type cartRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewCartRepository(db *database.TransactionManagerImpl) repository.CartRepository {
	return &cartRepositoryImpl{TransactionManagerImpl: db}
}

func (r *cartRepositoryImpl) AddToCart(ctx context.Context, cart *models.Cart) error {
	return r.GetTx(ctx).Create(cart).Error
}

func (r *cartRepositoryImpl) GetCartByUserID(ctx context.Context, userID uint) ([]models.Cart, error) {
	var carts []models.Cart

	err := r.GetTx(ctx).
		Preload("Variant.Product.Shop").
		Preload("Variant.Inventory").
		Preload("Variant.AttributeValues.AttributeValue.Attribute").
		Where("user_id = ?", userID).
		Find(&carts).Error
	return carts, err
}

func (r *cartRepositoryImpl) GetCartByUserAndVariant(ctx context.Context, userID uint, variantID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.GetTx(ctx).Where("user_id = ? AND variant_id = ?", userID, variantID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepositoryImpl) UpdateQuantity(ctx context.Context, cartID uint, qty int) error {
	return r.GetTx(ctx).Model(&models.Cart{}).Where("id = ?", cartID).Update("quantity", qty).Error
}

func (r *cartRepositoryImpl) DeleteCartItem(ctx context.Context, userID uint, cartID uint) error {
	return r.GetTx(ctx).Where("id = ? AND user_id = ?", cartID, userID).Delete(&models.Cart{}).Error
}

func (r *cartRepositoryImpl) ClearCart(ctx context.Context, userID uint) error {
	return r.GetTx(ctx).
		Unscoped().
		Where("user_id = ?", userID).
		Delete(&models.Cart{}).Error
}

func (r *cartRepositoryImpl) GetCartByIDs(ctx context.Context, userID uint, ids []uint) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.GetTx(ctx).
		Preload("Variant.Product.Shop").
		Preload("Variant.Inventory").
		Preload("Variant.AttributeValues.AttributeValue.Attribute").
		Where("user_id = ? AND id IN ?", userID, ids).
		Find(&carts).Error
	return carts, err
}

func (r *cartRepositoryImpl) DeleteMultiple(ctx context.Context, userID uint, ids []uint) error {
	return r.GetTx(ctx).
		Unscoped().
		Where("user_id = ? AND id IN ?", userID, ids).
		Delete(&models.Cart{}).Error
}
