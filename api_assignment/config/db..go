package config

import (
	"api_assignment/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	DB *gorm.DB
}

func init() {
	InitDB()
}
func InitDB() error {
	var err error
	dsn := "host=localhost user=username password=password dbname=assignment port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	if DB == nil {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		DB.Create(&models.SuperAdmin{Name: "nishant", Email: "nishant@example.com", Password: "password@1234"})
	}
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.SuperAdmin{}, &models.LoginActivity{})

	if err != nil {
		return err
	}

	return nil
}
