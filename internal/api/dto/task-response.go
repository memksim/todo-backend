package dto

type TaskResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	ExpiresAt  string `json:"expires_at"`
	IsComplete bool   `json:"is_complete"`
}
