package dto

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
	ExpiresAt   int64  `json:"expires_at"`
}
