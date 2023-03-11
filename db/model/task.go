package model

import "time"

type Task struct {
	Id               int64      `json:"id" db:"ID"`
	UserId           int64      `json:"user_id" db:"USERID"`
	Title            string     `json:"title" db:"TITLE"`
	Description      string     `json:"description" db:"DESCRIPTION"`
	Content          string     `json:"content" db:"CONTENT"`
	Hours            int64      `json:"hours" db:"HOURS"`
	PlannedStartDate int64      `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   int64      `json:"planned_end_date" db:"PLANNEDENDDATE"`
	ActualStartDate  int64      `json:"actual_start_date" db:"ACTUALSTARTDATE"`
	ActualEndDate    int64      `json:"actual_end_date" db:"ACTUALENDDATE"`
	CreatedAt        *time.Time `json:"created_at" db:"CREATEDAT"`
	UpdatedAt        *time.Time `json:"updated_at" db:"UPDATEDAT"`
	Status           int64      `json:"status" db:"STATUS"`
}

type TaskCreate struct {
	UserId           int64      `json:"user_id" db:"USERID"`
	Title            string     `json:"title" db:"TITLE"`
	Description      string     `json:"description" db:"DESCRIPTION"`
	Content          string     `json:"content" db:"CONTENT"`
	Hours            int64      `json:"hours" db:"HOURS"`
	PlannedStartDate int64      `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   int64      `json:"planned_end_date" db:"PLANNEDENDDATE"`
	CreatedBy        int64      `json:"created_by" db:"CREATEDBY"`
	CreatedAt        *time.Time `json:"created_at" db:"CREATEDAT"`
	Status           int64      `json:"status" db:"STATUS"`
}
