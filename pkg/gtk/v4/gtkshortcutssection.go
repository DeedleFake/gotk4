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
		{T: externglib.Type(C.gtk_shortcuts_section_get_type()), F: marshalShortcutsSection},
	})
}

// ShortcutsSection: a `GtkShortcutsSection` collects all the keyboard shortcuts
// and gestures for a major application mode.
//
// If your application needs multiple sections, you should give each section a
// unique [property@Gtk.ShortcutsSection:section-name] and a
// [property@Gtk.ShortcutsSection:title] that can be shown in the section
// selector of the [class@Gtk.ShortcutsWindow].
//
// The [property@Gtk.ShortcutsSection:max-height] property can be used to
// influence how the groups in the section are distributed over pages and
// columns.
//
// This widget is only meant to be used with [class@Gtk.ShortcutsWindow].
type ShortcutsSection interface {
	Box
}

// shortcutsSection implements the ShortcutsSection class.
type shortcutsSection struct {
	Box
}

// WrapShortcutsSection wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutsSection(obj *externglib.Object) ShortcutsSection {
	return shortcutsSection{
		Box: WrapBox(obj),
	}
}

func marshalShortcutsSection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsSection(obj), nil
}

func (s shortcutsSection) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s shortcutsSection) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s shortcutsSection) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s shortcutsSection) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s shortcutsSection) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s shortcutsSection) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s shortcutsSection) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b shortcutsSection) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (o shortcutsSection) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o shortcutsSection) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
