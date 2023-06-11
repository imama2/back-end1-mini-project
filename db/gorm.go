package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB(cfg *Config) (*gorm.DB, error) {
	DNS := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Protocol,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	//dsn := fmt.Sprintf("root:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=UTC", os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
		os.Exit(1)
	}
	return db, nil
}
