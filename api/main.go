package main

import (
	"github.com/t2469/labor-management-system.git/models"
	"github.com/t2469/labor-management-system.git/routes"
	"log"
)

func main() {
	models.InitDB()
	if err := models.DB.AutoMigrate(&models.User{}, &models.Attendance{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	router := routes.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
