package common

type App interface {
	RunServer() error
	RunDomainEventLoop() error
}
