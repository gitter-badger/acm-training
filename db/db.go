package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/uestc-acm/acm-training/config"
	"github.com/uestc-acm/acm-training/util"
)

var db *gorm.DB

// Init - initialize the database, and the destory process will be setup at main.go because
// it should be closed after the server being shutdown.
func Init() {
	var err error
	fmt.Println("Opening database at", config.GetDbFile())
	db, err = gorm.Open("sqlite3", config.GetDbFile())
	util.CheckArgument(err == nil, "Error occurs when opening database file:", err)
}

// DB - Gets the db instance of the server.
func DB() *gorm.DB {
	return db
}
