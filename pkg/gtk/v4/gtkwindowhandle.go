// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_window_handle_get_type()), F: marshalWindowHandler},
	})
}

// WindowHandle: GtkWindowHandle is a titlebar area widget.
//
// When added into a window, it can be dragged to move the window, and handles
// right click, double click and middle click as expected of a titlebar.
//
//
// CSS nodes
//
// GtkWindowHandle has a single CSS node with the name windowhandle.
//
//
// Accessibility
//
// GtkWindowHandle uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowHandle struct {
	Widget
}

var (
	_ Widgetter = (*WindowHandle)(nil)
)

func wrapWindowHandle(obj *externglib.Object) *WindowHandle {
	return &WindowHandle{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalWindowHandler(p uintptr) (interface{}, error) {
	return wrapWindowHandle(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowHandle creates a new GtkWindowHandle.
//
// The function returns the following values:
//
//    - windowHandle: new GtkWindowHandle.
//
func NewWindowHandle() *WindowHandle {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_window_handle_new()

	var _windowHandle *WindowHandle // out

	_windowHandle = wrapWindowHandle(externglib.Take(unsafe.Pointer(_cret)))

	return _windowHandle
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//    - widget (optional): child widget of self.
//
func (self *WindowHandle) Child() Widgetter {
	var _arg0 *C.GtkWindowHandle // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkWindowHandle)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_window_handle_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (self *WindowHandle) SetChild(child Widgetter) {
	var _arg0 *C.GtkWindowHandle // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkWindowHandle)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_window_handle_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}
