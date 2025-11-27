package storage

import "github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/models"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (models.Student, error)
}