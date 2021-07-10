// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
		{T: externglib.Type(C.gtk_button_accessible_get_type()), F: marshalButtonAccessible},
	})
}

type ButtonAccessible interface {
	gextras.Objector

	privateButtonAccessibleClass()
}

// ButtonAccessibleClass implements the ButtonAccessible interface.
type ButtonAccessibleClass struct {
	*externglib.Object
	ContainerAccessibleClass
	atk.ActionIface
	atk.ImageIface
}

var _ ButtonAccessible = (*ButtonAccessibleClass)(nil)

func wrapButtonAccessible(obj *externglib.Object) ButtonAccessible {
	return &ButtonAccessibleClass{
		Object: obj,
		ContainerAccessibleClass: ContainerAccessibleClass{
			WidgetAccessibleClass: WidgetAccessibleClass{
				AccessibleClass: AccessibleClass{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
		ActionIface: atk.ActionIface{
			Object: obj,
		},
		ImageIface: atk.ImageIface{
			Object: obj,
		},
	}
}

func marshalButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapButtonAccessible(obj), nil
}

func (*ButtonAccessibleClass) privateButtonAccessibleClass() {}
