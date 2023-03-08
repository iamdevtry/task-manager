package query

import (
	"context"

	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/db/model"
)

const listUsers = `SELECT * FROM users`

func (s *Store) ListUsers(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	err := s.db.Select(&users, listUsers)
	if err != nil {
		return nil, common.ErrCannotListEntity("users", err)
	}
	return users, nil
}

const getUser = `SELECT * FROM users WHERE id=:1`

func (s *Store) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	err := s.db.Get(user, getUser, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity("user", err)
	}
	return user, nil
}

const createUser = `BEGIN proc_adduser(:FIRSTNAME, :MIDDLENAME, :LASTNAME, :USERNAME, :MOBILE, :EMAIL, :PASSWORDHASH, :INTRO, :PROFILE); END;`

func (s *Store) Register(ctx context.Context, user model.UserCreate) error {
	_, err := s.db.Exec(createUser,
		user.FirstName,
		user.MiddleName,
		user.LastName,
		user.Username,
		user.Mobile,
		user.Email,
		user.Password,
		user.Intro,
		user.Profile,
	)
	if err != nil {
		return common.ErrCannotCreateEntity("user", err)
	}
	return nil
}

const loginUser = `SELECT * FROM users WHERE username=:1`

func (s *Store) Login(ctx context.Context, userLogin model.UserLogin) (*model.User, error) {
	user := model.User{}

	if err := s.db.Get(&user, loginUser, userLogin.Username); err != nil {
		return nil, common.ErrRecordNotFound
	}

	return &user, nil
}
