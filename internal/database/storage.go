package storage

import "github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/models"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (models.Student, error)
	GetStudents() ([]models.Student, error)
	DeleteStudent(id int64) (string, error)
	UpdateStudentById(id int64, name, email *string, age *int) (models.Student, error)
}
