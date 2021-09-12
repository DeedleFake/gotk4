// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_base_get_type()), F: marshalListBaser},
	})
}

// ListBase: GtkListBase is the abstract base class for GTK's list widgets.
type ListBase struct {
	Widget

	Orientable
	Scrollable
	*externglib.Object
}

// ListBaser describes ListBase's abstract methods.
type ListBaser interface {
	externglib.Objector

	privateListBase()
}

var _ ListBaser = (*ListBase)(nil)

func wrapListBase(obj *externglib.Object) *ListBase {
	return &ListBase{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		Scrollable: Scrollable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalListBaser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBase(obj), nil
}

func (*ListBase) privateListBase() {}
