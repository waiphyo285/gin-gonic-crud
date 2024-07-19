package main

import (
	"neohub.asia/mod/models"
	"neohub.asia/mod/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
