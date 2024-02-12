package controller

import (
	"golang-api/apiError"
	"golang-api/config"
	"golang-api/data/request"
	"golang-api/data/response"
	"golang-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	studentsService service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		studentsService: service,
	}
}

// Create			godoc
// @Summary			Create a student
// @Description		Save a student data in Db.
// @Param			student body request.CreateRequest true "create a student"
// @Produce			application/json
// @Students		students
// @Success			200 {object} response.HttpResponse{}
// @Router			/students [post]
func (controller *Controller) Create(ctx *gin.Context) {
	config.Logger.Info("Create")
	createRequest := request.CreateRequest{}
	err := ctx.ShouldBindJSON(&createRequest)
	apiError.HandleError(err)

	controller.studentsService.Create(createRequest)
	httpResponse := response.HttpResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, httpResponse)
}

// Update			godoc
// @Summary			Update a student
// @Description		Update student data.
// @Param			studentId path string true "update student by id"
// @Param			student body request.CreateRequest true "Update a student"
// @Students		student
// @Produce			application/json
// @Success			200 {object} response.HttpResponse{}
// @Router			/students/{studentId} [patch]
func (controller *Controller) Update(ctx *gin.Context) {
	config.Logger.Info("Update")
	updateRequest := request.UpdateRequest{}
	err := ctx.ShouldBindJSON(&updateRequest)
	apiError.HandleError(err)

	studentId := ctx.Param("studentId")
	id, err := strconv.Atoi(studentId)
	apiError.HandleError(err)
	updateRequest.Id = id

	controller.studentsService.Update(updateRequest)

	httpResponse := response.HttpResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, httpResponse)
}

// Delete			godoc
// @Summary			Delete a student
// @Param			studentId path string true "delete a student by id"
// @Description		Remove student data by id.
// @Produce			application/json
// @Students		student
// @Success			200 {object} response.HttpResponse{}
// @Router			/students/{studentId} [delete]
func (controller *Controller) Delete(ctx *gin.Context) {
	config.Logger.Info("Delete")
	studentId := ctx.Param("studentId")
	id, err := strconv.Atoi(studentId)
	apiError.HandleError(err)
	controller.studentsService.Delete(id)

	httpResponse := response.HttpResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, httpResponse)
}

// FindById 		godoc
// @Summary			Get Single student by id.
// @Param			studentId path string true "update student by id"
// @Description		Return the tahs whoes studentId valu mathes id.
// @Produce			application/json
// @Students		student
// @Success			200 {object} response.HttpResponse{}
// @Router			/students/{studentId} [get]
func (controller *Controller) FindById(ctx *gin.Context) {
	config.Logger.Info("FindById")
	studentId := ctx.Param("studentId")
	id, err := strconv.Atoi(studentId)
	apiError.HandleError(err)

	studentResponse := controller.studentsService.FindById(id)

	httpResponse := response.HttpResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   studentResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, httpResponse)
}

// FindAll 			godoc
// @Summary			Get All student.
// @Description		Return list of students.
// @Students		students
// @Success			200 {obejct} response.HttpResponse{}
// @Router			/students [get]
func (controller *Controller) FindAll(ctx *gin.Context) {
	config.Logger.Info("FindAll")
	studentResponse := controller.studentsService.FindAll()
	httpResponse := response.HttpResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   studentResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, httpResponse)

}
