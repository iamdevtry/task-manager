package model

type TaskMeta struct {
	Id      int64  `json:"id" db:"ID"`
	TaskId  int64  `json:"task_id" db:"TASKID"`
	Key     string `json:"key" db:"KEY"`
	Content string `json:"content" db:"CONTENT"`
}
