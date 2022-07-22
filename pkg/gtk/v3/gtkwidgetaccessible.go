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
	GTypeWidgetAccessible = coreglib.Type(C.gtk_widget_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWidgetAccessible, F: marshalWidgetAccessible},
	})
}

// WidgetAccessibleOverrider contains methods that are overridable.
type WidgetAccessibleOverrider interface {
}

type WidgetAccessible struct {
	_ [0]func() // equal guard
	Accessible

	atk.Component
}

var (
	_ coreglib.Objector = (*WidgetAccessible)(nil)
)

func initClassWidgetAccessible(gclass unsafe.Pointer, goval any) {
}

func wrapWidgetAccessible(obj *coreglib.Object) *WidgetAccessible {
	return &WidgetAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalWidgetAccessible(p uintptr) (interface{}, error) {
	return wrapWidgetAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
