// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeListBoxAccessible = coreglib.Type(C.gtk_list_box_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeListBoxAccessible, F: marshalListBoxAccessible},
	})
}

// ListBoxAccessibleOverrider contains methods that are overridable.
type ListBoxAccessibleOverrider interface {
}

type ListBoxAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Selection
}

var (
	_ coreglib.Objector = (*ListBoxAccessible)(nil)
)

func initClassListBoxAccessible(gclass unsafe.Pointer, goval any) {
}

func wrapListBoxAccessible(obj *coreglib.Object) *ListBoxAccessible {
	return &ListBoxAccessible{
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
	}
}

func marshalListBoxAccessible(p uintptr) (interface{}, error) {
	return wrapListBoxAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
