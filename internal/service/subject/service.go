package subject

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

func (s *Service) List() ([]Subject, error) {
	db := database.NewService()

	var subjects []Subject
	var id int
	var name string
	rows, _ := db.GetConn().Query(context.Background(), "select * from subjects")
	_, err := pgx.ForEachRow(rows, []any{&id, &name}, func() error {
		subjects = append(subjects, Subject{
			ID:   id,
			Name: name,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *Service) Get(idx int) (*Subject, error) {
	db := database.NewService()

	var id int
	var name string
	err := db.GetConn().QueryRow(context.Background(), "select * from subjects where id = $1", idx).Scan(&id, &name)
	if err != nil {
		return nil, err
	}

	return &Subject{
		ID:   id,
		Name: name,
	}, nil
}

func (s *Service) Add(req CreateSubjectRequest) (*Subject, error) {
	db := database.NewService()

	var subject Subject
	err := db.GetConn().QueryRow(
		context.Background(),
		"INSERT INTO subjects (name) VALUES ($1) RETURNING id, name",
		req.Name,
	).Scan(&subject.ID, &subject.Name)
	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (s *Service) Remove(idx int) (*Subject, error) {
	db := database.NewService()

	var subject Subject
	err := db.GetConn().QueryRow(
		context.Background(),
		"SELECT id, name FROM subjects WHERE id = $1",
		idx,
	).Scan(&subject.ID, &subject.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve subject with id %d: %v", idx, err)
	}

	_, err = db.GetConn().Exec(
		context.Background(),
		"DELETE FROM subjects WHERE id = $1",
		idx,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to delete subject: %v", err)
	}

	return &subject, nil
}

func (s *Service) Update(idx int, req CreateSubjectRequest) (*Subject, error) {
	db := database.NewService()

	var subject Subject
	err := db.GetConn().QueryRow(
		context.Background(),
		"SELECT id, name from subjects WHERE id = $1",
		idx,
	).Scan(&subject.ID, &subject.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve subject with id %d: %v", idx, err)
	}

	_, err = db.GetConn().Exec(
		context.Background(),
		"UPDATE subjects SET name = $1 WHERE id = $2",
		req.Name, idx,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update subject: %v", err)
	}

	subject.Name = req.Name

	return &subject, nil
}
