package store

type User struct {
	ID   int
	Name string
}

type UserStore interface {
	CreateUser(name string) User
}
