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
		{T: externglib.Type(C.g_file_monitor_get_type()), F: marshalFileMonitor},
	})
}

// FileMonitorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileMonitorOverrider interface {
	// Cancel cancels a file monitor.
	Cancel() bool
}

// FileMonitor monitors a file or directory for changes.
//
// To obtain a Monitor for a file or directory, use g_file_monitor(),
// g_file_monitor_file(), or g_file_monitor_directory().
//
// To get informed about changes to the file or directory you are monitoring,
// connect to the Monitor::changed signal. The signal will be emitted in the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that the monitor was created in (though if the global default main
// context is blocked, this may cause notifications to be blocked even if the
// thread-default context is still running).
type FileMonitor interface {
	gextras.Objector

	// Cancel cancels a file monitor.
	Cancel() bool
	// IsCancelled returns whether the monitor is canceled.
	IsCancelled() bool
	// SetRateLimit sets the rate limit to which the @monitor will report
	// consecutive change events to the same file.
	SetRateLimit(limitMsecs int)
}

// FileMonitorClass implements the FileMonitor interface.
type FileMonitorClass struct {
	*externglib.Object
}

var _ FileMonitor = (*FileMonitorClass)(nil)

func wrapFileMonitor(obj *externglib.Object) FileMonitor {
	return &FileMonitorClass{
		Object: obj,
	}
}

func marshalFileMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileMonitor(obj), nil
}

// Cancel cancels a file monitor.
func (monitor *FileMonitorClass) Cancel() bool {
	var _arg0 *C.GFileMonitor // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_file_monitor_cancel(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsCancelled returns whether the monitor is canceled.
func (monitor *FileMonitorClass) IsCancelled() bool {
	var _arg0 *C.GFileMonitor // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_file_monitor_is_cancelled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetRateLimit sets the rate limit to which the @monitor will report
// consecutive change events to the same file.
func (monitor *FileMonitorClass) SetRateLimit(limitMsecs int) {
	var _arg0 *C.GFileMonitor // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))
	_arg1 = C.gint(limitMsecs)

	C.g_file_monitor_set_rate_limit(_arg0, _arg1)
}
