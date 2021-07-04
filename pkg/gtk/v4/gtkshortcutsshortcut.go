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

// ShortcutType: gtkShortcutType specifies the kind of shortcut that is being
// described.
//
// More values may be added to this enumeration over time.
type ShortcutType int

const (
	// accelerator: the shortcut is a keyboard accelerator. The
	// GtkShortcutsShortcut:accelerator property will be used.
	ShortcutTypeAccelerator ShortcutType = 0
	// GesturePinch: the shortcut is a pinch gesture. GTK provides an icon and
	// subtitle.
	ShortcutTypeGesturePinch ShortcutType = 1
	// GestureStretch: the shortcut is a stretch gesture. GTK provides an icon
	// and subtitle.
	ShortcutTypeGestureStretch ShortcutType = 2
	// GestureRotateClockwise: the shortcut is a clockwise rotation gesture. GTK
	// provides an icon and subtitle.
	ShortcutTypeGestureRotateClockwise ShortcutType = 3
	// GestureRotateCounterclockwise: the shortcut is a counterclockwise
	// rotation gesture. GTK provides an icon and subtitle.
	ShortcutTypeGestureRotateCounterclockwise ShortcutType = 4
	// GestureTwoFingerSwipeLeft: the shortcut is a two-finger swipe gesture.
	// GTK provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeLeft ShortcutType = 5
	// GestureTwoFingerSwipeRight: the shortcut is a two-finger swipe gesture.
	// GTK provides an icon and subtitle.
	ShortcutTypeGestureTwoFingerSwipeRight ShortcutType = 6
	// gesture: the shortcut is a gesture. The GtkShortcutsShortcut:icon
	// property will be used.
	ShortcutTypeGesture ShortcutType = 7
	// GestureSwipeLeft: the shortcut is a swipe gesture. GTK provides an icon
	// and subtitle.
	ShortcutTypeGestureSwipeLeft ShortcutType = 8
	// GestureSwipeRight: the shortcut is a swipe gesture. GTK provides an icon
	// and subtitle.
	ShortcutTypeGestureSwipeRight ShortcutType = 9
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShortcutsShortcut: a `GtkShortcutsShortcut` represents a single keyboard
// shortcut or gesture with a short text.
//
// This widget is only meant to be used with `GtkShortcutsWindow`.
type ShortcutsShortcut interface {
	Widget
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

func (s shortcutsShortcut) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s shortcutsShortcut) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s shortcutsShortcut) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s shortcutsShortcut) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s shortcutsShortcut) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s shortcutsShortcut) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s shortcutsShortcut) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b shortcutsShortcut) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
