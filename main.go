package main

import (
	"os"
	"syscall"

	"github.com/nixys/nxs-go-appctx-example-counter/ctx"
	"github.com/nixys/nxs-go-appctx-example-counter/misc"
	"github.com/nixys/nxs-go-appctx-example-counter/routines/counter"

	appctx "github.com/nixys/nxs-go-appctx/v3"
)

func main() {

	err := appctx.Init(nil).
		RoutinesSet(
			map[string]appctx.RoutineParam{
				"counter": {
					Handler: counter.Runtime,
				},
			},
		).
		ValueInitHandlerSet(ctx.AppCtxInit).
		SignalsSet([]appctx.SignalsParam{
			{
				Signals: []os.Signal{
					syscall.SIGTERM,
				},
				Handler: sigHandlerTerm,
			},
		}).
		Run()
	if err != nil {
		switch err {
		case misc.ErrArgSuccessExit:
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
}

func sigHandlerTerm(sig appctx.Signal) {
	sig.Shutdown(nil)
}
