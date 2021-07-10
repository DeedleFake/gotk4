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
		{T: externglib.Type(C.gtk_menu_item_accessible_get_type()), F: marshalMenuItemAccessible},
	})
}

type MenuItemAccessible interface {
	gextras.Objector

	privateMenuItemAccessibleClass()
}

// MenuItemAccessibleClass implements the MenuItemAccessible interface.
type MenuItemAccessibleClass struct {
	ContainerAccessibleClass
	atk.ActionIface
}

var _ MenuItemAccessible = (*MenuItemAccessibleClass)(nil)

func wrapMenuItemAccessible(obj *externglib.Object) MenuItemAccessible {
	return &MenuItemAccessibleClass{
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
	}
}

func marshalMenuItemAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuItemAccessible(obj), nil
}

func (*MenuItemAccessibleClass) privateMenuItemAccessibleClass() {}
