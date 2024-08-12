package account

import (
	"context"
	"fmt"

	db "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Insert(ctx context.Context, account dbmodel.Account) (uint, error)
	FindByID(ctx context.Context, id uint) (dbmodel.Account, error)
}

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) AccountRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, account dbmodel.Account) (uint, error) {
	res := r.db.WithContext(ctx).Create(&account)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to insert account: %w", res.Error)
	}
	return account.ID, nil
}

func (r *repository) FindByID(ctx context.Context, id uint) (dbmodel.Account, error) {
	var account dbmodel.Account
	res := r.db.WithContext(ctx).First(&account, id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return account, fmt.Errorf("account with ID %d not found", id)
		}
		return account, fmt.Errorf("failed to find account: %w", res.Error)
	}
	return account, nil
}
