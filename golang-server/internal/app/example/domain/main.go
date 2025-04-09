package domain

var _ ExampleDomain = (*Domain)(nil)

type Domain struct {
}

func NewExampleDomain() *Domain {
	return &Domain{}
}
