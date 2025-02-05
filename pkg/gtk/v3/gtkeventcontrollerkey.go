// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_EventControllerKey_ConnectKeyReleased(gpointer, guint, guint, GdkModifierType, guintptr);
// extern void _gotk4_gtk3_EventControllerKey_ConnectIMUpdate(gpointer, guintptr);
// extern void _gotk4_gtk3_EventControllerKey_ConnectFocusOut(gpointer, guintptr);
// extern void _gotk4_gtk3_EventControllerKey_ConnectFocusIn(gpointer, guintptr);
// extern gboolean _gotk4_gtk3_EventControllerKey_ConnectModifiers(gpointer, GdkModifierType, guintptr);
// extern gboolean _gotk4_gtk3_EventControllerKey_ConnectKeyPressed(gpointer, guint, guint, GdkModifierType, guintptr);
import "C"

// GType values.
var (
	GTypeEventControllerKey = coreglib.Type(C.gtk_event_controller_key_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEventControllerKey, F: marshalEventControllerKey},
	})
}

// EventControllerKey is an event controller meant for situations where you need
// access to key events.
//
// This object was added in 3.24.
type EventControllerKey struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerKey)(nil)
)

func wrapEventControllerKey(obj *coreglib.Object) *EventControllerKey {
	return &EventControllerKey{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	return wrapEventControllerKey(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (controller *EventControllerKey) ConnectFocusIn(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "focus-in", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectFocusIn), f)
}

func (controller *EventControllerKey) ConnectFocusOut(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "focus-out", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectFocusOut), f)
}

func (controller *EventControllerKey) ConnectIMUpdate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "im-update", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectIMUpdate), f)
}

// ConnectKeyPressed: this signal is emitted whenever a key is pressed.
func (controller *EventControllerKey) ConnectKeyPressed(f func(keyval, keycode uint, state gdk.ModifierType) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "key-pressed", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectKeyPressed), f)
}

// ConnectKeyReleased: this signal is emitted whenever a key is released.
func (controller *EventControllerKey) ConnectKeyReleased(f func(keyval, keycode uint, state gdk.ModifierType)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "key-released", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectKeyReleased), f)
}

func (controller *EventControllerKey) ConnectModifiers(f func(object gdk.ModifierType) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "modifiers", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectModifiers), f)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewEventControllerKey(widget Widgetter) *EventControllerKey {
	var _arg1 *C.GtkWidget          // out
	var _cret *C.GtkEventController // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_event_controller_key_new(_arg1)
	runtime.KeepAlive(widget)

	var _eventControllerKey *EventControllerKey // out

	_eventControllerKey = wrapEventControllerKey(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerKey
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (controller *EventControllerKey) Forward(widget Widgetter) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(coreglib.InternObject(controller).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
func (controller *EventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret C.guint                  // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(coreglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_key_get_group(_arg0)
	runtime.KeepAlive(controller)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IMContext gets the IM context of a key controller.
//
// The function returns the following values:
//
//    - imContext: IM context.
//
func (controller *EventControllerKey) IMContext() IMContexter {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret *C.GtkIMContext          // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(coreglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)
	runtime.KeepAlive(controller)

	var _imContext IMContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.IMContexter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(IMContexter)
			return ok
		})
		rv, ok := casted.(IMContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.IMContexter")
		}
		_imContext = rv
	}

	return _imContext
}

// The function takes the following parameters:
//
func (controller *EventControllerKey) SetIMContext(imContext IMContexter) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(coreglib.InternObject(controller).Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(coreglib.InternObject(imContext).Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(imContext)
}
