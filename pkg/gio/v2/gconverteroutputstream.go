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
		{T: externglib.Type(C.g_converter_output_stream_get_type()), F: marshalConverterOutputStream},
	})
}

// ConverterOutputStream: converter output stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, OutputStream implements OutputStream.
type ConverterOutputStream interface {
	gextras.Objector

	// Converter gets the #GConverter that is used by @converter_stream.
	Converter() *ConverterInterface
}

// ConverterOutputStreamClass implements the ConverterOutputStream interface.
type ConverterOutputStreamClass struct {
	*externglib.Object
	FilterOutputStreamClass
	PollableOutputStreamInterface
}

var _ ConverterOutputStream = (*ConverterOutputStreamClass)(nil)

func wrapConverterOutputStream(obj *externglib.Object) ConverterOutputStream {
	return &ConverterOutputStreamClass{
		Object: obj,
		FilterOutputStreamClass: FilterOutputStreamClass{
			OutputStreamClass: OutputStreamClass{
				Object: obj,
			},
		},
		PollableOutputStreamInterface: PollableOutputStreamInterface{
			OutputStreamClass: OutputStreamClass{
				Object: obj,
			},
		},
	}
}

func marshalConverterOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConverterOutputStream(obj), nil
}

// NewConverterOutputStream creates a new converter output stream for the
// @base_stream.
func NewConverterOutputStream(baseStream OutputStream, converter Converter) *ConverterOutputStreamClass {
	var _arg1 *C.GOutputStream // out
	var _arg2 *C.GConverter    // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer((&OutputStream).Native()))
	_arg2 = (*C.GConverter)(unsafe.Pointer((&Converter).Native()))

	_cret = C.g_converter_output_stream_new(_arg1, _arg2)

	var _converterOutputStream *ConverterOutputStreamClass // out

	_converterOutputStream = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*ConverterOutputStreamClass)

	return _converterOutputStream
}

// Converter gets the #GConverter that is used by @converter_stream.
func (c *ConverterOutputStreamClass) Converter() *ConverterInterface {
	var _arg0 *C.GConverterOutputStream // out
	var _cret *C.GConverter             // in

	_arg0 = (*C.GConverterOutputStream)(unsafe.Pointer((&ConverterOutputStream).Native()))

	_cret = C.g_converter_output_stream_get_converter(_arg0)

	var _converter *ConverterInterface // out

	_converter = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ConverterInterface)

	return _converter
}
