package user

import (
	"context"
	"fmt"

	db "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(ctx context.Context, user dbmodel.User) (uint, error)
	FindByID(ctx context.Context, id uint) (dbmodel.User, error)
	FindByAccountID(ctx context.Context, accountId uint) (dbmodel.User, error)
}

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, user dbmodel.User) (uint, error) {
	res := r.db.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to insert account: %w", res.Error)
	}
	return user.ID, nil
}

func (r *repository) FindByID(ctx context.Context, id uint) (dbmodel.User, error) {
	var user dbmodel.User
	res := r.db.WithContext(ctx).First(&user, id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return user, fmt.Errorf("user with ID %d not found", id)
		}
		return user, fmt.Errorf("failed to find user: %w", res.Error)
	}
	return user, nil
}

func (r *repository) FindByAccountID(ctx context.Context, accountId uint) (dbmodel.User, error) {
	var user dbmodel.User
	res := r.db.WithContext(ctx).Where("account_id = ?", accountId).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return user, fmt.Errorf("user with account ID %d not found", accountId)
		}
		return user, fmt.Errorf("failed to find user: %w", res.Error)
	}
	return user, nil
}
