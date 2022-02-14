package database

import (
	"github.com/anfelo/go-auth-template/internal/users"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&users.User{}); err != nil {
		return err
	}
	return nil
}
