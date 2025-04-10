package gateway

// Use this when one app needs to call another app
type ExampleGateway interface {
	ExampleGet(id string) (string, error)
}
