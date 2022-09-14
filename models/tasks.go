package models

import (
	"time"
)

type Task struct {
	ID         string
	UserID     string
	TaskDetail string
	DueDate    time.Time
}
