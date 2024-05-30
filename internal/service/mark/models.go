package mark

import "time"

type Mark struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	SubjectID int       `json:"subject_id"`
	Mark      int       `json:"mark"`
	Created   time.Time `json:"created_at"`
}

type CreateMarkRequest struct {
	StudentID int `json:"student_id" binding:"required"`
	SubjectID int `json:"subject_id" binding:"required"`
	Mark      int `json:"mark" binding:"required"`
}
