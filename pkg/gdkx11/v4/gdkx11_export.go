// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
import "C"

//export _gotk4_gdkx114_X11Screen_ConnectWindowManagerChanged
func _gotk4_gdkx114_X11Screen_ConnectWindowManagerChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
