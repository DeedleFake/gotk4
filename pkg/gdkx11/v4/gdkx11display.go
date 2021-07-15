// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Displayyer},
	})
}

// X11SetSmClientID sets the SM_CLIENT_ID property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientId string) {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(smClientId)))

	C.gdk_x11_set_sm_client_id(_arg1)
}

type X11Display struct {
	gdk.Display
}

var _ gextras.Nativer = (*X11Display)(nil)

func wrapX11Display(obj *externglib.Object) *X11Display {
	return &X11Display{
		Display: gdk.Display{
			Object: obj,
		},
	}
}

func marshalX11Displayyer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Display(obj), nil
}

// ErrorTrapPop pops the error trap pushed by gdk_x11_display_error_trap_push().
// Will XSync() if necessary and will always block until the error is known to
// have occurred or not occurred, so the error code can be returned.
//
// If you don’t need to use the return value,
// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
func (display *X11Display) ErrorTrapPop() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_error_trap_pop(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ErrorTrapPopIgnored pops the error trap pushed by
// gdk_x11_display_error_trap_push(). Does not block to see if an error
// occurred; merely records the range of requests to ignore errors for, and
// ignores those errors if they arrive asynchronously.
func (display *X11Display) ErrorTrapPopIgnored() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(_arg0)
}

// ErrorTrapPush begins a range of X requests on display for which X error
// events will be ignored. Unignored errors (when no trap is pushed) will abort
// the application. Use gdk_x11_display_error_trap_pop() or
// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
// function.
func (display *X11Display) ErrorTrapPush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_error_trap_push(_arg0)
}

// DefaultGroup returns the default group leader surface for all toplevel
// surfaces on display. This surface is implicitly created by GDK. See
// gdk_x11_surface_set_group().
func (display *X11Display) DefaultGroup() *gdk.Surface {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_default_group(_arg0)

	var _surface *gdk.Surface // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_surface = &gdk.Surface{
			Object: obj,
		}
	}

	return _surface
}

// GlxVersion retrieves the version of the GLX implementation.
func (display *X11Display) GlxVersion() (major int, minor int, ok bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // in
	var _arg2 C.int         // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_glx_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out
	var _ok bool   // out

	_major = int(_arg1)
	_minor = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _major, _minor, _ok
}

// PrimaryMonitor gets the primary monitor for the display.
//
// The primary monitor is considered the monitor where the “main desktop” lives.
// While normal application surfaces typically allow the window manager to place
// the surfaces, specialized desktop applications such as panels should place
// themselves on the primary monitor.
//
// If no monitor is the designated primary monitor, any monitor (usually the
// first) may be returned.
func (display *X11Display) PrimaryMonitor() *gdk.Monitor {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_primary_monitor(_arg0)

	var _monitor *gdk.Monitor // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_monitor = &gdk.Monitor{
			Object: obj,
		}
	}

	return _monitor
}

// Screen retrieves the X11Screen of the display.
func (display *X11Display) Screen() *X11Screen {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkX11Screen // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_screen(_arg0)

	var _x11Screen *X11Screen // out

	_x11Screen = wrapX11Screen(externglib.Take(unsafe.Pointer(_cret)))

	return _x11Screen
}

// StartupNotificationID gets the startup notification ID for a display.
func (display *X11Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UserTime returns the timestamp of the last user interaction on display. The
// timestamp is taken from events caused by user interaction such as key presses
// or pointer movements. See gdk_x11_surface_set_user_time().
func (display *X11Display) UserTime() uint32 {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_user_time(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Grab: call XGrabServer() on display. To ungrab the display again, use
// gdk_x11_display_ungrab().
//
// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
func (display *X11Display) Grab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_grab(_arg0)
}

// SetCursorTheme sets the cursor theme from which the images for cursor should
// be taken.
//
// If the windowing system supports it, existing cursors created with
// gdk_cursor_new_from_name() are updated to reflect the theme change. Custom
// cursors constructed with gdk_cursor_new_from_texture() will have to be
// handled by the application (GTK applications can learn about cursor theme
// changes by listening for change notification for the corresponding Setting).
func (display *X11Display) SetCursorTheme(theme string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(theme)))
	_arg2 = C.int(size)

	C.gdk_x11_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID environment
// variable, but in some cases (such as the application not being launched using
// exec()) it can come from other sources.
//
// If the ID contains the string "_TIME" then the portion following that string
// is taken to be the X11 timestamp of the event that triggered the application
// to be launched and the GDK current event time is set accordingly.
//
// The startup ID is also what is used to signal that the startup is complete
// (for example, when opening a window or when calling
// gdk_display_notify_startup_complete()).
func (display *X11Display) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))

	C.gdk_x11_display_set_startup_notification_id(_arg0, _arg1)
}

// SetSurfaceScale forces a specific window scale for all windows on this
// display, instead of using the default or user configured scale. This is can
// be used to disable scaling support by setting scale to 1, or to
// programmatically set the window scale.
//
// Once the scale is set by this call it will not change in response to later
// user configuration changes.
func (display *X11Display) SetSurfaceScale(scale int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.int(scale)

	C.gdk_x11_display_set_surface_scale(_arg0, _arg1)
}

// StringToCompoundText: convert a string from the encoding of the current
// locale into a form suitable for storing in a window property.
func (display *X11Display) StringToCompoundText(str string) (encoding string, format int, ctext []byte, gint int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar
	var _arg5 C.int // in
	var _cret C.int // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))

	_cret = C.gdk_x11_display_string_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _gint int // out

	_encoding = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	_format = int(_arg3)
	defer C.free(unsafe.Pointer(_arg4))
	_ctext = make([]byte, _arg5)
	copy(_ctext, unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5))
	_gint = int(_cret)

	return _encoding, _format, _ctext, _gint
}

// Ungrab display after it has been grabbed with gdk_x11_display_grab().
func (display *X11Display) Ungrab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_x11_display_ungrab(_arg0)
}

// UTF8ToCompoundText converts from UTF-8 to compound text.
func (display *X11Display) UTF8ToCompoundText(str string) (string, int, []byte, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar
	var _arg5 C.int      // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))

	_cret = C.gdk_x11_display_utf8_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _ok bool // out

	_encoding = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	_format = int(_arg3)
	defer C.free(unsafe.Pointer(_arg4))
	_ctext = make([]byte, _arg5)
	copy(_ctext, unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5))
	if _cret != 0 {
		_ok = true
	}

	return _encoding, _format, _ctext, _ok
}

// X11DisplayOpen tries to open a new display to the X server given by
// display_name. If opening the display fails, NULL is returned.
func X11DisplayOpen(displayName string) *gdk.Display {
	var _arg1 *C.char       // out
	var _cret *C.GdkDisplay // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(displayName)))

	_cret = C.gdk_x11_display_open(_arg1)

	var _display *gdk.Display // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// X11DisplaySetProgramClass sets the program class.
//
// The X11 backend uses the program class to set the class name part of the
// WM_CLASS property on toplevel windows; see the ICCCM.
func X11DisplaySetProgramClass(display *gdk.Display, programClass string) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.char       // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(programClass)))

	C.gdk_x11_display_set_program_class(_arg1, _arg2)
}
