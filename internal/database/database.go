package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	POSTGRES = "postgres"
	NEON     = "neon"
	LOCAL    = "local"
)

func NewLoggerGorm() gormLogger.Interface {
	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormLogger.Config{
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			SlowThreshold:             200 * time.Millisecond,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	return newLogger
}

func NewPostgresDB(ctx context.Context) *gorm.DB {
	dsn := "host=" + config.DB_POSTGRES_CONTAINER.DB_HOST + " user=" + config.DB_POSTGRES_CONTAINER.DB_USER + " password=" + config.DB_POSTGRES_CONTAINER.DB_PASSWORD + " dbname=" + config.DB_POSTGRES_CONTAINER.DB_NAME + " port=" + config.DB_POSTGRES_CONTAINER.DB_PORT + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "NewPostgresDB",
			Layer:   logging.LoggerLayers.CONFIGURATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil
	}
	return db
}

func NewPostgresDBLocal(ctx context.Context) *gorm.DB {
	dsn := "host=" + config.DB_POSTGRES_LOCAL.DB_HOST + " user=" + config.DB_POSTGRES_LOCAL.DB_USER + " password=" + config.DB_POSTGRES_LOCAL.DB_PASSWORD + " dbname=" + config.DB_POSTGRES_LOCAL.DB_NAME + " port=" + config.DB_POSTGRES_LOCAL.DB_PORT + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "NewPostgresDBLocal",
			Layer:   logging.LoggerLayers.CONFIGURATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil
	}
	return db
}

func SetupDatabaseConnection(ctx context.Context, SGBD string) (*gorm.DB, *sql.DB) {
	var db *gorm.DB

	switch SGBD {
	case POSTGRES:
		db = NewPostgresDB(ctx)
	case NEON:
		db = NeonConnection(ctx)
	case LOCAL:
		db = NewPostgresDBLocal(ctx)
	default:
		logging.NewLogger(logging.Logger{
			Context: ctx,
			Code:    exceptions.RFC400_CODE,
			Message: "Invalid SGBD type provided: " + SGBD,
			From:    "SetupDatabaseConnection",
			Layer:   logging.LoggerLayers.CONFIGURATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})

		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "SetupDatabaseConnection",
			Layer:   logging.LoggerLayers.CONFIGURATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})

		panic("failed to connect database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	logging.NewLogger(logging.Logger{
		Context: ctx,
		Code:    exceptions.RFC200_CODE,
		Message: "Database connection successfully configured",
		From:    "SetupDatabaseConnection",
		Layer:   logging.LoggerLayers.CONFIGURATION,
		TypeLog: logging.LoggerTypes.INFO,
	})

	return db, sqlDB
}

func CheckConnection(db *gorm.DB) bool {
	sqlDB, err := db.DB()
	if err != nil {
		return false
	}

	if err := sqlDB.Ping(); err != nil {
		return false
	}

	return true
}

func Shutdown(ctx context.Context, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			Code:    exceptions.RFC500_CODE,
			Message: "Error getting DB instance: " + err.Error(),
			From:    "Shutdown",
			Layer:   logging.LoggerLayers.CONFIGURATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})

		return
	}

	if err := sqlDB.Close(); err != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			Code:    exceptions.RFC500_CODE,
			Message: "Error closing database connection: " + err.Error(),
			From:    "Shutdown",
			Layer:   logging.LoggerLayers.CONFIGURATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
	}
}

func NeonConnection(ctx context.Context) *gorm.DB {
	dsn := "postgresql://" + config.DB_NEON.DB_USER + ":" + config.DB_NEON.DB_PASSWORD + "@" + config.DB_NEON.DB_HOST + "/" + config.DB_NEON.DB_NAME + "?sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: NewLoggerGorm(),
	})

	if err != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "NeonConnection",
			Layer:   logging.LoggerLayers.CONFIGURATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil
	}
	return db
}
