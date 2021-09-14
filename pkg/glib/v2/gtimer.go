// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	_ "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib.h>
import "C"

// USEC_PER_SEC: number of microseconds in one second (1 million). This macro is
// provided for code readability.
const USEC_PER_SEC = 1000000

// Usleep pauses the current thread for the given number of microseconds.
//
// There are 1 million microseconds per second (represented by the USEC_PER_SEC
// macro). g_usleep() may have limited precision, depending on hardware and
// operating system; don't rely on the exact length of the sleep.
func Usleep(microseconds uint32) {
	var _arg1 C.gulong // out

	_arg1 = C.gulong(microseconds)

	C.g_usleep(_arg1)
	runtime.KeepAlive(microseconds)
}

// TimeValFromISO8601 converts a string containing an ISO 8601 encoded date and
// time to a Val and puts it into time_.
//
// iso_date must include year, month, day, hours, minutes, and seconds. It can
// optionally include fractions of a second and a time zone indicator. (In the
// absence of any time zone indication, the timestamp is assumed to be in local
// time.)
//
// Any leading or trailing space in iso_date is ignored.
//
// This function was deprecated, along with Val itself, in GLib 2.62. Equivalent
// functionality is available using code like:
//
//    GDateTime *dt = g_date_time_new_from_iso8601 (iso8601_string, NULL);
//    gint64 time_val = g_date_time_to_unix (dt);
//    g_date_time_unref (dt);
//
// Deprecated: Val is not year-2038-safe. Use g_date_time_new_from_iso8601()
// instead.
func TimeValFromISO8601(isoDate string) (TimeVal, bool) {
	var _arg1 *C.gchar   // out
	var _arg2 C.GTimeVal // in
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(isoDate)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_time_val_from_iso8601(_arg1, &_arg2)
	runtime.KeepAlive(isoDate)

	var _time_ TimeVal // out
	var _ok bool       // out

	_time_ = *(*TimeVal)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _time_, _ok
}
