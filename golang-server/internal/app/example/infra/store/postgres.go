package store

type PostgresStore struct {
	// DB *sql.DB
}

var _ ExampleStore = (*PostgresStore)(nil)

func NewPostgresStore() *PostgresStore {
	return &PostgresStore{}
}

func (p *PostgresStore) CreateExample(name string) Example {
	// Implement the logic to create a example in the database
	return Example{}
}
