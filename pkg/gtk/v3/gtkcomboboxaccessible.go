// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_combo_box_accessible_get_type()), F: marshalComboBoxAccessibler},
	})
}

type ComboBoxAccessible struct {
	ContainerAccessible

	atk.Action
	atk.Selection
	*externglib.Object
}

func wrapComboBoxAccessible(obj *externglib.Object) *ComboBoxAccessible {
	return &ComboBoxAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
		Action: atk.Action{
			Object: obj,
		},
		Selection: atk.Selection{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalComboBoxAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComboBoxAccessible(obj), nil
}

func (*ComboBoxAccessible) privateComboBoxAccessible() {}
