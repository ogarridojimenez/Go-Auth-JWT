package initializers

import "github.com/ogarridojimenez/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
