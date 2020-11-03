package ger

import (
	"context"
	"errors"
)

// Task is a interruptable function.
// This package functions expects task to return following errors to sygnal task lifecycle:
// 		- ErrStopped
//		- ErrRestart
type Task func(ctx context.Context) error

// NoCtx creates a task from an error returning function.
func NoCtx(t func() error) Task {
	return func(context.Context) error {
		return t()
	}
}

// NoError creates a task from a context consuming function.
func NoError(t func(ctx context.Context)) Task {
	return func(ctx context.Context) error {
		t(ctx)
		return nil
	}
}

// Proc creates a task from a procedure-like function.
func Proc(t func()) Task {
	return func(context.Context) error {
		t()
		return nil
	}
}

// AllForOne runs provided tasks with all-for-one strategy.
// IF any of the tasks fails with ErrRestart, then all tasks will be restarted.
func AllForOne(ctx context.Context, tasks ...Task) error {
	var supervisor = newSupervisor(tasks, func(ctx _TaskContext, err error) {
		var isRestart = errors.Is(err, ErrRestart)
		switch {
		case isRestart && ctx.IsRestarting():
			return
		case isRestart:
			ctx.RestartAll()
			return
		default:
			ctx.PushError(err)
			ctx.StopMe()
		}
	})
	defer freeSupervisor(supervisor)
	return supervisor.Run(ctx)
}

// OneForOne runs provided tasks with one-for-one strategy.
// If any of the tasks fails, then only this task will be restarted.
func OneForOne(ctx context.Context, tasks ...Task) error {
	var supervisor = newSupervisor(tasks, func(ctx _TaskContext, err error) {
		if errors.Is(err, ErrRestart) {
			ctx.RestartMe()
			return
		}
		ctx.PushError(err)
		ctx.StopMe()
	})
	defer freeSupervisor(supervisor)
	return supervisor.Run(ctx)
}

// OneForRest runst provided tasks with one-for-rest strategy.
// If any of the tasks fails, then all tasks, passed in the
// function after the failed one, will be restarted.
func OneForRest(ctx context.Context, tasks ...Task) error {
	var supervisor = newSupervisor(tasks, func(ctx _TaskContext, err error) {
		var isRestart = errors.Is(err, ErrRestart)
		switch {
		case isRestart && ctx.IsRestarting():
			return
		case isRestart:
			ctx.RestartRest()
			return
		default:
			ctx.PushError(err)
			ctx.StopMe()
		}
	})
	defer freeSupervisor(supervisor)
	return supervisor.Run(ctx)
}
