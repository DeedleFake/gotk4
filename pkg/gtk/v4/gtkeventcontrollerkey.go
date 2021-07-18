// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKeyer},
	})
}

// EventControllerKey: GtkEventControllerKey is an event controller that
// provides access to key events.
type EventControllerKey struct {
	EventController
}

var _ gextras.Nativer = (*EventControllerKey)(nil)

func wrapEventControllerKey(obj *externglib.Object) *EventControllerKey {
	return &EventControllerKey{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerKeyer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerKey(obj), nil
}

// NewEventControllerKey creates a new event controller that will handle key
// events.
func NewEventControllerKey() *EventControllerKey {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_key_new()

	var _eventControllerKey *EventControllerKey // out

	_eventControllerKey = wrapEventControllerKey(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerKey
}

// Forward forwards the current event of this controller to a widget.
//
// This function can only be used in handlers for the
// gtk.EventControllerKey::key-pressed, gtk.EventControllerKey::key-released or
// gtk.EventControllerKey::modifiers signals.
func (controller *EventControllerKey) Forward(widget Widgeter) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Group gets the key group of the current event of this controller.
//
// See gdk.KeyEvent.GetLayout().
func (controller *EventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret C.guint                  // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_group(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ImContext gets the input method context of the key controller.
func (controller *EventControllerKey) ImContext() IMContexter {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret *C.GtkIMContext          // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)

	var _imContext IMContexter // out

	_imContext = (*gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(IMContexter)

	return _imContext
}

// SetImContext sets the input method context of the key controller.
func (controller *EventControllerKey) SetImContext(imContext IMContexter) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer((imContext).(gextras.Nativer).Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
}
