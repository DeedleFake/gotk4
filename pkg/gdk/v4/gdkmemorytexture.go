// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_memory_format_get_type()), F: marshalMemoryFormat},
		{T: externglib.Type(C.gdk_memory_texture_get_type()), F: marshalMemoryTexture},
	})
}

// MemoryFormat: `GdkMemoryFormat` describes a format that bytes can have in
// memory.
//
// It describes formats by listing the contents of the memory passed to it. So
// GDK_MEMORY_A8R8G8B8 will be 1 byte (8 bits) of alpha, followed by a byte each
// of red, green and blue. It is not endian-dependent, so CAIRO_FORMAT_ARGB32 is
// represented by different `GdkMemoryFormats` on architectures with different
// endiannesses.
//
// Its naming is modelled after VkFormat (see
// https://www.khronos.org/registry/vulkan/specs/1.0/html/vkspec.htmlFormat for
// details).
type MemoryFormat int

const (
	// B8G8R8A8Premultiplied: 4 bytes; for blue, green, red, alpha. The color
	// values are premultiplied with the alpha value.
	MemoryFormatB8G8R8A8Premultiplied MemoryFormat = 0
	// A8R8G8B8Premultiplied: 4 bytes; for alpha, red, green, blue. The color
	// values are premultiplied with the alpha value.
	MemoryFormatA8R8G8B8Premultiplied MemoryFormat = 1
	// R8G8B8A8Premultiplied: 4 bytes; for red, green, blue, alpha The color
	// values are premultiplied with the alpha value.
	MemoryFormatR8G8B8A8Premultiplied MemoryFormat = 2
	// b8g8r8a8: 4 bytes; for blue, green, red, alpha.
	MemoryFormatB8G8R8A8 MemoryFormat = 3
	// a8r8g8b8: 4 bytes; for alpha, red, green, blue.
	MemoryFormatA8R8G8B8 MemoryFormat = 4
	// r8g8b8a8: 4 bytes; for red, green, blue, alpha.
	MemoryFormatR8G8B8A8 MemoryFormat = 5
	// a8b8g8r8: 4 bytes; for alpha, blue, green, red.
	MemoryFormatA8B8G8R8 MemoryFormat = 6
	// r8g8b8: 3 bytes; for red, green, blue. The data is opaque.
	MemoryFormatR8G8B8 MemoryFormat = 7
	// b8g8r8: 3 bytes; for blue, green, red. The data is opaque.
	MemoryFormatB8G8R8 MemoryFormat = 8
	// NFormats: the number of formats. This value will change as more formats
	// get added, so do not rely on its concrete integer.
	MemoryFormatNFormats MemoryFormat = 9
)

func marshalMemoryFormat(p uintptr) (interface{}, error) {
	return MemoryFormat(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// MemoryTexture: a `GdkTexture` representing image data in memory.
type MemoryTexture interface {
	Texture
}

// memoryTexture implements the MemoryTexture class.
type memoryTexture struct {
	Texture
}

// WrapMemoryTexture wraps a GObject to the right type. It is
// primarily used internally.
func WrapMemoryTexture(obj *externglib.Object) MemoryTexture {
	return memoryTexture{
		Texture: WrapTexture(obj),
	}
}

func marshalMemoryTexture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMemoryTexture(obj), nil
}

func (p memoryTexture) ComputeConcreteSize(specifiedWidth float64, specifiedHeight float64, defaultWidth float64, defaultHeight float64) (concreteWidth float64, concreteHeight float64) {
	return WrapPaintable(gextras.InternObject(p)).ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight)
}

func (p memoryTexture) CurrentImage() Paintable {
	return WrapPaintable(gextras.InternObject(p)).CurrentImage()
}

func (p memoryTexture) Flags() PaintableFlags {
	return WrapPaintable(gextras.InternObject(p)).Flags()
}

func (p memoryTexture) IntrinsicAspectRatio() float64 {
	return WrapPaintable(gextras.InternObject(p)).IntrinsicAspectRatio()
}

func (p memoryTexture) IntrinsicHeight() int {
	return WrapPaintable(gextras.InternObject(p)).IntrinsicHeight()
}

func (p memoryTexture) IntrinsicWidth() int {
	return WrapPaintable(gextras.InternObject(p)).IntrinsicWidth()
}

func (p memoryTexture) InvalidateContents() {
	WrapPaintable(gextras.InternObject(p)).InvalidateContents()
}

func (p memoryTexture) InvalidateSize() {
	WrapPaintable(gextras.InternObject(p)).InvalidateSize()
}

func (p memoryTexture) Snapshot(snapshot Snapshot, width float64, height float64) {
	WrapPaintable(gextras.InternObject(p)).Snapshot(snapshot, width, height)
}
