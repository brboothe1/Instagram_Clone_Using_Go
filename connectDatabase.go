package main

// connectDatabase.go uses environment variables to connect our program to the PostgreSQL database
// This adds security, and can be ported to new servers with an env.bat file

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)
var db *gorm.DB
var err error

func ConnectDatabase() {
	// Loading environment variables to connect to postgres database securely

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Database connection string using loaded environment variables
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// Opening connection to database
	db, err := gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database.")
	}

	// Close connection to database when the main function finishes
	defer db.Close()

	// AutoMigrate updates tables in the database with new column fields if they are added to their struct in structs.go
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Photo{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Like{})
	db.AutoMigrate(&Newsfeed{})
	db.AutoMigrate(&Story{})
	db.AutoMigrate(&DirectMessage{})
	db.AutoMigrate(&ChatRoom{})

	db.Create(user1)
	/*for idx := range photo {
		db.Create(&photo[idx])
	}*/


	var route = mux.NewRouter()
	AddRoutes(route)

	http.ListenAndServe(":8080", route)
}
