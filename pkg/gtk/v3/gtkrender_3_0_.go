// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// RenderBackground renders the background of an element.
//
// Typical background rendering, showing the effect of background-image,
// border-width and border-radius:
//
// ! (background.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderBackground(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_background(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}
