// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_type_get_type()), F: marshalShortcutType},
		{T: externglib.Type(C.gtk_shortcuts_shortcut_get_type()), F: marshalShortcutsShortcuter},
	})
}

// ShortcutType specifies the kind of shortcut that is being described.
//
// More values may be added to this enumeration over time.
type ShortcutType int

const (
	// ShortcutAccelerator: shortcut is a keyboard accelerator. The
	// GtkShortcutsShortcut:accelerator property will be used.
	ShortcutAccelerator ShortcutType = iota
	// ShortcutGesturePinch: shortcut is a pinch gesture. GTK provides an icon
	// and subtitle.
	ShortcutGesturePinch
	// ShortcutGestureStretch: shortcut is a stretch gesture. GTK provides an
	// icon and subtitle.
	ShortcutGestureStretch
	// ShortcutGestureRotateClockwise: shortcut is a clockwise rotation gesture.
	// GTK provides an icon and subtitle.
	ShortcutGestureRotateClockwise
	// ShortcutGestureRotateCounterclockwise: shortcut is a counterclockwise
	// rotation gesture. GTK provides an icon and subtitle.
	ShortcutGestureRotateCounterclockwise
	// ShortcutGestureTwoFingerSwipeLeft: shortcut is a two-finger swipe
	// gesture. GTK provides an icon and subtitle.
	ShortcutGestureTwoFingerSwipeLeft
	// ShortcutGestureTwoFingerSwipeRight: shortcut is a two-finger swipe
	// gesture. GTK provides an icon and subtitle.
	ShortcutGestureTwoFingerSwipeRight
	// ShortcutGesture: shortcut is a gesture. The GtkShortcutsShortcut:icon
	// property will be used.
	ShortcutGesture
	// ShortcutGestureSwipeLeft: shortcut is a swipe gesture. GTK provides an
	// icon and subtitle.
	ShortcutGestureSwipeLeft
	// ShortcutGestureSwipeRight: shortcut is a swipe gesture. GTK provides an
	// icon and subtitle.
	ShortcutGestureSwipeRight
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ShortcutType.
func (s ShortcutType) String() string {
	switch s {
	case ShortcutAccelerator:
		return "Accelerator"
	case ShortcutGesturePinch:
		return "GesturePinch"
	case ShortcutGestureStretch:
		return "GestureStretch"
	case ShortcutGestureRotateClockwise:
		return "GestureRotateClockwise"
	case ShortcutGestureRotateCounterclockwise:
		return "GestureRotateCounterclockwise"
	case ShortcutGestureTwoFingerSwipeLeft:
		return "GestureTwoFingerSwipeLeft"
	case ShortcutGestureTwoFingerSwipeRight:
		return "GestureTwoFingerSwipeRight"
	case ShortcutGesture:
		return "Gesture"
	case ShortcutGestureSwipeLeft:
		return "GestureSwipeLeft"
	case ShortcutGestureSwipeRight:
		return "GestureSwipeRight"
	default:
		return fmt.Sprintf("ShortcutType(%d)", s)
	}
}

// ShortcutsShortcut: GtkShortcutsShortcut represents a single keyboard shortcut
// or gesture with a short text.
//
// This widget is only meant to be used with GtkShortcutsWindow.
type ShortcutsShortcut struct {
	Widget
}

var _ gextras.Nativer = (*ShortcutsShortcut)(nil)

func wrapShortcutsShortcut(obj *externglib.Object) *ShortcutsShortcut {
	return &ShortcutsShortcut{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsShortcuter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutsShortcut(obj), nil
}

func (*ShortcutsShortcut) privateShortcutsShortcut() {}
