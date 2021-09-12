// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gio.h"
import "C"

// DBusMessageBytesNeeded: utility function to calculate how many bytes are
// needed to completely deserialize the D-Bus message stored at blob.
func DBusMessageBytesNeeded(blob []byte) (int, error) {
	var _arg1 *C.guchar // out
	var _arg2 C.gsize
	var _cret C.gssize  // in
	var _cerr *C.GError // in

	_arg2 = (C.gsize)(len(blob))
	if len(blob) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&blob[0]))
	}

	_cret = C.g_dbus_message_bytes_needed(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(blob)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}
