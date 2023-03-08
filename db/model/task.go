package model

type Task struct {
	Id               int64  `json:"id" db:"ID"`
	UserId           int64  `json:"user_id" db:"USERID"`
	Title            string `json:"title" db:"TITLE"`
	Description      string `json:"description" db:"DESCRIPTION"`
	Content          string `json:"content" db:"CONTENT"`
	Hours            int64  `json:"hours" db:"HOURS"`
	PlannedStartDate int64  `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   int64  `json:"planned_end_date" db:"PLANNEDENDDATE"`
	ActualStartDate  int64  `json:"actual_start_date" db:"ACTUALSTARTDATE"`
	ActualEndDate    int64  `json:"actual_end_date" db:"ACTUALENDDATE"`
	CreatedBy        int64  `json:"created_by" db:"CREATEDBY"`
	UpdatedBy        int64  `json:"updated_by" db:"UPDATEDBY"`
	CreatedAt        int64  `json:"created_at" db:"CREATEDAT"`
	UpdatedAt        int64  `json:"updated_at" db:"UPDATEDAT"`
	Status           int64  `json:"status" db:"STATUS"`
}

type TaskCreate struct {
	UserId           int64  `json:"user_id" db:"USERID"`
	Title            string `json:"title" db:"TITLE"`
	Description      string `json:"description" db:"DESCRIPTION"`
	Content          string `json:"content" db:"CONTENT"`
	Hours            int64  `json:"hours" db:"HOURS"`
	PlannedStartDate int64  `json:"planned_start_date" db:"PLANNEDSTARTDATE"`
	PlannedEndDate   int64  `json:"planned_end_date" db:"PLANNEDENDDATE"`
	CreatedBy        int64  `json:"created_by" db:"CREATEDBY"`
	CreatedAt        int64  `json:"created_at" db:"CREATEDAT"`
	Status           int64  `json:"status" db:"STATUS"`
}
