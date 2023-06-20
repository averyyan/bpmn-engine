package engine

import (
	"context"
	"time"
)

func newPIContext(ctx context.Context) *piContext {
	_ctx, cannel := context.WithCancel(ctx)
	return &piContext{
		ctx:    _ctx,
		cannel: cannel,
	}
}

type piContext struct {
	ctx    context.Context
	cannel context.CancelFunc
}

func (pic *piContext) Completed() {
	pic.cannel()
}

func (pic *piContext) Deadline() (deadline time.Time, ok bool) {
	return pic.ctx.Deadline()
}

func (pic *piContext) Done() <-chan struct{} {
	return pic.ctx.Done()
}

func (pic *piContext) Err() error {
	return pic.ctx.Err()
}

func (pic *piContext) Value(key any) any {
	return pic.ctx.Value(key)
}
