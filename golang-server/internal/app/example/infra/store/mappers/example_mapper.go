package mapper

import (
	"rz-server/internal/app/example/domain"
	store "rz-server/internal/app/example/infra/store"
)

func ToDBExample(example domain.Example) store.Example {
	return store.Example{
		Name: example.Name,
	}
}

func FromDBExample(example store.Example) domain.Example {
	return domain.Example{
		Name: example.Name,
	}
}
