package database

import (
	"github.com/minand-mohan/library-app-api/database/models"
	"github.com/minand-mohan/library-app-api/utils"
	"gorm.io/gorm"
)

func Migrate(repo *gorm.DB) {
	log := utils.NewLogger()
	log.Info("Migrating database")
	repo.AutoMigrate(&models.User{})
}
