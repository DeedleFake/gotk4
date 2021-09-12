// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_entry_accessible_get_type()), F: marshalEntryAccessibler},
	})
}

type EntryAccessible struct {
	WidgetAccessible

	atk.Action
	atk.EditableText
	atk.Text
	*externglib.Object
}

func wrapEntryAccessible(obj *externglib.Object) *EntryAccessible {
	return &EntryAccessible{
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
		Action: atk.Action{
			Object: obj,
		},
		EditableText: atk.EditableText{
			Object: obj,
		},
		Text: atk.Text{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalEntryAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEntryAccessible(obj), nil
}

func (*EntryAccessible) privateEntryAccessible() {}
