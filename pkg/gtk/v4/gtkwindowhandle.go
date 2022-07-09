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
import "C"

// GTypeWindowHandle returns the GType for the type WindowHandle.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeWindowHandle() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "WindowHandle").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalWindowHandle)
	return gtype
}

// WindowHandleOverrider contains methods that are overridable.
type WindowHandleOverrider interface {
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
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*WindowHandle)(nil)
)

func classInitWindowHandler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWindowHandle(obj *coreglib.Object) *WindowHandle {
	return &WindowHandle{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalWindowHandle(p uintptr) (interface{}, error) {
	return wrapWindowHandle(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowHandle creates a new GtkWindowHandle.
//
// The function returns the following values:
//
//    - windowHandle: new GtkWindowHandle.
//
func NewWindowHandle() *WindowHandle {
	_info := girepository.MustFind("Gtk", "WindowHandle")
	_gret := _info.InvokeClassMethod("new_WindowHandle", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _windowHandle *WindowHandle // out

	_windowHandle = wrapWindowHandle(coreglib.Take(unsafe.Pointer(_cret)))

	return _windowHandle
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//    - widget (optional): child widget of self.
//
func (self *WindowHandle) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "WindowHandle")
	_gret := _info.InvokeClassMethod("get_child", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	_info := girepository.MustFind("Gtk", "WindowHandle")
	_info.InvokeClassMethod("set_child", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}
