// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_aspect_frame_get_type()), F: marshalAspectFramer},
	})
}

// AspectFrame: GtkAspectFrame preserves the aspect ratio of its child.
//
// The frame can respect the aspect ratio of the child widget, or use its own
// aspect ratio.
//
//
// CSS nodes
//
// GtkAspectFrame uses a CSS node with name frame.
type AspectFrame struct {
	Widget
}

func wrapAspectFrame(obj *externglib.Object) *AspectFrame {
	return &AspectFrame{
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
	}
}

func marshalAspectFramer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAspectFrame(obj), nil
}

// NewAspectFrame: create a new GtkAspectFrame.
func NewAspectFrame(xalign float32, yalign float32, ratio float32, obeyChild bool) *AspectFrame {
	var _arg1 C.float      // out
	var _arg2 C.float      // out
	var _arg3 C.float      // out
	var _arg4 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.float(xalign)
	_arg2 = C.float(yalign)
	_arg3 = C.float(ratio)
	if obeyChild {
		_arg4 = C.TRUE
	}

	_cret = C.gtk_aspect_frame_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
	runtime.KeepAlive(ratio)
	runtime.KeepAlive(obeyChild)

	var _aspectFrame *AspectFrame // out

	_aspectFrame = wrapAspectFrame(externglib.Take(unsafe.Pointer(_cret)))

	return _aspectFrame
}

// Child gets the child widget of self.
func (self *AspectFrame) Child() Widgetter {
	var _arg0 *C.GtkAspectFrame // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_aspect_frame_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// ObeyChild returns whether the child's size request should override the set
// aspect ratio of the GtkAspectFrame.
func (self *AspectFrame) ObeyChild() bool {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_aspect_frame_get_obey_child(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ratio returns the desired aspect ratio of the child.
func (self *AspectFrame) Ratio() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_aspect_frame_get_ratio(_arg0)
	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// XAlign returns the horizontal alignment of the child within the allocation of
// the GtkAspectFrame.
func (self *AspectFrame) XAlign() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_aspect_frame_get_xalign(_arg0)
	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// YAlign returns the vertical alignment of the child within the allocation of
// the GtkAspectFrame.
func (self *AspectFrame) YAlign() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_aspect_frame_get_yalign(_arg0)
	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// SetChild sets the child widget of self.
func (self *AspectFrame) SetChild(child Widgetter) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_aspect_frame_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetObeyChild sets whether the aspect ratio of the child's size request should
// override the set aspect ratio of the GtkAspectFrame.
func (self *AspectFrame) SetObeyChild(obeyChild bool) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))
	if obeyChild {
		_arg1 = C.TRUE
	}

	C.gtk_aspect_frame_set_obey_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(obeyChild)
}

// SetRatio sets the desired aspect ratio of the child.
func (self *AspectFrame) SetRatio(ratio float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))
	_arg1 = C.float(ratio)

	C.gtk_aspect_frame_set_ratio(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ratio)
}

// SetXAlign sets the horizontal alignment of the child within the allocation of
// the GtkAspectFrame.
func (self *AspectFrame) SetXAlign(xalign float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))
	_arg1 = C.float(xalign)

	C.gtk_aspect_frame_set_xalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(xalign)
}

// SetYAlign sets the vertical alignment of the child within the allocation of
// the GtkAspectFrame.
func (self *AspectFrame) SetYAlign(yalign float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(self.Native()))
	_arg1 = C.float(yalign)

	C.gtk_aspect_frame_set_yalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(yalign)
}
