// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "glib.h"
import "C"

// PatternMatchSimple matches a string against a pattern given as a string. If
// this function is to be called in a loop, it's more efficient to compile the
// pattern once with g_pattern_spec_new() and call g_pattern_match_string()
// repeatedly.
func PatternMatchSimple(pattern string, _string string) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_pattern_match_simple(_arg1, _arg2)
	runtime.KeepAlive(pattern)
	runtime.KeepAlive(_string)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
