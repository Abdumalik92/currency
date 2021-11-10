package db

import (
	"curr/models"
	"curr/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var database *gorm.DB

func initDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", utils.AppSettings.PostgresParams.DataBase)
	if err != nil {
		log.Fatal("Couldn't connect to database", err.Error())
	}
	db.AutoMigrate(&models.Banks{}, &models.Currencies{}, &models.Colors{})
	// enabling gorm log mode, used for debugging
	db.LogMode(true)

	db.SingularTable(true)

	return db
}

//Creates connection to database
func StartDbConnection() {
	database = initDB()
}

//func for getting db conn globally
func GetDBConn() *gorm.DB {
	return database
}
