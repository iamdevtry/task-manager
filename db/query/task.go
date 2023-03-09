package query

import (
	"context"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const createTask = `BEGIN proc_addtask(:userid, :title, :description, :hours, :plannedstartdate, :plannedenddate, content); END;`

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

func (s *Store) ListTask(ctx context.Context, task model.Task) ([]model.Task, error) {
	tasks := []model.Task{}
	err := s.db.Select(&tasks, listTasks)
	if err != nil {
		return nil, common.ErrCannotListEntity("tasks", err)
	}
	return tasks, nil
}
