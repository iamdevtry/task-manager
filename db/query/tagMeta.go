package query

import (
	"context"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const listTaskMeta = `SELECT * FROM task_metas`

func (store *Store) ListTaskMeta(ctx context.Context) ([]model.TaskMeta, error) {
	taskMetas := []model.TaskMeta{}
	err := store.db.Select(&taskMetas, listTaskMeta)
	if err != nil {
		return nil, common.ErrCannotListEntity("taskMetas", err)
	}
	return taskMetas, nil
}

const getTaskMeta = `SELECT * FROM task_metas WHERE id = :id`

func (store *Store) GetTaskMeta(ctx context.Context, id int64) (model.TaskMeta, error) {
	var taskMeta model.TaskMeta
	err := store.db.Get(&taskMeta, getTaskMeta, id)
	if err != nil {
		return taskMeta, common.ErrCannotGetEntity("taskMeta", err)
	}
	return taskMeta, nil
}

const deleteTaskMeta = `DELETE FROM task_metas WHERE id = :id`

func (store *Store) DeleteTaskMeta(ctx context.Context, id int64) error {
	_, err := store.db.Exec(deleteTaskMeta, id)
	if err != nil {
		return common.ErrCannotDeletedEntity("taskMeta", err)
	}
	return nil
}
