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
//
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_display_get_type()), F: marshalWaylandDisplay},
	})
}

// WaylandDisplay: the Wayland implementation of `GdkDisplay`.
//
// Beyond the regular [class@Gdk.Display] API, the Wayland implementation
// provides access to Wayland objects such as the `wl_display` with
// [method@GdkWayland.WaylandDisplay.get_wl_display], the `wl_compositor` with
// [method@GdkWayland.WaylandDisplay.get_wl_compositor].
//
// You can find out what Wayland globals are supported by a display with
// [method@GdkWayland.WaylandDisplay.query_registry].
type WaylandDisplay interface {
	gextras.Objector

	// StartupNotificationID gets the startup notification ID for a Wayland
	// display, or nil if no ID has been defined.
	StartupNotificationID() string
	// QueryRegistry returns true if the the interface was found in the display
	// `wl_registry.global` handler.
	QueryRegistry(global string) bool
	// SetCursorTheme sets the cursor theme for the given @display.
	SetCursorTheme(name string, size int)
	// SetStartupNotificationID sets the startup notification ID for a display.
	//
	// This is usually taken from the value of the `DESKTOP_STARTUP_ID`
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// [method@Gdk.Display.notify_startup_complete]).
	SetStartupNotificationID(startupId string)
}

// WaylandDisplayClass implements the WaylandDisplay interface.
type WaylandDisplayClass struct {
	gdk.DisplayClass
}

var _ WaylandDisplay = (*WaylandDisplayClass)(nil)

func wrapWaylandDisplay(obj *externglib.Object) WaylandDisplay {
	return &WaylandDisplayClass{
		DisplayClass: gdk.DisplayClass{
			Object: obj,
		},
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWaylandDisplay(obj), nil
}

// StartupNotificationID gets the startup notification ID for a Wayland display,
// or nil if no ID has been defined.
func (display *WaylandDisplayClass) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_wayland_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// QueryRegistry returns true if the the interface was found in the display
// `wl_registry.global` handler.
func (display *WaylandDisplayClass) QueryRegistry(global string) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(C.CString(global))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_display_query_registry(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCursorTheme sets the cursor theme for the given @display.
func (display *WaylandDisplayClass) SetCursorTheme(name string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(size)

	C.gdk_wayland_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the `DESKTOP_STARTUP_ID` environment
// variable, but in some cases (such as the application not being launched using
// exec()) it can come from other sources.
//
// The startup ID is also what is used to signal that the startup is complete
// (for example, when opening a window or when calling
// [method@Gdk.Display.notify_startup_complete]).
func (display *WaylandDisplayClass) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_display_set_startup_notification_id(_arg0, _arg1)
}
