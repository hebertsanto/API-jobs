package models

type ApplyJob struct {
	ID         int    `json:"id"`
	UserName   string `json:"user_name"`
	Curriculum string `json:"curriculum"`
	UserID     int    `json:"user_id"`
	JobID      int    `json:"job_id"`
}
