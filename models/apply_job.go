package models

type ApplyJob struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	UserID   int    `json:"user_id" validate:"required"`
	JobID    int    `json:"job_id"`
}
