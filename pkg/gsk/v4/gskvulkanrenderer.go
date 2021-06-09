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
		{T: externglib.Type(C.gsk_vulkan_renderer_get_type()), F: marshalVulkanRenderer},
	})
}

// VulkanRenderer: a GSK renderer that is using Vulkan.
type VulkanRenderer interface {
	Renderer
}

// vulkanRenderer implements the VulkanRenderer interface.
type vulkanRenderer struct {
	Renderer
}

var _ VulkanRenderer = (*vulkanRenderer)(nil)

// WrapVulkanRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapVulkanRenderer(obj *externglib.Object) VulkanRenderer {
	return VulkanRenderer{
		Renderer: WrapRenderer(obj),
	}
}

func marshalVulkanRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVulkanRenderer(obj), nil
}

// NewVulkanRenderer constructs a class VulkanRenderer.
func NewVulkanRenderer() VulkanRenderer {
	var _cret C.GskVulkanRenderer

	cret = C.gsk_vulkan_renderer_new()

	var _vulkanRenderer VulkanRenderer

	_vulkanRenderer = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(VulkanRenderer)

	return _vulkanRenderer
}
