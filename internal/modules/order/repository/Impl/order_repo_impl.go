package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/order/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	*database.TransactionManagerImpl
}

func NewOrderRepository(db *database.TransactionManagerImpl) repository.OrderRepository {
	return &orderRepositoryImpl{
		TransactionManagerImpl: db,
	}
}

func (r *orderRepositoryImpl) CreateTransaction(ctx context.Context, tx *models.Transaction) error {
	return r.GetTx(ctx).Create(tx).Error
}

func (r *orderRepositoryImpl) CreateOrder(ctx context.Context, order *models.Order) error {
	return r.GetTx(ctx).Create(order).Error
}

func (r *orderRepositoryImpl) CreateOrderItems(ctx context.Context, items []models.OrderItem) error {
	return r.GetTx(ctx).Create(&items).Error
}

func (r *orderRepositoryImpl) ReserveStock(ctx context.Context, variantID uint, qty int) error {
	result := r.GetTx(ctx).Model(&models.Inventory{}).
		Where("variant_id = ? AND (total_stock - reserved_stock) >= ?", variantID, qty).
		Updates(map[string]interface{}{
			"reserved_stock": gorm.Expr("reserved_stock + ?", qty),
			"version":        gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insufficient stock")
	}
	return nil
}

func (r *orderRepositoryImpl) FinalizeStock(ctx context.Context, variantID uint, qty int) error {
	result := r.GetTx(ctx).Model(&models.Inventory{}).
		Where("variant_id = ? AND reserved_stock >= ?", variantID, qty).
		Updates(map[string]interface{}{
			"total_stock":    gorm.Expr("total_stock - ?", qty),
			"reserved_stock": gorm.Expr("reserved_stock - ?", qty),
			"version":        gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("failed to finalize stock: insufficient reserved amount for variant %d", variantID)
	}
	return nil
}

func (r *orderRepositoryImpl) ReleaseStock(ctx context.Context, variantID uint, qty int) error {
	return r.GetTx(ctx).Model(&models.Inventory{}).
		Where("variant_id = ? AND reserved_stock >= ?", variantID, qty).
		Updates(map[string]interface{}{
			"reserved_stock": gorm.Expr("reserved_stock - ?", qty),
			"version":        gorm.Expr("version + 1"),
		}).Error
}

func (r *orderRepositoryImpl) UpdateStatus(ctx context.Context, orderID uint, status string) error {
	return r.GetTx(ctx).Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *orderRepositoryImpl) UpdateStatusWithTracking(ctx context.Context, orderID uint, status string, tracking string) error {
	return r.GetTx(ctx).Model(&models.Order{}).Where("id = ?", orderID).
		Updates(map[string]interface{}{
			"status":          status,
			"tracking_number": tracking,
		}).Error
}

func (r *orderRepositoryImpl) UpdateTransactionStatus(ctx context.Context, txID uint, status string) error {
	return r.GetTx(ctx).Model(&models.Transaction{}).Where("id = ?", txID).Update("payment_status", status).Error
}

func (r *orderRepositoryImpl) UpdateTransactionToken(ctx context.Context, txID uint, token string) error {
	return r.GetTx(ctx).Model(&models.Transaction{}).Where("id = ?", txID).Update("snap_token", token).Error
}

func (r *orderRepositoryImpl) GetOrderByID(ctx context.Context, orderID uint) (*models.Order, error) {
	var order models.Order
	err := r.GetTx(ctx).
		Preload("Transaction").
		Preload("Shop").
		Preload("OrderItems.Variant.Product.Images").
		First(&order, orderID).Error
	return &order, err
}

func (r *orderRepositoryImpl) GetOrdersByUserID(ctx context.Context, userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.GetTx(ctx).
		Preload("Shop").
		Preload("OrderItems.Variant.Product.Images").
		Preload("OrderItems.Variant.AttributeValues.AttributeValue.Attribute").
		Joins("JOIN transactions ON transactions.id = orders.transaction_id").
		Where("transactions.user_id = ?", userID).
		Order("orders.created_at DESC").Find(&orders).Error
	return orders, err
}

func (r *orderRepositoryImpl) GetOrdersByShopID(ctx context.Context, shopID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.GetTx(ctx).
		Preload("OrderItems.Variant.Product").
		Preload("OrderItems.Variant.AttributeValues.AttributeValue.Attribute").
		Preload("Transaction.User").
		Where("shop_id = ?", shopID).
		Order("created_at DESC").Find(&orders).Error
	return orders, err
}

func (r *orderRepositoryImpl) GetOrdersByTransactionID(ctx context.Context, txID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.GetTx(ctx).Preload("OrderItems").Where("transaction_id = ?", txID).Find(&orders).Error
	return orders, err
}
