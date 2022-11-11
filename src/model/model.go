package model

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateConnection() *gorm.DB {
	conn, err := gorm.Open(sqlite.Open("mysqlitedatabase.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = conn
	return conn
}

func Migrations(conn *gorm.DB) {
	// Migrate the schema
	conn.AutoMigrate(Gorilla{})
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
