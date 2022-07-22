// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gchar* _gotk4_gio2_TlsPasswordClass_get_default_warning(GTlsPassword*);
// extern guchar* _gotk4_gio2_TlsPasswordClass_get_value(GTlsPassword*, gsize*);
import "C"

// GType values.
var (
	GTypeTLSPassword = coreglib.Type(C.g_tls_password_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTLSPassword, F: marshalTLSPassword},
	})
}

// TLSPasswordOverrider contains methods that are overridable.
type TLSPasswordOverrider interface {
	// The function returns the following values:
	//
	DefaultWarning() string
	// Value: get the password value. If length is not NULL then it will be
	// filled in with the length of the password value. (Note that the password
	// value is not nul-terminated, so you can only pass NULL for length in
	// contexts where you know the password will have a certain fixed length.).
	//
	// The function takes the following parameters:
	//
	//    - length (optional): location to place the length of the password.
	//
	// The function returns the following values:
	//
	//    - guint8: password value (owned by the password object).
	//
	Value(length *uint) *byte
}

// TLSPassword holds a password used in TLS.
type TLSPassword struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TLSPassword)(nil)
)

func classInitTLSPassworder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GTlsPasswordClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ DefaultWarning() string }); ok {
		pclass.get_default_warning = (*[0]byte)(C._gotk4_gio2_TlsPasswordClass_get_default_warning)
	}

	if _, ok := goval.(interface{ Value(length *uint) *byte }); ok {
		pclass.get_value = (*[0]byte)(C._gotk4_gio2_TlsPasswordClass_get_value)
	}
}

//export _gotk4_gio2_TlsPasswordClass_get_default_warning
func _gotk4_gio2_TlsPasswordClass_get_default_warning(arg0 *C.GTlsPassword) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DefaultWarning() string })

	utf8 := iface.DefaultWarning()

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gio2_TlsPasswordClass_get_value
func _gotk4_gio2_TlsPasswordClass_get_value(arg0 *C.GTlsPassword, arg1 *C.gsize) (cret *C.guchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Value(length *uint) *byte })

	var _length *uint // out

	if arg1 != nil {
		_length = (*uint)(unsafe.Pointer(arg1))
	}

	guint8 := iface.Value(_length)

	cret = (*C.guchar)(unsafe.Pointer(guint8))

	return cret
}

func wrapTLSPassword(obj *coreglib.Object) *TLSPassword {
	return &TLSPassword{
		Object: obj,
	}
}

func marshalTLSPassword(p uintptr) (interface{}, error) {
	return wrapTLSPassword(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTLSPassword: create a new Password object.
//
// The function takes the following parameters:
//
//    - flags: password flags.
//    - description of what the password is for.
//
// The function returns the following values:
//
//    - tlsPassword: newly allocated password object.
//
func NewTLSPassword(flags TLSPasswordFlags, description string) *TLSPassword {
	var _arg1 C.GTlsPasswordFlags // out
	var _arg2 *C.gchar            // out
	var _cret *C.GTlsPassword     // in

	_arg1 = C.GTlsPasswordFlags(flags)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_password_new(_arg1, _arg2)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(description)

	var _tlsPassword *TLSPassword // out

	_tlsPassword = wrapTLSPassword(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _tlsPassword
}

// Description: get a description string about what the password will be used
// for.
//
// The function returns the following values:
//
//    - utf8: description of the password.
//
func (password *TLSPassword) Description() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))

	_cret = C.g_tls_password_get_description(_arg0)
	runtime.KeepAlive(password)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Flags: get flags about the password.
//
// The function returns the following values:
//
//    - tlsPasswordFlags flags about the password.
//
func (password *TLSPassword) Flags() TLSPasswordFlags {
	var _arg0 *C.GTlsPassword     // out
	var _cret C.GTlsPasswordFlags // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))

	_cret = C.g_tls_password_get_flags(_arg0)
	runtime.KeepAlive(password)

	var _tlsPasswordFlags TLSPasswordFlags // out

	_tlsPasswordFlags = TLSPasswordFlags(_cret)

	return _tlsPasswordFlags
}

// Value: get the password value. If length is not NULL then it will be filled
// in with the length of the password value. (Note that the password value is
// not nul-terminated, so you can only pass NULL for length in contexts where
// you know the password will have a certain fixed length.).
//
// The function takes the following parameters:
//
//    - length (optional): location to place the length of the password.
//
// The function returns the following values:
//
//    - guint8: password value (owned by the password object).
//
func (password *TLSPassword) Value(length *uint) *byte {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gsize        // out
	var _cret *C.guchar       // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	if length != nil {
		_arg1 = (*C.gsize)(unsafe.Pointer(length))
	}

	_cret = C.g_tls_password_get_value(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(length)

	var _guint8 *byte // out

	_guint8 = (*byte)(unsafe.Pointer(_cret))

	return _guint8
}

// Warning: get a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
//
// The function returns the following values:
//
//    - utf8: warning.
//
func (password *TLSPassword) Warning() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))

	_cret = C.g_tls_password_get_warning(_arg0)
	runtime.KeepAlive(password)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetDescription: set a description string about what the password will be used
// for.
//
// The function takes the following parameters:
//
//    - description of the password.
//
func (password *TLSPassword) SetDescription(description string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_description(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(description)
}

// SetFlags: set flags about the password.
//
// The function takes the following parameters:
//
//    - flags about the password.
//
func (password *TLSPassword) SetFlags(flags TLSPasswordFlags) {
	var _arg0 *C.GTlsPassword     // out
	var _arg1 C.GTlsPasswordFlags // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	_arg1 = C.GTlsPasswordFlags(flags)

	C.g_tls_password_set_flags(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(flags)
}

// SetValue: set the value for this password. The value will be copied by the
// password object.
//
// Specify the length, for a non-nul-terminated password. Pass -1 as length if
// using a nul-terminated password, and length will be calculated automatically.
// (Note that the terminating nul is not considered part of the password in this
// case.).
//
// The function takes the following parameters:
//
//    - value: new password value.
//
func (password *TLSPassword) SetValue(value []byte) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.guchar       // out
	var _arg2 C.gssize

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	_arg2 = (C.gssize)(len(value))
	if len(value) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&value[0]))
	}

	C.g_tls_password_set_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(password)
	runtime.KeepAlive(value)
}

// SetWarning: set a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
//
// The function takes the following parameters:
//
//    - warning: user readable warning.
//
func (password *TLSPassword) SetWarning(warning string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(coreglib.InternObject(password).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(warning)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_warning(_arg0, _arg1)
	runtime.KeepAlive(password)
	runtime.KeepAlive(warning)
}
