// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_xi2_get_type()), F: marshalX11DeviceXI2},
	})
}

type X11DeviceXI2 interface {
	gdk.Device

	// AsDevice casts the class to the gdk.Device interface.
	AsDevice() gdk.Device

	// GetAssociatedDevice returns the associated device to @device, if @device
	// is of type GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or
	// keyboard.
	//
	// If @device is of type GDK_DEVICE_TYPE_SLAVE, it will return the master
	// device to which @device is attached to.
	//
	// If @device is of type GDK_DEVICE_TYPE_FLOATING, nil will be returned, as
	// there is no associated device.
	//
	// This method is inherited from gdk.Device
	GetAssociatedDevice() gdk.Device
	// GetAxes returns the axes currently available on the device.
	//
	// This method is inherited from gdk.Device
	GetAxes() gdk.AxisFlags
	// GetAxisUse returns the axis use for @index_.
	//
	// This method is inherited from gdk.Device
	GetAxisUse(index_ uint) gdk.AxisUse
	// GetDeviceType returns the device type for @device.
	//
	// This method is inherited from gdk.Device
	GetDeviceType() gdk.DeviceType
	// GetDisplay returns the Display to which @device pertains.
	//
	// This method is inherited from gdk.Device
	GetDisplay() gdk.Display
	// GetHasCursor determines whether the pointer follows device motion. This
	// is not meaningful for keyboard devices, which don't have a pointer.
	//
	// This method is inherited from gdk.Device
	GetHasCursor() bool
	// GetKey: if @index_ has a valid keyval, this function will return true and
	// fill in @keyval and @modifiers with the keyval settings.
	//
	// This method is inherited from gdk.Device
	GetKey(index_ uint) (uint, gdk.ModifierType, bool)
	// GetLastEventWindow gets information about which window the given pointer
	// device is in, based on events that have been received so far from the
	// display server. If another application has a pointer grab, or this
	// application has a grab with owner_events = false, nil may be returned
	// even if the pointer is physically over one of this application's windows.
	//
	// This method is inherited from gdk.Device
	GetLastEventWindow() gdk.Window
	// GetMode determines the mode of the device.
	//
	// This method is inherited from gdk.Device
	GetMode() gdk.InputMode
	// GetNAxes returns the number of axes the device currently has.
	//
	// This method is inherited from gdk.Device
	GetNAxes() int
	// GetNKeys returns the number of keys the device currently has.
	//
	// This method is inherited from gdk.Device
	GetNKeys() int
	// GetName determines the name of the device.
	//
	// This method is inherited from gdk.Device
	GetName() string
	// GetPosition gets the current location of @device. As a slave device
	// coordinates are those of its master pointer, This function may not be
	// called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is an
	// ongoing grab on them, see gdk_device_grab().
	//
	// This method is inherited from gdk.Device
	GetPosition() (screen gdk.Screen, x int, y int)
	// GetPositionDouble gets the current location of @device in double
	// precision. As a slave device's coordinates are those of its master
	// pointer, this function may not be called on devices of type
	// GDK_DEVICE_TYPE_SLAVE, unless there is an ongoing grab on them. See
	// gdk_device_grab().
	//
	// This method is inherited from gdk.Device
	GetPositionDouble() (screen gdk.Screen, x float64, y float64)
	// GetProductID returns the product ID of this device, or nil if this
	// information couldn't be obtained. This ID is retrieved from the device,
	// and is thus constant for it. See gdk_device_get_vendor_id() for more
	// information.
	//
	// This method is inherited from gdk.Device
	GetProductID() string
	// GetSeat returns the Seat the device belongs to.
	//
	// This method is inherited from gdk.Device
	GetSeat() gdk.Seat
	// GetSource determines the type of the device.
	//
	// This method is inherited from gdk.Device
	GetSource() gdk.InputSource
	// GetVendorID returns the vendor ID of this device, or nil if this
	// information couldn't be obtained. This ID is retrieved from the device,
	// and is thus constant for it.
	//
	// This function, together with gdk_device_get_product_id(), can be used to
	// eg. compose #GSettings paths to store settings for this device.
	//
	//     static GSettings *
	//     get_device_settings (GdkDevice *device)
	//     {
	//       const gchar *vendor, *product;
	//       GSettings *settings;
	//       GdkDevice *device;
	//       gchar *path;
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
	//
	// This method is inherited from gdk.Device
	GetVendorID() string
	// GetWindowAtPosition obtains the window underneath @device, returning the
	// location of the device in @win_x and @win_y. Returns nil if the window
	// tree under @device is not known to GDK (for example, belongs to another
	// application).
	//
	// As a slave device coordinates are those of its master pointer, This
	// function may not be called on devices of type GDK_DEVICE_TYPE_SLAVE,
	// unless there is an ongoing grab on them, see gdk_device_grab().
	//
	// This method is inherited from gdk.Device
	GetWindowAtPosition() (winX int, winY int, window gdk.Window)
	// GetWindowAtPositionDouble obtains the window underneath @device,
	// returning the location of the device in @win_x and @win_y in double
	// precision. Returns nil if the window tree under @device is not known to
	// GDK (for example, belongs to another application).
	//
	// As a slave device coordinates are those of its master pointer, This
	// function may not be called on devices of type GDK_DEVICE_TYPE_SLAVE,
	// unless there is an ongoing grab on them, see gdk_device_grab().
	//
	// This method is inherited from gdk.Device
	GetWindowAtPositionDouble() (winX float64, winY float64, window gdk.Window)
	// Grab grabs the device so that all events coming from this device are
	// passed to this application until the device is ungrabbed with
	// gdk_device_ungrab(), or the window becomes unviewable. This overrides any
	// previous grab on the device by this client.
	//
	// Note that @device and @window need to be on the same display.
	//
	// Device grabs are used for operations which need complete control over the
	// given device events (either pointer or keyboard). For example in GTK+
	// this is used for Drag and Drop operations, popup menus and such.
	//
	// Note that if the event mask of an X window has selected both button press
	// and button release events, then a button press event will cause an
	// automatic pointer grab until the button is released. X does this
	// automatically since most applications expect to receive button press and
	// release events in pairs. It is equivalent to a pointer grab on the window
	// with @owner_events set to true.
	//
	// If you set up anything at the time you take the grab that needs to be
	// cleaned up when the grab ends, you should handle the EventGrabBroken
	// events that are emitted when the grab ends unvoluntarily.
	//
	// Deprecated: since version 3.20.
	//
	// This method is inherited from gdk.Device
	Grab(window gdk.Window, grabOwnership gdk.GrabOwnership, ownerEvents bool, eventMask gdk.EventMask, cursor gdk.Cursor, time_ uint32) gdk.GrabStatus
	// SetAxisUse specifies how an axis of a device is used.
	//
	// This method is inherited from gdk.Device
	SetAxisUse(index_ uint, use gdk.AxisUse)
	// SetKey specifies the X key event to generate when a macro button of a
	// device is pressed.
	//
	// This method is inherited from gdk.Device
	SetKey(index_ uint, keyval uint, modifiers gdk.ModifierType)
	// SetMode sets a the mode of an input device. The mode controls if the
	// device is active and whether the device’s range is mapped to the entire
	// screen or to a single window.
	//
	// Note: This is only meaningful for floating devices, master devices (and
	// slaves connected to these) drive the pointer cursor, which is not limited
	// by the input mode.
	//
	// This method is inherited from gdk.Device
	SetMode(mode gdk.InputMode) bool
	// Ungrab: release any grab on @device.
	//
	// Deprecated: since version 3.20.
	//
	// This method is inherited from gdk.Device
	Ungrab(time_ uint32)
	// Warp warps @device in @display to the point @x,@y on the screen @screen,
	// unless the device is confined to a window by a grab, in which case it
	// will be moved as far as allowed by the grab. Warping the pointer creates
	// events as if the user had moved the mouse instantaneously to the
	// destination.
	//
	// Note that the pointer should normally be under the control of the user.
	// This function was added to cover some rare use cases like keyboard
	// navigation support for the color picker in the ColorSelectionDialog.
	//
	// This method is inherited from gdk.Device
	Warp(screen gdk.Screen, x int, y int)
}

// x11DeviceXI2 implements the X11DeviceXI2 interface.
type x11DeviceXI2 struct {
	*externglib.Object
}

var _ X11DeviceXI2 = (*x11DeviceXI2)(nil)

// WrapX11DeviceXI2 wraps a GObject to a type that implements
// interface X11DeviceXI2. It is primarily used internally.
func WrapX11DeviceXI2(obj *externglib.Object) X11DeviceXI2 {
	return x11DeviceXI2{obj}
}

func marshalX11DeviceXI2(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceXI2(obj), nil
}

func (x x11DeviceXI2) AsDevice() gdk.Device {
	return gdk.WrapDevice(gextras.InternObject(x))
}

func (d x11DeviceXI2) GetAssociatedDevice() gdk.Device {
	return gdk.WrapDevice(gextras.InternObject(d)).GetAssociatedDevice()
}

func (d x11DeviceXI2) GetAxes() gdk.AxisFlags {
	return gdk.WrapDevice(gextras.InternObject(d)).GetAxes()
}

func (d x11DeviceXI2) GetAxisUse(index_ uint) gdk.AxisUse {
	return gdk.WrapDevice(gextras.InternObject(d)).GetAxisUse(index_)
}

func (d x11DeviceXI2) GetDeviceType() gdk.DeviceType {
	return gdk.WrapDevice(gextras.InternObject(d)).GetDeviceType()
}

func (d x11DeviceXI2) GetDisplay() gdk.Display {
	return gdk.WrapDevice(gextras.InternObject(d)).GetDisplay()
}

func (d x11DeviceXI2) GetHasCursor() bool {
	return gdk.WrapDevice(gextras.InternObject(d)).GetHasCursor()
}

func (d x11DeviceXI2) GetKey(index_ uint) (uint, gdk.ModifierType, bool) {
	return gdk.WrapDevice(gextras.InternObject(d)).GetKey(index_)
}

func (d x11DeviceXI2) GetLastEventWindow() gdk.Window {
	return gdk.WrapDevice(gextras.InternObject(d)).GetLastEventWindow()
}

func (d x11DeviceXI2) GetMode() gdk.InputMode {
	return gdk.WrapDevice(gextras.InternObject(d)).GetMode()
}

func (d x11DeviceXI2) GetNAxes() int {
	return gdk.WrapDevice(gextras.InternObject(d)).GetNAxes()
}

func (d x11DeviceXI2) GetNKeys() int {
	return gdk.WrapDevice(gextras.InternObject(d)).GetNKeys()
}

func (d x11DeviceXI2) GetName() string {
	return gdk.WrapDevice(gextras.InternObject(d)).GetName()
}

func (d x11DeviceXI2) GetPosition() (screen gdk.Screen, x int, y int) {
	return gdk.WrapDevice(gextras.InternObject(d)).GetPosition()
}

func (d x11DeviceXI2) GetPositionDouble() (screen gdk.Screen, x float64, y float64) {
	return gdk.WrapDevice(gextras.InternObject(d)).GetPositionDouble()
}

func (d x11DeviceXI2) GetProductID() string {
	return gdk.WrapDevice(gextras.InternObject(d)).GetProductID()
}

func (d x11DeviceXI2) GetSeat() gdk.Seat {
	return gdk.WrapDevice(gextras.InternObject(d)).GetSeat()
}

func (d x11DeviceXI2) GetSource() gdk.InputSource {
	return gdk.WrapDevice(gextras.InternObject(d)).GetSource()
}

func (d x11DeviceXI2) GetVendorID() string {
	return gdk.WrapDevice(gextras.InternObject(d)).GetVendorID()
}

func (d x11DeviceXI2) GetWindowAtPosition() (winX int, winY int, window gdk.Window) {
	return gdk.WrapDevice(gextras.InternObject(d)).GetWindowAtPosition()
}

func (d x11DeviceXI2) GetWindowAtPositionDouble() (winX float64, winY float64, window gdk.Window) {
	return gdk.WrapDevice(gextras.InternObject(d)).GetWindowAtPositionDouble()
}

func (d x11DeviceXI2) Grab(window gdk.Window, grabOwnership gdk.GrabOwnership, ownerEvents bool, eventMask gdk.EventMask, cursor gdk.Cursor, time_ uint32) gdk.GrabStatus {
	return gdk.WrapDevice(gextras.InternObject(d)).Grab(window, grabOwnership, ownerEvents, eventMask, cursor, time_)
}

func (d x11DeviceXI2) SetAxisUse(index_ uint, use gdk.AxisUse) {
	gdk.WrapDevice(gextras.InternObject(d)).SetAxisUse(index_, use)
}

func (d x11DeviceXI2) SetKey(index_ uint, keyval uint, modifiers gdk.ModifierType) {
	gdk.WrapDevice(gextras.InternObject(d)).SetKey(index_, keyval, modifiers)
}

func (d x11DeviceXI2) SetMode(mode gdk.InputMode) bool {
	return gdk.WrapDevice(gextras.InternObject(d)).SetMode(mode)
}

func (d x11DeviceXI2) Ungrab(time_ uint32) {
	gdk.WrapDevice(gextras.InternObject(d)).Ungrab(time_)
}

func (d x11DeviceXI2) Warp(screen gdk.Screen, x int, y int) {
	gdk.WrapDevice(gextras.InternObject(d)).Warp(screen, x, y)
}
