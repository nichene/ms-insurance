package postgres

import (
	"context"
	"fmt"
	"log"
	"ms-insurance/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(ctx context.Context, cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("DB - Error openning connection to database", err)
		return nil, err
	}

	sqldb, _ := db.DB()
	err = sqldb.PingContext(ctx)
	if err != nil {
		log.Fatal("DB - Error on ping to database", err)
		return nil, err
	}

	log.Default().Print("DB - Database loaded")
	return db, nil
}
