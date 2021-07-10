// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_aspect_frame_get_type()), F: marshalAspectFrame},
	})
}

// AspectFrame: the AspectFrame is useful when you want pack a widget so that it
// can resize but always retains the same aspect ratio. For instance, one might
// be drawing a small preview of a larger image. AspectFrame derives from Frame,
// so it can draw a label and a frame around the child. The frame will be
// “shrink-wrapped” to the size of the child.
//
//
// CSS nodes
//
// GtkAspectFrame uses a CSS node with name frame.
type AspectFrame interface {
	gextras.Objector

	// Set parameters for an existing AspectFrame.
	Set(xalign float32, yalign float32, ratio float32, obeyChild bool)
}

// AspectFrameClass implements the AspectFrame interface.
type AspectFrameClass struct {
	*externglib.Object
	FrameClass
	BuildableIface
}

var _ AspectFrame = (*AspectFrameClass)(nil)

func wrapAspectFrame(obj *externglib.Object) AspectFrame {
	return &AspectFrameClass{
		Object: obj,
		FrameClass: FrameClass{
			Object: obj,
			BinClass: BinClass{
				Object: obj,
				ContainerClass: ContainerClass{
					Object: obj,
					WidgetClass: WidgetClass{
						Object: obj,
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						BuildableIface: BuildableIface{
							Object: obj,
						},
					},
					BuildableIface: BuildableIface{
						Object: obj,
					},
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
	}
}

func marshalAspectFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAspectFrame(obj), nil
}

// NewAspectFrame: create a new AspectFrame.
func NewAspectFrame(label string, xalign float32, yalign float32, ratio float32, obeyChild bool) *AspectFrameClass {
	var _arg1 *C.gchar     // out
	var _arg2 C.gfloat     // out
	var _arg3 C.gfloat     // out
	var _arg4 C.gfloat     // out
	var _arg5 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gfloat(xalign)
	_arg3 = C.gfloat(yalign)
	_arg4 = C.gfloat(ratio)
	if obeyChild {
		_arg5 = C.TRUE
	}

	_cret = C.gtk_aspect_frame_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _aspectFrame *AspectFrameClass // out

	_aspectFrame = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*AspectFrameClass)

	return _aspectFrame
}

// Set parameters for an existing AspectFrame.
func (aspectFrame *AspectFrameClass) Set(xalign float32, yalign float32, ratio float32, obeyChild bool) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.gfloat          // out
	var _arg2 C.gfloat          // out
	var _arg3 C.gfloat          // out
	var _arg4 C.gboolean        // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(aspectFrame.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)
	_arg3 = C.gfloat(ratio)
	if obeyChild {
		_arg4 = C.TRUE
	}

	C.gtk_aspect_frame_set(_arg0, _arg1, _arg2, _arg3, _arg4)
}
