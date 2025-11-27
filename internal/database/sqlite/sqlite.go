package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/config"
	"github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// * Here we will store the connection of our database
type Sqlite struct {
	Db *sql.DB
}

// * Let's create a function which creates the instance of our struct and returns it
// Here we return Sqlite instance
func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT, 
	email TEXT, 
	age INTEGER 
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	// Step 1: Prepare query
	statement, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES(?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (models.Student, error) {

	// Step 1 : Prepare query
	statement, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")

	if err != nil {
		return models.Student{}, err
	}

	// Make sure to close the satement
	defer statement.Close()

	var student models.Student

	err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}

		return models.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]models.Student, error) {
	// Step 1: Prepare query
	statement, err := s.Db.Prepare("SELECT id, name, email, age FROM students ")

	if err != nil {
		return nil, err
	}

	// Make sure to close the statement
	defer statement.Close()

	rows, err := statement.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var student models.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
