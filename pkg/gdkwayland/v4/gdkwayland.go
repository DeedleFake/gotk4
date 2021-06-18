// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

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
		{T: externglib.Type(C.gdk_wayland_device_get_type()), F: marshalWaylandDevice},
		{T: externglib.Type(C.gdk_wayland_display_get_type()), F: marshalWaylandDisplay},
		{T: externglib.Type(C.gdk_wayland_monitor_get_type()), F: marshalWaylandMonitor},
		{T: externglib.Type(C.gdk_wayland_popup_get_type()), F: marshalWaylandPopup},
		{T: externglib.Type(C.gdk_wayland_seat_get_type()), F: marshalWaylandSeat},
		{T: externglib.Type(C.gdk_wayland_surface_get_type()), F: marshalWaylandSurface},
		{T: externglib.Type(C.gdk_wayland_toplevel_get_type()), F: marshalWaylandToplevel},
	})
}

// WaylandDevice: the Wayland implementation of `GdkDevice`.
//
// Beyond the regular [class@Gdk.Device] API, the Wayland implementation
// provides access to Wayland objects such as the `wl_seat` with
// [method@GdkWayland.WaylandDevice.get_wl_seat], the `wl_keyboard` with
// [method@GdkWayland.WaylandDevice.get_wl_keyboard] and the `wl_pointer` with
// [method@GdkWayland.WaylandDevice.get_wl_pointer].
type WaylandDevice interface {
	gdk.Device

	// NodePath returns the `/dev/input/event*` path of this device.
	//
	// For `GdkDevice`s that possibly coalesce multiple hardware devices (eg.
	// mouse, keyboard, touch,...), this function will return nil.
	//
	// This is most notably implemented for devices of type GDK_SOURCE_PEN,
	// GDK_SOURCE_TABLET_PAD.
	NodePath() string
}

// waylandDevice implements the WaylandDevice class.
type waylandDevice struct {
	gdk.Device
}

var _ WaylandDevice = (*waylandDevice)(nil)

// WrapWaylandDevice wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandDevice(obj *externglib.Object) WaylandDevice {
	return waylandDevice{
		Device: gdk.WrapDevice(obj),
	}
}

func marshalWaylandDevice(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDevice(obj), nil
}

func (d waylandDevice) NodePath() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_wayland_device_get_node_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
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
	gdk.Display

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

// waylandDisplay implements the WaylandDisplay class.
type waylandDisplay struct {
	gdk.Display
}

var _ WaylandDisplay = (*waylandDisplay)(nil)

// WrapWaylandDisplay wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandDisplay(obj *externglib.Object) WaylandDisplay {
	return waylandDisplay{
		Display: gdk.WrapDisplay(obj),
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDisplay(obj), nil
}

func (d waylandDisplay) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_wayland_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d waylandDisplay) QueryRegistry(global string) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(global))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_display_query_registry(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d waylandDisplay) SetCursorTheme(name string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.int)(size)

	C.gdk_wayland_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

func (d waylandDisplay) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_display_set_startup_notification_id(_arg0, _arg1)
}

// WaylandMonitor: the Wayland implementation of `GdkMonitor`.
//
// Beyond the [class@Gdk.Monitor] API, the Wayland implementation offers access
// to the Wayland `wl_output` object with
// [method@GdkWayland.WaylandMonitor.get_wl_output].
type WaylandMonitor interface {
	gdk.Monitor
}

// waylandMonitor implements the WaylandMonitor class.
type waylandMonitor struct {
	gdk.Monitor
}

var _ WaylandMonitor = (*waylandMonitor)(nil)

// WrapWaylandMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandMonitor(obj *externglib.Object) WaylandMonitor {
	return waylandMonitor{
		Monitor: gdk.WrapMonitor(obj),
	}
}

func marshalWaylandMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandMonitor(obj), nil
}

// WaylandPopup: the Wayland implementation of `GdkPopup`.
type WaylandPopup interface {
	WaylandSurface
}

// waylandPopup implements the WaylandPopup class.
type waylandPopup struct {
	WaylandSurface
}

var _ WaylandPopup = (*waylandPopup)(nil)

// WrapWaylandPopup wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandPopup(obj *externglib.Object) WaylandPopup {
	return waylandPopup{
		WaylandSurface: WrapWaylandSurface(obj),
	}
}

func marshalWaylandPopup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandPopup(obj), nil
}

// WaylandSeat: the Wayland implementation of `GdkSeat`.
//
// Beyond the regular [class@Gdk.Seat] API, the Wayland implementation provides
// access to the Wayland `wl_seat` object with
// [method@GdkWayland.WaylandSeat.get_wl_seat].
type WaylandSeat interface {
	gdk.Seat
}

// waylandSeat implements the WaylandSeat class.
type waylandSeat struct {
	gdk.Seat
}

var _ WaylandSeat = (*waylandSeat)(nil)

// WrapWaylandSeat wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandSeat(obj *externglib.Object) WaylandSeat {
	return waylandSeat{
		Seat: gdk.WrapSeat(obj),
	}
}

func marshalWaylandSeat(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandSeat(obj), nil
}

// WaylandSurface: the Wayland implementation of `GdkSurface`.
//
// Beyond the [class@Gdk.Surface] API, the Wayland implementation offers access
// to the Wayland `wl_surface` object with
// [method@GdkWayland.WaylandSurface.get_wl_surface].
type WaylandSurface interface {
	gdk.Surface
}

// waylandSurface implements the WaylandSurface class.
type waylandSurface struct {
	gdk.Surface
}

var _ WaylandSurface = (*waylandSurface)(nil)

// WrapWaylandSurface wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandSurface(obj *externglib.Object) WaylandSurface {
	return waylandSurface{
		Surface: gdk.WrapSurface(obj),
	}
}

func marshalWaylandSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandSurface(obj), nil
}

// WaylandToplevel: the Wayland implementation of `GdkToplevel`.
//
// Beyond the [interface@Gdk.Toplevel] API, the Wayland implementation has API
// to set up cross-process parent-child relationships between surfaces with
// [method@GdkWayland.WaylandToplevel.export_handle] and
// [method@GdkWayland.WaylandToplevel.set_transient_for_exported].
type WaylandToplevel interface {
	WaylandSurface

	// SetApplicationID sets the application id on a `GdkToplevel`.
	SetApplicationID(applicationId string)
	// SetTransientForExported marks @toplevel as transient for the surface to
	// which the given @parent_handle_str refers.
	//
	// Typically, the handle will originate from a
	// [method@GdkWayland.WaylandToplevel.export_handle] call in another
	// process.
	//
	// Note that this API depends on an unstable Wayland protocol, and thus may
	// require changes in the future.
	SetTransientForExported(parentHandleStr string) bool
	// UnexportHandle destroys the handle that was obtained with
	// gdk_wayland_toplevel_export_handle().
	//
	// It is an error to call this function on a surface that does not have a
	// handle.
	//
	// Note that this API depends on an unstable Wayland protocol, and thus may
	// require changes in the future.
	UnexportHandle()
}

// waylandToplevel implements the WaylandToplevel class.
type waylandToplevel struct {
	WaylandSurface
}

var _ WaylandToplevel = (*waylandToplevel)(nil)

// WrapWaylandToplevel wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandToplevel(obj *externglib.Object) WaylandToplevel {
	return waylandToplevel{
		WaylandSurface: WrapWaylandSurface(obj),
	}
}

func marshalWaylandToplevel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandToplevel(obj), nil
}

func (t waylandToplevel) SetApplicationID(applicationId string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_toplevel_set_application_id(_arg0, _arg1)
}

func (t waylandToplevel) SetTransientForExported(parentHandleStr string) bool {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(parentHandleStr))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_toplevel_set_transient_for_exported(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t waylandToplevel) UnexportHandle() {
	var _arg0 *C.GdkToplevel // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	C.gdk_wayland_toplevel_unexport_handle(_arg0)
}
