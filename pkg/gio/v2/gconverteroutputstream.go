// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_converter_output_stream_get_type()), F: marshalConverterOutputStreamer},
	})
}

// ConverterOutputStream: converter output stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, OutputStream implements OutputStream.
type ConverterOutputStream struct {
	FilterOutputStream

	*externglib.Object
	OutputStream
	PollableOutputStream
}

var (
	_ FilterOutputStreamer = (*ConverterOutputStream)(nil)
	_ externglib.Objector  = (*ConverterOutputStream)(nil)
	_ OutputStreamer       = (*ConverterOutputStream)(nil)
)

func wrapConverterOutputStream(obj *externglib.Object) *ConverterOutputStream {
	return &ConverterOutputStream{
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Object: obj,
		OutputStream: OutputStream{
			Object: obj,
		},
		PollableOutputStream: PollableOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
	}
}

func marshalConverterOutputStreamer(p uintptr) (interface{}, error) {
	return wrapConverterOutputStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConverterOutputStream creates a new converter output stream for the
// base_stream.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//    - converter: #GConverter.
//
// The function returns the following values:
//
//    - converterOutputStream: new Stream.
//
func NewConverterOutputStream(baseStream OutputStreamer, converter Converterer) *ConverterOutputStream {
	var _arg1 *C.GOutputStream // out
	var _arg2 *C.GConverter    // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))
	_arg2 = (*C.GConverter)(unsafe.Pointer(converter.Native()))

	_cret = C.g_converter_output_stream_new(_arg1, _arg2)
	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(converter)

	var _converterOutputStream *ConverterOutputStream // out

	_converterOutputStream = wrapConverterOutputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _converterOutputStream
}

// Converter gets the #GConverter that is used by converter_stream.
//
// The function returns the following values:
//
//    - converter of the converter output stream.
//
func (converterStream *ConverterOutputStream) Converter() Converterer {
	var _arg0 *C.GConverterOutputStream // out
	var _cret *C.GConverter             // in

	_arg0 = (*C.GConverterOutputStream)(unsafe.Pointer(converterStream.Native()))

	_cret = C.g_converter_output_stream_get_converter(_arg0)
	runtime.KeepAlive(converterStream)

	var _converter Converterer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Converterer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Converterer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Converterer")
		}
		_converter = rv
	}

	return _converter
}
