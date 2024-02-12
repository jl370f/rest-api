package repository

import (
	"errors"
	"golang-api/apiError"
	"golang-api/data/request"
	"golang-api/model"

	"gorm.io/gorm"
)

type Repository interface {
	Save(student model.Students)
	Update(student model.Students)
	Delete(studentId int)
	FindById(studentId int) (student model.Students, err error)
	FindAll() []model.Students
}

type repository struct {
	Db *gorm.DB
}

func NewRepository(Db *gorm.DB) Repository {
	return &repository{Db: Db}
}

// Delete
func (s *repository) Delete(studentId int) {
	var student model.Students
	result := s.Db.Where("id = ?", studentId).Delete(&student)
	apiError.HandleError(result.Error)
}

// FindAll
func (s *repository) FindAll() []model.Students {
	var student []model.Students
	result := s.Db.Find(&student)
	apiError.HandleError(result.Error)
	return student
}

// FindById
func (s *repository) FindById(studentId int) (student model.Students, err error) {
	result := s.Db.Find(&student, studentId)
	if result != nil {
		return student, nil
	} else {
		return student, errors.New("student is not found")
	}
}

// Save
func (s *repository) Save(student model.Students) {
	result := s.Db.Create(&student)
	apiError.HandleError(result.Error)
}

// Update
func (s *repository) Update(student model.Students) {
	var update = request.UpdateRequest{
		Id:   student.Id,
		Name: student.Name,
	}
	result := s.Db.Model(&student).Updates(update)
	apiError.HandleError(result.Error)
}
