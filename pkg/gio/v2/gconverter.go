// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_converter_get_type()), F: marshalConverter},
	})
}

// ConverterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ConverterOverrider interface {
	// Reset resets all internal state in the converter, making it behave as if
	// it was just created. If the converter has any internal state that would
	// produce output then that output is lost.
	Reset()
}

// Converter is implemented by objects that convert binary data in various ways.
// The conversion can be stateful and may fail at any place.
//
// Some example conversions are: character set conversion, compression,
// decompression and regular expression replace.
type Converter interface {
	gextras.Objector

	// Reset resets all internal state in the converter, making it behave as if
	// it was just created. If the converter has any internal state that would
	// produce output then that output is lost.
	Reset()
}

// ConverterIface implements the Converter interface.
type ConverterIface struct {
	*externglib.Object
}

var _ Converter = (*ConverterIface)(nil)

func wrapConverter(obj *externglib.Object) Converter {
	return &ConverterIface{
		Object: obj,
	}
}

func marshalConverter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConverter(obj), nil
}

// Reset resets all internal state in the converter, making it behave as if it
// was just created. If the converter has any internal state that would produce
// output then that output is lost.
func (converter *ConverterIface) Reset() {
	var _arg0 *C.GConverter // out

	_arg0 = (*C.GConverter)(unsafe.Pointer(converter.Native()))

	C.g_converter_reset(_arg0)
}
