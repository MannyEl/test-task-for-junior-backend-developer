package handlers

import (
	"time"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
	taskdomain "example.com/taskservice/internal/domain/task"
)

type Recurrence struct {
	StartDate     time.Time              `json:"start_date,omitempty"`
	EndDate       time.Time              `json:"end_date,omitempty"`
	IntervalDays  int                    `json:"interval_days,omitempty"`
	MonthDays     []int                  `json:"month_days,omitempty"`
	SpecificDates []time.Time            `json:"specific_dates,omitempty"`
	EvenOdd       scheduledomain.EvenOdd `json:"even_odd,omitempty"`
}

type taskMutationDTO struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
	Recurrence  Recurrence        `json:"recurrence,omitempty"`
}

type taskDTO struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
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
