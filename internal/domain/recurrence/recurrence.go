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

type Recurrence struct {
	ID            int             `json:"id"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	StartDate     *time.Time      `json:"start_date,omitempty"`
	EndDate       *time.Time      `json:"end_date,omitempty"`
	IntervalDays  *int            `json:"interval_days,omitempty"`
	MonthDays     *[]int          `json:"month_days,omitempty"`
	SpecificDates *[]time.Time    `json:"specific_dates,omitempty"`
	EvenOdd       *EvenOdd        `json:"even_odd,omitempty"`
	Type          *RecurrenceType `json:"type,omitempty"`
	Active        bool            `json:"active"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}
