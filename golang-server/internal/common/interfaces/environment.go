package interfaces

type SPACE string

const (
	SERVER   SPACE = "server"
	DATABASE SPACE = "database"
)

type Environment interface {
	GetEnv(space SPACE, key string) string
}
