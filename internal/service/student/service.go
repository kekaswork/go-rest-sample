package student

import (
	"context"

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
	var id int
	var firstName string
	var lastName string
	rows, _ := db.GetConn().Query(context.Background(), "select * from students")
	_, err := pgx.ForEachRow(rows, []any{&id, &firstName, &lastName}, func() error {
		students = append(students, Student{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *Service) Get(idx int) (*Student, error) {
	db := database.NewService()

	var id int
	var firstName string
	var lastName string
	err := db.GetConn().QueryRow(context.Background(), "select * from students where id = $1", idx).Scan(&id, &firstName, &lastName)
	if err != nil {
		return nil, err
	}

	return &Student{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
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

func (s *Service) Remove(id int) (*Student, error) {
	// todo

	return &Student{}, nil
}

func (s *Service) Update(id int) (*Student, error) {
	// todo

	return &Student{}, nil
}
