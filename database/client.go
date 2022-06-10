package database

import (
	"JobWorker/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

//Config to maintain DB configuration properties
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(config Config) error {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)

	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}
func Migrate(table *model.Person) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
