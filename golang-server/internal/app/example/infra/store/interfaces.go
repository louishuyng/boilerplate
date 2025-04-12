package store

import example_store_data "rz-server/internal/app/example/infra/store/sql/example/data"

type ExampleStore interface {
	CreateExample(name string) example_store_data.ExampleStoreData
}
