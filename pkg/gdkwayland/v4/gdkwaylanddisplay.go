// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_display_get_type()), F: marshalWaylandDisplayer},
	})
}

// WaylandDisplay: wayland implementation of GdkDisplay.
//
// Beyond the regular gdk.Display API, the Wayland implementation provides
// access to Wayland objects such as the wl_display with
// gdkwayland.WaylandDisplay.GetWlDisplay(), the wl_compositor with
// gdkwayland.WaylandDisplay.GetWlCompositor().
//
// You can find out what Wayland globals are supported by a display with
// gdkwayland.WaylandDisplay.QueryRegistry().
type WaylandDisplay struct {
	gdk.Display
}

var _ gextras.Nativer = (*WaylandDisplay)(nil)

func wrapWaylandDisplay(obj *externglib.Object) *WaylandDisplay {
	return &WaylandDisplay{
		Display: gdk.Display{
			Object: obj,
		},
	}
}

func marshalWaylandDisplayer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWaylandDisplay(obj), nil
}

// StartupNotificationID gets the startup notification ID for a Wayland display,
// or NULL if no ID has been defined.
func (display *WaylandDisplay) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_wayland_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// QueryRegistry returns TRUE if the the interface was found in the display
// wl_registry.global handler.
func (display *WaylandDisplay) QueryRegistry(global string) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(global)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_display_query_registry(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCursorTheme sets the cursor theme for the given display.
func (display *WaylandDisplay) SetCursorTheme(name string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(size)

	C.gdk_wayland_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID environment
// variable, but in some cases (such as the application not being launched using
// exec()) it can come from other sources.
//
// The startup ID is also what is used to signal that the startup is complete
// (for example, when opening a window or when calling
// gdk.Display.NotifyStartupComplete()).
func (display *WaylandDisplay) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_display_set_startup_notification_id(_arg0, _arg1)
}
