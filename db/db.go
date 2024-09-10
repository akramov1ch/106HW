package db

import (
	"log"

	"106HW/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("filemetadata.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.FileMetadata{})
}

func SaveFileMetadata(metadata models.FileMetadata) {

	db.Create(&metadata)
}

func GetFileMetadata() []models.FileMetadata {
	var files []models.FileMetadata
	db.Find(&files)
	return files
}
