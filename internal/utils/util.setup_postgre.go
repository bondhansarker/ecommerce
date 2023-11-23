package utils

import (
	"context"
	"time"

	"github.com/bondhansarker/ecommerce/internal/config"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"github.com/bondhansarker/ecommerce/internal/datasources/drivers"
	"github.com/jmoiron/sqlx"
)

var dbClient *sqlx.DB

func SetupPostgresConnection() {
	var dsn string
	switch config.AppConfig.Environment {
	case constants.EnvironmentDevelopment:
		dsn = config.AppConfig.DBPostgreDsn
	case constants.EnvironmentProduction:
		dsn = config.AppConfig.DBPostgreURL
	}

	// Setup sqlx config of postgreSQL
	config := drivers.SQLXConfig{
		DriverName:     config.AppConfig.DBPostgreDriver,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}
	var err error
	// Initialize postgreSQL connection with sqlx
	dbClient, err = config.InitializeSQLXDatabase()
	if err != nil {
		panic(err)
	}
}

func GetDbClient() *sqlx.DB {
	return dbClient
}

func BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return dbClient.BeginTxx(ctx, nil)
}

func CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func RollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}
