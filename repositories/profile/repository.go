package profile

import (
	"context"
	"fmt"

	db "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Insert(ctx context.Context, profile dbmodel.Profile) (uint, error)
	FindByID(ctx context.Context, id uint) (dbmodel.Profile, error)
	FindByAccountID(ctx context.Context, accountId uint) (dbmodel.Profile, error)
}

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) ProfileRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, profile dbmodel.Profile) (uint, error) {
	res := r.db.WithContext(ctx).Create(&profile)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to insert account: %w", res.Error)
	}
	return profile.ID, nil
}

func (r *repository) FindByID(ctx context.Context, id uint) (dbmodel.Profile, error) {
	var profile dbmodel.Profile
	res := r.db.WithContext(ctx).First(&profile, id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return profile, fmt.Errorf("profile with ID %d not found", id)
		}
		return profile, fmt.Errorf("failed to find profile: %w", res.Error)
	}
	return profile, nil
}

func (r *repository) FindByAccountID(ctx context.Context, accountId uint) (dbmodel.Profile, error) {
	var profile dbmodel.Profile
	res := r.db.WithContext(ctx).Where("account_id = ?", accountId).First(&profile)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return profile, fmt.Errorf("profile with account ID %d not found", accountId)
		}
		return profile, fmt.Errorf("failed to find profile: %w", res.Error)
	}
	return profile, nil
}
