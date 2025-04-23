package util

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
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

func NewPostgresDB() *gorm.DB {
	dsn := "host=" + config.DB_POSTGRES_CONTAINER.DB_HOST + " user=" + config.DB_POSTGRES_CONTAINER.DB_USER + " password=" + config.DB_POSTGRES_CONTAINER.DB_PASSWORD + " dbname=" + config.DB_POSTGRES_CONTAINER.DB_NAME + " port=" + config.DB_POSTGRES_CONTAINER.DB_PORT + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	return db
}

func NewPostgresDBLocal() *gorm.DB {
	dsn := "host=" + config.DB_POSTGRES_LOCAL.DB_HOST + " user=" + config.DB_POSTGRES_LOCAL.DB_USER + " password=" + config.DB_POSTGRES_LOCAL.DB_PASSWORD + " dbname=" + config.DB_POSTGRES_LOCAL.DB_NAME + " port=" + config.DB_POSTGRES_LOCAL.DB_PORT + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	return db
}

func SetupDatabaseConnection(ctx context.Context, SGBD string) (*gorm.DB, *sql.DB) {
	var db *gorm.DB

	switch SGBD {
	case POSTGRES:
		db = NewPostgresDB()
	case NEON:
		db = NeonConnection()
	case LOCAL:
		db = NewPostgresDBLocal()
	default:
		NewLogger(Logger{
			Context: ctx,
			Code:    exceptions.RFC400_CODE,
			Message: "Invalid SGBD type provided: " + SGBD,
			From:    "SetupDatabaseConnection",
			Layer:   LoggerLayers.CONFIGURATION,
			TypeLog: LoggerTypes.ERROR,
		})

		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		NewLogger(Logger{
			Context: ctx,
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "SetupDatabaseConnection",
			Layer:   LoggerLayers.CONFIGURATION,
			TypeLog: LoggerTypes.ERROR,
		})

		panic("failed to connect database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	NewLogger(Logger{
		Context: ctx,
		Code:    exceptions.RFC200_CODE,
		Message: "Database connection successfully configured",
		From:    "SetupDatabaseConnection",
		Layer:   LoggerLayers.CONFIGURATION,
		TypeLog: LoggerTypes.INFO,
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

func Shutdown(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting DB instance: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}

func NeonConnection() *gorm.DB {
	dsn := "postgresql://" + config.DB_NEON.DB_USER + ":" + config.DB_NEON.DB_PASSWORD + "@" + config.DB_NEON.DB_HOST + "/" + config.DB_NEON.DB_NAME + "?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: NewLoggerGorm(),
	})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	return db
}
