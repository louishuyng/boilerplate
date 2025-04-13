package user_sql_store

import (
	"context"
	"rz-server/internal/app/user/infra/store"
	sql_store "rz-server/internal/app/user/infra/store/sql"
	user_store_data "rz-server/internal/app/user/infra/store/sql/user/data"

	repository "rz-server/internal/app/user/infra/store/sql/repository"
)

var _ store.UserStore = (*UserStore)(nil)

type UserStore struct {
	Queries *repository.Queries
}

func New(store *sql_store.Repository) *UserStore {
	return &UserStore{
		Queries: store.Queries,
	}
}

func (s *UserStore) CreateUser(body user_store_data.CreateUserBody) *user_store_data.Data {
	user, err := s.Queries.CreateUser(context.Background(), repository.CreateUserParams{
		Email:       body.Email,
		Password:    body.Password,
		DisplayName: body.DisplayName,
	})

	if err != nil {
		panic(err)
	}

	return user_store_data.New(user.ID, user.DisplayName, user.Email, user.Password)
}

func (s *UserStore) GetUserByEmail(email string) *user_store_data.Data {
	user, err := s.Queries.GetUserByEmail(context.Background(), email)

	if err != nil {
		return nil
	}

	return user_store_data.New(user.ID, user.DisplayName, user.Email, user.Password)
}
