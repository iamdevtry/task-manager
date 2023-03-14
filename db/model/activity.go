package model

import "time"

type Activity struct {
	Id               int64      `json:"id" db:"ID"`
	UserId           int64      `json:"user_id" db:"USERID"`
	TaskId           int64      `json:"task_id" db:"TASKID"`
	Title            string     `json:"title" db:"TITLE"`
	Description      string     `json:"description" db:"DESCRIPTION"`
	Status           int64      `json:"status" db:"STATUS"`
	Hours            float32    `json:"hours" db:"HOURS"`
	CreatedAt        *time.Time `json:"created_at" db:"CREATEDAT"`
	UpdatedAt        *time.Time `json:"updated_at" db:"UPDATEDAT"`
	PlannedStartDate *time.Time `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   *time.Time `json:"planned_end_date" db:"PLANNEDENDDATE"`
	ActualStartDate  *time.Time `json:"actual_start_date" db:"ACTUALSTARTDATE"`
	ActualEndDate    *time.Time `json:"actual_end_date" db:"ACTUALENDDATE"`
	Content          string     `json:"content" db:"CONTENT"`
}

type ActivityCreate struct {
	UserId           int64   `json:"user_id" db:"USERID"`
	TaskId           int64   `json:"task_id" db:"TASKID"`
	Title            string  `json:"title" db:"TITLE"`
	Description      string  `json:"description" db:"DESCRIPTION"`
	Hours            float32 `json:"hours" db:"HOURS"`
	PlannedStartDate string  `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   string  `json:"planned_end_date" db:"PLANNEDENDDATE"`
	Content          string  `json:"content" db:"CONTENT"`
}
