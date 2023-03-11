package query

import (
	"context"

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

func (store *Store) GetTag(ctx context.Context, id int64) (model.Tag, error) {
	var tag model.Tag
	err := store.db.Get(&tag, getTag, id)
	if err != nil {
		return tag, common.ErrCannotGetEntity("tag", err)
	}
	return tag, nil
}

const deleteTag = `DELETE FROM tags WHERE id = :id`

func (store *Store) DeleteTag(ctx context.Context, id int64) error {
	_, err := store.db.Exec(deleteTag, id)
	if err != nil {
		return common.ErrCannotDeletedEntity("tag", err)
	}
	return nil
}
