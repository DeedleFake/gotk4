// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// NewAttrStrikethroughColor: create a new strikethrough color attribute.
//
// This attribute modifies the color of strikethrough lines. If not set,
// strikethrough lines will use the foreground color.
//
// The function takes the following parameters:
//
//    - red value (ranging from 0 to 65535).
//    - green value.
//    - blue value.
//
// The function returns the following values:
//
//    - attribute: newly allocated PangoAttribute, which should be freed with
//      pango.Attribute.Destroy().
//
func NewAttrStrikethroughColor(red, green, blue uint16) *Attribute {
	var _arg1 C.guint16         // out
	var _arg2 C.guint16         // out
	var _arg3 C.guint16         // out
	var _cret *C.PangoAttribute // in

	_arg1 = C.guint16(red)
	_arg2 = C.guint16(green)
	_arg3 = C.guint16(blue)

	_cret = C.pango_attr_strikethrough_color_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(red)
	runtime.KeepAlive(green)
	runtime.KeepAlive(blue)

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attribute)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attribute_destroy((*C.PangoAttribute)(intern.C))
		},
	)

	return _attribute
}

// NewAttrUnderlineColor: create a new underline color attribute.
//
// This attribute modifies the color of underlines. If not set, underlines will
// use the foreground color.
//
// The function takes the following parameters:
//
//    - red value (ranging from 0 to 65535).
//    - green value.
//    - blue value.
//
// The function returns the following values:
//
//    - attribute: newly allocated PangoAttribute, which should be freed with
//      pango.Attribute.Destroy().
//
func NewAttrUnderlineColor(red, green, blue uint16) *Attribute {
	var _arg1 C.guint16         // out
	var _arg2 C.guint16         // out
	var _arg3 C.guint16         // out
	var _cret *C.PangoAttribute // in

	_arg1 = C.guint16(red)
	_arg2 = C.guint16(green)
	_arg3 = C.guint16(blue)

	_cret = C.pango_attr_underline_color_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(red)
	runtime.KeepAlive(green)
	runtime.KeepAlive(blue)

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attribute)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attribute_destroy((*C.PangoAttribute)(intern.C))
		},
	)

	return _attribute
}

// NewAttrSizeAbsolute: create a new font-size attribute in device units.
//
// The function takes the following parameters:
//
//    - size: font size, in PANGO_SCALEths of a device unit.
//
// The function returns the following values:
//
//    - attribute: newly allocated PangoAttribute, which should be freed with
//      pango.Attribute.Destroy().
//
func NewAttrSizeAbsolute(size int) *Attribute {
	var _arg1 C.int             // out
	var _cret *C.PangoAttribute // in

	_arg1 = C.int(size)

	_cret = C.pango_attr_size_new_absolute(_arg1)
	runtime.KeepAlive(size)

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attribute)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attribute_destroy((*C.PangoAttribute)(intern.C))
		},
	)

	return _attribute
}
