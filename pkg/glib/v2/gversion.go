// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	_ "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "glib.h"
import "C"

// CheckVersion checks that the GLib library in use is compatible with the given
// version. Generally you would pass in the constants IB_MAJOR_VERSION,
// IB_MINOR_VERSION, IB_MICRO_VERSION as the three arguments to this function;
// that produces a check that the library in use is compatible with the version
// of GLib the application or module was compiled against.
//
// Compatibility is defined by two things: first the version of the running
// library is newer than the version
// required_major.required_minor.required_micro. Second the running library must
// be binary compatible with the version
// required_major.required_minor.required_micro (same major version.)
func CheckVersion(requiredMajor uint, requiredMinor uint, requiredMicro uint) string {
	var _arg1 C.guint  // out
	var _arg2 C.guint  // out
	var _arg3 C.guint  // out
	var _cret *C.gchar // in

	_arg1 = C.guint(requiredMajor)
	_arg2 = C.guint(requiredMinor)
	_arg3 = C.guint(requiredMicro)

	_cret = C.glib_check_version(_arg1, _arg2, _arg3)
	runtime.KeepAlive(requiredMajor)
	runtime.KeepAlive(requiredMinor)
	runtime.KeepAlive(requiredMicro)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
