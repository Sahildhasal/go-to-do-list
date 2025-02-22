package models

import "time"

type TodoResponse struct {
	ID          int       `json:"taskId"`
	Name        string    `json:"taskName"`
	Description string    `json:"taskDescription"`
	Date        time.Time `json:"dueDate" time_format:"2006-01-02"`
}

type TodoRequest struct {
	Name        string `json:"taskName"`
	Description string `json:"taskDescription"`
	Date        string `json:"dueDate"`
}
