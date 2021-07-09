// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_gl_renderer_get_type()), F: marshalGLRenderer},
	})
}

// GLRenderer: GSK renderer that is using OpenGL.
type GLRenderer interface {
	gextras.Objector

	privateGLRendererClass()
}

// GLRendererClass implements the GLRenderer interface.
type GLRendererClass struct {
	RendererClass
}

var _ GLRenderer = (*GLRendererClass)(nil)

func wrapGLRenderer(obj *externglib.Object) GLRenderer {
	return &GLRendererClass{
		RendererClass: RendererClass{
			Object: obj,
		},
	}
}

func marshalGLRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGLRenderer(obj), nil
}

// NewGLRenderer creates a new Renderer using OpenGL. This is the default
// renderer used by GTK.
func NewGLRenderer() *GLRendererClass {
	var _cret *C.GskRenderer // in

	_cret = C.gsk_gl_renderer_new()

	var _glRenderer *GLRendererClass // out

	_glRenderer = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*GLRendererClass)

	return _glRenderer
}

func (*GLRendererClass) privateGLRendererClass() {}
