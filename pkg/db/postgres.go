package db

import (
	"database/sql"
	"fmt"

	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/internal/inventory"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"github.com/ryanadiputraa/inventra/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	maxOpenConns    = 60
	connMaxLifeTime = 120
	maxIdleConn     = 30
	connMaxIdleTime = 20
)

func NewPostgres(c config.Config) (db *gorm.DB, sqlDB *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return
	}

	sqlDB, err = db.DB()
	if err != nil {
		return
	}

	if err = sqlDB.Ping(); err != nil {
		return
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifeTime)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime)

	if err = runMigration(db); err != nil {
		return
	}

	return
}

func runMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&organization.Organization{},
		&organization.Member{},
		&inventory.Item{},
		&inventory.ItemPrice{},
	)
}
