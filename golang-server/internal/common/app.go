package common

type App interface {
	RegisterAPI() error
	RegisterDomainEvent() error
}
