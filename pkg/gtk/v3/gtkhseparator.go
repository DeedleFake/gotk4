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
		{T: externglib.Type(C.gtk_hseparator_get_type()), F: marshalHSeparator},
	})
}

// HSeparator: the HSeparator widget is a horizontal separator, used to group
// the widgets within a window. It displays a horizontal line with a shadow to
// make it appear sunken into the interface.
//
// > The HSeparator widget is not used as a separator within menus. > To create
// a separator in a menu create an empty SeparatorMenuItem > widget using
// gtk_separator_menu_item_new() and add it to the menu with >
// gtk_menu_shell_append().
//
// GtkHSeparator has been deprecated, use Separator instead.
type HSeparator interface {
	gextras.Objector

	privateHSeparatorClass()
}

// HSeparatorClass implements the HSeparator interface.
type HSeparatorClass struct {
	*externglib.Object
	SeparatorClass
	BuildableIface
	OrientableIface
}

var _ HSeparator = (*HSeparatorClass)(nil)

func wrapHSeparator(obj *externglib.Object) HSeparator {
	return &HSeparatorClass{
		Object: obj,
		SeparatorClass: SeparatorClass{
			Object: obj,
			WidgetClass: WidgetClass{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
			OrientableIface: OrientableIface{
				Object: obj,
			},
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		OrientableIface: OrientableIface{
			Object: obj,
		},
	}
}

func marshalHSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHSeparator(obj), nil
}

// NewHSeparator creates a new HSeparator.
//
// Deprecated: Use gtk_separator_new() with GTK_ORIENTATION_HORIZONTAL instead.
func NewHSeparator() *HSeparatorClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hseparator_new()

	var _hSeparator *HSeparatorClass // out

	_hSeparator = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*HSeparatorClass)

	return _hSeparator
}

func (*HSeparatorClass) privateHSeparatorClass() {}
