// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gio.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_filter_output_stream_get_type()), F: marshalFilterOutputStreamer},
	})
}

// FilterOutputStream: base class for output stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterOutputStream struct {
	OutputStream
}

// FilterOutputStreamer describes FilterOutputStream's abstract methods.
type FilterOutputStreamer interface {
	externglib.Objector

	// BaseStream gets the base stream for the filter stream.
	BaseStream() OutputStreamer
	// CloseBaseStream returns whether the base stream will be closed when
	// stream is closed.
	CloseBaseStream() bool
	// SetCloseBaseStream sets whether the base stream will be closed when
	// stream is closed.
	SetCloseBaseStream(closeBase bool)
}

var _ FilterOutputStreamer = (*FilterOutputStream)(nil)

func wrapFilterOutputStream(obj *externglib.Object) *FilterOutputStream {
	return &FilterOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
	}
}

func marshalFilterOutputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFilterOutputStream(obj), nil
}

// BaseStream gets the base stream for the filter stream.
func (stream *FilterOutputStream) BaseStream() OutputStreamer {
	var _arg0 *C.GFilterOutputStream // out
	var _cret *C.GOutputStream       // in

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_filter_output_stream_get_base_stream(_arg0)
	runtime.KeepAlive(stream)

	var _outputStream OutputStreamer // out

	_outputStream = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(OutputStreamer)

	return _outputStream
}

// CloseBaseStream returns whether the base stream will be closed when stream is
// closed.
func (stream *FilterOutputStream) CloseBaseStream() bool {
	var _arg0 *C.GFilterOutputStream // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_filter_output_stream_get_close_base_stream(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when stream is
// closed.
func (stream *FilterOutputStream) SetCloseBaseStream(closeBase bool) {
	var _arg0 *C.GFilterOutputStream // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(stream.Native()))
	if closeBase {
		_arg1 = C.TRUE
	}

	C.g_filter_output_stream_set_close_base_stream(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(closeBase)
}
