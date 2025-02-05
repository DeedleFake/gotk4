// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeToolPaletteDragTargets = coreglib.Type(C.gtk_tool_palette_drag_targets_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeToolPaletteDragTargets, F: marshalToolPaletteDragTargets},
	})
}

// ToolPaletteDragTargets flags used to specify the supported drag targets.
type ToolPaletteDragTargets C.guint

const (
	// ToolPaletteDragItems: support drag of items.
	ToolPaletteDragItems ToolPaletteDragTargets = 0b1
	// ToolPaletteDragGroups: support drag of groups.
	ToolPaletteDragGroups ToolPaletteDragTargets = 0b10
)

func marshalToolPaletteDragTargets(p uintptr) (interface{}, error) {
	return ToolPaletteDragTargets(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ToolPaletteDragTargets.
func (t ToolPaletteDragTargets) String() string {
	if t == 0 {
		return "ToolPaletteDragTargets(0)"
	}

	var builder strings.Builder
	builder.Grow(42)

	for t != 0 {
		next := t & (t - 1)
		bit := t - next

		switch bit {
		case ToolPaletteDragItems:
			builder.WriteString("Items|")
		case ToolPaletteDragGroups:
			builder.WriteString("Groups|")
		default:
			builder.WriteString(fmt.Sprintf("ToolPaletteDragTargets(0b%b)|", bit))
		}

		t = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if t contains other.
func (t ToolPaletteDragTargets) Has(other ToolPaletteDragTargets) bool {
	return (t & other) == other
}

// ToolPaletteClass: instance of this type is always passed by reference.
type ToolPaletteClass struct {
	*toolPaletteClass
}

// toolPaletteClass is the struct that's finalized.
type toolPaletteClass struct {
	native *C.GtkToolPaletteClass
}

// ParentClass: parent class.
func (t *ToolPaletteClass) ParentClass() *ContainerClass {
	valptr := &t.native.parent_class
	var _v *ContainerClass // out
	_v = (*ContainerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
