package handlers

import (
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type EvenOdd string

const (
	Even EvenOdd = "even"
	Odd  EvenOdd = "odd"
)

type Settings struct {
	StartDate     time.Time   `json:"start_date"`
	EndDate       time.Time   `json:"end_date"`
	IntervalDays  int         `json:"interval_days"`
	MonthDays     []int       `json:"month_days"`
	SpecificDates []time.Time `json:"specific_dates"`
	EvenOdd       EvenOdd     `json:"even_odd"`
}

type taskMutationDTO struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
}

type taskDTO struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Settings
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	return taskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
