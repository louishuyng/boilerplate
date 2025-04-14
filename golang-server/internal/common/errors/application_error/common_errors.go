package application_error

import (
	"fmt"
)

const (
	STORE_SQL_ERROR    = 0
	EXTERNAL_API_ERROR = 1
)

var DEFAULT_ERRORS = map[string]*Error{
	fmt.Sprintf("%d", STORE_SQL_ERROR):    NewError("app", "common", STORE_SQL_ERROR, InternalError, "Store SQL error"),
	fmt.Sprintf("%d", EXTERNAL_API_ERROR): NewError("app", "common", EXTERNAL_API_ERROR, InternalError, "External API error"),
}
