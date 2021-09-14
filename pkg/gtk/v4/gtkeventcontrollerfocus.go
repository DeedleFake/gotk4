// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_focus_get_type()), F: marshalEventControllerFocusser},
	})
}

// EventControllerFocus: GtkEventControllerFocus is an event controller to keep
// track of keyboard focus.
//
// The event controller offers gtk.EventControllerFocus::enter and
// gtk.EventControllerFocus::leave signals, as well as
// gtk.EventControllerFocus:is-focus and gtk.EventControllerFocus:contains-focus
// properties which are updated to reflect focus changes inside the widget
// hierarchy that is rooted at the controllers widget.
type EventControllerFocus struct {
	EventController
}

func wrapEventControllerFocus(obj *externglib.Object) *EventControllerFocus {
	return &EventControllerFocus{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerFocusser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerFocus(obj), nil
}

// NewEventControllerFocus creates a new event controller that will handle focus
// events.
func NewEventControllerFocus() *EventControllerFocus {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_focus_new()

	var _eventControllerFocus *EventControllerFocus // out

	_eventControllerFocus = wrapEventControllerFocus(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerFocus
}

// ContainsFocus returns TRUE if focus is within self or one of its children.
func (self *EventControllerFocus) ContainsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_focus_contains_focus(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsFocus returns TRUE if focus is within self, but not one of its children.
func (self *EventControllerFocus) IsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_focus_is_focus(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
