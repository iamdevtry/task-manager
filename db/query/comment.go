package query

import (
	"context"
	"database/sql"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const addComment = `BEGIN proc_addcomment(:taskId, :activityId, :title, :content, :inserted_id); END;`

func (store *Store) AddComment(ctx context.Context, comment model.CommentCreate) (*model.Comment, error) {
	var err error
	var commentId int64
	commentCreated := &model.Comment{}
	if comment.TaskId == 0 {
		_, err = store.db.Exec(addComment, nil, comment.ActivityId, comment.Title, comment.Content, sql.Out{Dest: &commentId})
	}

	if comment.ActivityId == 0 {
		_, err = store.db.Exec(addComment, comment.TaskId, nil, comment.Title, comment.Content, sql.Out{Dest: &commentId})
	}

	if err != nil {
		return nil, common.ErrCannotCreateEntity("comment", err)
	}
	commentCreated = &model.Comment{
		Id:         commentId,
		TaskId:     &comment.TaskId,
		ActivityId: &comment.ActivityId,
		Title:      comment.Title,
		Content:    &comment.Content,
	}

	return commentCreated, nil
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

const getCommentByActivityId = `SELECT * FROM comments WHERE activityId = :id`

func (store *Store) ListCommentsByActivityId(ctx context.Context, id int64) ([]model.Comment, error) {
	var comment []model.Comment
	err := store.db.Select(&comment, getCommentByActivityId, id)
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
