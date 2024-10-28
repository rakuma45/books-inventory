package db

import (
	"books-project/models"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "database/sql"
	_ "github.com/lib/pq"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)
	return db
}

func Migrate(db *gorm.DB)  {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}
	if db.Find(&data).RecordNotFound(){
		seederBook(db)
	}
}

func seederBook(db *gorm.DB)  {
	data := []models.Books{{ 
		Title: "Book 1",
		Author: "Author 1",
		Description: "Description 1",
		Stock: 10,
	 },{ 
		Title: "Book 2",
		Author: "Author 2",
		Description: "Description 2",
		Stock: 15,
	 },{ 
		Title: "Book 3",
		Author: "Author 3",
		Description: "Description 3",
		Stock: 7,
	 },{ 
		Title: "Book 4",
		Author: "Author 4",
		Description: "Description 4",
		Stock: 12,
	 }}

	 for _,v:= range data {
		db.Create(&v)
	 }
}