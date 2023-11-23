package constants

import "errors"

var (
	// global
	ErrUnexpected = errors.New("unexpected error")

	// entity
	ErrUserNotFound  = errors.New("user not found")
	ErrDBNoRowsFound = errors.New("sql: no rows in result set")

	// config
	ErrLoadConfig  = errors.New("failed to load config file")
	ErrParseConfig = errors.New("failed to parse env to config struct")
	ErrEmptyVar    = errors.New("required variable environment is empty")
)
