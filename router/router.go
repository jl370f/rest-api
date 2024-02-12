package router

import (
	"golang-api/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(studentsController *controller.Controller) *gin.Engine {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")
	baseRouter.GET("/students", studentsController.FindAll)
	baseRouter.GET("/students/:studentId", studentsController.FindById)
	baseRouter.POST("/students", studentsController.Create)
	baseRouter.PATCH("/students/:studentId", studentsController.Update)
	baseRouter.DELETE("/students/:studentId", studentsController.Delete)

	return router
}
