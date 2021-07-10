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
		{T: externglib.Type(C.gtk_check_menu_item_accessible_get_type()), F: marshalCheckMenuItemAccessible},
	})
}

type CheckMenuItemAccessible interface {
	gextras.Objector

	privateCheckMenuItemAccessibleClass()
}

// CheckMenuItemAccessibleClass implements the CheckMenuItemAccessible interface.
type CheckMenuItemAccessibleClass struct {
	MenuItemAccessibleClass
	atk.ActionIface
}

var _ CheckMenuItemAccessible = (*CheckMenuItemAccessibleClass)(nil)

func wrapCheckMenuItemAccessible(obj *externglib.Object) CheckMenuItemAccessible {
	return &CheckMenuItemAccessibleClass{
		MenuItemAccessibleClass: MenuItemAccessibleClass{
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
		},
		ActionIface: atk.ActionIface{
			Object: obj,
		},
	}
}

func marshalCheckMenuItemAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCheckMenuItemAccessible(obj), nil
}

func (*CheckMenuItemAccessibleClass) privateCheckMenuItemAccessibleClass() {}
