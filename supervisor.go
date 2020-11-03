package ger

import (
	"context"
	"sync"

	"go.uber.org/multierr"
)

type _TaskContext interface {
	context.Context

	StopMe()
	RestartMe()
	RestartAll()
	RestartRest()

	PushError(error)

	IsRestarting() bool
	IsRunning() bool
}

type taskSupervisor struct {
	isStarted bool
	wg        sync.WaitGroup
	tasks     []*taskContext
	errors    chan error
	strategy  restartStrategy
	debugLog  func(string, ...interface{})
}

type restartStrategy func(ctx _TaskContext, err error)

var supPool = &sync.Pool{
	New: func() interface{} {
		return &taskSupervisor{}
	},
}

func newSupervisor(tasks []Task, strategy restartStrategy) *taskSupervisor {
	var sup, ok = supPool.Get().(*taskSupervisor)
	if !ok {
		sup = &taskSupervisor{}
	}
	sup.init(tasks, strategy)
	return sup
}

func freeSupervisor(sup *taskSupervisor) {
	sup.reset()
	supPool.Put(sup)
}

func (supervisor *taskSupervisor) init(tasks []Task, strategy restartStrategy) {
	supervisor.errors = make(chan error, len(tasks))
	supervisor.strategy = strategy
	supervisor.tasks = make([]*taskContext, 0, len(tasks))
	for i, t := range tasks {
		supervisor.tasks = append(supervisor.tasks, &taskContext{
			id:      i,
			Context: context.Background(),
			task:    t,
			sup:     supervisor,
		})
	}
}

func (supervisor *taskSupervisor) reset() {
	*supervisor = taskSupervisor{}
}

func (supervisor *taskSupervisor) getStrategy() restartStrategy {
	return supervisor.strategy
}

func (supervisor *taskSupervisor) getErrors() chan error {
	return supervisor.errors
}

func (supervisor *taskSupervisor) done() {
	supervisor.wg.Done()
}

func (supervisor *taskSupervisor) debugf(format string, args ...interface{}) {
	if supervisor.debugLog != nil {
		supervisor.debugLog(format, args...)
	}
}

func (supervisor *taskSupervisor) StopAll() {
	for _, t := range supervisor.tasks {
		t.StopMe()
	}
}

func (supervisor *taskSupervisor) RestartAll() {
	for _, t := range supervisor.tasks {
		t.RestartMe()
	}
}

func (supervisor *taskSupervisor) restartRest(id int) {
	for _, t := range supervisor.tasks[id:] {
		t.RestartMe()
	}
}

func (supervisor *taskSupervisor) Run(ctx context.Context) error {
	supervisor.wg.Add(len(supervisor.tasks))
	supervisor.debugf("starting tasks")
	for _, task := range supervisor.tasks {
		task.Context = ctx
		go task.startMe()
	}
	supervisor.isStarted = true
	return supervisor.CollectErrors(ctx)
}

func (supervisor *taskSupervisor) CollectErrors(ctx context.Context) error {
	if !supervisor.isStarted {
		return ErrSupervisorIsNotRunning
	}
	go func() {
		supervisor.wg.Wait()
		close(supervisor.errors)
	}()

	var errRun error
	for {
		var ok, err = supervisor.awaitError(ctx)
		if !ok {
			break
		}
		multierr.AppendInto(&errRun, err)
	}
	return errRun
}

func (supervisor *taskSupervisor) awaitError(ctx context.Context) (next bool, _ error) {
	select {
	case err, ok := <-supervisor.errors:
		return ok, err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

func (supervisor *taskSupervisor) SetLog(lg func(string, ...interface{})) {
	supervisor.debugLog = lg
}
