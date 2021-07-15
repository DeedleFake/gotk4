// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_custom_layout_get_type()), F: marshalCustomLayouter},
	})
}

// CustomLayout: GtkCustomLayout uses closures for size negotiation.
//
// A GtkCustomLayout uses closures matching to the old GtkWidget virtual
// functions for size negotiation, as a convenience API to ease the porting
// towards the corresponding `GtkLayoutManager virtual functions.
type CustomLayout struct {
	LayoutManager
}

var _ gextras.Nativer = (*CustomLayout)(nil)

func wrapCustomLayout(obj *externglib.Object) *CustomLayout {
	return &CustomLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalCustomLayouter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCustomLayout(obj), nil
}

func (*CustomLayout) privateCustomLayout() {}
