// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_scroll_get_type()), F: marshalEventControllerScroll},
	})
}

// EventControllerScroll: `GtkEventControllerScroll` is an event controller that
// handles scroll events.
//
// It is capable of handling both discrete and continuous scroll events from
// mice or touchpads, abstracting them both with the
// [signal@Gtk.EventControllerScroll::scroll] signal. Deltas in the discrete
// case are multiples of 1.
//
// In the case of continuous scroll events, `GtkEventControllerScroll` encloses
// all [signal@Gtk.EventControllerScroll::scroll] emissions between two
// [signal@Gtk.EventControllerScroll::scroll-begin] and
// [signal@Gtk.EventControllerScroll::scroll-end] signals.
//
// The behavior of the event controller can be modified by the flags given at
// creation time, or modified at a later point through
// [method@Gtk.EventControllerScroll.set_flags] (e.g. because the scrolling
// conditions of the widget changed).
//
// The controller can be set up to emit motion for either/both vertical and
// horizontal scroll events through GTK_EVENT_CONTROLLER_SCROLL_VERTICAL,
// GTK_EVENT_CONTROLLER_SCROLL_HORIZONTAL and
// GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES. If any axis is disabled, the
// respective [signal@Gtk.EventControllerScroll::scroll] delta will be 0.
// Vertical scroll events will be translated to horizontal motion for the
// devices incapable of horizontal scrolling.
//
// The event controller can also be forced to emit discrete events on all
// devices through GTK_EVENT_CONTROLLER_SCROLL_DISCRETE. This can be used to
// implement discrete actions triggered through scroll events (e.g. switching
// across combobox options).
//
// The GTK_EVENT_CONTROLLER_SCROLL_KINETIC flag toggles the emission of the
// [signal@Gtk.EventControllerScroll::decelerate] signal, emitted at the end of
// scrolling with two X/Y velocity arguments that are consistent with the motion
// that was received.
type EventControllerScroll interface {
	EventController

	// Flags gets the flags conditioning the scroll controller behavior.
	Flags() EventControllerScrollFlags
	// SetFlags sets the flags conditioning scroll controller behavior.
	SetFlags(flags EventControllerScrollFlags)
}

// eventControllerScroll implements the EventControllerScroll interface.
type eventControllerScroll struct {
	EventController
}

var _ EventControllerScroll = (*eventControllerScroll)(nil)

// WrapEventControllerScroll wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerScroll(obj *externglib.Object) EventControllerScroll {
	return EventControllerScroll{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerScroll(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerScroll(obj), nil
}

// NewEventControllerScroll constructs a class EventControllerScroll.
func NewEventControllerScroll(flags EventControllerScrollFlags) EventControllerScroll {
	var _arg1 C.GtkEventControllerScrollFlags

	_arg1 = (C.GtkEventControllerScrollFlags)(flags)

	var _cret C.GtkEventControllerScroll

	cret = C.gtk_event_controller_scroll_new(_arg1)

	var _eventControllerScroll EventControllerScroll

	_eventControllerScroll = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(EventControllerScroll)

	return _eventControllerScroll
}

// Flags gets the flags conditioning the scroll controller behavior.
func (s eventControllerScroll) Flags() EventControllerScrollFlags {
	var _arg0 *C.GtkEventControllerScroll

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(s.Native()))

	var _cret C.GtkEventControllerScrollFlags

	cret = C.gtk_event_controller_scroll_get_flags(_arg0)

	var _eventControllerScrollFlags EventControllerScrollFlags

	_eventControllerScrollFlags = EventControllerScrollFlags(_cret)

	return _eventControllerScrollFlags
}

// SetFlags sets the flags conditioning scroll controller behavior.
func (s eventControllerScroll) SetFlags(flags EventControllerScrollFlags) {
	var _arg0 *C.GtkEventControllerScroll
	var _arg1 C.GtkEventControllerScrollFlags

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkEventControllerScrollFlags)(flags)

	C.gtk_event_controller_scroll_set_flags(_arg0, _arg1)
}
