package account

import (
	"context"
	"testing"

	"github.com/jaysyanshar/godate-rest/config"
	database "github.com/jaysyanshar/godate-rest/internal/database"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
	"github.com/stretchr/testify/assert"
)

var (
	cfg config.Config = config.Config{
		DbDriver: "sqlite3",
		DbName:   ":memory:",
	}
	testDb *database.Database
	repo   AccountRepository
)

func setup() {
	var err error
	testDb, err = database.Connect(&cfg)
	if err != nil {
		panic("failed to connect to database")
	}
	testDb.AutoMigrate(&dbmodel.Account{}, &dbmodel.User{})
	repo = NewRepository(testDb)
}

func teardown() {
	testDb.Migrator().DropTable(&dbmodel.Account{}, &dbmodel.User{})
	testDb.Close()
	testDb = nil
	repo = nil
}

func TestAccountRepository_Insert(t *testing.T) {
	setup()
	defer teardown()

	account := dbmodel.Account{
		Email:    "test@mail.com",
		Password: "password",
	}

	id, err := repo.Insert(context.Background(), account)
	assert.NoError(t, err)
	assert.Greater(t, id, uint(0))

	testDb.WithContext(context.Background()).Delete(&account, id)
}

func TestAccountRepository_FindByID(t *testing.T) {
	setup()
	defer teardown()

	account := dbmodel.Account{
		Email:    "test@mail.com",
		Password: "password",
	}

	id, err := repo.Insert(context.Background(), account)
	assert.NoError(t, err)
	assert.Greater(t, id, uint(0))

	foundAccount, err := repo.FindByID(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, account.Email, foundAccount.Email)
	assert.Equal(t, account.Password, foundAccount.Password)
}
