// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_get_type()), F: marshalDisplayer},
	})
}

// Display: GdkDisplay objects are the GDK representation of a workstation.
//
// Their purpose are two-fold:
//
// - To manage and provide information about input devices (pointers, keyboards,
// etc)
//
// - To manage and provide information about output devices (monitors,
// projectors, etc)
//
// Most of the input device handling has been factored out into separate
// gdk.Seat objects. Every display has a one or more seats, which can be
// accessed with gdk.Display.GetDefaultSeat() and gdk.Display.ListSeats().
//
// Output devices are represented by gdk.Monitor objects, which can be accessed
// with gdk.Display.GetMonitorAtSurface() and similar APIs.
type Display struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Display)(nil)

func wrapDisplay(obj *externglib.Object) *Display {
	return &Display{
		Object: obj,
	}
}

func marshalDisplayer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDisplay(obj), nil
}

// Beep emits a short beep on display
func (display *Display) Beep() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_beep(_arg0)
}

// Close closes the connection to the windowing system for the given display.
//
// This cleans up associated resources.
func (display *Display) Close() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_close(_arg0)
}

// DeviceIsGrabbed returns TRUE if there is an ongoing grab on device for
// display.
func (display *Display) DeviceIsGrabbed(device Devicer) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkDevice  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((device).(gextras.Nativer).Native()))

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
func (display *Display) Flush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_flush(_arg0)
}

// AppLaunchContext returns a GdkAppLaunchContext suitable for launching
// applications on the given display.
func (display *Display) AppLaunchContext() *AppLaunchContext {
	var _arg0 *C.GdkDisplay          // out
	var _cret *C.GdkAppLaunchContext // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_app_launch_context(_arg0)

	var _appLaunchContext *AppLaunchContext // out

	_appLaunchContext = wrapAppLaunchContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appLaunchContext
}

// Clipboard gets the clipboard used for copy/paste operations.
func (display *Display) Clipboard() *Clipboard {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_clipboard(_arg0)

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(_cret)))

	return _clipboard
}

// DefaultSeat returns the default GdkSeat for this display.
//
// Note that a display may not have a seat. In this case, this function will
// return NULL.
func (display *Display) DefaultSeat() Seater {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSeat    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_default_seat(_arg0)

	var _seat Seater // out

	_seat = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Seater)

	return _seat
}

// MonitorAtSurface gets the monitor in which the largest area of surface
// resides.
//
// Returns a monitor close to surface if it is outside of all monitors.
func (display *Display) MonitorAtSurface(surface Surfacer) *Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkSurface // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer((surface).(gextras.Nativer).Native()))

	_cret = C.gdk_display_get_monitor_at_surface(_arg0, _arg1)

	var _monitor *Monitor // out

	_monitor = wrapMonitor(externglib.Take(unsafe.Pointer(_cret)))

	return _monitor
}

// Monitors gets the list of monitors associated with this display.
//
// Subsequent calls to this function will always return the same list for the
// same display.
//
// You can listen to the GListModel::items-changed signal on this list to
// monitor changes to the monitor of this display.
func (self *Display) Monitors() gio.ListModeller {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GListModel // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_display_get_monitors(_arg0)

	var _listModel gio.ListModeller // out

	_listModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// Name gets the name of the display.
func (display *Display) Name() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PrimaryClipboard gets the clipboard used for the primary selection.
//
// On backends where the primary clipboard is not supported natively, GDK
// emulates this clipboard locally.
func (display *Display) PrimaryClipboard() *Clipboard {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_primary_clipboard(_arg0)

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(_cret)))

	return _clipboard
}

// Setting retrieves a desktop-wide setting such as double-click time for the
// display.
func (display *Display) Setting(name string, value *externglib.Value) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.GValue     // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
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
// or NULL if no ID has been defined.
func (display *Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IsClosed finds out if the display has been closed.
func (display *Display) IsClosed() bool {
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
// Check gdk.Display.IsRGBA() for whether the display supports an alpha channel.
//
// On X11 this function returns whether a compositing manager is compositing on
// display.
//
// On modern displays, this value is always TRUE.
func (display *Display) IsComposited() bool {
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

// IsRGBA returns whether surfaces on this display are created with an alpha
// channel.
//
// Even if a TRUE is returned, it is possible that the surface’s alpha channel
// won’t be honored when displaying the surface on the screen: in particular,
// for X an appropriate windowing manager and compositing manager must be
// running to provide appropriate display. Use gdk.Display.IsComposited() to
// check if that is the case.
//
// On modern displays, this value is always TRUE.
func (display *Display) IsRGBA() bool {
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

// ListSeats returns the list of seats known to display.
func (display *Display) ListSeats() []Seater {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GList      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_list_seats(_arg0)

	var _list []Seater // out

	_list = make([]Seater, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkSeat)(v)
		var dst Seater // out
		dst = (gextras.CastObject(externglib.Take(unsafe.Pointer(src)))).(Seater)
		_list = append(_list, dst)
	})

	return _list
}

// MapKeycode returns the keyvals bound to keycode.
//
// The Nth GdkKeymapKey in keys is bound to the Nth keyval in keyvals.
//
// When a keycode is pressed by the user, the keyval from this list of entries
// is selected by considering the effective keyboard group and level.
//
// Free the returned arrays with g_free().
func (display *Display) MapKeycode(keycode uint) ([]KeymapKey, []uint, bool) {
	var _arg0 *C.GdkDisplay   // out
	var _arg1 C.guint         // out
	var _arg2 *C.GdkKeymapKey // in
	var _arg4 C.int           // in
	var _arg3 *C.guint        // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint(keycode)

	_cret = C.gdk_display_map_keycode(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _keys []KeymapKey // out
	var _keyvals []uint   // out
	var _ok bool          // out

	defer C.free(unsafe.Pointer(_arg2))
	_keys = make([]KeymapKey, _arg4)
	copy(_keys, unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg4))
	defer C.free(unsafe.Pointer(_arg3))
	_keyvals = make([]uint, _arg4)
	copy(_keyvals, unsafe.Slice((*uint)(unsafe.Pointer(_arg3)), _arg4))
	if _cret != 0 {
		_ok = true
	}

	return _keys, _keyvals, _ok
}

// MapKeyval obtains a list of keycode/group/level combinations that will
// generate keyval.
//
// Groups and levels are two kinds of keyboard mode; in general, the level
// determines whether the top or bottom symbol on a key is used, and the group
// determines whether the left or right symbol is used.
//
// On US keyboards, the shift key changes the keyboard level, and there are no
// groups. A group switch key might convert a keyboard between Hebrew to English
// modes, for example.
//
// GdkEventKey contains a group field that indicates the active keyboard group.
// The level is computed from the modifier mask.
//
// The returned array should be freed with g_free().
func (display *Display) MapKeyval(keyval uint) ([]KeymapKey, bool) {
	var _arg0 *C.GdkDisplay   // out
	var _arg1 C.guint         // out
	var _arg2 *C.GdkKeymapKey // in
	var _arg3 C.int           // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint(keyval)

	_cret = C.gdk_display_map_keyval(_arg0, _arg1, &_arg2, &_arg3)

	var _keys []KeymapKey // out
	var _ok bool          // out

	defer C.free(unsafe.Pointer(_arg2))
	_keys = make([]KeymapKey, _arg3)
	copy(_keys, unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg3))
	if _cret != 0 {
		_ok = true
	}

	return _keys, _ok
}

// NotifyStartupComplete indicates to the GUI environment that the application
// has finished loading, using a given identifier.
//
// GTK will call this function automatically for gtk.Window with custom
// startup-notification identifier unless
// gtk.Window.SetAutoStartupNotification() is called to disable that feature.
func (display *Display) NotifyStartupComplete(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_display_notify_startup_complete(_arg0, _arg1)
}

// PutEvent appends the given event onto the front of the event queue for
// display.
//
// This function is only useful in very special situations and should not be
// used by applications.
func (display *Display) PutEvent(event Eventer) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkEvent   // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer((event).(gextras.Nativer).Native()))

	C.gdk_display_put_event(_arg0, _arg1)
}

// SupportsInputShapes returns TRUE if the display supports input shapes.
//
// This means that gdk.Surface.SetInputRegion() can be used to modify the input
// shape of surfaces on display.
//
// On modern displays, this value is always TRUE.
func (display *Display) SupportsInputShapes() bool {
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
// current state of the program. Calling gdk.Display.Sync() before
// gdkx11.Display.ErrorTrapPop() makes sure that any errors generated from
// earlier requests are handled before the error trap is removed.
//
// This is most useful for X11. On windowing systems where requests are handled
// synchronously, this function will do nothing.
func (display *Display) Sync() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_sync(_arg0)
}

// TranslateKey translates the contents of a GdkEventKey into a keyval,
// effective group, and level.
//
// Modifiers that affected the translation and are thus unavailable for
// application use are returned in consumed_modifiers.
//
// The effective_group is the group that was actually used for the translation;
// some keys such as Enter are not affected by the active keyboard group. The
// level is derived from state.
//
// consumed_modifiers gives modifiers that should be masked out from state when
// comparing this key press to a keyboard shortcut. For instance, on a US
// keyboard, the plus symbol is shifted, so when comparing a key press to a
// <Control>plus accelerator <Shift> should be masked out.
//
// This function should rarely be needed, since GdkEventKey already contains the
// translated keyval. It is exported for the benefit of virtualized test
// environments.
func (display *Display) TranslateKey(keycode uint, state ModifierType, group int) (keyval uint, effectiveGroup int, level int, consumed ModifierType, ok bool) {
	var _arg0 *C.GdkDisplay     // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _arg3 C.int             // out
	var _arg4 C.guint           // in
	var _arg5 C.int             // in
	var _arg6 C.int             // in
	var _arg7 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint(keycode)
	_arg2 = C.GdkModifierType(state)
	_arg3 = C.int(group)

	_cret = C.gdk_display_translate_key(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)

	var _keyval uint           // out
	var _effectiveGroup int    // out
	var _level int             // out
	var _consumed ModifierType // out
	var _ok bool               // out

	_keyval = uint(_arg4)
	_effectiveGroup = int(_arg5)
	_level = int(_arg6)
	_consumed = ModifierType(_arg7)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _effectiveGroup, _level, _consumed, _ok
}

// DisplayGetDefault gets the default GdkDisplay.
//
// This is a convenience function for: gdk_display_manager_get_default_display
// (gdk_display_manager_get ()).
func DisplayGetDefault() *Display {
	var _cret *C.GdkDisplay // in

	_cret = C.gdk_display_get_default()

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// DisplayOpen opens a display.
func DisplayOpen(displayName string) *Display {
	var _arg1 *C.char       // out
	var _cret *C.GdkDisplay // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(displayName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_display_open(_arg1)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}
