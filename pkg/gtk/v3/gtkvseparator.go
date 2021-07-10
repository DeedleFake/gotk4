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
		{T: externglib.Type(C.gtk_vseparator_get_type()), F: marshalVSeparator},
	})
}

// VSeparator: the VSeparator widget is a vertical separator, used to group the
// widgets within a window. It displays a vertical line with a shadow to make it
// appear sunken into the interface.
//
// GtkVSeparator has been deprecated, use Separator instead.
type VSeparator interface {
	gextras.Objector

	privateVSeparatorClass()
}

// VSeparatorClass implements the VSeparator interface.
type VSeparatorClass struct {
	*externglib.Object
	SeparatorClass
	BuildableIface
	OrientableIface
}

var _ VSeparator = (*VSeparatorClass)(nil)

func wrapVSeparator(obj *externglib.Object) VSeparator {
	return &VSeparatorClass{
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

func marshalVSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVSeparator(obj), nil
}

// NewVSeparator creates a new VSeparator.
//
// Deprecated: Use gtk_separator_new() with GTK_ORIENTATION_VERTICAL instead.
func NewVSeparator() *VSeparatorClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vseparator_new()

	var _vSeparator *VSeparatorClass // out

	_vSeparator = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*VSeparatorClass)

	return _vSeparator
}

func (*VSeparatorClass) privateVSeparatorClass() {}
