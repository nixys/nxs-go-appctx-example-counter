package counter

import (
	"time"

	"github.com/nixys/nxs-go-appctx-example-counter/ctx"
	"github.com/nixys/nxs-go-appctx-example-counter/modules/counter"

	appctx "github.com/nixys/nxs-go-appctx/v3"
)

func Runtime(app appctx.App) error {

	cc := app.ValueGet().(*ctx.Ctx)

	count := counter.Init()

	count.Inc().Print(cc.Log)

	timer := time.NewTimer(cc.CounterInterval)

	for {
		select {
		case <-app.SelfCtxDone():
			cc.Log.Debugf("counter: shutdown")
			return nil
		case <-timer.C:
			count.Inc().Print(cc.Log)
			timer.Reset(cc.CounterInterval)
		}
	}
}
