package initializers

import (
	"crud/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDb(){
  var err error

   dsn := "golang:golang@tcp(0.0.0.0:3306)/apigolang?charset=utf8mb4&parseTime=True&loc=Local"

  DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil{
    panic("Erro to connect db")
  }
  err = DB.AutoMigrate(&models.UserModel{})
    if err != nil {
        panic("failed to migrate database schema")
  }

   fmt.Println("Database tables created successfully")
}
