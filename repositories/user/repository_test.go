package user

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
	repo   UserRepository
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
	if testDb != nil {
		testDb.Migrator().DropTable(&dbmodel.Account{}, &dbmodel.User{})
		testDb = nil
	}
	repo = nil
}

func TestInsertUser(t *testing.T) {
	setup()
	defer teardown()

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
}

func TestFindByAccountID(t *testing.T) {
	setup()
	defer teardown()

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
}
