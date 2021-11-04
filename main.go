// Copyright (c) 2021 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/golang/glog"
)

var (
	duration = 30 * time.Second
	amountX  = 1
	amountY  = 1
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())
	_ = flag.Set("logtostderr", "true")
	flag.Parse()

	glog.V(0).Infof("started")
	ctx := contextWithSig(context.Background())

	var x, y int
	for {
		select {
		case <-ctx.Done():
			glog.V(0).Infof("done")
			return
		case <-time.NewTimer(duration).C:
			lastX := x
			lastY := y
			x, y = robotgo.GetMousePos()

			if x != lastX || y != lastY {
				glog.V(0).Infof("skip move")
				continue
			}
			robotgo.MoveMouse(x+amountX, y+amountY)
			robotgo.MoveMouse(x, y)
			glog.V(0).Infof("moved")
		}
	}
}

func contextWithSig(ctx context.Context) context.Context {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-signalCh:
		case <-ctx.Done():
		}
	}()

	return ctxWithCancel
}
