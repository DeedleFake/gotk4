// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// BindingsActivateEvent looks up key bindings for object to find one matching
// event, and if one was found, activate it.
//
// The function takes the following parameters:
//
//    - object (generally must be a widget).
//    - event: EventKey.
//
// The function returns the following values:
//
//    - ok: TRUE if a matching key binding was found.
//
func BindingsActivateEvent(object *coreglib.Object, event *gdk.EventKey) bool {
	var _arg1 *C.GObject     // out
	var _arg2 *C.GdkEventKey // out
	var _cret C.gboolean     // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.GdkEventKey)(gextras.StructNative(unsafe.Pointer(event)))

	_cret = C.gtk_bindings_activate_event(_arg1, _arg2)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
