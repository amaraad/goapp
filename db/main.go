package database

import (
	"fmt"

	"github.com/amaraad/goapp/helpers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
	sslmode  string
}

var config = dbConfig{"localhost", 5432, "postgres", "goapp_db", "postgres", "disable"}

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.host, config.port, config.user, config.dbname, config.password, config.sslmode)
}

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  getDatabaseUrl(),
		PreferSimpleProtocol: false,
	}), &gorm.Config{})
	helpers.HandleErr(err)
	// database.DB().SetMaxIdleConns(10)
	// database.DB().SetMaxOpenConns(100)
	DB = database

}
