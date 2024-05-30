package student

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kekaswork/go-rest-sample/internal/database"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() ([]Student, error) {
	db := database.NewService()

	var students []Student
	var student Student
	rows, _ := db.GetConn().Query(context.Background(), "select * from students")
	_, err := pgx.ForEachRow(rows, []any{&student.ID, &student.FirstName, &student.LastName}, func() error {
		students = append(students, student)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *Service) Get(idx int) (*Student, error) {
	db := database.NewService()

	var student Student
	err := db.GetConn().QueryRow(
		context.Background(),
		"select * from students where id = $1", idx,
	).Scan(&student.ID, &student.FirstName, &student.LastName)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Service) Add(req CreateStudentRequest) (*Student, error) {
	db := database.NewService()

	var student Student
	err := db.GetConn().QueryRow(
		context.Background(),
		"INSERT INTO students (first_name, last_name) VALUES ($1, $2) RETURNING id, first_name, last_name",
		req.FirstName, req.LastName,
	).Scan(&student.ID, &student.FirstName, &student.LastName)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Service) Remove(idx int) (*Student, error) {
	db := database.NewService()

	var student Student
	err := db.GetConn().QueryRow(
		context.Background(),
		"SELECT id, first_name, last_name FROM students WHERE id = $1",
		idx,
	).Scan(&student.ID, &student.FirstName, &student.LastName)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve student with id %d: %v", idx, err)
	}

	_, err = db.GetConn().Exec(
		context.Background(),
		"DELETE FROM students WHERE id = $1",
		idx,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to delete student: %v", err)
	}

	return &student, nil
}

func (s *Service) Update(idx int, req CreateStudentRequest) (*Student, error) {
	db := database.NewService()

	var student Student
	err := db.GetConn().QueryRow(
		context.Background(),
		"SELECT id, first_name, last_name FROM students WHERE id = $1",
		idx,
	).Scan(&student.ID, &student.FirstName, &student.LastName)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve student with id %d: %v", idx, err)
	}

	_, err = db.GetConn().Exec(
		context.Background(),
		"UPDATE students SET first_name = $1, last_name = $2 WHERE id = $3",
		req.FirstName, req.LastName, idx,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update student: %v", err)
	}

	student.FirstName = req.FirstName
	student.LastName = req.LastName

	return &student, nil
}
