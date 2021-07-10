// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKey},
	})
}

// EventControllerKey is an event controller meant for situations where you need
// access to key events.
//
// This object was added in 3.24.
type EventControllerKey interface {
	gextras.Objector

	Forward(widget Widget) bool
	Group() uint
	// ImContext gets the IM context of a key controller.
	ImContext() *IMContextClass
	SetImContext(imContext IMContext)
}

// EventControllerKeyClass implements the EventControllerKey interface.
type EventControllerKeyClass struct {
	EventControllerClass
}

var _ EventControllerKey = (*EventControllerKeyClass)(nil)

func wrapEventControllerKey(obj *externglib.Object) EventControllerKey {
	return &EventControllerKeyClass{
		EventControllerClass: EventControllerClass{
			Object: obj,
		},
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerKey(obj), nil
}

func NewEventControllerKey(widget Widget) *EventControllerKeyClass {
	var _arg1 *C.GtkWidget          // out
	var _cret *C.GtkEventController // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_key_new(_arg1)

	var _eventControllerKey *EventControllerKeyClass // out

	_eventControllerKey = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*EventControllerKeyClass)

	return _eventControllerKey
}

func (controller *EventControllerKeyClass) Forward(widget Widget) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (controller *EventControllerKeyClass) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret C.guint                  // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_group(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ImContext gets the IM context of a key controller.
func (controller *EventControllerKeyClass) ImContext() *IMContextClass {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret *C.GtkIMContext          // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)

	var _imContext *IMContextClass // out

	_imContext = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*IMContextClass)

	return _imContext
}

func (controller *EventControllerKeyClass) SetImContext(imContext IMContext) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
}
