package db

import (
	"fmt"
	"github.com/weeb-vip/anime-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDatabase(cfg config.DBConfig) *DB {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s&interpolateParams=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase, cfg.SSLMode)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", cfg.Host, cfg.User, cfg.Password, cfg.DataBase, cfg.Port)
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{DB: db}
}
