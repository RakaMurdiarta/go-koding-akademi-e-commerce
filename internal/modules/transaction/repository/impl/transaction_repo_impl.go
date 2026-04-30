package impl

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/transaction/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
)

type transactionRepoImpl struct {
	*database.TransactionManagerImpl
}

func NewTransactionRepo(db *database.TransactionManagerImpl) repository.ITransaction {
	return &transactionRepoImpl{
		TransactionManagerImpl: db,
	}
}

func (t *transactionRepoImpl) GetByID(ctx context.Context, ID uint) (*models.Transaction, error) {

	var transaction models.Transaction

	err := t.GetTx(ctx).
		First(&transaction, ID).
		Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil

}
