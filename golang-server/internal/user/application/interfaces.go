package application

type User struct {
	ID   int
	Name string
}

type UserService interface {
	CreateUser(createUserCommand CreateUserCommand) (*User, error)
}
