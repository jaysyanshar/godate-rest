package user

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
	repo   UserRepository
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

func TestInsertUser(t *testing.T) {
	// Create a new user
	user := dbmodel.User{
		FirstName: "John",
		LastName:  "Doe",
		AccountID: 1,
	}

	// Insert the user into the repository
	insertedID, err := repo.Insert(context.Background(), user)
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	// Retrieve the user from the repository by ID
	insertedUser, err := repo.FindByID(context.Background(), insertedID)
	assert.NoError(t, err)
	assert.Equal(t, user.FirstName, insertedUser.FirstName)
	assert.Equal(t, user.LastName, insertedUser.LastName)

	defer end()
}

func TestFindByAccountID(t *testing.T) {
	// Create a new user
	user := dbmodel.User{
		FirstName: "John",
		LastName:  "Doe",
		AccountID: 1,
	}

	// Insert the user into the repository
	insertedID, err := repo.Insert(context.Background(), user)
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	// Retrieve the user from the repository by account ID
	insertedUser, err := repo.FindByAccountID(context.Background(), user.AccountID)
	assert.NoError(t, err)
	assert.Equal(t, user.FirstName, insertedUser.FirstName)
	assert.Equal(t, user.LastName, insertedUser.LastName)

	defer end()
}
