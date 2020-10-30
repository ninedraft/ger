// Code generated by moq;
// github.com/matryer/moq

package ger

import (
	"sync"
)

// Ensure, that tSupMock does implement tSup.
// If this is not the case, regenerate this file with moq.
var _ tSup = &tSupMock{}

// tSupMock is a mock implementation of tSup.
//
//     func TestSomethingThatUsestSup(t *testing.T) {
//
//         // make and configure a mocked tSup
//         mockedtSup := &tSupMock{
//             RestartAllFunc: func()  {
//                     panic("mock out the RestartAll method")
//             },
//             StopAllFunc: func()  {
//                     panic("mock out the StopAll method")
//             },
//             debugfFunc: func(msg string, args ...interface{})  {
//                     panic("mock out the debugf method")
//             },
//             doneFunc: func()  {
//                     panic("mock out the done method")
//             },
//             getErrorsFunc: func() chan error {
//                     panic("mock out the getErrors method")
//             },
//             getStrategyFunc: func() restartStrategy {
//                     panic("mock out the getStrategy method")
//             },
//             restartRestFunc: func(id int)  {
//                     panic("mock out the restartRest method")
//             },
//         }
//
//         // use mockedtSup in code that requires tSup
//         // and then make assertions.
//
//     }
type tSupMock struct {
	// RestartAllFunc mocks the RestartAll method.
	RestartAllFunc func()

	// StopAllFunc mocks the StopAll method.
	StopAllFunc func()

	// debugfFunc mocks the debugf method.
	debugfFunc func(msg string, args ...interface{})

	// doneFunc mocks the done method.
	doneFunc func()

	// getErrorsFunc mocks the getErrors method.
	getErrorsFunc func() chan error

	// getStrategyFunc mocks the getStrategy method.
	getStrategyFunc func() restartStrategy

	// restartRestFunc mocks the restartRest method.
	restartRestFunc func(id int)

	// calls tracks calls to the methods.
	calls struct {
		// RestartAll holds details about calls to the RestartAll method.
		RestartAll []struct {
		}
		// StopAll holds details about calls to the StopAll method.
		StopAll []struct {
		}
		// debugf holds details about calls to the debugf method.
		debugf []struct {
			// Msg is the msg argument value.
			Msg string
			// Args is the args argument value.
			Args []interface{}
		}
		// done holds details about calls to the done method.
		done []struct {
		}
		// getErrors holds details about calls to the getErrors method.
		getErrors []struct {
		}
		// getStrategy holds details about calls to the getStrategy method.
		getStrategy []struct {
		}
		// restartRest holds details about calls to the restartRest method.
		restartRest []struct {
			// ID is the id argument value.
			ID int
		}
	}
	lockRestartAll  sync.RWMutex
	lockStopAll     sync.RWMutex
	lockdebugf      sync.RWMutex
	lockdone        sync.RWMutex
	lockgetErrors   sync.RWMutex
	lockgetStrategy sync.RWMutex
	lockrestartRest sync.RWMutex
}

// RestartAll calls RestartAllFunc.
func (mock *tSupMock) RestartAll() {
	if mock.RestartAllFunc == nil {
		panic("tSupMock.RestartAllFunc: method is nil but tSup.RestartAll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRestartAll.Lock()
	mock.calls.RestartAll = append(mock.calls.RestartAll, callInfo)
	mock.lockRestartAll.Unlock()
	mock.RestartAllFunc()
}

// RestartAllCalls gets all the calls that were made to RestartAll.
// Check the length with:
//     len(mockedtSup.RestartAllCalls())
func (mock *tSupMock) RestartAllCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRestartAll.RLock()
	calls = mock.calls.RestartAll
	mock.lockRestartAll.RUnlock()
	return calls
}

// StopAll calls StopAllFunc.
func (mock *tSupMock) StopAll() {
	if mock.StopAllFunc == nil {
		panic("tSupMock.StopAllFunc: method is nil but tSup.StopAll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStopAll.Lock()
	mock.calls.StopAll = append(mock.calls.StopAll, callInfo)
	mock.lockStopAll.Unlock()
	mock.StopAllFunc()
}

// StopAllCalls gets all the calls that were made to StopAll.
// Check the length with:
//     len(mockedtSup.StopAllCalls())
func (mock *tSupMock) StopAllCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStopAll.RLock()
	calls = mock.calls.StopAll
	mock.lockStopAll.RUnlock()
	return calls
}

// debugf calls debugfFunc.
func (mock *tSupMock) debugf(msg string, args ...interface{}) {
	if mock.debugfFunc == nil {
		panic("tSupMock.debugfFunc: method is nil but tSup.debugf was just called")
	}
	callInfo := struct {
		Msg  string
		Args []interface{}
	}{
		Msg:  msg,
		Args: args,
	}
	mock.lockdebugf.Lock()
	mock.calls.debugf = append(mock.calls.debugf, callInfo)
	mock.lockdebugf.Unlock()
	mock.debugfFunc(msg, args...)
}

// debugfCalls gets all the calls that were made to debugf.
// Check the length with:
//     len(mockedtSup.debugfCalls())
func (mock *tSupMock) debugfCalls() []struct {
	Msg  string
	Args []interface{}
} {
	var calls []struct {
		Msg  string
		Args []interface{}
	}
	mock.lockdebugf.RLock()
	calls = mock.calls.debugf
	mock.lockdebugf.RUnlock()
	return calls
}

// done calls doneFunc.
func (mock *tSupMock) done() {
	if mock.doneFunc == nil {
		panic("tSupMock.doneFunc: method is nil but tSup.done was just called")
	}
	callInfo := struct {
	}{}
	mock.lockdone.Lock()
	mock.calls.done = append(mock.calls.done, callInfo)
	mock.lockdone.Unlock()
	mock.doneFunc()
}

// doneCalls gets all the calls that were made to done.
// Check the length with:
//     len(mockedtSup.doneCalls())
func (mock *tSupMock) doneCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockdone.RLock()
	calls = mock.calls.done
	mock.lockdone.RUnlock()
	return calls
}

// getErrors calls getErrorsFunc.
func (mock *tSupMock) getErrors() chan error {
	if mock.getErrorsFunc == nil {
		panic("tSupMock.getErrorsFunc: method is nil but tSup.getErrors was just called")
	}
	callInfo := struct {
	}{}
	mock.lockgetErrors.Lock()
	mock.calls.getErrors = append(mock.calls.getErrors, callInfo)
	mock.lockgetErrors.Unlock()
	return mock.getErrorsFunc()
}

// getErrorsCalls gets all the calls that were made to getErrors.
// Check the length with:
//     len(mockedtSup.getErrorsCalls())
func (mock *tSupMock) getErrorsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockgetErrors.RLock()
	calls = mock.calls.getErrors
	mock.lockgetErrors.RUnlock()
	return calls
}

// getStrategy calls getStrategyFunc.
func (mock *tSupMock) getStrategy() restartStrategy {
	if mock.getStrategyFunc == nil {
		panic("tSupMock.getStrategyFunc: method is nil but tSup.getStrategy was just called")
	}
	callInfo := struct {
	}{}
	mock.lockgetStrategy.Lock()
	mock.calls.getStrategy = append(mock.calls.getStrategy, callInfo)
	mock.lockgetStrategy.Unlock()
	return mock.getStrategyFunc()
}

// getStrategyCalls gets all the calls that were made to getStrategy.
// Check the length with:
//     len(mockedtSup.getStrategyCalls())
func (mock *tSupMock) getStrategyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockgetStrategy.RLock()
	calls = mock.calls.getStrategy
	mock.lockgetStrategy.RUnlock()
	return calls
}

// restartRest calls restartRestFunc.
func (mock *tSupMock) restartRest(id int) {
	if mock.restartRestFunc == nil {
		panic("tSupMock.restartRestFunc: method is nil but tSup.restartRest was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	mock.lockrestartRest.Lock()
	mock.calls.restartRest = append(mock.calls.restartRest, callInfo)
	mock.lockrestartRest.Unlock()
	mock.restartRestFunc(id)
}

// restartRestCalls gets all the calls that were made to restartRest.
// Check the length with:
//     len(mockedtSup.restartRestCalls())
func (mock *tSupMock) restartRestCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockrestartRest.RLock()
	calls = mock.calls.restartRest
	mock.lockrestartRest.RUnlock()
	return calls
}
