package store

type Example struct {
	ID   int
	Name string
}

type ExampleStore interface {
	CreateExample(name string) Example
}
