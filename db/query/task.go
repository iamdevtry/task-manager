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
