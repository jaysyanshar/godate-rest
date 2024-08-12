package db

import (
	"fmt"
	"sync"

	"github.com/jaysyanshar/godate-rest/config"
	"gorm.io/gorm"

	// database drivers
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
)

var once sync.Once
var db *Database

type Database struct {
	*gorm.DB
}

func Connect(cfg *config.Config) (db *Database, err error) {
	dialector, err := buildDialector(cfg)
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		// connect to the database using GORM
		gormDb, e := gorm.Open(dialector, &gorm.Config{})
		if e != nil {
			err = fmt.Errorf("failed to open database: %w", e)
			return
		}

		db = &Database{gormDb}
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Close the database connection
func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Close()
}

func buildDialector(cfg *config.Config) (gorm.Dialector, error) {
	dataSourceName := buildDataSourceName(cfg)
	switch cfg.DbDriver {
	case "mysql":
		return mysql.Open(dataSourceName), nil
	case "postgres":
		return postgres.Open(dataSourceName), nil
	case "sqlite3":
		return sqlite.Open(dataSourceName), nil
	case "mssql":
		return sqlserver.Open(dataSourceName), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.DbDriver)
	}
}

func buildDataSourceName(cfg *config.Config) string {
	switch cfg.DbDriver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	case "sqlite3":
		return cfg.DbName
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	default:
		return ""
	}
}
