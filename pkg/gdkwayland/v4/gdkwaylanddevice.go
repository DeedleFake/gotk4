// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gdk_wayland_device_get_type()), F: marshalWaylandDevice},
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

	// AsDevice casts the class to the gdk.Device interface.
	AsDevice() gdk.Device

	// GetCapsLockState retrieves whether the Caps Lock modifier of the keyboard
	// is locked.
	//
	// This is only relevant for keyboard devices.
	//
	// This method is inherited from gdk.Device
	GetCapsLockState() bool
	// GetDeviceTool retrieves the current tool for @device.
	//
	// This method is inherited from gdk.Device
	GetDeviceTool() gdk.DeviceTool
	// GetDirection returns the direction of effective layout of the keyboard.
	//
	// This is only relevant for keyboard devices.
	//
	// The direction of a layout is the direction of the majority of its
	// symbols. See [func@Pango.unichar_direction].
	//
	// This method is inherited from gdk.Device
	GetDirection() pango.Direction
	// GetDisplay returns the `GdkDisplay` to which @device pertains.
	//
	// This method is inherited from gdk.Device
	GetDisplay() gdk.Display
	// GetHasCursor determines whether the pointer follows device motion.
	//
	// This is not meaningful for keyboard devices, which don't have a pointer.
	//
	// This method is inherited from gdk.Device
	GetHasCursor() bool
	// GetModifierState retrieves the current modifier state of the keyboard.
	//
	// This is only relevant for keyboard devices.
	//
	// This method is inherited from gdk.Device
	GetModifierState() gdk.ModifierType
	// GetName: the name of the device, suitable for showing in a user
	// interface.
	//
	// This method is inherited from gdk.Device
	GetName() string
	// GetNumLockState retrieves whether the Num Lock modifier of the keyboard
	// is locked.
	//
	// This is only relevant for keyboard devices.
	//
	// This method is inherited from gdk.Device
	GetNumLockState() bool
	// GetNumTouches retrieves the number of touch points associated to @device.
	//
	// This method is inherited from gdk.Device
	GetNumTouches() uint
	// GetProductID returns the product ID of this device.
	//
	// This ID is retrieved from the device, and does not change. See
	// [method@Gdk.Device.get_vendor_id] for more information.
	//
	// This method is inherited from gdk.Device
	GetProductID() string
	// GetScrollLockState retrieves whether the Scroll Lock modifier of the
	// keyboard is locked.
	//
	// This is only relevant for keyboard devices.
	//
	// This method is inherited from gdk.Device
	GetScrollLockState() bool
	// GetSeat returns the `GdkSeat` the device belongs to.
	//
	// This method is inherited from gdk.Device
	GetSeat() gdk.Seat
	// GetSource determines the type of the device.
	//
	// This method is inherited from gdk.Device
	GetSource() gdk.InputSource
	// GetSurfaceAtPosition obtains the surface underneath @device, returning
	// the location of the device in @win_x and @win_y
	//
	// Returns nil if the surface tree under @device is not known to GDK (for
	// example, belongs to another application).
	//
	// This method is inherited from gdk.Device
	GetSurfaceAtPosition() (winX float64, winY float64, surface gdk.Surface)
	// GetTimestamp returns the timestamp of the last activity for this device.
	//
	// In practice, this means the timestamp of the last event that was received
	// from the OS for this device. (GTK may occasionally produce events for a
	// device that are not received from the OS, and will not update the
	// timestamp).
	//
	// This method is inherited from gdk.Device
	GetTimestamp() uint32
	// GetVendorID returns the vendor ID of this device.
	//
	// This ID is retrieved from the device, and does not change.
	//
	// This function, together with [method@Gdk.Device.get_product_id], can be
	// used to eg. compose `GSettings` paths to store settings for this device.
	//
	// “`c static GSettings * get_device_settings (GdkDevice *device) { const
	// char *vendor, *product; GSettings *settings; GdkDevice *device; char
	// *path;
	//
	//      vendor = gdk_device_get_vendor_id (device);
	//      product = gdk_device_get_product_id (device);
	//
	//      path = g_strdup_printf ("/org/example/app/devices/s:s/", vendor, product);
	//      settings = g_settings_new_with_path (DEVICE_SCHEMA, path);
	//      g_free (path);
	//
	//      return settings;
	//    }
	//
	// “`
	//
	// This method is inherited from gdk.Device
	GetVendorID() string
	// HasBidiLayouts determines if layouts for both right-to-left and
	// left-to-right languages are in use on the keyboard.
	//
	// This is only relevant for keyboard devices.
	//
	// This method is inherited from gdk.Device
	HasBidiLayouts() bool

	// NodePath returns the `/dev/input/event*` path of this device.
	//
	// For `GdkDevice`s that possibly coalesce multiple hardware devices (eg.
	// mouse, keyboard, touch,...), this function will return nil.
	//
	// This is most notably implemented for devices of type GDK_SOURCE_PEN,
	// GDK_SOURCE_TABLET_PAD.
	NodePath() string
}

// waylandDevice implements the WaylandDevice interface.
type waylandDevice struct {
	*externglib.Object
}

var _ WaylandDevice = (*waylandDevice)(nil)

// WrapWaylandDevice wraps a GObject to a type that implements
// interface WaylandDevice. It is primarily used internally.
func WrapWaylandDevice(obj *externglib.Object) WaylandDevice {
	return waylandDevice{obj}
}

func marshalWaylandDevice(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDevice(obj), nil
}

func (w waylandDevice) AsDevice() gdk.Device {
	return gdk.WrapDevice(gextras.InternObject(w))
}

func (d waylandDevice) GetCapsLockState() bool {
	return gdk.WrapDevice(gextras.InternObject(d)).GetCapsLockState()
}

func (d waylandDevice) GetDeviceTool() gdk.DeviceTool {
	return gdk.WrapDevice(gextras.InternObject(d)).GetDeviceTool()
}

func (d waylandDevice) GetDirection() pango.Direction {
	return gdk.WrapDevice(gextras.InternObject(d)).GetDirection()
}

func (d waylandDevice) GetDisplay() gdk.Display {
	return gdk.WrapDevice(gextras.InternObject(d)).GetDisplay()
}

func (d waylandDevice) GetHasCursor() bool {
	return gdk.WrapDevice(gextras.InternObject(d)).GetHasCursor()
}

func (d waylandDevice) GetModifierState() gdk.ModifierType {
	return gdk.WrapDevice(gextras.InternObject(d)).GetModifierState()
}

func (d waylandDevice) GetName() string {
	return gdk.WrapDevice(gextras.InternObject(d)).GetName()
}

func (d waylandDevice) GetNumLockState() bool {
	return gdk.WrapDevice(gextras.InternObject(d)).GetNumLockState()
}

func (d waylandDevice) GetNumTouches() uint {
	return gdk.WrapDevice(gextras.InternObject(d)).GetNumTouches()
}

func (d waylandDevice) GetProductID() string {
	return gdk.WrapDevice(gextras.InternObject(d)).GetProductID()
}

func (d waylandDevice) GetScrollLockState() bool {
	return gdk.WrapDevice(gextras.InternObject(d)).GetScrollLockState()
}

func (d waylandDevice) GetSeat() gdk.Seat {
	return gdk.WrapDevice(gextras.InternObject(d)).GetSeat()
}

func (d waylandDevice) GetSource() gdk.InputSource {
	return gdk.WrapDevice(gextras.InternObject(d)).GetSource()
}

func (d waylandDevice) GetSurfaceAtPosition() (winX float64, winY float64, surface gdk.Surface) {
	return gdk.WrapDevice(gextras.InternObject(d)).GetSurfaceAtPosition()
}

func (d waylandDevice) GetTimestamp() uint32 {
	return gdk.WrapDevice(gextras.InternObject(d)).GetTimestamp()
}

func (d waylandDevice) GetVendorID() string {
	return gdk.WrapDevice(gextras.InternObject(d)).GetVendorID()
}

func (d waylandDevice) HasBidiLayouts() bool {
	return gdk.WrapDevice(gextras.InternObject(d)).HasBidiLayouts()
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
