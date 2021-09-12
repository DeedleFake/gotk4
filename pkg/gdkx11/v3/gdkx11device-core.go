// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gdkx11.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_core_get_type()), F: marshalX11DeviceCorer},
	})
}

type X11DeviceCore struct {
	gdk.Device
}

func wrapX11DeviceCore(obj *externglib.Object) *X11DeviceCore {
	return &X11DeviceCore{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalX11DeviceCorer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11DeviceCore(obj), nil
}

func (*X11DeviceCore) privateX11DeviceCore() {}
