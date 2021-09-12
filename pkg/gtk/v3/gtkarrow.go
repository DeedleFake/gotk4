// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_arrow_get_type()), F: marshalArrower},
	})
}

// Arrow should be used to draw simple arrows that need to point in one of the
// four cardinal directions (up, down, left, or right). The style of the arrow
// can be one of shadow in, shadow out, etched in, or etched out. Note that
// these directions and style types may be amended in versions of GTK+ to come.
//
// GtkArrow will fill any space alloted to it, but since it is inherited from
// Misc, it can be padded and/or aligned, to fill exactly the space the
// programmer desires.
//
// Arrows are created with a call to gtk_arrow_new(). The direction or style of
// an arrow can be changed after creation by using gtk_arrow_set().
//
// GtkArrow has been deprecated; you can simply use a Image with a suitable icon
// name, such as “pan-down-symbolic“. When replacing GtkArrow by an image, pay
// attention to the fact that GtkArrow is doing automatic flipping between
// K_ARROW_LEFT and K_ARROW_RIGHT, depending on the text direction. To get the
// same effect with an image, use the icon names “pan-start-symbolic“ and
// “pan-end-symbolic“, which react to the text direction.
type Arrow struct {
	Misc
}

func wrapArrow(obj *externglib.Object) *Arrow {
	return &Arrow{
		Misc: Misc{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalArrower(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapArrow(obj), nil
}

// NewArrow creates a new Arrow widget.
//
// Deprecated: Use a Image with a suitable icon.
func NewArrow(arrowType ArrowType, shadowType ShadowType) *Arrow {
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out
	var _cret *C.GtkWidget    // in

	_arg1 = C.GtkArrowType(arrowType)
	_arg2 = C.GtkShadowType(shadowType)

	_cret = C.gtk_arrow_new(_arg1, _arg2)
	runtime.KeepAlive(arrowType)
	runtime.KeepAlive(shadowType)

	var _arrow *Arrow // out

	_arrow = wrapArrow(externglib.Take(unsafe.Pointer(_cret)))

	return _arrow
}

// Set sets the direction and style of the Arrow, arrow.
//
// Deprecated: Use a Image with a suitable icon.
func (arrow *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	var _arg0 *C.GtkArrow     // out
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out

	_arg0 = (*C.GtkArrow)(unsafe.Pointer(arrow.Native()))
	_arg1 = C.GtkArrowType(arrowType)
	_arg2 = C.GtkShadowType(shadowType)

	C.gtk_arrow_set(_arg0, _arg1, _arg2)
	runtime.KeepAlive(arrow)
	runtime.KeepAlive(arrowType)
	runtime.KeepAlive(shadowType)
}
