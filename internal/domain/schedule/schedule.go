package schedule

import "time"

type EvenOdd string

const (
	Even EvenOdd = "even"
	Odd  EvenOdd = "odd"
)

type RecurrenceType string

const (
	Interval    RecurrenceType = "interval"
	MonthlyDays RecurrenceType = "monthly_days"
)

type Schedule struct {
	ID              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	StartDate       time.Time      `json:"start_date"`
	EndDate         time.Time      `json:"end_date"`
	IntervalDays    int            `json:"interval_days"`
	MonthDays       []int          `json:"month_days"`
	SpecificDates   []time.Time    `json:"specific_dates"`
	EvenOdd         EvenOdd        `json:"even_odd"`
	Type            RecurrenceType `json:"type"`
	Active          bool           `json:"active"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	LastGeneratedAt time.Time      `json:"last_generated_at"`
}
