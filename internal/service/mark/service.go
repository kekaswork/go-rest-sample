package mark

import (
	"context"
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
	var id, studentID, subjectID, mark int
	var createdAt time.Time
	rows, _ := db.GetConn().Query(context.Background(), "select * from marks")
	_, err := pgx.ForEachRow(rows, []any{&id, &studentID, &subjectID, &mark, &createdAt}, func() error {
		marks = append(marks, Mark{
			ID:        id,
			StudentID: studentID,
			SubjectID: subjectID,
			Mark:      mark,
			Created:   createdAt,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return marks, nil
}

func (s *Service) Get(idx int) (*Mark, error) {
	db := database.NewService()

	var id, studentID, subjectID, mark int
	var createdAt time.Time
	err := db.GetConn().QueryRow(context.Background(), "select * from marks where id = $1", idx).Scan(&id, &studentID, &subjectID, &mark, &createdAt)
	if err != nil {
		return nil, err
	}

	return &Mark{
		ID:        id,
		StudentID: studentID,
		SubjectID: subjectID,
		Mark:      mark,
		Created:   createdAt,
	}, nil
}

func (s *Service) Add(id int) (*Mark, error) {
	// todo

	return &Mark{}, nil
}

func (s *Service) Remove(id int) (*Mark, error) {
	// todo

	return &Mark{}, nil
}

func (s *Service) Update(id int) (*Mark, error) {
	// todo

	return &Mark{}, nil
}
