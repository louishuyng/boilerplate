package interfaces

type ResourceMapper[T any] interface {
	ToResource() T
}
