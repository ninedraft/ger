package ger

import (
	"context"
	"errors"
)

type Task func(ctx context.Context) error

func NoCtx(t func() error) Task {
	return func(context.Context) error {
		return t()
	}
}

func NoError(t func(ctx context.Context)) Task {
	return func(ctx context.Context) error {
		t(ctx)
		return nil
	}
}

func Proc(t func()) Task {
	return func(context.Context) error {
		t()
		return nil
	}
}

func AllForOne(ctx context.Context, tasks ...Task) error {
	var supervisor = newSupervisor(tasks, func(ctx TaskContext, err error) {
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
	return supervisor.Run(ctx)
}

func OneForOne(ctx context.Context, tasks ...Task) error {
	var supervisor = newSupervisor(tasks, func(ctx TaskContext, err error) {
		if errors.Is(err, ErrRestart) {
			ctx.RestartMe()
			return
		}
		ctx.PushError(err)
		ctx.StopMe()
	})
	return supervisor.Run(ctx)
}

func OneForRest(ctx context.Context, tasks ...Task) error {
	var supervisor = newSupervisor(tasks, func(ctx TaskContext, err error) {
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
	return supervisor.Run(ctx)
}
