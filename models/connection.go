package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Fungsi untuk inisialisasi koneksi database
func InitDB() {
	dsn := "user=postgres password=1sampai10! dbname=SampTest host=localhost port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database: ", err)
	} else {
		fmt.Println("Connected to PostgreSQL database")
	}
}

// Fungsi untuk mendapatkan instance database *gorm.DB
func GetDB() *gorm.DB {
	return DB
}
