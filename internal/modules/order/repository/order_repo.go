package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type OrderRepository interface {
	CreateTransaction(ctx context.Context, tx *models.Transaction) error
	CreateOrder(ctx context.Context, order *models.Order) error
	CreateOrderItems(ctx context.Context, items []models.OrderItem) error

	ReserveStock(ctx context.Context, variantID uint, qty int) error
	FinalizeStock(ctx context.Context, variantID uint, qty int) error
	ReleaseStock(ctx context.Context, variantID uint, qty int) error

	GetOrderByID(ctx context.Context, orderID uint) (*models.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uint) ([]models.Order, error)
	GetOrdersByShopID(ctx context.Context, shopID uint) ([]models.Order, error)
	GetOrdersByTransactionID(ctx context.Context, txID uint) ([]models.Order, error)

	UpdateStatus(ctx context.Context, orderID uint, status string) error
	UpdateStatusWithTracking(ctx context.Context, orderID uint, status string, tracking string) error
	UpdateTransactionStatus(ctx context.Context, txID uint, status string) error
	UpdateTransactionToken(ctx context.Context, txID uint, token string) error

	database.TransactionManager
}
