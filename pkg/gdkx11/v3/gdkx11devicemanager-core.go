// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_manager_core_get_type()), F: marshalX11DeviceManagerCorer},
	})
}

type X11DeviceManagerCore struct {
	gdk.DeviceManager
}

func wrapX11DeviceManagerCore(obj *externglib.Object) *X11DeviceManagerCore {
	return &X11DeviceManagerCore{
		DeviceManager: gdk.DeviceManager{
			Object: obj,
		},
	}
}

func marshalX11DeviceManagerCorer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11DeviceManagerCore(obj), nil
}

func (*X11DeviceManagerCore) privateX11DeviceManagerCore() {}
