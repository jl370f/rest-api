package service

import (
	"golang-api/apiError"
	"golang-api/data/request"
	"golang-api/data/response"
	"golang-api/model"
	"golang-api/repository"
)

type Service interface {
	Create(student request.CreateRequest)
	Update(student request.UpdateRequest)
	Delete(studentId int)
	FindById(studentId int) response.StudentResponse
	FindAll() []response.StudentResponse
}

type service struct {
	Repository repository.Repository
}

func NewService(studentRepository repository.Repository) Service {
	return &service{
		Repository: studentRepository,
	}
}

// Create Service
func (s *service) Create(student request.CreateRequest) {
	studentModel := model.Students{
		Name: student.Name,
	}
	s.Repository.Save(studentModel)
}

// Delete Service
func (s *service) Delete(studentId int) {
	s.Repository.Delete(studentId)
}

// FindAll Service
func (s *service) FindAll() []response.StudentResponse {
	result := s.Repository.FindAll()

	students := []response.StudentResponse{}
	for _, value := range result {
		s := response.StudentResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		students = append(students, s)
	}

	return students
}

// FindById Service
func (s *service) FindById(studentId int) response.StudentResponse {
	studentData, err := s.Repository.FindById(studentId)
	apiError.HandleError(err)

	studentResponse := response.StudentResponse{
		Id:   studentData.Id,
		Name: studentData.Name,
	}
	return studentResponse
}

// Update Service
func (s *service) Update(student request.UpdateRequest) {
	studentData, err := s.Repository.FindById(student.Id)
	apiError.HandleError(err)
	studentData.Name = student.Name
	s.Repository.Update(studentData)
}
