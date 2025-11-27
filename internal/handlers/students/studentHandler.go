package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	storage "github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/database"
	"github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/models"
	"github.com/bhavishaya-khandelwal-dianapps/Student-Management/internal/utils/response"
	"github.com/go-playground/validator"
)

// * 1. Function to create a new student
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a student")

		var student models.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request Validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully", slog.String("UserId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})

		// response.WriteJson(w, http.StatusCreated, map[string]string{"success": "true"})
	}
}

// * 2. Function to get student by id
func GetStudentById(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		id := req.PathValue("id")

		slog.Info("getting a student", slog.String("id = ", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(res, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Error("error getting user", slog.String("id", id))
			response.WriteJson(res, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(res, http.StatusOK, student)
	}
}

// * 3. Function to list all the students
func GetStudents(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		slog.Info("Getting all students")

		students, err := storage.GetStudents()

		if err != nil {
			response.WriteJson(res, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(res, http.StatusOK, students)
	}
}

// * 4. Function to delete student
func DeleteStudent(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(res, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		msg, err := storage.DeleteStudent(intId)

		if err != nil {
			slog.Error("Error deleting user", slog.String("id", id))
			response.WriteJson(res, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(res, http.StatusOK, msg)
	}
}
