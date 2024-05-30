package report

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kekaswork/go-rest-sample/internal/database"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Report(req ReportRequest) (*Report, error) {
	db := database.NewService()

	var resultItem ReportItem
	var resultList []ReportItem
	var results Report
	var subject string
	rows, err := db.GetConn().Query(
		context.Background(),
		`SELECT 
			st.id AS student_id, 
			st.first_name, 
			st.last_name, 
			s.name AS subject_name, 
			AVG(m.mark) AS avg_mark
		FROM 
			marks AS m
		JOIN 
			students AS st ON st.id = m.student_id
		JOIN 
			subjects AS s ON m.subject_id = s.id
		WHERE 
			st.last_name LIKE $1
			AND s.ID = $2
		GROUP BY 
			st.id, st.first_name, st.last_name, s.name
		HAVING 
		AVG(m.mark) > $3;`,
		req.Prefix+"%", req.SubjectID, req.MinAvgMark,
	)
	if err != nil {
		return nil, err
	}
	_, err = pgx.ForEachRow(rows, []any{&resultItem.StudentID, &resultItem.FirstName, &resultItem.LastName, &subject, &resultItem.AverageMark}, func() error {
		resultList = append(resultList, resultItem)
		return nil
	})
	if err != nil {
		return nil, err
	}
	results.Subject = subject
	results.Results = resultList

	return &results, nil
}
