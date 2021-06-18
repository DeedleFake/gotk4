// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_app_launch_context_get_type()), F: marshalX11AppLaunchContext},
		{T: externglib.Type(C.gdk_x11_cursor_get_type()), F: marshalX11Cursor},
		{T: externglib.Type(C.gdk_x11_device_core_get_type()), F: marshalX11DeviceCore},
		{T: externglib.Type(C.gdk_x11_device_manager_core_get_type()), F: marshalX11DeviceManagerCore},
		{T: externglib.Type(C.gdk_x11_device_manager_xi2_get_type()), F: marshalX11DeviceManagerXI2},
		{T: externglib.Type(C.gdk_x11_device_xi2_get_type()), F: marshalX11DeviceXI2},
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Display},
		{T: externglib.Type(C.gdk_x11_display_manager_get_type()), F: marshalX11DisplayManager},
		{T: externglib.Type(C.gdk_x11_drag_context_get_type()), F: marshalX11DragContext},
		{T: externglib.Type(C.gdk_x11_gl_context_get_type()), F: marshalX11GLContext},
		{T: externglib.Type(C.gdk_x11_keymap_get_type()), F: marshalX11Keymap},
		{T: externglib.Type(C.gdk_x11_monitor_get_type()), F: marshalX11Monitor},
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screen},
		{T: externglib.Type(C.gdk_x11_visual_get_type()), F: marshalX11Visual},
		{T: externglib.Type(C.gdk_x11_window_get_type()), F: marshalX11Window},
	})
}

// X11DeviceGetID returns the device ID as seen by XInput2.
//
// > If gdk_disable_multidevice() has been called, this function > will
// respectively return 2/3 for the core pointer and keyboard, > (matching the
// IDs for the Virtual Core Pointer and Keyboard in > XInput 2), but calling
// this function on any slave devices (i.e. > those managed via XInput 1.x),
// will return 0.
func X11DeviceGetID(device X11DeviceCore) int {
	var _arg1 *C.GdkDevice // out
	var _cret C.gint       // in

	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_x11_device_get_id(_arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager X11DeviceManagerCore, deviceId int) X11DeviceCore {
	var _arg1 *C.GdkDeviceManager // out
	var _arg2 C.gint              // out
	var _cret *C.GdkDevice        // in

	_arg1 = (*C.GdkDeviceManager)(unsafe.Pointer(deviceManager.Native()))
	_arg2 = (C.gint)(deviceId)

	_cret = C.gdk_x11_device_manager_lookup(_arg1, _arg2)

	var _x11DeviceCore X11DeviceCore // out

	_x11DeviceCore = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(X11DeviceCore)

	return _x11DeviceCore
}

// X11FreeCompoundText frees the data returned from
// gdk_x11_display_string_to_compound_text().
func X11FreeCompoundText(ctext *byte) {
	var _arg1 *C.guchar // out

	_arg1 = (*C.guchar)(unsafe.Pointer(ctext))

	C.gdk_x11_free_compound_text(_arg1)
}

// X11FreeTextList frees the array of strings created by
// gdk_x11_display_text_property_to_text_list().
func X11FreeTextList(list *string) {
	var _arg1 **C.gchar // out

	_arg1 = (**C.gchar)(C.CString(list))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_free_text_list(_arg1)
}

// X11GetDefaultScreen gets the default GTK+ screen number.
func X11GetDefaultScreen() int {
	var _cret C.gint // in

	_cret = C.gdk_x11_get_default_screen()

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// X11GetParentRelativePattern: used with gdk_window_set_background_pattern() to
// inherit background from parent window. Useful for imitating transparency when
// compositing is not available. Otherwise behaves like a transparent pattern.
func X11GetParentRelativePattern() *cairo.Pattern {
	var _cret *C.cairo_pattern_t // in

	_cret = C.gdk_x11_get_parent_relative_pattern()

	var _pattern *cairo.Pattern // out

	_pattern = cairo.WrapPattern(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_pattern, func(v *cairo.Pattern) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _pattern
}

// X11GetServerTime: routine to get the current X server time stamp.
func X11GetServerTime(window X11Window) uint32 {
	var _arg1 *C.GdkWindow // out
	var _cret C.guint32    // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gdk_x11_get_server_time(_arg1)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// X11GrabServer: call gdk_x11_display_grab() on the default display. To ungrab
// the server again, use gdk_x11_ungrab_server().
//
// gdk_x11_grab_server()/gdk_x11_ungrab_server() calls can be nested.
func X11GrabServer() {
	C.gdk_x11_grab_server()
}

// X11RegisterStandardEventType registers interest in receiving extension events
// with type codes between @event_base and `event_base + n_events - 1`. The
// registered events must have the window field in the same place as core X
// events (this is not the case for e.g. XKB extension events).
//
// If an event type is registered, events of this type will go through global
// and window-specific filters (see gdk_window_add_filter()). Unregistered
// events will only go through global filters. GDK may register the events of
// some X extensions on its own.
//
// This function should only be needed in unusual circumstances, e.g. when
// filtering XInput extension events on the root window.
func X11RegisterStandardEventType(display X11Display, eventBase int, nEvents int) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (C.gint)(eventBase)
	_arg3 = (C.gint)(nEvents)

	C.gdk_x11_register_standard_event_type(_arg1, _arg2, _arg3)
}

// X11SetSmClientID sets the `SM_CLIENT_ID` property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientId string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(smClientId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_set_sm_client_id(_arg1)
}

// X11UngrabServer: ungrab the default display after it has been grabbed with
// gdk_x11_grab_server().
func X11UngrabServer() {
	C.gdk_x11_ungrab_server()
}

type X11AppLaunchContext interface {
	gdk.AppLaunchContext
}

// x11AppLaunchContext implements the X11AppLaunchContext class.
type x11AppLaunchContext struct {
	gdk.AppLaunchContext
}

var _ X11AppLaunchContext = (*x11AppLaunchContext)(nil)

// WrapX11AppLaunchContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11AppLaunchContext(obj *externglib.Object) X11AppLaunchContext {
	return x11AppLaunchContext{
		AppLaunchContext: gdk.WrapAppLaunchContext(obj),
	}
}

func marshalX11AppLaunchContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11AppLaunchContext(obj), nil
}

type X11Cursor interface {
	gdk.Cursor
}

// x11Cursor implements the X11Cursor class.
type x11Cursor struct {
	gdk.Cursor
}

var _ X11Cursor = (*x11Cursor)(nil)

// WrapX11Cursor wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Cursor(obj *externglib.Object) X11Cursor {
	return x11Cursor{
		Cursor: gdk.WrapCursor(obj),
	}
}

func marshalX11Cursor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Cursor(obj), nil
}

type X11DeviceCore interface {
	gdk.Device
}

// x11DeviceCore implements the X11DeviceCore class.
type x11DeviceCore struct {
	gdk.Device
}

var _ X11DeviceCore = (*x11DeviceCore)(nil)

// WrapX11DeviceCore wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceCore(obj *externglib.Object) X11DeviceCore {
	return x11DeviceCore{
		Device: gdk.WrapDevice(obj),
	}
}

func marshalX11DeviceCore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceCore(obj), nil
}

type X11DeviceManagerCore interface {
	gdk.DeviceManager
}

// x11DeviceManagerCore implements the X11DeviceManagerCore class.
type x11DeviceManagerCore struct {
	gdk.DeviceManager
}

var _ X11DeviceManagerCore = (*x11DeviceManagerCore)(nil)

// WrapX11DeviceManagerCore wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceManagerCore(obj *externglib.Object) X11DeviceManagerCore {
	return x11DeviceManagerCore{
		DeviceManager: gdk.WrapDeviceManager(obj),
	}
}

func marshalX11DeviceManagerCore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceManagerCore(obj), nil
}

type X11DeviceManagerXI2 interface {
	X11DeviceManagerCore
}

// x11DeviceManagerXI2 implements the X11DeviceManagerXI2 class.
type x11DeviceManagerXI2 struct {
	X11DeviceManagerCore
}

var _ X11DeviceManagerXI2 = (*x11DeviceManagerXI2)(nil)

// WrapX11DeviceManagerXI2 wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceManagerXI2(obj *externglib.Object) X11DeviceManagerXI2 {
	return x11DeviceManagerXI2{
		X11DeviceManagerCore: WrapX11DeviceManagerCore(obj),
	}
}

func marshalX11DeviceManagerXI2(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceManagerXI2(obj), nil
}

type X11DeviceXI2 interface {
	gdk.Device
}

// x11DeviceXI2 implements the X11DeviceXI2 class.
type x11DeviceXI2 struct {
	gdk.Device
}

var _ X11DeviceXI2 = (*x11DeviceXI2)(nil)

// WrapX11DeviceXI2 wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceXI2(obj *externglib.Object) X11DeviceXI2 {
	return x11DeviceXI2{
		Device: gdk.WrapDevice(obj),
	}
}

func marshalX11DeviceXI2(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceXI2(obj), nil
}

type X11Display interface {
	gdk.Display

	// ErrorTrapPop pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
	// always block until the error is known to have occurred or not occurred,
	// so the error code can be returned.
	//
	// If you don’t need to use the return value,
	// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
	//
	// See gdk_error_trap_pop() for the all-displays-at-once equivalent.
	ErrorTrapPop() int
	// ErrorTrapPopIgnored pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Does not block to see if an error
	// occurred; merely records the range of requests to ignore errors for, and
	// ignores those errors if they arrive asynchronously.
	//
	// See gdk_error_trap_pop_ignored() for the all-displays-at-once equivalent.
	ErrorTrapPopIgnored()
	// ErrorTrapPush begins a range of X requests on @display for which X error
	// events will be ignored. Unignored errors (when no trap is pushed) will
	// abort the application. Use gdk_x11_display_error_trap_pop() or
	// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
	// function.
	//
	// See also gdk_error_trap_push() to push a trap on all displays.
	ErrorTrapPush()
	// StartupNotificationID gets the startup notification ID for a display.
	StartupNotificationID() string
	// UserTime returns the timestamp of the last user interaction on @display.
	// The timestamp is taken from events caused by user interaction such as key
	// presses or pointer movements. See gdk_x11_window_set_user_time().
	UserTime() uint32
	// Grab: call XGrabServer() on @display. To ungrab the display again, use
	// gdk_x11_display_ungrab().
	//
	// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
	Grab()
	// SetCursorTheme sets the cursor theme from which the images for cursor
	// should be taken.
	//
	// If the windowing system supports it, existing cursors created with
	// gdk_cursor_new(), gdk_cursor_new_for_display() and
	// gdk_cursor_new_from_name() are updated to reflect the theme change.
	// Custom cursors constructed with gdk_cursor_new_from_pixbuf() will have to
	// be handled by the application (GTK+ applications can learn about cursor
	// theme changes by listening for change notification for the corresponding
	// Setting).
	SetCursorTheme(theme string, size int)
	// SetStartupNotificationID sets the startup notification ID for a display.
	//
	// This is usually taken from the value of the DESKTOP_STARTUP_ID
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// If the ID contains the string "_TIME" then the portion following that
	// string is taken to be the X11 timestamp of the event that triggered the
	// application to be launched and the GDK current event time is set
	// accordingly.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// gdk_notify_startup_complete()).
	SetStartupNotificationID(startupId string)
	// SetWindowScale forces a specific window scale for all windows on this
	// display, instead of using the default or user configured scale. This is
	// can be used to disable scaling support by setting @scale to 1, or to
	// programmatically set the window scale.
	//
	// Once the scale is set by this call it will not change in response to
	// later user configuration changes.
	SetWindowScale(scale int)
	// StringToCompoundText: convert a string from the encoding of the current
	// locale into a form suitable for storing in a window property.
	StringToCompoundText(str string) (encoding gdk.Atom, format int, ctext []byte, gint int)
	// TextPropertyToTextList: convert a text string from the encoding as it is
	// stored in a property into an array of strings in the encoding of the
	// current locale. (The elements of the array represent the nul-separated
	// elements of the original text string.)
	TextPropertyToTextList(encoding *gdk.Atom, format int, text *byte, length int, list **string) int
	// Ungrab: ungrab @display after it has been grabbed with
	// gdk_x11_display_grab().
	Ungrab()
	// UTF8ToCompoundText converts from UTF-8 to compound text.
	UTF8ToCompoundText(str string) (gdk.Atom, int, []byte, bool)
}

// x11Display implements the X11Display class.
type x11Display struct {
	gdk.Display
}

var _ X11Display = (*x11Display)(nil)

// WrapX11Display wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Display(obj *externglib.Object) X11Display {
	return x11Display{
		Display: gdk.WrapDisplay(obj),
	}
}

func marshalX11Display(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Display(obj), nil
}

func (d x11Display) ErrorTrapPop() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gint        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_error_trap_pop(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

func (d x11Display) ErrorTrapPopIgnored() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(_arg0)
}

func (d x11Display) ErrorTrapPush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_push(_arg0)
}

func (d x11Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d x11Display) UserTime() uint32 {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_user_time(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (d x11Display) Grab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_grab(_arg0)
}

func (d x11Display) SetCursorTheme(theme string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(theme))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.gint)(size)

	C.gdk_x11_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

func (d x11Display) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_display_set_startup_notification_id(_arg0, _arg1)
}

func (d x11Display) SetWindowScale(scale int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.gint)(scale)

	C.gdk_x11_display_set_window_scale(_arg0, _arg1)
}

func (d x11Display) StringToCompoundText(str string) (encoding gdk.Atom, format int, ctext []byte, gint int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out
	var _encoding gdk.Atom
	var _arg3 C.gint // in
	var _arg4 *C.guchar
	var _arg5 C.gint // in
	var _cret C.gint // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_string_to_compound_text(_arg0, _arg1, (*C.GdkAtom)(unsafe.Pointer(&_encoding)), &_arg3, &_arg4, &_arg5)

	var _format int // out
	var _ctext []byte
	var _gint int // out

	_format = (int)(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_gint = (int)(_cret)

	return _encoding, _format, _ctext, _gint
}

func (d x11Display) TextPropertyToTextList(encoding *gdk.Atom, format int, text *byte, length int, list **string) int {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.GdkAtom     // out
	var _arg2 C.gint        // out
	var _arg3 *C.guchar     // out
	var _arg4 C.gint        // out
	var _arg5 ***C.gchar    // out
	var _cret C.gint        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.GdkAtom)(unsafe.Pointer(encoding.Native()))
	_arg2 = (C.gint)(format)
	_arg3 = (*C.guchar)(unsafe.Pointer(text))
	_arg4 = (C.gint)(length)
	_arg5 = (***C.gchar)(C.CString(list))
	defer C.free(unsafe.Pointer(_arg5))

	_cret = C.gdk_x11_display_text_property_to_text_list(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

func (d x11Display) Ungrab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_ungrab(_arg0)
}

func (d x11Display) UTF8ToCompoundText(str string) (gdk.Atom, int, []byte, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out
	var _encoding gdk.Atom
	var _arg3 C.gint // in
	var _arg4 *C.guchar
	var _arg5 C.gint     // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_utf8_to_compound_text(_arg0, _arg1, (*C.GdkAtom)(unsafe.Pointer(&_encoding)), &_arg3, &_arg4, &_arg5)

	var _format int // out
	var _ctext []byte
	var _ok bool // out

	_format = (int)(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _encoding, _format, _ctext, _ok
}

type X11DisplayManager interface {
	gdk.DisplayManager
}

// x11DisplayManager implements the X11DisplayManager class.
type x11DisplayManager struct {
	gdk.DisplayManager
}

var _ X11DisplayManager = (*x11DisplayManager)(nil)

// WrapX11DisplayManager wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DisplayManager(obj *externglib.Object) X11DisplayManager {
	return x11DisplayManager{
		DisplayManager: gdk.WrapDisplayManager(obj),
	}
}

func marshalX11DisplayManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DisplayManager(obj), nil
}

type X11DragContext interface {
	gdk.DragContext
}

// x11DragContext implements the X11DragContext class.
type x11DragContext struct {
	gdk.DragContext
}

var _ X11DragContext = (*x11DragContext)(nil)

// WrapX11DragContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DragContext(obj *externglib.Object) X11DragContext {
	return x11DragContext{
		DragContext: gdk.WrapDragContext(obj),
	}
}

func marshalX11DragContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DragContext(obj), nil
}

type X11GLContext interface {
	gdk.GLContext
}

// x11GLContext implements the X11GLContext class.
type x11GLContext struct {
	gdk.GLContext
}

var _ X11GLContext = (*x11GLContext)(nil)

// WrapX11GLContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11GLContext(obj *externglib.Object) X11GLContext {
	return x11GLContext{
		GLContext: gdk.WrapGLContext(obj),
	}
}

func marshalX11GLContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11GLContext(obj), nil
}

type X11Keymap interface {
	gdk.Keymap

	// GroupForState extracts the group from the state field sent in an X Key
	// event. This is only needed for code processing raw X events, since
	// EventKey directly includes an is_modifier field.
	GroupForState(state uint) int
	// KeyIsModifier determines whether a particular key code represents a key
	// that is a modifier. That is, it’s a key that normally just affects the
	// keyboard state and the behavior of other keys rather than producing a
	// direct effect itself. This is only needed for code processing raw X
	// events, since EventKey directly includes an is_modifier field.
	KeyIsModifier(keycode uint) bool
}

// x11Keymap implements the X11Keymap class.
type x11Keymap struct {
	gdk.Keymap
}

var _ X11Keymap = (*x11Keymap)(nil)

// WrapX11Keymap wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Keymap(obj *externglib.Object) X11Keymap {
	return x11Keymap{
		Keymap: gdk.WrapKeymap(obj),
	}
}

func marshalX11Keymap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Keymap(obj), nil
}

func (k x11Keymap) GroupForState(state uint) int {
	var _arg0 *C.GdkKeymap // out
	var _arg1 C.guint      // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(k.Native()))
	_arg1 = (C.guint)(state)

	_cret = C.gdk_x11_keymap_get_group_for_state(_arg0, _arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

func (k x11Keymap) KeyIsModifier(keycode uint) bool {
	var _arg0 *C.GdkKeymap // out
	var _arg1 C.guint      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(k.Native()))
	_arg1 = (C.guint)(keycode)

	_cret = C.gdk_x11_keymap_key_is_modifier(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

type X11Monitor interface {
	gdk.Monitor
}

// x11Monitor implements the X11Monitor class.
type x11Monitor struct {
	gdk.Monitor
}

var _ X11Monitor = (*x11Monitor)(nil)

// WrapX11Monitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Monitor(obj *externglib.Object) X11Monitor {
	return x11Monitor{
		Monitor: gdk.WrapMonitor(obj),
	}
}

func marshalX11Monitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Monitor(obj), nil
}

type X11Screen interface {
	gdk.Screen

	// CurrentDesktop returns the current workspace for @screen when running
	// under a window manager that supports multiple workspaces, as described in
	// the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	CurrentDesktop() uint32
	// NumberOfDesktops returns the number of workspaces for @screen when
	// running under a window manager that supports multiple workspaces, as
	// described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	NumberOfDesktops() uint32
	// ScreenNumber returns the index of a Screen.
	ScreenNumber() int
	// WindowManagerName returns the name of the window manager for @screen.
	WindowManagerName() string
	// SupportsNetWmHint: this function is specific to the X11 backend of GDK,
	// and indicates whether the window manager supports a certain hint from the
	// Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	//
	// When using this function, keep in mind that the window manager can change
	// over time; so you shouldn’t use this function in a way that impacts
	// persistent application state. A common bug is that your application can
	// start up before the window manager does when the user logs in, and before
	// the window manager starts gdk_x11_screen_supports_net_wm_hint() will
	// return false for every property. You can monitor the
	// window_manager_changed signal on Screen to detect a window manager
	// change.
	SupportsNetWmHint(property *gdk.Atom) bool
}

// x11Screen implements the X11Screen class.
type x11Screen struct {
	gdk.Screen
}

var _ X11Screen = (*x11Screen)(nil)

// WrapX11Screen wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Screen(obj *externglib.Object) X11Screen {
	return x11Screen{
		Screen: gdk.WrapScreen(obj),
	}
}

func marshalX11Screen(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Screen(obj), nil
}

func (s x11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkScreen // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (s x11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkScreen // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (s x11Screen) ScreenNumber() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.int        // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

func (s x11Screen) WindowManagerName() string {
	var _arg0 *C.GdkScreen // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s x11Screen) SupportsNetWmHint(property *gdk.Atom) bool {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.GdkAtom    // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GdkAtom)(unsafe.Pointer(property.Native()))

	_cret = C.gdk_x11_screen_supports_net_wm_hint(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

type X11Visual interface {
	gdk.Visual
}

// x11Visual implements the X11Visual class.
type x11Visual struct {
	gdk.Visual
}

var _ X11Visual = (*x11Visual)(nil)

// WrapX11Visual wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Visual(obj *externglib.Object) X11Visual {
	return x11Visual{
		Visual: gdk.WrapVisual(obj),
	}
}

func marshalX11Visual(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Visual(obj), nil
}

type X11Window interface {
	gdk.Window

	// Desktop gets the number of the workspace @window is on.
	Desktop() uint32
	// MoveToCurrentDesktop moves the window to the correct workspace when
	// running under a window manager that supports multiple workspaces, as
	// described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification. Will not do
	// anything if the window is already on all workspaces.
	MoveToCurrentDesktop()
	// MoveToDesktop moves the window to the given workspace when running unde a
	// window manager that supports multiple workspaces, as described in the
	// Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	MoveToDesktop(desktop uint32)
	// SetFrameExtents: this is the same as gdk_window_set_shadow_width() but it
	// only works on GdkX11Window.
	SetFrameExtents(left int, right int, top int, bottom int)
	// SetFrameSyncEnabled: this function can be used to disable frame
	// synchronization for a window. Normally frame synchronziation will be
	// enabled or disabled based on whether the system has a compositor that
	// supports frame synchronization, but if the window is not directly managed
	// by the window manager, then frame synchronziation may need to be
	// disabled. This is the case for a window embedded via the XEMBED protocol.
	SetFrameSyncEnabled(frameSyncEnabled bool)
	// SetHideTitlebarWhenMaximized: set a hint for the window manager,
	// requesting that the titlebar should be hidden when the window is
	// maximized.
	//
	// Note that this property is automatically updated by GTK+, so this
	// function should only be used by applications which do not use GTK+ to
	// create toplevel windows.
	SetHideTitlebarWhenMaximized(hideTitlebarWhenMaximized bool)
	// SetThemeVariant: GTK+ applications can request a dark theme variant. In
	// order to make other applications - namely window managers using GTK+ for
	// themeing - aware of this choice, GTK+ uses this function to export the
	// requested theme variant as _GTK_THEME_VARIANT property on toplevel
	// windows.
	//
	// Note that this property is automatically updated by GTK+, so this
	// function should only be used by applications which do not use GTK+ to
	// create toplevel windows.
	SetThemeVariant(variant string)
	// SetUserTime: the application can use this call to update the
	// _NET_WM_USER_TIME property on a toplevel window. This property stores an
	// Xserver time which represents the time of the last user input event
	// received for this window. This property may be used by the window manager
	// to alter the focus, stacking, and/or placement behavior of windows when
	// they are mapped depending on whether the new window was created by a user
	// action or is a "pop-up" window activated by a timer or some other event.
	//
	// Note that this property is automatically updated by GDK, so this function
	// should only be used by applications which handle input events bypassing
	// GDK.
	SetUserTime(timestamp uint32)
	// SetUTF8Property: this function modifies or removes an arbitrary X11
	// window property of type UTF8_STRING. If the given @window is not a
	// toplevel window, it is ignored.
	SetUTF8Property(name string, value string)
}

// x11Window implements the X11Window class.
type x11Window struct {
	gdk.Window
}

var _ X11Window = (*x11Window)(nil)

// WrapX11Window wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Window(obj *externglib.Object) X11Window {
	return x11Window{
		Window: gdk.WrapWindow(obj),
	}
}

func marshalX11Window(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Window(obj), nil
}

func (w x11Window) Desktop() uint32 {
	var _arg0 *C.GdkWindow // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gdk_x11_window_get_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (w x11Window) MoveToCurrentDesktop() {
	var _arg0 *C.GdkWindow // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))

	C.gdk_x11_window_move_to_current_desktop(_arg0)
}

func (w x11Window) MoveToDesktop(desktop uint32) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.guint32    // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (C.guint32)(desktop)

	C.gdk_x11_window_move_to_desktop(_arg0, _arg1)
}

func (w x11Window) SetFrameExtents(left int, right int, top int, bottom int) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _arg4 C.int        // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (C.int)(left)
	_arg2 = (C.int)(right)
	_arg3 = (C.int)(top)
	_arg4 = (C.int)(bottom)

	C.gdk_x11_window_set_frame_extents(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (w x11Window) SetFrameSyncEnabled(frameSyncEnabled bool) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))
	if frameSyncEnabled {
		_arg1 = C.TRUE
	}

	C.gdk_x11_window_set_frame_sync_enabled(_arg0, _arg1)
}

func (w x11Window) SetHideTitlebarWhenMaximized(hideTitlebarWhenMaximized bool) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))
	if hideTitlebarWhenMaximized {
		_arg1 = C.TRUE
	}

	C.gdk_x11_window_set_hide_titlebar_when_maximized(_arg0, _arg1)
}

func (w x11Window) SetThemeVariant(variant string) {
	var _arg0 *C.GdkWindow // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.char)(C.CString(variant))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_window_set_theme_variant(_arg0, _arg1)
}

func (w x11Window) SetUserTime(timestamp uint32) {
	var _arg0 *C.GdkWindow // out
	var _arg1 C.guint32    // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (C.guint32)(timestamp)

	C.gdk_x11_window_set_user_time(_arg0, _arg1)
}

func (w x11Window) SetUTF8Property(name string, value string) {
	var _arg0 *C.GdkWindow // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	C.gdk_x11_window_set_utf8_property(_arg0, _arg1, _arg2)
}
