// app/database/database.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path/to/your/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB = db
	DB.AutoMigrate(&models.Dentista{}, &models.Paciente{}, &models.Turno{})
}
