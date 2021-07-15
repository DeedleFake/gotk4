// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.g_charset_converter_get_type()), F: marshalCharsetConverterer},
	})
}

// CharsetConverter is an implementation of #GConverter based on GIConv.
type CharsetConverter struct {
	*externglib.Object

	Converter
	Initable
}

var _ gextras.Nativer = (*CharsetConverter)(nil)

func wrapCharsetConverter(obj *externglib.Object) *CharsetConverter {
	return &CharsetConverter{
		Object: obj,
		Converter: Converter{
			Object: obj,
		},
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalCharsetConverterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCharsetConverter(obj), nil
}

// NewCharsetConverter creates a new Converter.
func NewCharsetConverter(toCharset string, fromCharset string) (*CharsetConverter, error) {
	var _arg1 *C.gchar             // out
	var _arg2 *C.gchar             // out
	var _cret *C.GCharsetConverter // in
	var _cerr *C.GError            // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(toCharset)))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(fromCharset)))

	_cret = C.g_charset_converter_new(_arg1, _arg2, &_cerr)

	var _charsetConverter *CharsetConverter // out
	var _goerr error                        // out

	_charsetConverter = wrapCharsetConverter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _charsetConverter, _goerr
}

// NumFallbacks gets the number of fallbacks that converter has applied so far.
func (converter *CharsetConverter) NumFallbacks() uint {
	var _arg0 *C.GCharsetConverter // out
	var _cret C.guint              // in

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(converter.Native()))

	_cret = C.g_charset_converter_get_num_fallbacks(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// UseFallback gets the Converter:use-fallback property.
func (converter *CharsetConverter) UseFallback() bool {
	var _arg0 *C.GCharsetConverter // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(converter.Native()))

	_cret = C.g_charset_converter_get_use_fallback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetUseFallback sets the Converter:use-fallback property.
func (converter *CharsetConverter) SetUseFallback(useFallback bool) {
	var _arg0 *C.GCharsetConverter // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GCharsetConverter)(unsafe.Pointer(converter.Native()))
	if useFallback {
		_arg1 = C.TRUE
	}

	C.g_charset_converter_set_use_fallback(_arg0, _arg1)
}
