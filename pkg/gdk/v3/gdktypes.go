// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

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
		{T: externglib.Type(C.gdk_axis_use_get_type()), F: marshalAxisUse},
		{T: externglib.Type(C.gdk_byte_order_get_type()), F: marshalByteOrder},
		{T: externglib.Type(C.gdk_gl_error_get_type()), F: marshalGLError},
		{T: externglib.Type(C.gdk_grab_ownership_get_type()), F: marshalGrabOwnership},
		{T: externglib.Type(C.gdk_grab_status_get_type()), F: marshalGrabStatus},
		{T: externglib.Type(C.gdk_modifier_intent_get_type()), F: marshalModifierIntent},
		{T: externglib.Type(C.gdk_window_type_hint_get_type()), F: marshalWindowTypeHint},
		{T: externglib.Type(C.gdk_axis_flags_get_type()), F: marshalAxisFlags},
		{T: externglib.Type(C.gdk_event_mask_get_type()), F: marshalEventMask},
		{T: externglib.Type(C.gdk_modifier_type_get_type()), F: marshalModifierType},
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// AxisUse: enumeration describing the way in which a device axis (valuator)
// maps onto the predefined valuator types that GTK+ understands.
//
// Note that the X and Y axes are not really needed; pointer devices report
// their location via the x/y members of events regardless. Whether X and Y are
// present as axes depends on the GDK backend.
type AxisUse int

const (
	// Ignore: the axis is ignored.
	AxisUseIgnore AxisUse = iota
	// X: the axis is used as the x axis.
	AxisUseX
	// Y: the axis is used as the y axis.
	AxisUseY
	// Pressure: the axis is used for pressure information.
	AxisUsePressure
	// Xtilt: the axis is used for x tilt information.
	AxisUseXtilt
	// Ytilt: the axis is used for y tilt information.
	AxisUseYtilt
	// Wheel: the axis is used for wheel information.
	AxisUseWheel
	// Distance: the axis is used for pen/tablet distance information. (Since:
	// 3.22)
	AxisUseDistance
	// Rotation: the axis is used for pen rotation information. (Since: 3.22)
	AxisUseRotation
	// Slider: the axis is used for pen slider information. (Since: 3.22)
	AxisUseSlider
	// Last: constant equal to the numerically highest axis value.
	AxisUseLast
)

func marshalAxisUse(p uintptr) (interface{}, error) {
	return AxisUse(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ByteOrder: set of values describing the possible byte-orders for storing
// pixel values in memory.
type ByteOrder int

const (
	// LsbFirst: the values are stored with the least-significant byte first.
	// For instance, the 32-bit value 0xffeecc would be stored in memory as
	// 0xcc, 0xee, 0xff, 0x00.
	ByteOrderLsbFirst ByteOrder = iota
	// MsbFirst: the values are stored with the most-significant byte first. For
	// instance, the 32-bit value 0xffeecc would be stored in memory as 0x00,
	// 0xff, 0xee, 0xcc.
	ByteOrderMsbFirst
)

func marshalByteOrder(p uintptr) (interface{}, error) {
	return ByteOrder(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GLError: error enumeration for GLContext.
type GLError int

const (
	// NotAvailable: openGL support is not available
	GLErrorNotAvailable GLError = iota
	// UnsupportedFormat: the requested visual format is not supported
	GLErrorUnsupportedFormat
	// UnsupportedProfile: the requested profile is not supported
	GLErrorUnsupportedProfile
)

func marshalGLError(p uintptr) (interface{}, error) {
	return GLError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GrabOwnership defines how device grabs interact with other devices.
type GrabOwnership int

const (
	// None: all other devices’ events are allowed.
	GrabOwnershipNone GrabOwnership = iota
	// Window: other devices’ events are blocked for the grab window.
	GrabOwnershipWindow
	// Application: other devices’ events are blocked for the whole application.
	GrabOwnershipApplication
)

func marshalGrabOwnership(p uintptr) (interface{}, error) {
	return GrabOwnership(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GrabStatus: returned by gdk_device_grab(), gdk_pointer_grab() and
// gdk_keyboard_grab() to indicate success or the reason for the failure of the
// grab attempt.
type GrabStatus int

const (
	// Success: the resource was successfully grabbed.
	GrabStatusSuccess GrabStatus = iota
	// AlreadyGrabbed: the resource is actively grabbed by another client.
	GrabStatusAlreadyGrabbed
	// InvalidTime: the resource was grabbed more recently than the specified
	// time.
	GrabStatusInvalidTime
	// NotViewable: the grab window or the @confine_to window are not viewable.
	GrabStatusNotViewable
	// Frozen: the resource is frozen by an active grab of another client.
	GrabStatusFrozen
	// Failed: the grab failed for some other reason. Since 3.16
	GrabStatusFailed
)

func marshalGrabStatus(p uintptr) (interface{}, error) {
	return GrabStatus(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ModifierIntent: this enum is used with gdk_keymap_get_modifier_mask() in
// order to determine what modifiers the currently used windowing system backend
// uses for particular purposes. For example, on X11/Windows, the Control key is
// used for invoking menu shortcuts (accelerators), whereas on Apple computers
// it’s the Command key (which correspond to GDK_CONTROL_MASK and GDK_MOD2_MASK,
// respectively).
type ModifierIntent int

const (
	// PrimaryAccelerator: the primary modifier used to invoke menu
	// accelerators.
	ModifierIntentPrimaryAccelerator ModifierIntent = iota
	// ContextMenu: the modifier used to invoke context menus. Note that mouse
	// button 3 always triggers context menus. When this modifier is not 0, it
	// additionally triggers context menus when used with mouse button 1.
	ModifierIntentContextMenu
	// ExtendSelection: the modifier used to extend selections using
	// `modifier`-click or `modifier`-cursor-key
	ModifierIntentExtendSelection
	// ModifySelection: the modifier used to modify selections, which in most
	// cases means toggling the clicked item into or out of the selection.
	ModifierIntentModifySelection
	// NoTextInput: when any of these modifiers is pressed, the key event cannot
	// produce a symbol directly. This is meant to be used for input methods,
	// and for use cases like typeahead search.
	ModifierIntentNoTextInput
	// ShiftGroup: the modifier that switches between keyboard groups (AltGr on
	// X11/Windows and Option/Alt on OS X).
	ModifierIntentShiftGroup
	// DefaultModMask: the set of modifier masks accepted as modifiers in
	// accelerators. Needed because Command is mapped to MOD2 on OSX, which is
	// widely used, but on X11 MOD2 is NumLock and using that for a mod key is
	// problematic at best. Ref:
	// https://bugzilla.gnome.org/show_bug.cgi?id=736125.
	ModifierIntentDefaultModMask
)

func marshalModifierIntent(p uintptr) (interface{}, error) {
	return ModifierIntent(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// WindowTypeHint: these are hints for the window manager that indicate what
// type of function the window has. The window manager can use this when
// determining decoration and behaviour of the window. The hint must be set
// before mapping the window.
//
// See the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec) specification for more details
// about window types.
type WindowTypeHint int

const (
	// Normal: normal toplevel window.
	WindowTypeHintNormal WindowTypeHint = iota
	// Dialog: dialog window.
	WindowTypeHintDialog
	// Menu: window used to implement a menu; GTK+ uses this hint only for
	// torn-off menus, see TearoffMenuItem.
	WindowTypeHintMenu
	// Toolbar: window used to implement toolbars.
	WindowTypeHintToolbar
	// Splashscreen: window used to display a splash screen during application
	// startup.
	WindowTypeHintSplashscreen
	// Utility: utility windows which are not detached toolbars or dialogs.
	WindowTypeHintUtility
	// Dock: used for creating dock or panel windows.
	WindowTypeHintDock
	// Desktop: used for creating the desktop background window.
	WindowTypeHintDesktop
	// DropdownMenu: menu that belongs to a menubar.
	WindowTypeHintDropdownMenu
	// PopupMenu: menu that does not belong to a menubar, e.g. a context menu.
	WindowTypeHintPopupMenu
	// Tooltip: tooltip.
	WindowTypeHintTooltip
	// Notification - typically a “bubble” that belongs to a status icon.
	WindowTypeHintNotification
	// Combo: popup from a combo box.
	WindowTypeHintCombo
	// Dnd: window that is used to implement a DND cursor.
	WindowTypeHintDnd
)

func marshalWindowTypeHint(p uintptr) (interface{}, error) {
	return WindowTypeHint(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AxisFlags flags describing the current capabilities of a device/tool.
type AxisFlags int

const (
	// AxisFlagsX: x axis is present
	AxisFlagsX AxisFlags = 0b10
	// AxisFlagsY: y axis is present
	AxisFlagsY AxisFlags = 0b100
	// AxisFlagsPressure: pressure axis is present
	AxisFlagsPressure AxisFlags = 0b1000
	// AxisFlagsXtilt: x tilt axis is present
	AxisFlagsXtilt AxisFlags = 0b10000
	// AxisFlagsYtilt: y tilt axis is present
	AxisFlagsYtilt AxisFlags = 0b100000
	// AxisFlagsWheel: wheel axis is present
	AxisFlagsWheel AxisFlags = 0b1000000
	// AxisFlagsDistance: distance axis is present
	AxisFlagsDistance AxisFlags = 0b10000000
	// AxisFlagsRotation z-axis rotation is present
	AxisFlagsRotation AxisFlags = 0b100000000
	// AxisFlagsSlider: slider axis is present
	AxisFlagsSlider AxisFlags = 0b1000000000
)

func marshalAxisFlags(p uintptr) (interface{}, error) {
	return AxisFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EventMask: set of bit-flags to indicate which events a window is to receive.
// Most of these masks map onto one or more of the EventType event types above.
//
// See the [input handling overview][chap-input-handling] for details of [event
// masks][event-masks] and [event propagation][event-propagation].
//
// GDK_POINTER_MOTION_HINT_MASK is deprecated. It is a special mask to reduce
// the number of GDK_MOTION_NOTIFY events received. When using
// GDK_POINTER_MOTION_HINT_MASK, fewer GDK_MOTION_NOTIFY events will be sent,
// some of which are marked as a hint (the is_hint member is true). To receive
// more motion events after a motion hint event, the application needs to asks
// for more, by calling gdk_event_request_motions().
//
// Since GTK 3.8, motion events are already compressed by default, independent
// of this mechanism. This compression can be disabled with
// gdk_window_set_event_compression(). See the documentation of that function
// for details.
//
// If GDK_TOUCH_MASK is enabled, the window will receive touch events from
// touch-enabled devices. Those will come as sequences of EventTouch with type
// GDK_TOUCH_UPDATE, enclosed by two events with type GDK_TOUCH_BEGIN and
// GDK_TOUCH_END (or GDK_TOUCH_CANCEL). gdk_event_get_event_sequence() returns
// the event sequence for these events, so different sequences may be
// distinguished.
type EventMask int

const (
	// EventMaskExposureMask: receive expose events
	EventMaskExposureMask EventMask = 0b10
	// EventMaskPointerMotionMask: receive all pointer motion events
	EventMaskPointerMotionMask EventMask = 0b100
	// EventMaskPointerMotionHintMask: deprecated. see the explanation above
	EventMaskPointerMotionHintMask EventMask = 0b1000
	// EventMaskButtonMotionMask: receive pointer motion events while any button
	// is pressed
	EventMaskButtonMotionMask EventMask = 0b10000
	// EventMaskButton1MotionMask: receive pointer motion events while 1 button
	// is pressed
	EventMaskButton1MotionMask EventMask = 0b100000
	// EventMaskButton2MotionMask: receive pointer motion events while 2 button
	// is pressed
	EventMaskButton2MotionMask EventMask = 0b1000000
	// EventMaskButton3MotionMask: receive pointer motion events while 3 button
	// is pressed
	EventMaskButton3MotionMask EventMask = 0b10000000
	// EventMaskButtonPressMask: receive button press events
	EventMaskButtonPressMask EventMask = 0b100000000
	// EventMaskButtonReleaseMask: receive button release events
	EventMaskButtonReleaseMask EventMask = 0b1000000000
	// EventMaskKeyPressMask: receive key press events
	EventMaskKeyPressMask EventMask = 0b10000000000
	// EventMaskKeyReleaseMask: receive key release events
	EventMaskKeyReleaseMask EventMask = 0b100000000000
	// EventMaskEnterNotifyMask: receive window enter events
	EventMaskEnterNotifyMask EventMask = 0b1000000000000
	// EventMaskLeaveNotifyMask: receive window leave events
	EventMaskLeaveNotifyMask EventMask = 0b10000000000000
	// EventMaskFocusChangeMask: receive focus change events
	EventMaskFocusChangeMask EventMask = 0b100000000000000
	// EventMaskStructureMask: receive events about window configuration change
	EventMaskStructureMask EventMask = 0b1000000000000000
	// EventMaskPropertyChangeMask: receive property change events
	EventMaskPropertyChangeMask EventMask = 0b10000000000000000
	// EventMaskVisibilityNotifyMask: receive visibility change events
	EventMaskVisibilityNotifyMask EventMask = 0b100000000000000000
	// EventMaskProximityInMask: receive proximity in events
	EventMaskProximityInMask EventMask = 0b1000000000000000000
	// EventMaskProximityOutMask: receive proximity out events
	EventMaskProximityOutMask EventMask = 0b10000000000000000000
	// EventMaskSubstructureMask: receive events about window configuration
	// changes of child windows
	EventMaskSubstructureMask EventMask = 0b100000000000000000000
	// EventMaskScrollMask: receive scroll events
	EventMaskScrollMask EventMask = 0b1000000000000000000000
	// EventMaskTouchMask: receive touch events. Since 3.4
	EventMaskTouchMask EventMask = 0b10000000000000000000000
	// EventMaskSmoothScrollMask: receive smooth scrolling events. Since 3.4
	EventMaskSmoothScrollMask EventMask = 0b100000000000000000000000
	// EventMaskTouchpadGestureMask: receive touchpad gesture events. Since 3.18
	EventMaskTouchpadGestureMask EventMask = 0b1000000000000000000000000
	// EventMaskTabletPadMask: receive tablet pad events. Since 3.22
	EventMaskTabletPadMask EventMask = 0b10000000000000000000000000
	// EventMaskAllEventsMask: the combination of all the above event masks.
	EventMaskAllEventsMask EventMask = 0b11111111111111111111111110
)

func marshalEventMask(p uintptr) (interface{}, error) {
	return EventMask(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ModifierType: set of bit-flags to indicate the state of modifier keys and
// mouse buttons in various event types. Typical modifier keys are Shift,
// Control, Meta, Super, Hyper, Alt, Compose, Apple, CapsLock or ShiftLock.
//
// Like the X Window System, GDK supports 8 modifier keys and 5 mouse buttons.
//
// Since 2.10, GDK recognizes which of the Meta, Super or Hyper keys are mapped
// to Mod2 - Mod5, and indicates this by setting GDK_SUPER_MASK, GDK_HYPER_MASK
// or GDK_META_MASK in the state field of key events.
//
// Note that GDK may add internal values to events which include reserved values
// such as GDK_MODIFIER_RESERVED_13_MASK. Your code should preserve and ignore
// them. You can use GDK_MODIFIER_MASK to remove all reserved values.
//
// Also note that the GDK X backend interprets button press events for button
// 4-7 as scroll events, so GDK_BUTTON4_MASK and GDK_BUTTON5_MASK will never be
// set.
type ModifierType int

const (
	// ModifierTypeShiftMask: the Shift key.
	ModifierTypeShiftMask ModifierType = 0b1
	// ModifierTypeLockMask: lock key (depending on the modifier mapping of the
	// X server this may either be CapsLock or ShiftLock).
	ModifierTypeLockMask ModifierType = 0b10
	// ModifierTypeControlMask: the Control key.
	ModifierTypeControlMask ModifierType = 0b100
	// ModifierTypeMod1Mask: the fourth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier, but
	// normally it is the Alt key).
	ModifierTypeMod1Mask ModifierType = 0b1000
	// ModifierTypeMod2Mask: the fifth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier).
	ModifierTypeMod2Mask ModifierType = 0b10000
	// ModifierTypeMod3Mask: the sixth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier).
	ModifierTypeMod3Mask ModifierType = 0b100000
	// ModifierTypeMod4Mask: the seventh modifier key (it depends on the
	// modifier mapping of the X server which key is interpreted as this
	// modifier).
	ModifierTypeMod4Mask ModifierType = 0b1000000
	// ModifierTypeMod5Mask: the eighth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier).
	ModifierTypeMod5Mask ModifierType = 0b10000000
	// ModifierTypeButton1Mask: the first mouse button.
	ModifierTypeButton1Mask ModifierType = 0b100000000
	// ModifierTypeButton2Mask: the second mouse button.
	ModifierTypeButton2Mask ModifierType = 0b1000000000
	// ModifierTypeButton3Mask: the third mouse button.
	ModifierTypeButton3Mask ModifierType = 0b10000000000
	// ModifierTypeButton4Mask: the fourth mouse button.
	ModifierTypeButton4Mask ModifierType = 0b100000000000
	// ModifierTypeButton5Mask: the fifth mouse button.
	ModifierTypeButton5Mask ModifierType = 0b1000000000000
	// ModifierTypeModifierReserved13Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved13Mask ModifierType = 0b10000000000000
	// ModifierTypeModifierReserved14Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved14Mask ModifierType = 0b100000000000000
	// ModifierTypeModifierReserved15Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved15Mask ModifierType = 0b1000000000000000
	// ModifierTypeModifierReserved16Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved16Mask ModifierType = 0b10000000000000000
	// ModifierTypeModifierReserved17Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved17Mask ModifierType = 0b100000000000000000
	// ModifierTypeModifierReserved18Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved18Mask ModifierType = 0b1000000000000000000
	// ModifierTypeModifierReserved19Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved19Mask ModifierType = 0b10000000000000000000
	// ModifierTypeModifierReserved20Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved20Mask ModifierType = 0b100000000000000000000
	// ModifierTypeModifierReserved21Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved21Mask ModifierType = 0b1000000000000000000000
	// ModifierTypeModifierReserved22Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved22Mask ModifierType = 0b10000000000000000000000
	// ModifierTypeModifierReserved23Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved23Mask ModifierType = 0b100000000000000000000000
	// ModifierTypeModifierReserved24Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved24Mask ModifierType = 0b1000000000000000000000000
	// ModifierTypeModifierReserved25Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved25Mask ModifierType = 0b10000000000000000000000000
	// ModifierTypeSuperMask: the Super modifier. Since 2.10
	ModifierTypeSuperMask ModifierType = 0b100000000000000000000000000
	// ModifierTypeHyperMask: the Hyper modifier. Since 2.10
	ModifierTypeHyperMask ModifierType = 0b1000000000000000000000000000
	// ModifierTypeMetaMask: the Meta modifier. Since 2.10
	ModifierTypeMetaMask ModifierType = 0b10000000000000000000000000000
	// ModifierTypeModifierReserved29Mask: reserved bit flag; do not use in your
	// own code
	ModifierTypeModifierReserved29Mask ModifierType = 0b100000000000000000000000000000
	// ModifierTypeReleaseMask: not used in GDK itself. GTK+ uses it to
	// differentiate between (keyval, modifiers) pairs from key press and
	// release events.
	ModifierTypeReleaseMask ModifierType = 0b1000000000000000000000000000000
	// ModifierTypeModifierMask: mask covering all modifier types.
	ModifierTypeModifierMask ModifierType = 0b1011100000000000001111111111111
)

func marshalModifierType(p uintptr) (interface{}, error) {
	return ModifierType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Point defines the x and y coordinates of a point.
type Point struct {
	native C.GdkPoint
}

// WrapPoint wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPoint(ptr unsafe.Pointer) *Point {
	return (*Point)(ptr)
}

// Native returns the underlying C source pointer.
func (p *Point) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Rectangle defines the position and size of a rectangle. It is identical to
// #cairo_rectangle_int_t.
type Rectangle struct {
	native C.GdkRectangle
}

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	return (*Rectangle)(ptr)
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Rectangle)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Equal checks if the two given rectangles are equal.
func (rect1 *Rectangle) Equal(rect2 *Rectangle) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(unsafe.Pointer(rect1))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect2))

	_cret = C.gdk_rectangle_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Intersect calculates the intersection of two rectangles. It is allowed for
// @dest to be the same as either @src1 or @src2. If the rectangles do not
// intersect, @dest’s width and height is set to 0 and its x and y values are
// undefined. If you are only interested in whether the rectangles intersect,
// but not in the intersecting area itself, pass nil for @dest.
func (src1 *Rectangle) Intersect(src2 *Rectangle) (Rectangle, bool) {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(unsafe.Pointer(src1))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2))

	_cret = C.gdk_rectangle_intersect(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out
	var _ok bool        // out

	_dest = *(*Rectangle)(unsafe.Pointer((&_arg2)))
	if _cret != 0 {
		_ok = true
	}

	return _dest, _ok
}

// Union calculates the union of two rectangles. The union of rectangles @src1
// and @src2 is the smallest rectangle which includes both @src1 and @src2
// within it. It is allowed for @dest to be the same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (src1 *Rectangle) Union(src2 *Rectangle) Rectangle {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in

	_arg0 = (*C.GdkRectangle)(unsafe.Pointer(src1))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2))

	C.gdk_rectangle_union(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out

	_dest = *(*Rectangle)(unsafe.Pointer((&_arg2)))

	return _dest
}
