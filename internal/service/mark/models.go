package mark

import "time"

type Mark struct {
	ID        int
	StudentID int
	SubjectID int
	Mark      int
	Created   time.Time
}
