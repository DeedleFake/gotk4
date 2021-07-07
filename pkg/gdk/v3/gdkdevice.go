// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_type_get_type()), F: marshalDeviceType},
		{T: externglib.Type(C.gdk_input_mode_get_type()), F: marshalInputMode},
		{T: externglib.Type(C.gdk_input_source_get_type()), F: marshalInputSource},
		{T: externglib.Type(C.gdk_device_get_type()), F: marshalDevice},
	})
}

// DeviceType indicates the device type. See
// [above][GdkDeviceManager.description] for more information about the meaning
// of these device types.
type DeviceType int

const (
	// Master: device is a master (or virtual) device. There will be an
	// associated focus indicator on the screen.
	DeviceTypeMaster DeviceType = iota
	// Slave: device is a slave (or physical) device.
	DeviceTypeSlave
	// Floating: device is a physical device, currently not attached to any
	// virtual device.
	DeviceTypeFloating
)

func marshalDeviceType(p uintptr) (interface{}, error) {
	return DeviceType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// InputMode: enumeration that describes the mode of an input device.
type InputMode int

const (
	// Disabled: the device is disabled and will not report any events.
	ModeDisabled InputMode = iota
	// Screen: the device is enabled. The device’s coordinate space maps to the
	// entire screen.
	ModeScreen
	// Window: the device is enabled. The device’s coordinate space is mapped to
	// a single window. The manner in which this window is chosen is undefined,
	// but it will typically be the same way in which the focus window for key
	// events is determined.
	ModeWindow
)

func marshalInputMode(p uintptr) (interface{}, error) {
	return InputMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// InputSource: enumeration describing the type of an input device in general
// terms.
type InputSource int

const (
	// Mouse: the device is a mouse. (This will be reported for the core
	// pointer, even if it is something else, such as a trackball.)
	SourceMouse InputSource = iota
	// Pen: the device is a stylus of a graphics tablet or similar device.
	SourcePen
	// Eraser: the device is an eraser. Typically, this would be the other end
	// of a stylus on a graphics tablet.
	SourceEraser
	// Cursor: the device is a graphics tablet “puck” or similar device.
	SourceCursor
	// Keyboard: the device is a keyboard.
	SourceKeyboard
	// Touchscreen: the device is a direct-input touch device, such as a
	// touchscreen or tablet. This device type has been added in 3.4.
	SourceTouchscreen
	// Touchpad: the device is an indirect touch device, such as a touchpad.
	// This device type has been added in 3.4.
	SourceTouchpad
	// Trackpoint: the device is a trackpoint. This device type has been added
	// in 3.22
	SourceTrackpoint
	// TabletPad: the device is a "pad", a collection of buttons, rings and
	// strips found in drawing tablets. This device type has been added in 3.22.
	SourceTabletPad
)

func marshalInputSource(p uintptr) (interface{}, error) {
	return InputSource(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Device: the Device object represents a single input device, such as a
// keyboard, a mouse, a touchpad, etc.
//
// See the DeviceManager documentation for more information about the various
// kinds of master and slave devices, and their relationships.
type Device interface {
	gextras.Objector

	// AssociatedDevice returns the associated device to @device, if @device is
	// of type GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or
	// keyboard.
	//
	// If @device is of type GDK_DEVICE_TYPE_SLAVE, it will return the master
	// device to which @device is attached to.
	//
	// If @device is of type GDK_DEVICE_TYPE_FLOATING, nil will be returned, as
	// there is no associated device.
	AssociatedDevice() Device
	// Axes returns the axes currently available on the device.
	Axes() AxisFlags
	// AxisUse returns the axis use for @index_.
	AxisUse(index_ uint) AxisUse
	// DeviceType returns the device type for @device.
	DeviceType() DeviceType
	// Display returns the Display to which @device pertains.
	Display() Display
	// HasCursor determines whether the pointer follows device motion. This is
	// not meaningful for keyboard devices, which don't have a pointer.
	HasCursor() bool
	// Key: if @index_ has a valid keyval, this function will return true and
	// fill in @keyval and @modifiers with the keyval settings.
	Key(index_ uint) (uint, ModifierType, bool)
	// LastEventWindow gets information about which window the given pointer
	// device is in, based on events that have been received so far from the
	// display server. If another application has a pointer grab, or this
	// application has a grab with owner_events = false, nil may be returned
	// even if the pointer is physically over one of this application's windows.
	LastEventWindow() Window
	// Mode determines the mode of the device.
	Mode() InputMode
	// NAxes returns the number of axes the device currently has.
	NAxes() int
	// NKeys returns the number of keys the device currently has.
	NKeys() int
	// Name determines the name of the device.
	Name() string
	// Position gets the current location of @device. As a slave device
	// coordinates are those of its master pointer, This function may not be
	// called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is an
	// ongoing grab on them, see gdk_device_grab().
	Position() (screen Screen, x int, y int)
	// PositionDouble gets the current location of @device in double precision.
	// As a slave device's coordinates are those of its master pointer, this
	// function may not be called on devices of type GDK_DEVICE_TYPE_SLAVE,
	// unless there is an ongoing grab on them. See gdk_device_grab().
	PositionDouble() (screen Screen, x float64, y float64)
	// ProductID returns the product ID of this device, or nil if this
	// information couldn't be obtained. This ID is retrieved from the device,
	// and is thus constant for it. See gdk_device_get_vendor_id() for more
	// information.
	ProductID() string
	// Seat returns the Seat the device belongs to.
	Seat() Seat
	// Source determines the type of the device.
	Source() InputSource
	// VendorID returns the vendor ID of this device, or nil if this information
	// couldn't be obtained. This ID is retrieved from the device, and is thus
	// constant for it.
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
	VendorID() string
	// WindowAtPosition obtains the window underneath @device, returning the
	// location of the device in @win_x and @win_y. Returns nil if the window
	// tree under @device is not known to GDK (for example, belongs to another
	// application).
	//
	// As a slave device coordinates are those of its master pointer, This
	// function may not be called on devices of type GDK_DEVICE_TYPE_SLAVE,
	// unless there is an ongoing grab on them, see gdk_device_grab().
	WindowAtPosition() (winX int, winY int, window Window)
	// WindowAtPositionDouble obtains the window underneath @device, returning
	// the location of the device in @win_x and @win_y in double precision.
	// Returns nil if the window tree under @device is not known to GDK (for
	// example, belongs to another application).
	//
	// As a slave device coordinates are those of its master pointer, This
	// function may not be called on devices of type GDK_DEVICE_TYPE_SLAVE,
	// unless there is an ongoing grab on them, see gdk_device_grab().
	WindowAtPositionDouble() (winX float64, winY float64, window Window)
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
	Grab(window Window, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor Cursor, time_ uint32) GrabStatus
	// SetAxisUse specifies how an axis of a device is used.
	SetAxisUse(index_ uint, use AxisUse)
	// SetKey specifies the X key event to generate when a macro button of a
	// device is pressed.
	SetKey(index_ uint, keyval uint, modifiers ModifierType)
	// SetMode sets a the mode of an input device. The mode controls if the
	// device is active and whether the device’s range is mapped to the entire
	// screen or to a single window.
	//
	// Note: This is only meaningful for floating devices, master devices (and
	// slaves connected to these) drive the pointer cursor, which is not limited
	// by the input mode.
	SetMode(mode InputMode) bool
	// Ungrab: release any grab on @device.
	//
	// Deprecated: since version 3.20.
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
	Warp(screen Screen, x int, y int)
}

// device implements the Device interface.
type device struct {
	*externglib.Object
}

var _ Device = (*device)(nil)

// WrapDevice wraps a GObject to a type that implements
// interface Device. It is primarily used internally.
func WrapDevice(obj *externglib.Object) Device {
	return device{obj}
}

func marshalDevice(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDevice(obj), nil
}

func (d device) AssociatedDevice() Device {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_associated_device(_arg0)

	var _ret Device // out

	_ret = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _ret
}

func (d device) Axes() AxisFlags {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkAxisFlags // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_axes(_arg0)

	var _axisFlags AxisFlags // out

	_axisFlags = AxisFlags(_cret)

	return _axisFlags
}

func (d device) AxisUse(index_ uint) AxisUse {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint      // out
	var _cret C.GdkAxisUse // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_axis_use(_arg0, _arg1)

	var _axisUse AxisUse // out

	_axisUse = AxisUse(_cret)

	return _axisUse
}

func (d device) DeviceType() DeviceType {
	var _arg0 *C.GdkDevice    // out
	var _cret C.GdkDeviceType // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_device_type(_arg0)

	var _deviceType DeviceType // out

	_deviceType = DeviceType(_cret)

	return _deviceType
}

func (d device) Display() Display {
	var _arg0 *C.GdkDevice  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (d device) HasCursor() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_has_cursor(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d device) Key(index_ uint) (uint, ModifierType, bool) {
	var _arg0 *C.GdkDevice      // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // in
	var _arg3 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_key(_arg0, _arg1, &_arg2, &_arg3)

	var _keyval uint            // out
	var _modifiers ModifierType // out
	var _ok bool                // out

	_keyval = uint(_arg2)
	_modifiers = ModifierType(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _modifiers, _ok
}

func (d device) LastEventWindow() Window {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_last_event_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (d device) Mode() InputMode {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkInputMode // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_mode(_arg0)

	var _inputMode InputMode // out

	_inputMode = InputMode(_cret)

	return _inputMode
}

func (d device) NAxes() int {
	var _arg0 *C.GdkDevice // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_n_axes(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d device) NKeys() int {
	var _arg0 *C.GdkDevice // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_n_keys(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d device) Name() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d device) Position() (screen Screen, x int, y int) {
	var _arg0 *C.GdkDevice // out
	var _arg1 *C.GdkScreen // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	C.gdk_device_get_position(_arg0, &_arg1, &_arg2, &_arg3)

	var _screen Screen // out
	var _x int         // out
	var _y int         // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1))).(Screen)
	_x = int(_arg2)
	_y = int(_arg3)

	return _screen, _x, _y
}

func (d device) PositionDouble() (screen Screen, x float64, y float64) {
	var _arg0 *C.GdkDevice // out
	var _arg1 *C.GdkScreen // in
	var _arg2 C.gdouble    // in
	var _arg3 C.gdouble    // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	C.gdk_device_get_position_double(_arg0, &_arg1, &_arg2, &_arg3)

	var _screen Screen // out
	var _x float64     // out
	var _y float64     // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1))).(Screen)
	_x = float64(_arg2)
	_y = float64(_arg3)

	return _screen, _x, _y
}

func (d device) ProductID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_product_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d device) Seat() Seat {
	var _arg0 *C.GdkDevice // out
	var _cret *C.GdkSeat   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_seat(_arg0)

	var _seat Seat // out

	_seat = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Seat)

	return _seat
}

func (d device) Source() InputSource {
	var _arg0 *C.GdkDevice     // out
	var _cret C.GdkInputSource // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_source(_arg0)

	var _inputSource InputSource // out

	_inputSource = InputSource(_cret)

	return _inputSource
}

func (d device) VendorID() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_vendor_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d device) WindowAtPosition() (winX int, winY int, window Window) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.gint       // in
	var _arg2 C.gint       // in
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_window_at_position(_arg0, &_arg1, &_arg2)

	var _winX int      // out
	var _winY int      // out
	var _window Window // out

	_winX = int(_arg1)
	_winY = int(_arg2)
	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _winX, _winY, _window
}

func (d device) WindowAtPositionDouble() (winX float64, winY float64, window Window) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.gdouble    // in
	var _arg2 C.gdouble    // in
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_device_get_window_at_position_double(_arg0, &_arg1, &_arg2)

	var _winX float64  // out
	var _winY float64  // out
	var _window Window // out

	_winX = float64(_arg1)
	_winY = float64(_arg2)
	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _winX, _winY, _window
}

func (d device) Grab(window Window, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor Cursor, time_ uint32) GrabStatus {
	var _arg0 *C.GdkDevice       // out
	var _arg1 *C.GdkWindow       // out
	var _arg2 C.GdkGrabOwnership // out
	var _arg3 C.gboolean         // out
	var _arg4 C.GdkEventMask     // out
	var _arg5 *C.GdkCursor       // out
	var _arg6 C.guint32          // out
	var _cret C.GdkGrabStatus    // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GdkGrabOwnership(grabOwnership)
	if ownerEvents {
		_arg3 = C.TRUE
	}
	_arg4 = C.GdkEventMask(eventMask)
	_arg5 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))
	_arg6 = C.guint32(time_)

	_cret = C.gdk_device_grab(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _grabStatus GrabStatus // out

	_grabStatus = GrabStatus(_cret)

	return _grabStatus
}

func (d device) SetAxisUse(index_ uint, use AxisUse) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint      // out
	var _arg2 C.GdkAxisUse // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(index_)
	_arg2 = C.GdkAxisUse(use)

	C.gdk_device_set_axis_use(_arg0, _arg1, _arg2)
}

func (d device) SetKey(index_ uint, keyval uint, modifiers ModifierType) {
	var _arg0 *C.GdkDevice      // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(index_)
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	C.gdk_device_set_key(_arg0, _arg1, _arg2, _arg3)
}

func (d device) SetMode(mode InputMode) bool {
	var _arg0 *C.GdkDevice   // out
	var _arg1 C.GdkInputMode // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = C.GdkInputMode(mode)

	_cret = C.gdk_device_set_mode(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d device) Ungrab(time_ uint32) {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint32    // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint32(time_)

	C.gdk_device_ungrab(_arg0, _arg1)
}

func (d device) Warp(screen Screen, x int, y int) {
	var _arg0 *C.GdkDevice // out
	var _arg1 *C.GdkScreen // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gdk_device_warp(_arg0, _arg1, _arg2, _arg3)
}

// TimeCoord stores a single event in a motion history.
type TimeCoord struct {
	native C.GdkTimeCoord
}

// WrapTimeCoord wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTimeCoord(ptr unsafe.Pointer) *TimeCoord {
	return (*TimeCoord)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TimeCoord) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Time: the timestamp for this event.
func (t *TimeCoord) Time() uint32 {
	var v uint32 // out
	v = uint32(t.native.time)
	return v
}

// Axes: the values of the device’s axes.
func (t *TimeCoord) Axes() [128]float64 {
	var v [128]float64
	v = *(*[128]float64)(unsafe.Pointer(&t.native.axes))
	return v
}
