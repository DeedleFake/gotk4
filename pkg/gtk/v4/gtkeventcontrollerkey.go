// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_EventControllerKey_ConnectKeyPressed(gpointer, guint, guint, GdkModifierType, guintptr);
// extern gboolean _gotk4_gtk4_EventControllerKey_ConnectModifiers(gpointer, GdkModifierType, guintptr);
// extern void _gotk4_gtk4_EventControllerKey_ConnectIMUpdate(gpointer, guintptr);
// extern void _gotk4_gtk4_EventControllerKey_ConnectKeyReleased(gpointer, guint, guint, GdkModifierType, guintptr);
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

// EventControllerKeyOverrider contains methods that are overridable.
type EventControllerKeyOverrider interface {
}

// EventControllerKey: GtkEventControllerKey is an event controller that
// provides access to key events.
type EventControllerKey struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerKey)(nil)
)

func initClassEventControllerKey(gclass unsafe.Pointer, goval any) {
}

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

//export _gotk4_gtk4_EventControllerKey_ConnectIMUpdate
func _gotk4_gtk4_EventControllerKey_ConnectIMUpdate(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectIMUpdate is emitted whenever the input method context filters away a
// keypress and prevents the controller receiving it.
//
// See gtk.EventControllerKey.SetIMContext() and gtk.IMContext.FilterKeypress().
func (controller *EventControllerKey) ConnectIMUpdate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "im-update", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerKey_ConnectIMUpdate), f)
}

//export _gotk4_gtk4_EventControllerKey_ConnectKeyPressed
func _gotk4_gtk4_EventControllerKey_ConnectKeyPressed(arg0 C.gpointer, arg1 C.guint, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.guintptr) (cret C.gboolean) {
	var f func(keyval, keycode uint, state gdk.ModifierType) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(keyval, keycode uint, state gdk.ModifierType) (ok bool))
	}

	var _keyval uint            // out
	var _keycode uint           // out
	var _state gdk.ModifierType // out

	_keyval = uint(arg1)
	_keycode = uint(arg2)
	_state = gdk.ModifierType(arg3)

	ok := f(_keyval, _keycode, _state)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectKeyPressed is emitted whenever a key is pressed.
func (controller *EventControllerKey) ConnectKeyPressed(f func(keyval, keycode uint, state gdk.ModifierType) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "key-pressed", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerKey_ConnectKeyPressed), f)
}

//export _gotk4_gtk4_EventControllerKey_ConnectKeyReleased
func _gotk4_gtk4_EventControllerKey_ConnectKeyReleased(arg0 C.gpointer, arg1 C.guint, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.guintptr) {
	var f func(keyval, keycode uint, state gdk.ModifierType)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(keyval, keycode uint, state gdk.ModifierType))
	}

	var _keyval uint            // out
	var _keycode uint           // out
	var _state gdk.ModifierType // out

	_keyval = uint(arg1)
	_keycode = uint(arg2)
	_state = gdk.ModifierType(arg3)

	f(_keyval, _keycode, _state)
}

// ConnectKeyReleased is emitted whenever a key is released.
func (controller *EventControllerKey) ConnectKeyReleased(f func(keyval, keycode uint, state gdk.ModifierType)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "key-released", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerKey_ConnectKeyReleased), f)
}

//export _gotk4_gtk4_EventControllerKey_ConnectModifiers
func _gotk4_gtk4_EventControllerKey_ConnectModifiers(arg0 C.gpointer, arg1 C.GdkModifierType, arg2 C.guintptr) (cret C.gboolean) {
	var f func(keyval gdk.ModifierType) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(keyval gdk.ModifierType) (ok bool))
	}

	var _keyval gdk.ModifierType // out

	_keyval = gdk.ModifierType(arg1)

	ok := f(_keyval)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectModifiers is emitted whenever the state of modifier keys and pointer
// buttons change.
func (controller *EventControllerKey) ConnectModifiers(f func(keyval gdk.ModifierType) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "modifiers", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerKey_ConnectModifiers), f)
}

// NewEventControllerKey creates a new event controller that will handle key
// events.
//
// The function returns the following values:
//
//    - eventControllerKey: new GtkEventControllerKey.
//
func NewEventControllerKey() *EventControllerKey {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_key_new()

	var _eventControllerKey *EventControllerKey // out

	_eventControllerKey = wrapEventControllerKey(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerKey
}

// Forward forwards the current event of this controller to a widget.
//
// This function can only be used in handlers for the
// gtk.EventControllerKey::key-pressed, gtk.EventControllerKey::key-released or
// gtk.EventControllerKey::modifiers signals.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget.
//
// The function returns the following values:
//
//    - ok: whether the widget handled the event.
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

// Group gets the key group of the current event of this controller.
//
// See gdk.KeyEvent.GetLayout().
//
// The function returns the following values:
//
//    - guint: key group.
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

// IMContext gets the input method context of the key controller.
//
// The function returns the following values:
//
//    - imContext: GtkIMContext.
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

// SetIMContext sets the input method context of the key controller.
//
// The function takes the following parameters:
//
//    - imContext: GtkIMContext.
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
