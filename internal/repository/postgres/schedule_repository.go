package postgres

import (
	"context"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

func (r *Repository) CreateSchedule(ctx context.Context, schedule *scheduledomain.Schedule) (*scheduledomain.Schedule, error) {
	const query = `
		INSERT INTO task_generation_rules (title, description, start_date, end_date, interval_days, month_days, specific_dates, even_odd, created_at, updated_at, last_generated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`
	// TODO: add scan func
	err := r.pool.QueryRow(ctx, query, schedule.Title, schedule.Description, schedule.StartDate, schedule.EndDate, schedule.IntervalDays, schedule.MonthDays, schedule.SpecificDates, schedule.EvenOdd, schedule.CreatedAt, schedule.UpdatedAt, schedule.LastGeneratedAt).Scan(&schedule.ID)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}
