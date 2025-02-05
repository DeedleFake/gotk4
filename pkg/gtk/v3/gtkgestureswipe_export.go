// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_GestureSwipe_ConnectSwipe
func _gotk4_gtk3_GestureSwipe_ConnectSwipe(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(velocityX, velocityY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(velocityX, velocityY float64))
	}

	var _velocityX float64 // out
	var _velocityY float64 // out

	_velocityX = float64(arg1)
	_velocityY = float64(arg2)

	f(_velocityX, _velocityY)
}
