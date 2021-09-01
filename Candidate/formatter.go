package Candidate

type CreateCandidateInput struct {
	PositionID string `json:"positionid"  binding:"required"`
	CategoryID string `json:"categoryid"  binding:"required"`
	UserID     string `json:"userid"  binding:"required"`
}

type DetailCandidateInput struct {
	ID string `uri:"id"  binding:"required"`
}
