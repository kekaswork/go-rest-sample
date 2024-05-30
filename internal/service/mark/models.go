package mark

import "time"

type Mark struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	SubjectID int       `json:"subject_id"`
	Mark      int       `json:"mark"`
	Created   time.Time `json:"created_at"`
}
