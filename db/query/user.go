package query

import (
	"context"

	"github.com/iamdevtry/task-manager/db/model"
)

const listUsers = `SELECT * FROM users`

func (s *Store) ListUsers(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	err := s.db.Select(&users, listUsers)
	if err != nil {
		return nil, err
	}
	return users, nil
}

const getUser = `SELECT * FROM users WHERE id=:1`

func (s *Store) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	err := s.db.Get(user, getUser, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
