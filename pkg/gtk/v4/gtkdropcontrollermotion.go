// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_DropControllerMotion_ConnectEnter(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_DropControllerMotion_ConnectLeave(gpointer, guintptr);
// extern void _gotk4_gtk4_DropControllerMotion_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_controller_motion_get_type()), F: marshalDropControllerMotioner},
	})
}

// DropControllerMotionOverrider contains methods that are overridable.
type DropControllerMotionOverrider interface {
}

// DropControllerMotion: GtkDropControllerMotion is an event controller tracking
// the pointer during Drag-and-Drop operations.
//
// It is modeled after gtk.EventControllerMotion so if you have used that, this
// should feel really familiar.
//
// This controller is not able to accept drops, use gtk.DropTarget for that
// purpose.
type DropControllerMotion struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*DropControllerMotion)(nil)
)

func classInitDropControllerMotioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDropControllerMotion(obj *externglib.Object) *DropControllerMotion {
	return &DropControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalDropControllerMotioner(p uintptr) (interface{}, error) {
	return wrapDropControllerMotion(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_DropControllerMotion_ConnectEnter
func _gotk4_gtk4_DropControllerMotion_ConnectEnter(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
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

// ConnectEnter signals that the pointer has entered the widget.
func (self *DropControllerMotion) ConnectEnter(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "enter", false, unsafe.Pointer(C._gotk4_gtk4_DropControllerMotion_ConnectEnter), f)
}

//export _gotk4_gtk4_DropControllerMotion_ConnectLeave
func _gotk4_gtk4_DropControllerMotion_ConnectLeave(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLeave signals that the pointer has left the widget.
func (self *DropControllerMotion) ConnectLeave(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "leave", false, unsafe.Pointer(C._gotk4_gtk4_DropControllerMotion_ConnectLeave), f)
}

//export _gotk4_gtk4_DropControllerMotion_ConnectMotion
func _gotk4_gtk4_DropControllerMotion_ConnectMotion(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
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

// ConnectMotion: emitted when the pointer moves inside the widget.
func (self *DropControllerMotion) ConnectMotion(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "motion", false, unsafe.Pointer(C._gotk4_gtk4_DropControllerMotion_ConnectMotion), f)
}

// NewDropControllerMotion creates a new event controller that will handle
// pointer motion events during drag and drop.
//
// The function returns the following values:
//
//    - dropControllerMotion: new GtkDropControllerMotion.
//
func NewDropControllerMotion() *DropControllerMotion {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_drop_controller_motion_new()

	var _dropControllerMotion *DropControllerMotion // out

	_dropControllerMotion = wrapDropControllerMotion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dropControllerMotion
}

// ContainsPointer returns if a Drag-and-Drop operation is within the widget
// self or one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if a dragging pointer is within self or one of its children.
//
func (self *DropControllerMotion) ContainsPointer() bool {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_contains_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Drop returns the GdkDrop of a current Drag-and-Drop operation over the widget
// of self.
//
// The function returns the following values:
//
//    - drop (optional): GdkDrop currently happening within self or NULL if none.
//
func (self *DropControllerMotion) Drop() gdk.Dropper {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret *C.GdkDrop                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_get_drop(_arg0)
	runtime.KeepAlive(self)

	var _drop gdk.Dropper // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gdk.Dropper)
				return ok
			})
			rv, ok := casted.(gdk.Dropper)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dropper")
			}
			_drop = rv
		}
	}

	return _drop
}

// IsPointer returns if a Drag-and-Drop operation is within the widget self, not
// one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if a dragging pointer is within self but not one of its
//      children.
//
func (self *DropControllerMotion) IsPointer() bool {
	var _arg0 *C.GtkDropControllerMotion // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_controller_motion_is_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
