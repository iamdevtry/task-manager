package query

import (
	"context"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const listActivities = `SELECT * FROM activities`

func (store *Store) ListActivity(ctx context.Context) ([]model.Activity, error) {
	activities := []model.Activity{}
	err := store.db.Select(&activities, listActivities)
	if err != nil {
		return nil, common.ErrCannotListEntity("activities", err)
	}
	return activities, nil
}

const getActivity = `SELECT * FROM activities WHERE id = :id`

func (store *Store) GetActivity(ctx context.Context, id int64) (model.Activity, error) {
	var activity model.Activity
	err := store.db.Get(&activity, getActivity, id)
	if err != nil {
		return activity, common.ErrCannotGetEntity("activity", err)
	}
	return activity, nil
}

const deleteActivity = `DELETE FROM activities WHERE id = :id`

func (store *Store) DeleteActivity(ctx context.Context, id int64) error {
	_, err := store.db.Exec(deleteActivity, id)
	if err != nil {
		return common.ErrCannotDeletedEntity("activity", err)
	}
	return nil
}

const addActivity = `BEGIN proc_addactivity(:UserId, :TaskId, :Title, :Description, :Hours, :PlannedStartDate, :PlannedEndDate,:Content); END;`

func (store *Store) CreateActivity(ctx context.Context, activity model.ActivityCreate) error {
	_, err := store.db.Exec(addActivity,
		activity.UserId,
		activity.TaskId,
		activity.Title,
		activity.Description,
		activity.Hours,
		activity.PlannedStartDate,
		activity.PlannedEndDate,
		activity.Content,
	)
	if err != nil {
		return common.ErrCannotCreateEntity("activity", err)
	}
	return nil
}
