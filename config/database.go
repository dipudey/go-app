package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type DatabaseConfig struct {
	DatabaseConnection string `env:"DB_CONNECTION" env-default:"mysql"`
	Host               string `env:"DB_HOST" env-default:"localhost"`
	Port               int    `env:"DB_PORT" env-default:"8889"`
	User               string `env:"DB_USER" env-default:"root"`
	Password           string `env:"DB_PASSWORD" env-default:"root"`
	Name               string `env:"DB_NAME" env-default:"goapp"`
}

// ConnectDB establishes a connection to the MySQL database using GORM.
func (dc *DatabaseConfig) ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.User,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	log.Println("MySQL connected successfully")

	return db, nil
}
