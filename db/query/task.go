package query

import (
	"context"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const createTask = `BEGIN proc_addtask(:userid, :title, :description, :hours, :plannedstartdate, :plannedenddate, :content); END;`

func (s *Store) CreateTask(ctx context.Context, task model.TaskCreate) error {
	_, err := s.db.Exec(createTask,
		task.UserId,
		task.Title,
		task.Description,
		task.Hours,
		task.PlannedStartDate,
		task.PlannedEndDate,
		task.Content,
	)
	if err != nil {
		return common.ErrCannotCreateEntity("task", err)
	}
	return nil
}

const listTasks = `SELECT * FROM tasks`

func (s *Store) ListTask(ctx context.Context) ([]model.Task, error) {
	tasks := []model.Task{}
	err := s.db.Select(&tasks, listTasks)
	if err != nil {
		return nil, common.ErrCannotListEntity("tasks", err)
	}
	return tasks, nil
}

const getTask = `SELECT * FROM tasks WHERE id = :id`

func (s *Store) GetTask(ctx context.Context, id int64) (model.Task, error) {
	var task model.Task
	err := s.db.Get(&task, getTask, id)
	if err != nil {
		return task, common.ErrCannotGetEntity("task", err)
	}
	return task, nil
}

const getTaskByUserId = `SELECT * FROM tasks WHERE userid = :userid`

func (s *Store) GetTaskByUserId(ctx context.Context, userid int64) ([]model.Task, error) {
	tasks := []model.Task{}
	err := s.db.Select(&tasks, getTaskByUserId, userid)
	if err != nil {
		return nil, common.ErrCannotGetEntity("task", err)
	}
	return tasks, nil
}

const deleteTask = `DELETE FROM tasks WHERE id = :id`

func (s *Store) DeleteTask(ctx context.Context, id int64) error {
	_, err := s.db.Exec(deleteTask, id)
	if err != nil {
		return common.ErrCannotDeletedEntity("task", err)
	}
	return nil
}
