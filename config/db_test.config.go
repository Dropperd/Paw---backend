package config

import (
	"websiteapi/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db_Test *gorm.DB

func ConnectDB_TEST() {
	var err error

	dbUser := "root"
	dbPass := "UFP12UFP23"
	dbHost := "84.90.137.202"
	dbPort := "1355"
	dbDatabase := "PAWB_PD_TEST"
	
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	Db_Test, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = Db_Test.AutoMigrate(&entity.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func CloseDB_TEST() {
	db, err := Db_Test.DB()
	if err != nil {
		panic("failed to close database")
	}
	err = db.Close()
	if err != nil {
		return
	}
}
