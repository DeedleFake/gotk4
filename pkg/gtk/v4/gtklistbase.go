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
		{T: externglib.Type(C.gtk_list_base_get_type()), F: marshalListBaser},
	})
}

// ListBase: GtkListBase is the abstract base class for GTK's list widgets.
type ListBase struct {
	Widget

	Orientable
	Scrollable
}

var _ gextras.Nativer = (*ListBase)(nil)

// ListBaser describes ListBase's abstract methods.
type ListBaser interface {
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
		},
		Orientable: Orientable{
			Object: obj,
		},
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalListBaser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBase(obj), nil
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ListBase) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

func (*ListBase) privateListBase() {}
