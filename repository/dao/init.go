package dao

import (
	"github.com/asynccnu/be-infoSum/repository/model"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&model.InfoSum{})
}
