package migrations

import (
	"github.com/amaraad/goapp/db"
	"github.com/amaraad/goapp/graph/model"
	_ "github.com/amaraad/goapp/graph/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Refactor Migrate
func Migrate() {
	Choice := &model.Choice{}
	Question := &model.Question{}
	database.DB.AutoMigrate(&Choice, &Question)

}
// Delete Migrate transactions