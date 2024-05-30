package subject

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateSubjectRequest struct {
	Name string `json:"name" binding:"required"`
}
