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
		{T: externglib.Type(C.gtk_boolean_cell_accessible_get_type()), F: marshalBooleanCellAccessibler},
	})
}

// BooleanCellAccessibler describes BooleanCellAccessible's methods.
type BooleanCellAccessibler interface {
	gextras.Objector

	privateBooleanCellAccessible()
}

type BooleanCellAccessible struct {
	*externglib.Object

	RendererCellAccessible
	atk.Action
	atk.Component
}

var _ BooleanCellAccessibler = (*BooleanCellAccessible)(nil)

func wrapBooleanCellAccessibler(obj *externglib.Object) BooleanCellAccessibler {
	return &BooleanCellAccessible{
		Object: obj,
		RendererCellAccessible: RendererCellAccessible{
			Object: obj,
			CellAccessible: CellAccessible{
				Object: obj,
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Action: atk.Action{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
			},
			Action: atk.Action{
				Object: obj,
			},
			Component: atk.Component{
				Object: obj,
			},
		},
		Action: atk.Action{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalBooleanCellAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBooleanCellAccessibler(obj), nil
}

func (*BooleanCellAccessible) privateBooleanCellAccessible() {}
