// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_GestureLongPress_ConnectCancelled(gpointer, guintptr);
// extern void _gotk4_gtk3_GestureLongPress_ConnectPressed(gpointer, gdouble, gdouble, guintptr);
import "C"

// GTypeGestureLongPress returns the GType for the type GestureLongPress.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGestureLongPress() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "GestureLongPress").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGestureLongPress)
	return gtype
}

// GestureLongPress is a Gesture implementation able to recognize long presses,
// triggering the GestureLongPress::pressed after the timeout is exceeded.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the GestureLongPress::cancelled signal will
// be emitted.
type GestureLongPress struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureLongPress)(nil)
)

func wrapGestureLongPress(obj *coreglib.Object) *GestureLongPress {
	return &GestureLongPress{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureLongPress(p uintptr) (interface{}, error) {
	return wrapGestureLongPress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_GestureLongPress_ConnectCancelled
func _gotk4_gtk3_GestureLongPress_ConnectCancelled(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectCancelled: this signal is emitted whenever a press moved too far, or
// was released before GestureLongPress::pressed happened.
func (v *GestureLongPress) ConnectCancelled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "cancelled", false, unsafe.Pointer(C._gotk4_gtk3_GestureLongPress_ConnectCancelled), f)
}

//export _gotk4_gtk3_GestureLongPress_ConnectPressed
func _gotk4_gtk3_GestureLongPress_ConnectPressed(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectPressed: this signal is emitted whenever a press goes
// unmoved/unreleased longer than what the GTK+ defaults tell.
func (v *GestureLongPress) ConnectPressed(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "pressed", false, unsafe.Pointer(C._gotk4_gtk3_GestureLongPress_ConnectPressed), f)
}

// NewGestureLongPress returns a newly created Gesture that recognizes long
// presses.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureLongPress: newly created GestureLongPress.
//
func NewGestureLongPress(widget Widgetter) *GestureLongPress {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_info := girepository.MustFind("Gtk", "GestureLongPress")
	_gret := _info.InvokeClassMethod("new_GestureLongPress", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _gestureLongPress *GestureLongPress // out

	_gestureLongPress = wrapGestureLongPress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureLongPress
}
