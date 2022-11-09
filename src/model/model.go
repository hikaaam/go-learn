package model

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mysqlitedatabase.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrations() {
	// Migrate the schema
	DB().AutoMigrate(Gorilla{})
}

//Gorilla Table
type Gorilla struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Weight    int    `json:"weight"`
	Gender    bool   `json:"gender"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
