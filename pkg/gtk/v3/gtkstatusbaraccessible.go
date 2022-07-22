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
	GTypeStatusbarAccessible = coreglib.Type(C.gtk_statusbar_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStatusbarAccessible, F: marshalStatusbarAccessible},
	})
}

// StatusbarAccessibleOverrider contains methods that are overridable.
type StatusbarAccessibleOverrider interface {
}

type StatusbarAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*StatusbarAccessible)(nil)
)

func initClassStatusbarAccessible(gclass unsafe.Pointer, goval any) {
}

func wrapStatusbarAccessible(obj *coreglib.Object) *StatusbarAccessible {
	return &StatusbarAccessible{
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
	}
}

func marshalStatusbarAccessible(p uintptr) (interface{}, error) {
	return wrapStatusbarAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
