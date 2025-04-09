package interfaces

type App interface {
	RegisterAPI() error
	RegisterDomainEvent() error
}
