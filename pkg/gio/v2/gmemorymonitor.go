// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_memory_monitor_get_type()), F: marshalMemoryMonitorrer},
	})
}

// MEMORY_MONITOR_EXTENSION_POINT_NAME: extension point for memory usage
// monitoring functionality. See [Extending GIO][extending-gio].
const MEMORY_MONITOR_EXTENSION_POINT_NAME = "gio-memory-monitor"

// MemoryMonitorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MemoryMonitorOverrider interface {
	LowMemoryWarning(level MemoryMonitorWarningLevel)
}

// MemoryMonitor will monitor system memory and suggest to the application when
// to free memory so as to leave more room for other applications. It is
// implemented on Linux using the Low Memory Monitor
// (https://gitlab.freedesktop.org/hadess/low-memory-monitor/) (API
// documentation (https://hadess.pages.freedesktop.org/low-memory-monitor/)).
//
// There is also an implementation for use inside Flatpak sandboxes.
//
// Possible actions to take when the signal is received are:
//
// - Free caches
//
// - Save files that haven't been looked at in a while to disk, ready to be
// reopened when needed
//
// - Run a garbage collection cycle
//
// - Try and compress fragmented allocations
//
// - Exit on idle if the process has no reason to stay around
//
// - Call malloc_trim(3) (man:malloc_trim) to return cached heap pages to the
// kernel (if supported by your libc)
//
// Note that some actions may not always improve system performance, and so
// should be profiled for your application. malloc_trim(), for example, may make
// future heap allocations slower (due to releasing cached heap pages back to
// the kernel).
//
// See MonitorWarningLevel for details on the various warning levels.
//
//    static void
//    warning_cb (GMemoryMonitor *m, GMemoryMonitorWarningLevel level)
//    {
//      g_debug ("Warning level: d", level);
//      if (warning_level > G_MEMORY_MONITOR_WARNING_LEVEL_LOW)
//        drop_caches ();
//    }
//
//    static GMemoryMonitor *
//    monitor_low_memory (void)
//    {
//      GMemoryMonitor *m;
//      m = g_memory_monitor_dup_default ();
//      g_signal_connect (G_OBJECT (m), "low-memory-warning",
//                        G_CALLBACK (warning_cb), NULL);
//      return m;
//    }
//
// Don't forget to disconnect the Monitor::low-memory-warning signal, and unref
// the Monitor itself when exiting.
type MemoryMonitor struct {
	Initable
}

// MemoryMonitorrer describes MemoryMonitor's abstract methods.
type MemoryMonitorrer interface {
	externglib.Objector

	privateMemoryMonitor()
}

var _ MemoryMonitorrer = (*MemoryMonitor)(nil)

func wrapMemoryMonitor(obj *externglib.Object) *MemoryMonitor {
	return &MemoryMonitor{
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalMemoryMonitorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMemoryMonitor(obj), nil
}

func (*MemoryMonitor) privateMemoryMonitor() {}

// MemoryMonitorDupDefault gets a reference to the default Monitor for the
// system.
func MemoryMonitorDupDefault() MemoryMonitorrer {
	var _cret *C.GMemoryMonitor // in

	_cret = C.g_memory_monitor_dup_default()

	var _memoryMonitor MemoryMonitorrer // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(MemoryMonitorrer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.MemoryMonitorrer")
		}
		_memoryMonitor = rv
	}

	return _memoryMonitor
}
