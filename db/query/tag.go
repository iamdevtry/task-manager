package query

import (
	"context"
	"database/sql"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const listTags = `SELECT * FROM tags`

func (store *Store) ListTag(ctx context.Context) ([]model.Tag, error) {
	tags := []model.Tag{}
	err := store.db.Select(&tags, listTags)
	if err != nil {
		return nil, common.ErrCannotListEntity("tags", err)
	}
	return tags, nil
}

const getTag = `SELECT * FROM tags WHERE id = :id`

func (store *Store) GetTag(ctx context.Context, id int64) (*model.Tag, error) {
	var tag model.Tag
	err := store.db.Get(&tag, getTag, id)
	if err != nil {
		return &tag, common.ErrCannotGetEntity("tag", err)
	}
	return &tag, nil
}

const deleteTag = `DELETE FROM tags WHERE id = :id`

func (store *Store) DeleteTag(ctx context.Context, id int64) error {
	_, err := store.db.Exec(deleteTag, id)
	if err != nil {
		return common.ErrCannotDeletedEntity("tag", err)
	}
	return nil
}

const addTag = `BEGIN proc_addtag(:title, :slug, :inserted_id); END;`

func (store *Store) AddTag(ctx context.Context, tag model.Tag) (id *int64, err error) {
	var resultId int64
	_, err = store.db.Exec(addTag, tag.Title, tag.Slug, sql.Out{Dest: &resultId})
	if err != nil {
		return nil, common.ErrCannotCreateEntity("tag", err)
	}

	return &resultId, nil
}

const addTaskToTag = `BEGIN proc_addtasktotag(:TaskId, :TagId); END;`

func (store *Store) AddTaskToTag(ctx context.Context, taskId int64, tagId int64) error {
	_, err := store.db.Exec(addTaskToTag, taskId, tagId)
	if err != nil {
		return common.ErrCannotCreateEntity("task to tag", err)
	}
	return nil
}
