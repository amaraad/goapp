package database

import (
	"fmt"

	"github.com/amaraad/goapp/graph/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	db, err := gorm.Open("postgres", getDatabaseUrl())
	return db, err
}

func RunMigrations(db *gorm.DB) {
	if !db.HasTable(&model.Question{}) {
		db.CreateTable(&model.Question{})
	}
	if !db.HasTable(&model.Choice{}) {
		db.CreateTable(&model.Choice{})
		db.Model(&model.Choice{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
	}
}