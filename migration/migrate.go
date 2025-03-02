package migration

import (
	"auth_api/database"
	"auth_api/models"
	"fmt"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}
}