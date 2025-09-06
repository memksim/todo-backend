package dto

// CreateTaskRequest
// expires_at uses the RFC-3339 format
type CreateTaskRequest struct {
	Title     string `json:"title"`
	ExpiresAt string `json:"expires_at"`
}

type EditTaskRequest struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	ExpiresAt  string `json:"expires_at"`
	IsComplete bool   `json:"is_complete"`
}

type RemoveTaskRequest struct {
	Id string `json:"id"`
}
