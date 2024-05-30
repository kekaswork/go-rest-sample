package mark

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kekaswork/go-rest-sample/internal/database"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() ([]Mark, error) {
	db := database.NewService()

	var marks []Mark
	var mark Mark
	rows, _ := db.GetConn().Query(context.Background(), "select * from marks")
	_, err := pgx.ForEachRow(rows, []any{&mark.ID, &mark.StudentID, &mark.SubjectID, &mark.Mark, &mark.Created}, func() error {
		marks = append(marks, mark)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return marks, nil
}

func (s *Service) Get(idx int) (*Mark, error) {
	db := database.NewService()

	var mark Mark
	err := db.GetConn().QueryRow(
		context.Background(),
		"select * from marks where id = $1", idx,
	).Scan(&mark.ID, &mark.StudentID, &mark.SubjectID, &mark.Mark, &mark.Created)
	if err != nil {
		return nil, err
	}

	return &mark, nil
}

func (s *Service) Add(req CreateMarkRequest) (*Mark, error) {
	db := database.NewService()

	var mark Mark
	err := db.GetConn().QueryRow(
		context.Background(),
		"INSERT INTO marks (student_id, subject_id, mark, created_at) VALUES ($1, $2, $3, $4) RETURNING id, student_id, subject_id, mark, created_at",
		req.StudentID, req.SubjectID, req.Mark, time.Now(),
	).Scan(&mark.ID, &mark.StudentID, &mark.SubjectID, &mark.Mark, &mark.Created)
	if err != nil {
		return nil, err
	}

	return &mark, nil
}

func (s *Service) Remove(idx int) (*Mark, error) {
	db := database.NewService()

	var mark Mark
	err := db.GetConn().QueryRow(
		context.Background(),
		"SELECT id, student_id, subject_id, mark, created_at FROM marks WHERE id = $1",
		idx,
	).Scan(&mark.ID, &mark.StudentID, &mark.SubjectID, &mark.Mark, &mark.Created)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve mark with id %d: %v", idx, err)
	}

	_, err = db.GetConn().Exec(
		context.Background(),
		"DELETE FROM marks WHERE id = $1",
		idx,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to delete mark: %v", err)
	}

	return &mark, nil
}

func (s *Service) Update(idx int, req CreateMarkRequest) (*Mark, error) {
	db := database.NewService()

	var mark Mark
	err := db.GetConn().QueryRow(
		context.Background(),
		"SELECT id, student_id, subject_id, mark, created_at FROM marks WHERE id = $1",
		idx,
	).Scan(&mark.ID, &mark.StudentID, &mark.SubjectID, &mark.Mark, &mark.Created)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve mark with id %d: %v", idx, err)
	}

	_, err = db.GetConn().Exec(
		context.Background(),
		"UPDATE marks SET student_id = $1, subject_id = $2, mark = $3 WHERE id = $4",
		req.StudentID, req.SubjectID, req.Mark, idx,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update mark: %v", err)
	}

	mark.StudentID = req.StudentID
	mark.SubjectID = req.SubjectID
	mark.Mark = req.Mark

	return &mark, nil
}
