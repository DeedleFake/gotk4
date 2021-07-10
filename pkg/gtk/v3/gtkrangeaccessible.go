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
		{T: externglib.Type(C.gtk_range_accessible_get_type()), F: marshalRangeAccessible},
	})
}

type RangeAccessible interface {
	gextras.Objector

	privateRangeAccessibleClass()
}

// RangeAccessibleClass implements the RangeAccessible interface.
type RangeAccessibleClass struct {
	WidgetAccessibleClass
}

var _ RangeAccessible = (*RangeAccessibleClass)(nil)

func wrapRangeAccessible(obj *externglib.Object) RangeAccessible {
	return &RangeAccessibleClass{
		WidgetAccessibleClass: WidgetAccessibleClass{
			AccessibleClass: AccessibleClass{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
		},
	}
}

func marshalRangeAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRangeAccessible(obj), nil
}

func (*RangeAccessibleClass) privateRangeAccessibleClass() {}
