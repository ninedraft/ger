package ger

//go:generate stringer -type cError -linecomment -output errors_string.go
type cError uint8

const (
	// ErrRestart means the task must be restarted.
	ErrRestart cError = iota + 1 // restart me

	// ErrAllCanceled means that all supervised tasks are canceled.
	ErrAllCanceled // all tasks are canceled

	// ErrStopped means that the current task is stopped.
	ErrStopped // task is stopped

	// ErrSupervisorIsNotRunning means that you trying to wait on non-running supervisor.
	ErrSupervisorIsNotRunning // supervisor is not running
)

func (i cError) Error() string { return i.String() }
