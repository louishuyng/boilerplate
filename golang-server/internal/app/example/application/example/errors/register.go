package example_errors

func (e *ExampleError) RegisterAllErrors() {
	e.RegisterNotFoundError(EXAMPLE_NOT_FOUND, "example not found")
}
