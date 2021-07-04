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
		{T: externglib.Type(C.g_tls_password_get_type()), F: marshalTLSPassword},
	})
}

// TLSPassword holds a password used in TLS.
type TLSPassword interface {
	gextras.Objector

	Description() string

	Flags() TLSPasswordFlags

	Value(length *uint) *byte

	Warning() string

	SetDescriptionTLSPassword(description string)

	SetFlagsTLSPassword(flags TLSPasswordFlags)

	SetValueTLSPassword(value []byte)

	SetWarningTLSPassword(warning string)
}

// tlsPassword implements the TLSPassword class.
type tlsPassword struct {
	gextras.Objector
}

// WrapTLSPassword wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSPassword(obj *externglib.Object) TLSPassword {
	return tlsPassword{
		Objector: obj,
	}
}

func marshalTLSPassword(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSPassword(obj), nil
}

func NewTLSPassword(flags TLSPasswordFlags, description string) TLSPassword {
	var _arg1 C.GTlsPasswordFlags // out
	var _arg2 *C.gchar            // out
	var _cret *C.GTlsPassword     // in

	_arg1 = C.GTlsPasswordFlags(flags)
	_arg2 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_password_new(_arg1, _arg2)

	var _tlsPassword TLSPassword // out

	_tlsPassword = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSPassword)

	return _tlsPassword
}

func (p tlsPassword) Description() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))

	_cret = C.g_tls_password_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p tlsPassword) Flags() TLSPasswordFlags {
	var _arg0 *C.GTlsPassword     // out
	var _cret C.GTlsPasswordFlags // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))

	_cret = C.g_tls_password_get_flags(_arg0)

	var _tlsPasswordFlags TLSPasswordFlags // out

	_tlsPasswordFlags = TLSPasswordFlags(_cret)

	return _tlsPasswordFlags
}

func (p tlsPassword) Value(length *uint) *byte {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gsize        // out
	var _cret *C.guchar       // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gsize)(unsafe.Pointer(length))

	_cret = C.g_tls_password_get_value(_arg0, _arg1)

	var _guint8 *byte // out

	_guint8 = (*byte)(unsafe.Pointer(_cret))

	return _guint8
}

func (p tlsPassword) Warning() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))

	_cret = C.g_tls_password_get_warning(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p tlsPassword) SetDescriptionTLSPassword(description string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_description(_arg0, _arg1)
}

func (p tlsPassword) SetFlagsTLSPassword(flags TLSPasswordFlags) {
	var _arg0 *C.GTlsPassword     // out
	var _arg1 C.GTlsPasswordFlags // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	_arg1 = C.GTlsPasswordFlags(flags)

	C.g_tls_password_set_flags(_arg0, _arg1)
}

func (p tlsPassword) SetValueTLSPassword(value []byte) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.guchar
	var _arg2 C.gssize

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	_arg2 = C.gssize(len(value))
	_arg1 = (*C.guchar)(unsafe.Pointer(&value[0]))

	C.g_tls_password_set_value(_arg0, _arg1, _arg2)
}

func (p tlsPassword) SetWarningTLSPassword(warning string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(warning))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_warning(_arg0, _arg1)
}
