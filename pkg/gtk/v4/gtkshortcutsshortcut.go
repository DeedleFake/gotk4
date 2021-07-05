// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_type_get_type()), F: marshalShortcutType},
		{T: externglib.Type(C.gtk_shortcuts_shortcut_get_type()), F: marshalShortcutsShortcut},
	})
}

// ShortcutType specifies the kind of shortcut that is being described.
//
// More values may be added to this enumeration over time.
type ShortcutType int

const (
	// accelerator: the shortcut is a keyboard accelerator. The
	// GtkShortcutsShortcut:accelerator property will be used.
	ShortcutTypeAccelerator ShortcutType = iota
	// GesturePinch: the shortcut is a pinch gesture. GTK provides an icon and
	// subtitle.
	ShortcutTypeGesturePinch
	// GestureStretch: the shortcut is a stretch gesture. GTK provides an icon
	// and subtitle.
	ShortcutTypeGestureStretch
	// GestureRotateClockwise: the shortcut is a clockwise rotation gesture. GTK
	// provides an icon and subtitle.
	ShortcutTypeGestureRotateClockwise
	// GestureRotateCounterclockwise: the shortcut is a counterclockwise
	// rotation gesture. GTK provides an icon and subtitle.
	ShortcutTypeGestureRotateCounterclockwise
	// GestureTwoFingerSwipeLeft: the shortcut is a two-finger swipe gesture.
	// GTK provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeLeft
	// GestureTwoFingerSwipeRight: the shortcut is a two-finger swipe gesture.
	// GTK provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeRight
	// gesture: the shortcut is a gesture. The GtkShortcutsShortcut:icon
	// property will be used.
	ShortcutTypeGesture
	// GestureSwipeLeft: the shortcut is a swipe gesture. GTK provides an icon
	// and subtitle.
	ShortcutTypeGestureSwipeLeft
	// GestureSwipeRight: the shortcut is a swipe gesture. GTK provides an icon
	// and subtitle.
	ShortcutTypeGestureSwipeRight
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShortcutsShortcut: `GtkShortcutsShortcut` represents a single keyboard
// shortcut or gesture with a short text.
//
// This widget is only meant to be used with `GtkShortcutsWindow`.
type ShortcutsShortcut interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
}

// shortcutsShortcut implements the ShortcutsShortcut class.
type shortcutsShortcut struct {
	Widget
}

// WrapShortcutsShortcut wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutsShortcut(obj *externglib.Object) ShortcutsShortcut {
	return shortcutsShortcut{
		Widget: WrapWidget(obj),
	}
}

func marshalShortcutsShortcut(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsShortcut(obj), nil
}

func (s shortcutsShortcut) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(s))
}

func (s shortcutsShortcut) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(s))
}

func (s shortcutsShortcut) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(s))
}
