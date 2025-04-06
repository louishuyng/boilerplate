package store

type PostgresStore struct {
	// DB *sql.DB
}

var _ UserStore = (*PostgresStore)(nil)

func NewPostgresStore() *PostgresStore {
	return &PostgresStore{}
}

func (p *PostgresStore) CreateUser(name string) User {
	// Implement the logic to create a user in the database
	return User{}
}
