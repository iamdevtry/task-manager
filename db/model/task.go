package model

import "time"

type Task struct {
	Id               int64      `json:"id" db:"ID"`
	UserId           int64      `json:"user_id" db:"USERID"`
	Title            string     `json:"title" db:"TITLE"`
	Description      string     `json:"description" db:"DESCRIPTION"`
	Content          string     `json:"content" db:"CONTENT"`
	Hours            float32    `json:"hours" db:"HOURS"`
	PlannedStartDate *time.Time `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   *time.Time `json:"planned_end_date" db:"PLANNEDENDDATE"`
	ActualStartDate  *time.Time `json:"actual_start_date" db:"ACTUALSTARTDATE"`
	ActualEndDate    *time.Time `json:"actual_end_date" db:"ACTUALENDDATE"`
	CreatedAt        *time.Time `json:"created_at" db:"CREATEDAT"`
	UpdatedAt        *time.Time `json:"updated_at" db:"UPDATEDAT"`
	Status           int64      `json:"status" db:"STATUS"`
}

type TaskCreate struct {
	UserId           int64   `json:"user_id" db:"USERID"`
	Title            string  `json:"title" db:"TITLE"`
	Description      string  `json:"description" db:"DESCRIPTION"`
	Content          string  `json:"content" db:"CONTENT"`
	Hours            float32 `json:"hours" db:"HOURS"`
	PlannedStartDate string  `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   string  `json:"planned_end_date" db:"PLANNEDENDDATE"`
}
