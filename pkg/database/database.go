package database

import (
	"fmt"
	"log"
	"time"

	"github.com/saddmm/coba-fiber/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	maxRetries    = 5
	retryInterval = 5 * time.Second
)

func ConnectDB(conf config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Host, conf.User, conf.Pass, conf.Name, conf.Port)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < maxRetries; i++ {
		log.Printf("Mencoba terhubung ke database... (Percobaan %d/%d)", i+1, maxRetries)

		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Gagal mendapatkan instance database: %v", err)
			time.Sleep(retryInterval)
			continue
		}
		err = sqlDB.Ping()
		if err == nil {
			log.Println("✅ Koneksi database berhasil!")
			return db, nil
		}

		log.Printf("Gagal terhubung: %v. Mencoba lagi dalam %v...", err, retryInterval)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("gagal terhubung ke database setelah %d percobaan: %w", maxRetries, err)
}
