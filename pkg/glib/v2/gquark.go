// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// Quark is a non-zero integer which uniquely identifies a particular string. A
// GQuark value of zero is associated to NULL.
type Quark = uint32

// InternStaticString returns a canonical representation for string. Interned
// strings can be compared for equality by comparing the pointers, instead of
// using strcmp(). g_intern_static_string() does not copy the string, therefore
// string must not be freed or modified.
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func InternStaticString(_string string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_intern_static_string(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// InternString returns a canonical representation for string. Interned strings
// can be compared for equality by comparing the pointers, instead of using
// strcmp().
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func InternString(_string string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_intern_string(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// QuarkFromStaticString gets the #GQuark identifying the given (static) string.
// If the string does not currently have an associated #GQuark, a new #GQuark is
// created, linked to the given string.
//
// Note that this function is identical to g_quark_from_string() except that if
// a new #GQuark is created the string itself is used rather than a copy. This
// saves memory, but can only be used if the string will continue to exist until
// the program terminates. It can be used with statically allocated strings in
// the main program, but not with statically allocated memory in dynamically
// loaded modules, if you expect to ever unload the module again (e.g. do not
// use this function in GTK+ theme engines).
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func QuarkFromStaticString(_string string) Quark {
	var _arg1 *C.gchar // out
	var _cret C.GQuark // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_quark_from_static_string(_arg1)

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// QuarkFromString gets the #GQuark identifying the given string. If the string
// does not currently have an associated #GQuark, a new #GQuark is created,
// using a copy of the string.
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func QuarkFromString(_string string) Quark {
	var _arg1 *C.gchar // out
	var _cret C.GQuark // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_quark_from_string(_arg1)

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}

// QuarkToString gets the string associated with the given #GQuark.
func QuarkToString(quark Quark) string {
	var _arg1 C.GQuark // out
	var _cret *C.gchar // in

	_arg1 = C.guint32(quark)

	_cret = C.g_quark_to_string(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// QuarkTryString gets the #GQuark associated with the given string, or 0 if
// string is NULL or it has no associated #GQuark.
//
// If you want the GQuark to be created if it doesn't already exist, use
// g_quark_from_string() or g_quark_from_static_string().
//
// This function must not be used before library constructors have finished
// running.
func QuarkTryString(_string string) Quark {
	var _arg1 *C.gchar // out
	var _cret C.GQuark // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_quark_try_string(_arg1)

	var _quark Quark // out

	_quark = uint32(_cret)

	return _quark
}
