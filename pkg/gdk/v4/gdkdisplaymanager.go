// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// SetAllowedBackends sets a list of backends that GDK should try to use.
//
// This can be useful if your application does not work with certain GDK
// backends.
//
// By default, GDK tries all included backends.
//
// For example:
//
// “`c gdk_set_allowed_backends ("wayland,macos,*"); “`
//
// instructs GDK to try the Wayland backend first, followed by the MacOs
// backend, and then all others.
//
// If the `GDK_BACKEND` environment variable is set, it determines what backends
// are tried in what order, while still respecting the set of allowed backends
// that are specified by this function.
//
// The possible backend names are:
//
//    - `broadway`
//    - `macos`
//    - `wayland`.
//    - `win32`
//    - `x11`
//
// You can also include a `*` in the list to try all remaining backends.
//
// This call must happen prior to functions that open a display, such as
// [func@Gdk.Display.open], `gtk_init()`, or `gtk_init_check()` in order to take
// effect.
func SetAllowedBackends(backends string) {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(backends))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_set_allowed_backends(_arg1)
}
