// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// SetAllowedBackends sets a list of backends that GDK should try to use.
//
// This can be be useful if your application does not work with certain GDK
// backends.
//
// By default, GDK tries all included backends.
//
// For example,
//
//    gdk_set_allowed_backends ("wayland,quartz,*");
//
// instructs GDK to try the Wayland backend first, followed by the Quartz
// backend, and then all others.
//
// If the GDK_BACKEND environment variable is set, it determines what backends
// are tried in what order, while still respecting the set of allowed backends
// that are specified by this function.
//
// The possible backend names are x11, win32, quartz, broadway, wayland. You can
// also include a * in the list to try all remaining backends.
//
// This call must happen prior to gdk_display_open(), gtk_init(),
// gtk_init_with_args() or gtk_init_check() in order to take effect.
//
// The function takes the following parameters:
//
//    - backends: comma-separated list of backends.
//
func SetAllowedBackends(backends string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(backends)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_set_allowed_backends(_arg1)
	runtime.KeepAlive(backends)
}
