package ger

import (
	"context"
	"fmt"
	"sync/atomic"

	"go.uber.org/multierr"
)

type taskContext struct {
	id    int
	state int32
	sup   tSup

	context.Context
	task Task
}

type tSup interface {
	RestartAll()
	StopAll()
	restartRest(id int)

	debugf(msg string, args ...interface{})
	done()
	getStrategy() restartStrategy
	getErrors() chan error
}

const (
	taskStopped int32 = iota
	taskStarting
	taskRunning
	taskRestarting
)

func (tc *taskContext) startMe() {
	defer tc.sup.done()
	tc.sup.debugf("starting task")
	var strategy = tc.sup.getStrategy()
	atomic.StoreInt32(&tc.state, taskStarting)
	for tc.IsRestarting() || tc.isStarting() {
		tc.cycle(strategy)
	}
}

func (tc *taskContext) cycle(strategy restartStrategy) {
	var errTask error
	defer func() {
		strategy(tc, errTask)
		tc.StopMe()
	}()
	atomic.StoreInt32(&tc.state, taskRunning)
	errTask = tc.task(tc)
}

func (tc *taskContext) RestartMe() {
	atomic.CompareAndSwapInt32(&tc.state, taskRunning, taskRestarting)
}

func (tc *taskContext) StopMe() {
	atomic.CompareAndSwapInt32(&tc.state, taskRunning, taskStopped)
}

func (tc *taskContext) RestartAll() {
	tc.sup.RestartAll()
}

func (tc *taskContext) RestartRest() {
	tc.sup.restartRest(tc.id)
}

func (tc *taskContext) Done() <-chan struct{} {
	if tc.IsStopped() || tc.IsRestarting() {
		return alwaysDone
	}
	return tc.Context.Done()
}

var alwaysDone = make(chan struct{})

func init() { close(alwaysDone) }

func (tc *taskContext) Err() error {
	switch {
	case tc.IsStopped():
		return errStopped
	case tc.IsRestarting():
		return errRestarting
	default:
		return tc.Context.Err()
	}
}

var (
	errStopped    = wraperrs("context", context.Canceled, ErrStopped)
	errRestarting = wraperrs("context", context.Canceled, ErrRestart)
)

func (tc *taskContext) IsStopped() bool {
	return atomic.LoadInt32(&tc.state) == taskStopped
}

func (tc *taskContext) isStarting() bool {
	return atomic.LoadInt32(&tc.state) == taskStarting
}

func (tc *taskContext) IsRunning() bool {
	return atomic.LoadInt32(&tc.state) == taskRunning
}

func (tc *taskContext) IsRestarting() bool {
	return atomic.LoadInt32(&tc.state) == taskRestarting
}

func (tc *taskContext) PushError(err error) {
	select {
	case <-tc.Done():
	case tc.sup.getErrors() <- err:
	}
}

func wraperrs(msg string, errs ...error) error {
	var err = multierr.Combine(errs...)
	return fmt.Errorf("%s: %w", msg, err)
}
