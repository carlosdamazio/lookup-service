package db

import (
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/carlosdamazio/lookup-service/internal/models"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func SetUp(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Query{}); err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	once.Do(func() {
		DB, _ = gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{})
		SetUp(DB)
	})
	return DB
}
