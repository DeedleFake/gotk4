// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_filter_input_stream_get_type()), F: marshalFilterInputStreamer},
	})
}

// FilterInputStream: base class for input stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterInputStream struct {
	InputStream
}

var _ gextras.Nativer = (*FilterInputStream)(nil)

// FilterInputStreamer describes FilterInputStream's abstract methods.
type FilterInputStreamer interface {
	// BaseStream gets the base stream for the filter stream.
	BaseStream() InputStreamer
	// CloseBaseStream returns whether the base stream will be closed when
	// stream is closed.
	CloseBaseStream() bool
	// SetCloseBaseStream sets whether the base stream will be closed when
	// stream is closed.
	SetCloseBaseStream(closeBase bool)
}

var _ FilterInputStreamer = (*FilterInputStream)(nil)

func wrapFilterInputStream(obj *externglib.Object) *FilterInputStream {
	return &FilterInputStream{
		InputStream: InputStream{
			Object: obj,
		},
	}
}

func marshalFilterInputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFilterInputStream(obj), nil
}

// BaseStream gets the base stream for the filter stream.
func (stream *FilterInputStream) BaseStream() InputStreamer {
	var _arg0 *C.GFilterInputStream // out
	var _cret *C.GInputStream       // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_filter_input_stream_get_base_stream(_arg0)

	var _inputStream InputStreamer // out

	_inputStream = (*gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(InputStreamer)

	return _inputStream
}

// CloseBaseStream returns whether the base stream will be closed when stream is
// closed.
func (stream *FilterInputStream) CloseBaseStream() bool {
	var _arg0 *C.GFilterInputStream // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_filter_input_stream_get_close_base_stream(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when stream is
// closed.
func (stream *FilterInputStream) SetCloseBaseStream(closeBase bool) {
	var _arg0 *C.GFilterInputStream // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(stream.Native()))
	if closeBase {
		_arg1 = C.TRUE
	}

	C.g_filter_input_stream_set_close_base_stream(_arg0, _arg1)
}
