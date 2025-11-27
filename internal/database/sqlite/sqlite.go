package sqlite

import (
	"database/sql"

	"github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/config"
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
