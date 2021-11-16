package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"github.com/amaraad/goapp/graph/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
	sslmode string
}

var config = dbConfig{"localhost", 5432, "postgres", "goapp_db", "postgres", "disable"}

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.host, config.port, config.user, config.dbname, config.password, config.sslmode)
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDatabaseUrl()), &gorm.Config{})
	return db, err
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Choice{})
	db.AutoMigrate(&model.Question{})
}