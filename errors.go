package ger

//go:generate stringer -type cError -linecomment -output errors_string.go
type cError uint8

const (
	ErrRestart                cError = iota + 1 // restart me
	ErrAllCancelled                             // all tasks are cancelled
	ErrStopped                                  // task is stopped
	ErrSupervisorIsNotRunning                   // supervisor is not running
)

func (i cError) Error() string { return i.String() }
