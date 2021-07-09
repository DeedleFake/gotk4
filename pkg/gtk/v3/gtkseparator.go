// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_separator_get_type()), F: marshalSeparator},
	})
}

// Separator is a horizontal or vertical separator widget, depending on the
// value of the Orientable:orientation property, used to group the widgets
// within a window. It displays a line with a shadow to make it appear sunken
// into the interface.
//
//
// CSS nodes
//
// GtkSeparator has a single CSS node with name separator. The node gets one of
// the .horizontal or .vertical style classes.
type Separator interface {
	gextras.Objector

	privateSeparatorClass()
}

// SeparatorClass implements the Separator interface.
type SeparatorClass struct {
	*externglib.Object
	WidgetClass
	BuildableInterface
	OrientableInterface
}

var _ Separator = (*SeparatorClass)(nil)

func wrapSeparator(obj *externglib.Object) Separator {
	return &SeparatorClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSeparator(obj), nil
}

func (*SeparatorClass) privateSeparatorClass() {}
