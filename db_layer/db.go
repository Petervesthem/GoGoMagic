package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Card struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	SetName  string
	Rarity   string
	TypeLine string
	Count    uint // Additional count column
}

//id and count table

func main() {
	db, err := gorm.Open(sqlite.Open("mtg_cards.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	if db.AutoMigrate(&Card{}); err != nil {
		log.Fatal("failed to auto migrate db", err)
	}

	if err := db.Exec("ALTER TABLE cards ADD COLUMN count INTEGER DEFAULT 0").Error; err != nil {
		log.Fatal("failed to count to table", err)
	}

	log.Println("Successful migration")

}
