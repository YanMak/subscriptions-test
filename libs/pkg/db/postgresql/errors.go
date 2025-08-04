package db_postgresql

import (
	"errors"
)

var ErrDbOpenError = errors.New("DB open error")
var ErrDbPingError = errors.New("DB ping error")
