// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_get_type()), F: marshalDisplay},
	})
}

// Display: `GdkDisplay` objects are the GDK representation of a workstation.
//
// Their purpose are two-fold:
//
// - To manage and provide information about input devices (pointers, keyboards,
// etc) - To manage and provide information about output devices (monitors,
// projectors, etc)
//
// Most of the input device handling has been factored out into separate
// [class@Gdk.Seat] objects. Every display has a one or more seats, which can be
// accessed with [method@Gdk.Display.get_default_seat] and
// [method@Gdk.Display.list_seats].
//
// Output devices are represented by [class@Gdk.Monitor] objects, which can be
// accessed with [method@Gdk.Display.get_monitor_at_surface] and similar APIs.
type Display interface {
	gextras.Objector

	// Beep emits a short beep on @display
	Beep()
	// Close closes the connection to the windowing system for the given
	// display.
	//
	// This cleans up associated resources.
	Close()
	// DeviceIsGrabbed returns true if there is an ongoing grab on @device for
	// @display.
	DeviceIsGrabbed(device Device) bool
	// Flush flushes any requests queued for the windowing system.
	//
	// This happens automatically when the main loop blocks waiting for new
	// events, but if your application is drawing without returning control to
	// the main loop, you may need to call this function explicitly. A common
	// case where this function needs to be called is when an application is
	// executing drawing commands from a thread other than the thread where the
	// main loop is running.
	//
	// This is most useful for X11. On windowing systems where requests are
	// handled synchronously, this function will do nothing.
	Flush()
	// AppLaunchContext returns a `GdkAppLaunchContext` suitable for launching
	// applications on the given display.
	AppLaunchContext() *AppLaunchContextClass
	// Clipboard gets the clipboard used for copy/paste operations.
	Clipboard() *ClipboardClass
	// DefaultSeat returns the default `GdkSeat` for this display.
	//
	// Note that a display may not have a seat. In this case, this function will
	// return nil.
	DefaultSeat() *SeatClass
	// MonitorAtSurface gets the monitor in which the largest area of @surface
	// resides.
	//
	// Returns a monitor close to @surface if it is outside of all monitors.
	MonitorAtSurface(surface Surface) *MonitorClass
	// Name gets the name of the display.
	Name() string
	// PrimaryClipboard gets the clipboard used for the primary selection.
	//
	// On backends where the primary clipboard is not supported natively, GDK
	// emulates this clipboard locally.
	PrimaryClipboard() *ClipboardClass
	// Setting retrieves a desktop-wide setting such as double-click time for
	// the @display.
	Setting(name string, value *externglib.Value) bool
	// StartupNotificationID gets the startup notification ID for a Wayland
	// display, or nil if no ID has been defined.
	StartupNotificationID() string
	// IsClosed finds out if the display has been closed.
	IsClosed() bool
	// IsComposited returns whether surfaces can reasonably be expected to have
	// their alpha channel drawn correctly on the screen.
	//
	// Check [method@Gdk.Display.is_rgba] for whether the display supports an
	// alpha channel.
	//
	// On X11 this function returns whether a compositing manager is compositing
	// on @display.
	//
	// On modern displays, this value is always true.
	IsComposited() bool
	// IsRGBA returns whether surfaces on this @display are created with an
	// alpha channel.
	//
	// Even if a true is returned, it is possible that the surface’s alpha
	// channel won’t be honored when displaying the surface on the screen: in
	// particular, for X an appropriate windowing manager and compositing
	// manager must be running to provide appropriate display. Use
	// [method@Gdk.Display.is_composited] to check if that is the case.
	//
	// On modern displays, this value is always true.
	IsRGBA() bool
	// MapKeycode returns the keyvals bound to @keycode.
	//
	// The Nth `GdkKeymapKey` in @keys is bound to the Nth keyval in @keyvals.
	//
	// When a keycode is pressed by the user, the keyval from this list of
	// entries is selected by considering the effective keyboard group and
	// level.
	//
	// Free the returned arrays with g_free().
	MapKeycode(keycode uint) ([]KeymapKey, []uint, bool)
	// MapKeyval obtains a list of keycode/group/level combinations that will
	// generate @keyval.
	//
	// Groups and levels are two kinds of keyboard mode; in general, the level
	// determines whether the top or bottom symbol on a key is used, and the
	// group determines whether the left or right symbol is used.
	//
	// On US keyboards, the shift key changes the keyboard level, and there are
	// no groups. A group switch key might convert a keyboard between Hebrew to
	// English modes, for example.
	//
	// `GdkEventKey` contains a group field that indicates the active keyboard
	// group. The level is computed from the modifier mask.
	//
	// The returned array should be freed with g_free().
	MapKeyval(keyval uint) ([]KeymapKey, bool)
	// NotifyStartupComplete indicates to the GUI environment that the
	// application has finished loading, using a given identifier.
	//
	// GTK will call this function automatically for [class@Gtk.Window] with
	// custom startup-notification identifier unless
	// [method@Gtk.Window.set_auto_startup_notification] is called to disable
	// that feature.
	NotifyStartupComplete(startupId string)
	// PutEvent appends the given event onto the front of the event queue for
	// @display.
	//
	// This function is only useful in very special situations and should not be
	// used by applications.
	PutEvent(event Event)
	// SupportsInputShapes returns true if the display supports input shapes.
	//
	// This means that [method@Gdk.Surface.set_input_region] can be used to
	// modify the input shape of surfaces on @display.
	//
	// On modern displays, this value is always true.
	SupportsInputShapes() bool
	// Sync flushes any requests queued for the windowing system and waits until
	// all requests have been handled.
	//
	// This is often used for making sure that the display is synchronized with
	// the current state of the program. Calling [method@Gdk.Display.sync]
	// before [method@GdkX11.Display.error_trap_pop] makes sure that any errors
	// generated from earlier requests are handled before the error trap is
	// removed.
	//
	// This is most useful for X11. On windowing systems where requests are
	// handled synchronously, this function will do nothing.
	Sync()
}

// DisplayClass implements the Display interface.
type DisplayClass struct {
	*externglib.Object
}

var _ Display = (*DisplayClass)(nil)

func wrapDisplay(obj *externglib.Object) Display {
	return &DisplayClass{
		Object: obj,
	}
}

func marshalDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDisplay(obj), nil
}

// Beep emits a short beep on @display
func (display *DisplayClass) Beep() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_beep(_arg0)
}

// Close closes the connection to the windowing system for the given display.
//
// This cleans up associated resources.
func (display *DisplayClass) Close() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_close(_arg0)
}

// DeviceIsGrabbed returns true if there is an ongoing grab on @device for
// @display.
func (display *DisplayClass) DeviceIsGrabbed(device Device) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkDevice  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_display_device_is_grabbed(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Flush flushes any requests queued for the windowing system.
//
// This happens automatically when the main loop blocks waiting for new events,
// but if your application is drawing without returning control to the main
// loop, you may need to call this function explicitly. A common case where this
// function needs to be called is when an application is executing drawing
// commands from a thread other than the thread where the main loop is running.
//
// This is most useful for X11. On windowing systems where requests are handled
// synchronously, this function will do nothing.
func (display *DisplayClass) Flush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_flush(_arg0)
}

// AppLaunchContext returns a `GdkAppLaunchContext` suitable for launching
// applications on the given display.
func (display *DisplayClass) AppLaunchContext() *AppLaunchContextClass {
	var _arg0 *C.GdkDisplay          // out
	var _cret *C.GdkAppLaunchContext // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_app_launch_context(_arg0)

	var _appLaunchContext *AppLaunchContextClass // out

	_appLaunchContext = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*AppLaunchContextClass)

	return _appLaunchContext
}

// Clipboard gets the clipboard used for copy/paste operations.
func (display *DisplayClass) Clipboard() *ClipboardClass {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_clipboard(_arg0)

	var _clipboard *ClipboardClass // out

	_clipboard = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ClipboardClass)

	return _clipboard
}

// DefaultSeat returns the default `GdkSeat` for this display.
//
// Note that a display may not have a seat. In this case, this function will
// return nil.
func (display *DisplayClass) DefaultSeat() *SeatClass {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSeat    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_default_seat(_arg0)

	var _seat *SeatClass // out

	_seat = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*SeatClass)

	return _seat
}

// MonitorAtSurface gets the monitor in which the largest area of @surface
// resides.
//
// Returns a monitor close to @surface if it is outside of all monitors.
func (display *DisplayClass) MonitorAtSurface(surface Surface) *MonitorClass {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkSurface // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_display_get_monitor_at_surface(_arg0, _arg1)

	var _monitor *MonitorClass // out

	_monitor = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MonitorClass)

	return _monitor
}

// Name gets the name of the display.
func (display *DisplayClass) Name() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PrimaryClipboard gets the clipboard used for the primary selection.
//
// On backends where the primary clipboard is not supported natively, GDK
// emulates this clipboard locally.
func (display *DisplayClass) PrimaryClipboard() *ClipboardClass {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_primary_clipboard(_arg0)

	var _clipboard *ClipboardClass // out

	_clipboard = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ClipboardClass)

	return _clipboard
}

// Setting retrieves a desktop-wide setting such as double-click time for the
// @display.
func (display *DisplayClass) Setting(name string, value *externglib.Value) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.GValue     // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gdk_display_get_setting(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartupNotificationID gets the startup notification ID for a Wayland display,
// or nil if no ID has been defined.
func (display *DisplayClass) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IsClosed finds out if the display has been closed.
func (display *DisplayClass) IsClosed() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_is_closed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsComposited returns whether surfaces can reasonably be expected to have
// their alpha channel drawn correctly on the screen.
//
// Check [method@Gdk.Display.is_rgba] for whether the display supports an alpha
// channel.
//
// On X11 this function returns whether a compositing manager is compositing on
// @display.
//
// On modern displays, this value is always true.
func (display *DisplayClass) IsComposited() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_is_composited(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRGBA returns whether surfaces on this @display are created with an alpha
// channel.
//
// Even if a true is returned, it is possible that the surface’s alpha channel
// won’t be honored when displaying the surface on the screen: in particular,
// for X an appropriate windowing manager and compositing manager must be
// running to provide appropriate display. Use
// [method@Gdk.Display.is_composited] to check if that is the case.
//
// On modern displays, this value is always true.
func (display *DisplayClass) IsRGBA() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_is_rgba(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MapKeycode returns the keyvals bound to @keycode.
//
// The Nth `GdkKeymapKey` in @keys is bound to the Nth keyval in @keyvals.
//
// When a keycode is pressed by the user, the keyval from this list of entries
// is selected by considering the effective keyboard group and level.
//
// Free the returned arrays with g_free().
func (display *DisplayClass) MapKeycode(keycode uint) ([]KeymapKey, []uint, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out
	var _arg2 *C.GdkKeymapKey
	var _arg4 C.int // in
	var _arg3 *C.guint
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint(keycode)

	_cret = C.gdk_display_map_keycode(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _keys []KeymapKey
	var _keyvals []uint
	var _ok bool // out

	_keys = unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg4)
	runtime.SetFinalizer(&_keys, func(v *[]KeymapKey) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_keyvals = unsafe.Slice((*uint)(unsafe.Pointer(_arg3)), _arg4)
	runtime.SetFinalizer(&_keyvals, func(v *[]uint) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _keys, _keyvals, _ok
}

// MapKeyval obtains a list of keycode/group/level combinations that will
// generate @keyval.
//
// Groups and levels are two kinds of keyboard mode; in general, the level
// determines whether the top or bottom symbol on a key is used, and the group
// determines whether the left or right symbol is used.
//
// On US keyboards, the shift key changes the keyboard level, and there are no
// groups. A group switch key might convert a keyboard between Hebrew to English
// modes, for example.
//
// `GdkEventKey` contains a group field that indicates the active keyboard
// group. The level is computed from the modifier mask.
//
// The returned array should be freed with g_free().
func (display *DisplayClass) MapKeyval(keyval uint) ([]KeymapKey, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out
	var _arg2 *C.GdkKeymapKey
	var _arg3 C.int      // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint(keyval)

	_cret = C.gdk_display_map_keyval(_arg0, _arg1, &_arg2, &_arg3)

	var _keys []KeymapKey
	var _ok bool // out

	_keys = unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg3)
	runtime.SetFinalizer(&_keys, func(v *[]KeymapKey) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _keys, _ok
}

// NotifyStartupComplete indicates to the GUI environment that the application
// has finished loading, using a given identifier.
//
// GTK will call this function automatically for [class@Gtk.Window] with custom
// startup-notification identifier unless
// [method@Gtk.Window.set_auto_startup_notification] is called to disable that
// feature.
func (display *DisplayClass) NotifyStartupComplete(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_display_notify_startup_complete(_arg0, _arg1)
}

// PutEvent appends the given event onto the front of the event queue for
// @display.
//
// This function is only useful in very special situations and should not be
// used by applications.
func (display *DisplayClass) PutEvent(event Event) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkEvent   // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))

	C.gdk_display_put_event(_arg0, _arg1)
}

// SupportsInputShapes returns true if the display supports input shapes.
//
// This means that [method@Gdk.Surface.set_input_region] can be used to modify
// the input shape of surfaces on @display.
//
// On modern displays, this value is always true.
func (display *DisplayClass) SupportsInputShapes() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_input_shapes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sync flushes any requests queued for the windowing system and waits until all
// requests have been handled.
//
// This is often used for making sure that the display is synchronized with the
// current state of the program. Calling [method@Gdk.Display.sync] before
// [method@GdkX11.Display.error_trap_pop] makes sure that any errors generated
// from earlier requests are handled before the error trap is removed.
//
// This is most useful for X11. On windowing systems where requests are handled
// synchronously, this function will do nothing.
func (display *DisplayClass) Sync() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_sync(_arg0)
}
