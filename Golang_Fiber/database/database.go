package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

 var err error


 dsn := "root:@tcp(127.0.0.1:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
 dialect := mysql.Open(dsn)
 DB, err = gorm.Open(dialect, &gorm.Config{})

 if err != nil {
  log.Fatal("Failed to Connect to database...")
 }

 fmt.Println("Connecting to database...")
}