package db

import (
	"os"
	"testing"

	"github.com/jaysyanshar/godate-rest/config"
)

func TestConnect(t *testing.T) {
	cfg := &config.Config{
		DbDriver: "sqlite3",
		DbName:   "test_db",
	}

	db, err := Connect(cfg)
	if err != nil {
		t.Errorf("Failed to connect to the database: %v", err)
	}

	defer os.Remove("test_db")
	defer db.Close()
}

func TestBuildDataSourceName(t *testing.T) {
	testCases := []struct {
		name                   string
		dbDriver               string
		expectedDataSourceName string
		expectError            bool
	}{
		{
			name:                   "MySQL",
			dbDriver:               "mysql",
			expectedDataSourceName: "username:password@tcp(localhost:5432)/test_db",
			expectError:            false,
		},
		{
			name:                   "PostgreSQL",
			dbDriver:               "postgres",
			expectedDataSourceName: "host=localhost port=5432 user=username password=password dbname=test_db sslmode=disable",
			expectError:            false,
		},
		{
			name:                   "SQL Server",
			dbDriver:               "mssql",
			expectedDataSourceName: "sqlserver://username:password@localhost:5432?database=test_db",
			expectError:            false,
		},
		{
			name:                   "SQLite3",
			dbDriver:               "sqlite3",
			expectedDataSourceName: "test_db",
			expectError:            false,
		},
		{
			name:                   "Unknown",
			dbDriver:               "unknown",
			expectedDataSourceName: "",
			expectError:            true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &config.Config{
				DbDriver:   tc.dbDriver,
				DbHost:     "localhost",
				DbPort:     "5432",
				DbUser:     "username",
				DbPassword: "password",
				DbName:     "test_db",
			}

			dataSourceName := buildDataSourceName(cfg)

			if dataSourceName != tc.expectedDataSourceName && !tc.expectError {
				t.Errorf("Unexpected data source name. Got: %s, Expected: %s", dataSourceName, tc.expectedDataSourceName)
			}
		})
	}
}

func TestBuildDialector(t *testing.T) {
	testCases := []struct {
		name        string
		dbDriver    string
		expected    string
		expectError bool
	}{
		{
			name:        "MySQL",
			dbDriver:    "mysql",
			expected:    "mysql",
			expectError: false,
		},
		{
			name:        "PostgreSQL",
			dbDriver:    "postgres",
			expected:    "postgres",
			expectError: false,
		},
		{
			name:        "SQL Server",
			dbDriver:    "mssql",
			expected:    "sqlserver",
			expectError: false,
		},
		{
			name:        "SQLite3",
			dbDriver:    "sqlite3",
			expected:    "sqlite",
			expectError: false,
		},
		{
			name:        "Unknown",
			dbDriver:    "unknown",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &config.Config{
				DbDriver: tc.dbDriver,
			}

			dialector, err := buildDialector(cfg)

			if err != nil && !tc.expectError {
				t.Errorf("Unexpected error. Got: %v", err)
			}

			if !tc.expectError && dialector != nil && dialector.Name() != tc.expected {
				t.Errorf("Unexpected dialector. Got: %s, Expected: %s", dialector.Name(), tc.dbDriver)
			}
		})
	}
}
