// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_GestureStylus_ConnectDown(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureStylus_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureStylus_ConnectProximity(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureStylus_ConnectUp(gpointer, gdouble, gdouble, guintptr);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStylusser},
	})
}

// GestureStylusOverrider contains methods that are overridable.
type GestureStylusOverrider interface {
}

// GestureStylus: GtkGestureStylus is a GtkGesture specific to stylus input.
//
// The provided signals just relay the basic information of the stylus events.
type GestureStylus struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureStylus)(nil)
)

func classInitGestureStylusser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGestureStylus(obj *externglib.Object) *GestureStylus {
	return &GestureStylus{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureStylusser(p uintptr) (interface{}, error) {
	return wrapGestureStylus(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_GestureStylus_ConnectDown
func _gotk4_gtk4_GestureStylus_ConnectDown(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
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

// ConnectDown: emitted when the stylus touches the device.
func (gesture *GestureStylus) ConnectDown(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "down", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectDown), f)
}

//export _gotk4_gtk4_GestureStylus_ConnectMotion
func _gotk4_gtk4_GestureStylus_ConnectMotion(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
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

// ConnectMotion: emitted when the stylus moves while touching the device.
func (gesture *GestureStylus) ConnectMotion(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "motion", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectMotion), f)
}

//export _gotk4_gtk4_GestureStylus_ConnectProximity
func _gotk4_gtk4_GestureStylus_ConnectProximity(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
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

// ConnectProximity: emitted when the stylus is in proximity of the device.
func (gesture *GestureStylus) ConnectProximity(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "proximity", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectProximity), f)
}

//export _gotk4_gtk4_GestureStylus_ConnectUp
func _gotk4_gtk4_GestureStylus_ConnectUp(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
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

// ConnectUp: emitted when the stylus no longer touches the device.
func (gesture *GestureStylus) ConnectUp(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "up", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectUp), f)
}

// NewGestureStylus creates a new GtkGestureStylus.
//
// The function returns the following values:
//
//    - gestureStylus: newly created stylus gesture.
//
func NewGestureStylus() *GestureStylus {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_stylus_new()

	var _gestureStylus *GestureStylus // out

	_gestureStylus = wrapGestureStylus(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureStylus
}

// Axis returns the current value for the requested axis.
//
// This function must be called from the handler of one of the
// gtk.GestureStylus::down, gtk.GestureStylus::motion, gtk.GestureStylus::up or
// gtk.GestureStylus::proximity signals.
//
// The function takes the following parameters:
//
//    - axis: requested device axis.
//
// The function returns the following values:
//
//    - value: return location for the axis value.
//    - ok: TRUE if there is a current value for the axis.
//
func (gesture *GestureStylus) Axis(axis gdk.AxisUse) (float64, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 C.GdkAxisUse        // out
	var _arg2 C.double            // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.GdkAxisUse(axis)

	_cret = C.gtk_gesture_stylus_get_axis(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(axis)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// Backlog returns the accumulated backlog of tracking information.
//
// By default, GTK will limit rate of input events. On stylus input where
// accuracy of strokes is paramount, this function returns the accumulated
// coordinate/timing state before the emission of the current
// [Gtk.GestureStylus::motion] signal.
//
// This function may only be called within a gtk.GestureStylus::motion signal
// handler, the state given in this signal and obtainable through
// gtk.GestureStylus.GetAxis() express the latest (most up-to-date) state in
// motion history.
//
// The backlog is provided in chronological order.
//
// The function returns the following values:
//
//    - backlog coordinates and times for the backlog events.
//    - ok: TRUE if there is a backlog to unfold in the current state.
//
func (gesture *GestureStylus) Backlog() ([]gdk.TimeCoord, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 *C.GdkTimeCoord     // in
	var _arg2 C.guint             // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_stylus_get_backlog(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

	var _backlog []gdk.TimeCoord // out
	var _ok bool                 // out

	defer C.free(unsafe.Pointer(_arg1))
	{
		src := unsafe.Slice((*C.GdkTimeCoord)(_arg1), _arg2)
		_backlog = make([]gdk.TimeCoord, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_backlog[i] = *(*gdk.TimeCoord)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_backlog[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _backlog, _ok
}

// DeviceTool returns the GdkDeviceTool currently driving input through this
// gesture.
//
// This function must be called from the handler of one of the
// gtk.GestureStylus::down, gtk.GestureStylus::motion, gtk.GestureStylus::up or
// gtk.GestureStylus::proximity signals.
//
// The function returns the following values:
//
//    - deviceTool (optional): current stylus tool.
//
func (gesture *GestureStylus) DeviceTool() *gdk.DeviceTool {
	var _arg0 *C.GtkGestureStylus // out
	var _cret *C.GdkDeviceTool    // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_stylus_get_device_tool(_arg0)
	runtime.KeepAlive(gesture)

	var _deviceTool *gdk.DeviceTool // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_deviceTool = &gdk.DeviceTool{
				Object: obj,
			}
		}
	}

	return _deviceTool
}
