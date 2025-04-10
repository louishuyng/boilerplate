package interfaces

import "database/sql"

type CMD struct {
	Server       Server
	ConsumeEvent <-chan Event
	Util         *Util
	SqlDB        *sql.DB
	Environment  Environment
}
