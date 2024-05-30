package report

type Report struct {
	Subject string       `json:"subject"`
	Results []ReportItem `json:"results"`
}

type ReportItem struct {
	StudentID   int     `json:"student_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	AverageMark float32 `json:"average_mark"`
}

type ReportRequest struct {
	MinAvgMark float32 `json:"min_avg_mark" binding:"required"`
	SubjectID  int     `json:"subject_id" binding:"required"`
	Prefix     string  `json:"prefix"`
}
