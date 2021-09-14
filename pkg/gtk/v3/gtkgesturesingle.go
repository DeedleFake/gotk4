// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_single_get_type()), F: marshalGestureSingler},
	})
}

// GestureSingle is a subclass of Gesture, optimized (although not restricted)
// for dealing with mouse and single-touch gestures. Under interaction, these
// gestures stick to the first interacting sequence, which is accessible through
// gtk_gesture_single_get_current_sequence() while the gesture is being
// interacted with.
//
// By default gestures react to both GDK_BUTTON_PRIMARY and touch events,
// gtk_gesture_single_set_touch_only() can be used to change the touch behavior.
// Callers may also specify a different mouse button number to interact with
// through gtk_gesture_single_set_button(), or react to any mouse button by
// setting 0. While the gesture is active, the button being currently pressed
// can be known through gtk_gesture_single_get_current_button().
type GestureSingle struct {
	Gesture
}

func wrapGestureSingle(obj *externglib.Object) *GestureSingle {
	return &GestureSingle{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureSingler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureSingle(obj), nil
}

// Button returns the button number gesture listens for, or 0 if gesture reacts
// to any button press.
func (gesture *GestureSingle) Button() uint {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_single_get_button(_arg0)
	runtime.KeepAlive(gesture)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CurrentButton returns the button number currently interacting with gesture,
// or 0 if there is none.
func (gesture *GestureSingle) CurrentButton() uint {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_single_get_current_button(_arg0)
	runtime.KeepAlive(gesture)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CurrentSequence returns the event sequence currently interacting with
// gesture. This is only meaningful if gtk_gesture_is_active() returns TRUE.
func (gesture *GestureSingle) CurrentSequence() *gdk.EventSequence {
	var _arg0 *C.GtkGestureSingle // out
	var _cret *C.GdkEventSequence // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_single_get_current_sequence(_arg0)
	runtime.KeepAlive(gesture)

	var _eventSequence *gdk.EventSequence // out

	if _cret != nil {
		_eventSequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_eventSequence)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _eventSequence
}

// Exclusive gets whether a gesture is exclusive. For more information, see
// gtk_gesture_single_set_exclusive().
func (gesture *GestureSingle) Exclusive() bool {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_single_get_exclusive(_arg0)
	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TouchOnly returns TRUE if the gesture is only triggered by touch events.
func (gesture *GestureSingle) TouchOnly() bool {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_single_get_touch_only(_arg0)
	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetButton sets the button number gesture listens to. If non-0, every button
// press from a different button number will be ignored. Touch events implicitly
// match with button 1.
func (gesture *GestureSingle) SetButton(button uint) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.guint             // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.guint(button)

	C.gtk_gesture_single_set_button(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(button)
}

// SetExclusive sets whether gesture is exclusive. An exclusive gesture will
// only handle pointer and "pointer emulated" touch events, so at any given
// time, there is only one sequence able to interact with those.
func (gesture *GestureSingle) SetExclusive(exclusive bool) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))
	if exclusive {
		_arg1 = C.TRUE
	}

	C.gtk_gesture_single_set_exclusive(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(exclusive)
}

// SetTouchOnly: if touch_only is TRUE, gesture will only handle events of type
// K_TOUCH_BEGIN, K_TOUCH_UPDATE or K_TOUCH_END. If FALSE, mouse events will be
// handled too.
func (gesture *GestureSingle) SetTouchOnly(touchOnly bool) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(gesture.Native()))
	if touchOnly {
		_arg1 = C.TRUE
	}

	C.gtk_gesture_single_set_touch_only(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(touchOnly)
}
