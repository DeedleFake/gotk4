// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_input_source_get_type()), F: marshalInputSource},
		{T: externglib.Type(C.gdk_device_get_type()), F: marshalDevicer},
	})
}

// InputSource: enumeration describing the type of an input device in general
// terms.
type InputSource int

const (
	// SourceMouse: device is a mouse. (This will be reported for the core
	// pointer, even if it is something else, such as a trackball.)
	SourceMouse InputSource = iota
	// SourcePen: device is a stylus of a graphics tablet or similar device.
	SourcePen
	// SourceKeyboard: device is a keyboard.
	SourceKeyboard
	// SourceTouchscreen: device is a direct-input touch device, such as a
	// touchscreen or tablet
	SourceTouchscreen
	// SourceTouchpad: device is an indirect touch device, such as a touchpad
	SourceTouchpad
	// SourceTrackpoint: device is a trackpoint
	SourceTrackpoint
	// SourceTabletPad: device is a "pad", a collection of buttons, rings and
	// strips found in drawing tablets
	SourceTabletPad
)

func marshalInputSource(p uintptr) (interface{}, error) {
	return InputSource(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for InputSource.
func (i InputSource) String() string {
	switch i {
	case SourceMouse:
		return "Mouse"
	case SourcePen:
		return "Pen"
	case SourceKeyboard:
		return "Keyboard"
	case SourceTouchscreen:
		return "Touchscreen"
	case SourceTouchpad:
		return "Touchpad"
	case SourceTrackpoint:
		return "Trackpoint"
	case SourceTabletPad:
		return "TabletPad"
	default:
		return fmt.Sprintf("InputSource(%d)", i)
	}
}

// Device: GdkDevice object represents an input device, such as a keyboard, a
// mouse, or a touchpad.
//
// See the gdk.Seat documentation for more information about the various kinds
// of devices, and their relationships.
type Device struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Device)(nil)

// Devicer describes Device's abstract methods.
type Devicer interface {
	// CapsLockState retrieves whether the Caps Lock modifier of the keyboard is
	// locked.
	CapsLockState() bool
	// DeviceTool retrieves the current tool for device.
	DeviceTool() *DeviceTool
	// Direction returns the direction of effective layout of the keyboard.
	Direction() pango.Direction
	// Display returns the GdkDisplay to which device pertains.
	Display() *Display
	// HasCursor determines whether the pointer follows device motion.
	HasCursor() bool
	// ModifierState retrieves the current modifier state of the keyboard.
	ModifierState() ModifierType
	// Name: name of the device, suitable for showing in a user interface.
	Name() string
	// NumLockState retrieves whether the Num Lock modifier of the keyboard is
	// locked.
	NumLockState() bool
	// NumTouches retrieves the number of touch points associated to device.
	NumTouches() uint
	// ProductID returns the product ID of this device.
	ProductID() string
	// ScrollLockState retrieves whether the Scroll Lock modifier of the
	// keyboard is locked.
	ScrollLockState() bool
	// Seat returns the GdkSeat the device belongs to.
	Seat() Seater
	// Source determines the type of the device.
	Source() InputSource
	// SurfaceAtPosition obtains the surface underneath device, returning the
	// location of the device in win_x and win_y Returns NULL if the surface
	// tree under device is not known to GDK (for example, belongs to another
	// application).
	SurfaceAtPosition() (winX float64, winY float64, surface Surfacer)
	// Timestamp returns the timestamp of the last activity for this device.
	Timestamp() uint32
	// VendorID returns the vendor ID of this device.
	VendorID() string
	// HasBidiLayouts determines if layouts for both right-to-left and
	// left-to-right languages are in use on the keyboard.
	HasBidiLayouts() bool
}

var _ Devicer = (*Device)(nil)

func wrapDevice(obj *externglib.Object) *Device {
	return &Device{
		Object: obj,
	}
}

func marshalDevicer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDevice(obj), nil
}

// CapsLockState retrieves whether the Caps Lock modifier of the keyboard is
// locked.
//
// This is only relevant for keyboard devices.
func (device *Device) CapsLockState() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_caps_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DeviceTool retrieves the current tool for device.
func (device *Device) DeviceTool() *DeviceTool {
	var _arg0 *C.GdkDevice     // out
	var _cret *C.GdkDeviceTool // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_device_tool(_arg0)

	var _deviceTool *DeviceTool // out

	_deviceTool = wrapDeviceTool(externglib.Take(unsafe.Pointer(_cret)))

	return _deviceTool
}

// Direction returns the direction of effective layout of the keyboard.
//
// This is only relevant for keyboard devices.
//
// The direction of a layout is the direction of the majority of its symbols.
// See pango.UnicharDirection().
func (device *Device) Direction() pango.Direction {
	var _arg0 *C.GdkDevice     // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_direction(_arg0)

	var _direction pango.Direction // out

	_direction = pango.Direction(_cret)

	return _direction
}

// Display returns the GdkDisplay to which device pertains.
func (device *Device) Display() *Display {
	var _arg0 *C.GdkDevice  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_display(_arg0)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// HasCursor determines whether the pointer follows device motion.
//
// This is not meaningful for keyboard devices, which don't have a pointer.
func (device *Device) HasCursor() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_has_cursor(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ModifierState retrieves the current modifier state of the keyboard.
//
// This is only relevant for keyboard devices.
func (device *Device) ModifierState() ModifierType {
	var _arg0 *C.GdkDevice      // out
	var _cret C.GdkModifierType // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_modifier_state(_arg0)

	var _modifierType ModifierType // out

	_modifierType = ModifierType(_cret)

	return _modifierType
}

// Name: name of the device, suitable for showing in a user interface.
func (device *Device) Name() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NumLockState retrieves whether the Num Lock modifier of the keyboard is
// locked.
//
// This is only relevant for keyboard devices.
func (device *Device) NumLockState() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_num_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NumTouches retrieves the number of touch points associated to device.
func (device *Device) NumTouches() uint {
	var _arg0 *C.GdkDevice // out
	var _cret C.guint      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_num_touches(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ProductID returns the product ID of this device.
//
// This ID is retrieved from the device, and does not change. See
// gdk.Device.GetVendorID() for more information.
func (device *Device) ProductID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_product_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ScrollLockState retrieves whether the Scroll Lock modifier of the keyboard is
// locked.
//
// This is only relevant for keyboard devices.
func (device *Device) ScrollLockState() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_scroll_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Seat returns the GdkSeat the device belongs to.
func (device *Device) Seat() Seater {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkSeat   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_seat(_arg0)

	var _seat Seater // out

	_seat = (*gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Seater)

	return _seat
}

// Source determines the type of the device.
func (device *Device) Source() InputSource {
	var _arg0 *C.GdkDevice     // out
	var _cret C.GdkInputSource // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_source(_arg0)

	var _inputSource InputSource // out

	_inputSource = InputSource(_cret)

	return _inputSource
}

// SurfaceAtPosition obtains the surface underneath device, returning the
// location of the device in win_x and win_y
//
// Returns NULL if the surface tree under device is not known to GDK (for
// example, belongs to another application).
func (device *Device) SurfaceAtPosition() (winX float64, winY float64, surface Surfacer) {
	var _arg0 *C.GdkDevice  // out
	var _arg1 C.double      // in
	var _arg2 C.double      // in
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_surface_at_position(_arg0, &_arg1, &_arg2)

	var _winX float64     // out
	var _winY float64     // out
	var _surface Surfacer // out

	_winX = float64(_arg1)
	_winY = float64(_arg2)
	_surface = (*gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Surfacer)

	return _winX, _winY, _surface
}

// Timestamp returns the timestamp of the last activity for this device.
//
// In practice, this means the timestamp of the last event that was received
// from the OS for this device. (GTK may occasionally produce events for a
// device that are not received from the OS, and will not update the timestamp).
func (device *Device) Timestamp() uint32 {
	var _arg0 *C.GdkDevice // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_timestamp(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// VendorID returns the vendor ID of this device.
//
// This ID is retrieved from the device, and does not change.
//
// This function, together with gdk.Device.GetProductID(), can be used to eg.
// compose GSettings paths to store settings for this device.
//
//     static GSettings *
//     get_device_settings (GdkDevice *device)
//     {
//       const char *vendor, *product;
//       GSettings *settings;
//       GdkDevice *device;
//       char *path;
//
//       vendor = gdk_device_get_vendor_id (device);
//       product = gdk_device_get_product_id (device);
//
//       path = g_strdup_printf ("/org/example/app/devices/s:s/", vendor, product);
//       settings = g_settings_new_with_path (DEVICE_SCHEMA, path);
//       g_free (path);
//
//       return settings;
//     }
func (device *Device) VendorID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_get_vendor_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasBidiLayouts determines if layouts for both right-to-left and left-to-right
// languages are in use on the keyboard.
//
// This is only relevant for keyboard devices.
func (device *Device) HasBidiLayouts() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_device_has_bidi_layouts(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TimeCoord stores a single event in a motion history.
type TimeCoord struct {
	nocopy gextras.NoCopy
	native *C.GdkTimeCoord
}

// Time: timestamp for this event.
func (t *TimeCoord) Time() uint32 {
	var v uint32 // out
	v = uint32(t.native.time)
	return v
}

// Flags indicating what axes are present
func (t *TimeCoord) Flags() AxisFlags {
	var v AxisFlags // out
	v = AxisFlags(t.native.flags)
	return v
}

// Axes axis values
func (t *TimeCoord) Axes() [12]float64 {
	var v [12]float64
	v = *(*[12]float64)(unsafe.Pointer(&t.native.axes))
	return v
}
