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
	MemoryFormatB8G8R8A8Premultiplied MemoryFormat = iota
	// A8R8G8B8Premultiplied: 4 bytes; for alpha, red, green, blue. The color
	// values are premultiplied with the alpha value.
	MemoryFormatA8R8G8B8Premultiplied
	// R8G8B8A8Premultiplied: 4 bytes; for red, green, blue, alpha The color
	// values are premultiplied with the alpha value.
	MemoryFormatR8G8B8A8Premultiplied
	// B8G8R8A8: 4 bytes; for blue, green, red, alpha.
	MemoryFormatB8G8R8A8
	// A8R8G8B8: 4 bytes; for alpha, red, green, blue.
	MemoryFormatA8R8G8B8
	// R8G8B8A8: 4 bytes; for red, green, blue, alpha.
	MemoryFormatR8G8B8A8
	// A8B8G8R8: 4 bytes; for alpha, blue, green, red.
	MemoryFormatA8B8G8R8
	// R8G8B8: 3 bytes; for red, green, blue. The data is opaque.
	MemoryFormatR8G8B8
	// B8G8R8: 3 bytes; for blue, green, red. The data is opaque.
	MemoryFormatB8G8R8
	// NFormats: the number of formats. This value will change as more formats
	// get added, so do not rely on its concrete integer.
	MemoryFormatNFormats
)

func marshalMemoryFormat(p uintptr) (interface{}, error) {
	return MemoryFormat(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// MemoryTexture: `GdkTexture` representing image data in memory.
type MemoryTexture interface {
	gextras.Objector

	privateMemoryTextureClass()
}

// MemoryTextureClass implements the MemoryTexture interface.
type MemoryTextureClass struct {
	*externglib.Object
	TextureClass
	PaintableIface
}

var _ MemoryTexture = (*MemoryTextureClass)(nil)

func wrapMemoryTexture(obj *externglib.Object) MemoryTexture {
	return &MemoryTextureClass{
		Object: obj,
		TextureClass: TextureClass{
			Object: obj,
			PaintableIface: PaintableIface{
				Object: obj,
			},
		},
		PaintableIface: PaintableIface{
			Object: obj,
		},
	}
}

func marshalMemoryTexture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMemoryTexture(obj), nil
}

func (*MemoryTextureClass) privateMemoryTextureClass() {}
