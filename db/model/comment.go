package model

import "time"

type Comment struct {
	Id         int64      `json:"id" db:"ID"`
	TaskId     int64      `json:"task_id" db:"TASKID"`
	ActivityId int64      `json:"activity_id" db:"ACTIVITYID"`
	Title      string     `json:"title" db:"TITLE"`
	CreatedAt  *time.Time `json:"created_at" db:"CREATEDAT"`
	UpdatedAt  *time.Time `json:"updated_at" db:"UPDATEDAT"`
	Content    string     `json:"content" db:"CONTENT"`
}
