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
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_view_accessible_get_type()), F: marshalTreeViewAccessibler},
	})
}

type TreeViewAccessible struct {
	ContainerAccessible

	atk.Selection
	atk.Table
	CellAccessibleParent
}

var _ gextras.Nativer = (*TreeViewAccessible)(nil)

func wrapTreeViewAccessible(obj *externglib.Object) *TreeViewAccessible {
	return &TreeViewAccessible{
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
		Table: atk.Table{
			Object: obj,
		},
		CellAccessibleParent: CellAccessibleParent{
			Object: obj,
		},
	}
}

func marshalTreeViewAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeViewAccessible(obj), nil
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *TreeViewAccessible) Native() uintptr {
	return v.ContainerAccessible.WidgetAccessible.Accessible.ObjectClass.Object.Native()
}

func (*TreeViewAccessible) privateTreeViewAccessible() {}
