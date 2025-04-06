package domain

var _ UserEntity = (*User)(nil)

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{}
}

func (c *User) GetID() int {
	return 0
}
