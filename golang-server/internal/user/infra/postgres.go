package store

type Store struct {
	// DB *sql.DB
}

func NewPostgresStore() *Store {
	return &Store{}
}
