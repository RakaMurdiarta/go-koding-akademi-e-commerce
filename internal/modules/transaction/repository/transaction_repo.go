package repository

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
)

type ITransaction interface {
	GetByID(ctx context.Context, ID uint) (*models.Transaction, error)
}
