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
	GTypeScrolledWindowAccessible = coreglib.Type(C.gtk_scrolled_window_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScrolledWindowAccessible, F: marshalScrolledWindowAccessible},
	})
}

// ScrolledWindowAccessibleOverrider contains methods that are overridable.
type ScrolledWindowAccessibleOverrider interface {
}

type ScrolledWindowAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*ScrolledWindowAccessible)(nil)
)

func classInitScrolledWindowAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapScrolledWindowAccessible(obj *coreglib.Object) *ScrolledWindowAccessible {
	return &ScrolledWindowAccessible{
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

func marshalScrolledWindowAccessible(p uintptr) (interface{}, error) {
	return wrapScrolledWindowAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
