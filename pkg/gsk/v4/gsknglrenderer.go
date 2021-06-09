// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_ngl_renderer_get_type()), F: marshalNglRenderer},
	})
}

type NglRenderer interface {
	Renderer
}

// nglRenderer implements the NglRenderer interface.
type nglRenderer struct {
	Renderer
}

var _ NglRenderer = (*nglRenderer)(nil)

// WrapNglRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapNglRenderer(obj *externglib.Object) NglRenderer {
	return NglRenderer{
		Renderer: WrapRenderer(obj),
	}
}

func marshalNglRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNglRenderer(obj), nil
}

// NewNglRenderer constructs a class NglRenderer.
func NewNglRenderer() NglRenderer {
	var _cret C.GskNglRenderer

	cret = C.gsk_ngl_renderer_new()

	var _nglRenderer NglRenderer

	_nglRenderer = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(NglRenderer)

	return _nglRenderer
}
