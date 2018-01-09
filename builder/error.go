package builder

import "errors"

var (
	// ErrNotSupportType not supported SQL type error
	ErrNotSupportType = errors.New("not supported SQL type")
	// ErrNoNotInConditions no NOT IN params error
	ErrNoNotInConditions = errors.New("No NOT IN conditions")
	// ErrNoInConditions no IN params error
	ErrNoInConditions = errors.New("No IN conditions")
)
