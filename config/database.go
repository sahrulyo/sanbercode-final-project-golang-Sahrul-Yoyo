/*
package config

import (
	"fmt"

	"gorm.io/driver/postgres" // Import the PostgreSQL driver
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	// Sesuaikan nilai host, dbname, user, password sesuai dengan konfigurasi Anda
	connectionStr := "host=localhost user=user dbname=user sslmode=disable password=Ulyasar10389#"
	var err error
	DB, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Jika koneksi berhasil, cetak pesan keberhasilan
	fmt.Println("Successfully connected to database")
}
*/