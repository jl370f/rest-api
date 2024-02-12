package main

import (
	"golang-api/config"
	"golang-api/controller"
	_ "golang-api/docs"
	"golang-api/model"
	"golang-api/repository"
	"golang-api/router"
	"golang-api/service"
)

// @title 	API
// @version	1.0
// @description API using Gin framework

// @host 	localhost:8080
// @BasePath /api
func main() {

	config.Logger.Info("Starting the golang gin api app ...")
	db := config.DatabaseConnection()

	db.Table("students").AutoMigrate(&model.Students{})

	studentsRepository := repository.NewRepository(db)

	studentsService := service.NewService(studentsRepository)

	studentsController := controller.NewController(studentsService)

	routes := router.NewRouter(studentsController)
	routes.SetTrustedProxies(nil)
	routes.Run(":8080")
}
