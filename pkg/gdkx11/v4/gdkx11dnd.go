// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gdkx11.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_drag_get_type()), F: marshalX11Dragger},
	})
}

type X11Drag struct {
	gdk.Drag
}

func wrapX11Drag(obj *externglib.Object) *X11Drag {
	return &X11Drag{
		Drag: gdk.Drag{
			Object: obj,
		},
	}
}

func marshalX11Dragger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Drag(obj), nil
}

func (*X11Drag) privateX11Drag() {}
