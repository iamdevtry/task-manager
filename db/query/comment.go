package query

import (
	"context"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const addComment = `INSERT INTO comments (taskId, activityId, title, content) VALUES (:taskId, :activityId, :title, :content)`

func (store *Store) AddComment(ctx context.Context, comment model.CommentCreate) error {
	var err error
	if comment.TaskId == 0 {
		_, err = store.db.Exec(addComment, nil, comment.ActivityId, comment.Title, comment.Content)
	}

	if comment.ActivityId == 0 {
		_, err = store.db.Exec(addComment, comment.TaskId, nil, comment.Title, comment.Content)
	}

	if err != nil {
		return common.ErrCannotCreateEntity("comment", err)
	}

	return nil
}

const listComments = `SELECT * FROM comments`

func (store *Store) ListComment(ctx context.Context) ([]model.Comment, error) {
	comments := []model.Comment{}
	err := store.db.Select(&comments, listComments)
	if err != nil {
		return nil, common.ErrCannotListEntity("comments", err)
	}
	return comments, nil
}

const getComment = `SELECT * FROM comments WHERE id = :id`

func (store *Store) GetComment(ctx context.Context, id int64) (model.Comment, error) {
	var comment model.Comment
	err := store.db.Get(&comment, getComment, id)
	if err != nil {
		return comment, common.ErrCannotGetEntity("comment", err)
	}
	return comment, nil
}

const deleteComment = `DELETE FROM comments WHERE id = :id`

func (store *Store) DeleteComment(ctx context.Context, id int64) error {
	_, err := store.db.Exec(deleteComment, id)
	if err != nil {
		return common.ErrCannotDeletedEntity("comment", err)
	}
	return nil
}
