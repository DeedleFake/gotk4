// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// AppInfoMonitorGet gets the InfoMonitor for the current thread-default main
// context.
//
// The InfoMonitor will emit a "changed" signal in the thread-default main
// context whenever the list of installed applications (as reported by
// g_app_info_get_all()) may have changed.
//
// You must only call g_object_unref() on the return value from under the same
// main context as you created it.
//
// The function returns the following values:
//
//    - appInfoMonitor: reference to a InfoMonitor.
//
func AppInfoMonitorGet() *AppInfoMonitor {
	var _cret *C.GAppInfoMonitor // in

	_cret = C.g_app_info_monitor_get()

	var _appInfoMonitor *AppInfoMonitor // out

	_appInfoMonitor = wrapAppInfoMonitor(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appInfoMonitor
}
