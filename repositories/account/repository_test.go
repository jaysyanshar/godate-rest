package account

import (
	"context"
	"os"
	"testing"

	"github.com/jaysyanshar/godate-rest/config"
	database "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
	"github.com/stretchr/testify/assert"
)

var (
	cfg config.Config = config.Config{
		DbDriver: "sqlite3",
		DbName:   "test_db",
	}
	testDb *database.Database
	repo   AccountRepository
)

func init() {
	testDb, _ = database.Connect(&cfg)
	testDb.AutoMigrate(&dbmodel.Account{}, &dbmodel.User{})
	repo = NewRepository(testDb)
}

func end() {
	os.Remove("test_db")
	testDb.Close()
}

func TestAccountRepository_Insert(t *testing.T) {
	account := dbmodel.Account{
		Email:    "test@mail.com",
		Password: "password",
	}

	id, err := repo.Insert(context.Background(), account)

	assert.NoError(t, err)
	assert.Greater(t, id, uint(0))

	defer end()
}

func TestAccountRepository_FindByID(t *testing.T) {
	account := dbmodel.Account{
		Email:    "test@mail.com",
		Password: "password",
	}

	id, _ := repo.Insert(context.Background(), account)
	foundAccount, err := repo.FindByID(context.Background(), id)

	assert.NoError(t, err)
	assert.Equal(t, account.Email, foundAccount.Email)
	assert.Equal(t, account.Password, foundAccount.Password)

	defer end()
}
