package database

import "githubmbarkhanBusTracker/auth/models"

func AutoMigrate() error {
	err := DB.AutoMigrate(&models.User{})
	return err

}