package interfaces

type ServerApp interface {
	RegisterAPI() error
	RegisterDomainEvent() error
}
