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
		{T: externglib.Type(C.g_filter_input_stream_get_type()), F: marshalFilterInputStream},
	})
}

// FilterInputStream: base class for input stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterInputStream interface {
	gextras.Objector

	// BaseStream gets the base stream for the filter stream.
	BaseStream() *InputStreamClass
	// CloseBaseStream returns whether the base stream will be closed when
	// @stream is closed.
	CloseBaseStream() bool
	// SetCloseBaseStream sets whether the base stream will be closed when
	// @stream is closed.
	SetCloseBaseStream(closeBase bool)
}

// FilterInputStreamClass implements the FilterInputStream interface.
type FilterInputStreamClass struct {
	InputStreamClass
}

var _ FilterInputStream = (*FilterInputStreamClass)(nil)

func wrapFilterInputStream(obj *externglib.Object) FilterInputStream {
	return &FilterInputStreamClass{
		InputStreamClass: InputStreamClass{
			Object: obj,
		},
	}
}

func marshalFilterInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFilterInputStream(obj), nil
}

// BaseStream gets the base stream for the filter stream.
func (s *FilterInputStreamClass) BaseStream() *InputStreamClass {
	var _arg0 *C.GFilterInputStream // out
	var _cret *C.GInputStream       // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer((&FilterInputStream).Native()))

	_cret = C.g_filter_input_stream_get_base_stream(_arg0)

	var _inputStream *InputStreamClass // out

	_inputStream = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*InputStreamClass)

	return _inputStream
}

// CloseBaseStream returns whether the base stream will be closed when @stream
// is closed.
func (s *FilterInputStreamClass) CloseBaseStream() bool {
	var _arg0 *C.GFilterInputStream // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer((&FilterInputStream).Native()))

	_cret = C.g_filter_input_stream_get_close_base_stream(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when @stream
// is closed.
func (s *FilterInputStreamClass) SetCloseBaseStream(closeBase bool) {
	var _arg0 *C.GFilterInputStream // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer((&FilterInputStream).Native()))
	if closeBase {
		_arg1 = C.TRUE
	}

	C.g_filter_input_stream_set_close_base_stream(_arg0, _arg1)
}
