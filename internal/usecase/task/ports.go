package task

import (
	"context"
	"time"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
	taskdomain "example.com/taskservice/internal/domain/task"
)

type Repository interface {
	Create(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error)
	CreateSchedule(ctx context.Context, task *scheduledomain.Schedule) (*scheduledomain.Schedule, error)
	GetByID(ctx context.Context, id int64) (*taskdomain.Task, error)
	Update(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]taskdomain.Task, error)
}

type Usecase interface {
	Create(ctx context.Context, input CreateInput) (*taskdomain.Task, error)
	GetByID(ctx context.Context, id int64) (*taskdomain.Task, error)
	Update(ctx context.Context, id int64, input UpdateInput) (*taskdomain.Task, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]taskdomain.Task, error)
}

type Recurrence struct {
	StartDate     time.Time
	EndDate       time.Time
	IntervalDays  int
	MonthDays     []int
	SpecificDates []time.Time
	EvenOdd       scheduledomain.EvenOdd
}

type CreateInput struct {
	Title       string
	Description string
	Status      taskdomain.Status
	Recurrence  Recurrence
}

type UpdateInput struct {
	Title       string
	Description string
	Status      taskdomain.Status
}
