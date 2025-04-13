package user_store_data

import "github.com/google/uuid"

type Data struct {
	Id          uuid.UUID
	DisplayName string
	Email       string
	Password    string
}

func New(id uuid.UUID, displayName string, email string, password string) *Data {
	return &Data{
		Id:          id,
		DisplayName: displayName,
		Email:       email,
		Password:    password,
	}
}

type CreateUserBody struct {
	DisplayName string
	Email       string
	Password    string
}
