// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_zoom_get_type()), F: marshalGestureZoom},
	})
}

// GestureZoom: `GtkGestureZoom` is a `GtkGesture` for 2-finger pinch/zoom
// gestures.
//
// Whenever the distance between both tracked sequences changes, the
// [signal@Gtk.GestureZoom::scale-changed] signal is emitted to report the scale
// factor.
type GestureZoom interface {
	gextras.Objector

	// ScaleDelta gets the scale delta.
	//
	// If @gesture is active, this function returns the zooming difference since
	// the gesture was recognized (hence the starting point is considered 1:1).
	// If @gesture is not active, 1 is returned.
	ScaleDelta() float64
}

// GestureZoomClass implements the GestureZoom interface.
type GestureZoomClass struct {
	GestureClass
}

var _ GestureZoom = (*GestureZoomClass)(nil)

func wrapGestureZoom(obj *externglib.Object) GestureZoom {
	return &GestureZoomClass{
		GestureClass: GestureClass{
			EventControllerClass: EventControllerClass{
				Object: obj,
			},
		},
	}
}

func marshalGestureZoom(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureZoom(obj), nil
}

// NewGestureZoom returns a newly created `GtkGesture` that recognizes
// pinch/zoom gestures.
func NewGestureZoom() *GestureZoomClass {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_zoom_new()

	var _gestureZoom *GestureZoomClass // out

	_gestureZoom = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*GestureZoomClass)

	return _gestureZoom
}

// ScaleDelta gets the scale delta.
//
// If @gesture is active, this function returns the zooming difference since the
// gesture was recognized (hence the starting point is considered 1:1). If
// @gesture is not active, 1 is returned.
func (gesture *GestureZoomClass) ScaleDelta() float64 {
	var _arg0 *C.GtkGestureZoom // out
	var _cret C.double          // in

	_arg0 = (*C.GtkGestureZoom)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_zoom_get_scale_delta(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
