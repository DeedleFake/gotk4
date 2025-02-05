// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_GestureRotate_ConnectAngleChanged(gpointer, gdouble, gdouble, guintptr);
import "C"

// GType values.
var (
	GTypeGestureRotate = coreglib.Type(C.gtk_gesture_rotate_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGestureRotate, F: marshalGestureRotate},
	})
}

// GestureRotate: GtkGestureRotate is a GtkGesture for 2-finger rotations.
//
// Whenever the angle between both handled sequences changes, the
// gtk.GestureRotate::angle-changed signal is emitted.
type GestureRotate struct {
	_ [0]func() // equal guard
	Gesture
}

var (
	_ Gesturer = (*GestureRotate)(nil)
)

func wrapGestureRotate(obj *coreglib.Object) *GestureRotate {
	return &GestureRotate{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureRotate(p uintptr) (interface{}, error) {
	return wrapGestureRotate(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAngleChanged is emitted when the angle between both tracked points
// changes.
func (gesture *GestureRotate) ConnectAngleChanged(f func(angle, angleDelta float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "angle-changed", false, unsafe.Pointer(C._gotk4_gtk4_GestureRotate_ConnectAngleChanged), f)
}

// NewGestureRotate returns a newly created GtkGesture that recognizes 2-touch
// rotation gestures.
//
// The function returns the following values:
//
//    - gestureRotate: newly created GtkGestureRotate.
//
func NewGestureRotate() *GestureRotate {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_rotate_new()

	var _gestureRotate *GestureRotate // out

	_gestureRotate = wrapGestureRotate(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureRotate
}

// AngleDelta gets the angle delta in radians.
//
// If gesture is active, this function returns the angle difference in radians
// since the gesture was first recognized. If gesture is not active, 0 is
// returned.
//
// The function returns the following values:
//
//    - gdouble: angle delta in radians.
//
func (gesture *GestureRotate) AngleDelta() float64 {
	var _arg0 *C.GtkGestureRotate // out
	var _cret C.double            // in

	_arg0 = (*C.GtkGestureRotate)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_cret = C.gtk_gesture_rotate_get_angle_delta(_arg0)
	runtime.KeepAlive(gesture)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
