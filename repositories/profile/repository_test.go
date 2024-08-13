package profile

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
	repo   ProfileRepository
)

func setup() {
	var err error
	testDb, err = database.Connect(&cfg)
	if err != nil {
		panic("failed to connect to database")
	}
	testDb.AutoMigrate(&dbmodel.Account{}, &dbmodel.Profile{})
	repo = NewRepository(testDb)
}

func teardown() {
	if testDb != nil {
		testDb.Migrator().DropTable(&dbmodel.Account{}, &dbmodel.Profile{})
		testDb = nil
	}
	repo = nil
}

func TestInsertProfile(t *testing.T) {
	setup()
	defer teardown()

	// Create a new profile
	profile := dbmodel.Profile{
		FirstName: "John",
		LastName:  "Doe",
		AccountID: 1,
	}

	// Insert the profile into the repository
	insertedID, err := repo.Insert(context.Background(), profile)
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	// Retrieve the profile from the repository by ID
	insertedProfile, err := repo.FindByID(context.Background(), insertedID)
	assert.NoError(t, err)
	assert.Equal(t, profile.FirstName, insertedProfile.FirstName)
	assert.Equal(t, profile.LastName, insertedProfile.LastName)
}

func TestFindByAccountID(t *testing.T) {
	setup()
	defer teardown()

	// Create a new profile
	profile := dbmodel.Profile{
		FirstName: "John",
		LastName:  "Doe",
		AccountID: 1,
	}

	// Insert the profile into the repository
	insertedID, err := repo.Insert(context.Background(), profile)
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	// Retrieve the profile from the repository by account ID
	insertedProfile, err := repo.FindByAccountID(context.Background(), profile.AccountID)
	assert.NoError(t, err)
	assert.Equal(t, profile.FirstName, insertedProfile.FirstName)
	assert.Equal(t, profile.LastName, insertedProfile.LastName)
}
