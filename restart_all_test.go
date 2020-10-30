package ger_test

import (
	"context"
	"sync/atomic"
	"testing"

	"github.com/ninedraft/ger"

	"github.com/stretchr/testify/assert"
)

func TestRestartAll(test *testing.T) {
	var ctx = context.Background()
	var counter = new(int32)
	_ = ger.AllForOne(ctx,
		func(ctx context.Context) error {
			var c = atomic.AddInt32(counter, 1)
			if c < 10 {
				return ger.ErrRestart
			}
			return nil
		},
	)
	assert.EqualValues(test, 10, atomic.LoadInt32(counter), "restart counter")
}
