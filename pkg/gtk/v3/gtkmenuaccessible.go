// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_accessible_get_type()), F: marshalMenuAccessibler},
	})
}

type MenuAccessible struct {
	MenuShellAccessible
}

func wrapMenuAccessible(obj *externglib.Object) *MenuAccessible {
	return &MenuAccessible{
		MenuShellAccessible: MenuShellAccessible{
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
			Selection: atk.Selection{
				Object: obj,
			},
		},
	}
}

func marshalMenuAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuAccessible(obj), nil
}

func (*MenuAccessible) privateMenuAccessible() {}
