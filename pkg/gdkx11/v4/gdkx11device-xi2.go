// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_xi2_get_type()), F: marshalX11DeviceXI2er},
	})
}

type X11DeviceType int

const (
	X11DeviceTypeLogical X11DeviceType = iota
	X11DeviceTypePhysical
	X11DeviceTypeFloating
)

type X11DeviceXI2 struct {
	gdk.Device
}

var _ gextras.Nativer = (*X11DeviceXI2)(nil)

func wrapX11DeviceXI2(obj *externglib.Object) *X11DeviceXI2 {
	return &X11DeviceXI2{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalX11DeviceXI2er(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11DeviceXI2(obj), nil
}

func (*X11DeviceXI2) privateX11DeviceXI2() {}
