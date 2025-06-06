//go:build windows && !go1.21

package screenshot

import (
	"image"
	"syscall"
	"unsafe"

	"github.com/tailscale/win"
)

func NumActiveDisplays() int {
	var count int
	count = 0
	ptr := unsafe.Pointer(&count)
	enumDisplayMonitors(win.HDC(0), nil, syscall.NewCallback(countupMonitorCallback), uintptr(ptr))
	return count
}

func GetDisplayBounds(displayIndex int) image.Rectangle {
	var ctx getMonitorBoundsContext
	ctx.Index = displayIndex
	ctx.Count = 0
	ptr := unsafe.Pointer(&ctx)
	enumDisplayMonitors(win.HDC(0), nil, syscall.NewCallback(getMonitorBoundsCallback), uintptr(ptr))
	return image.Rect(
		int(ctx.Rect.Left), int(ctx.Rect.Top),
		int(ctx.Rect.Right), int(ctx.Rect.Bottom))
}
