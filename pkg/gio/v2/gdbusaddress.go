// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// DBusAddressEscapeValue: escape string so it can appear in a D-Bus address as
// the value part of a key-value pair.
//
// For instance, if string is /run/bus-for-:0, this function would return
// /run/bus-for-3A0, which could be used in a D-Bus address like
// unix:nonce-tcp:host=127.0.0.1,port=42,noncefile=/run/bus-for-3A0.
func DBusAddressEscapeValue(_string string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))

	_cret = C.g_dbus_address_escape_value(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusAddressGetForBusSync: synchronously looks up the D-Bus address for the
// well-known message bus instance specified by bus_type. This may involve using
// various platform specific mechanisms.
//
// The returned address will be in the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
func DBusAddressGetForBusSync(ctx context.Context, busType BusType) (string, error) {
	var _arg2 *C.GCancellable // out
	var _arg1 C.GBusType      // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)

	_cret = C.g_dbus_address_get_for_bus_sync(_arg1, _arg2, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

// DBusAddressGetStream: asynchronously connects to an endpoint specified by
// address and sets up the connection so it is in a state to run the client-side
// of the D-Bus authentication conversation. address must be in the D-Bus
// address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// When the operation is finished, callback will be invoked. You can then call
// g_dbus_address_get_stream_finish() to get the result of the operation.
//
// This is an asynchronous failable function. See
// g_dbus_address_get_stream_sync() for the synchronous version.
func DBusAddressGetStream(ctx context.Context, address string, callback AsyncReadyCallback) {
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(address)))
	_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg4 = C.gpointer(gbox.AssignOnce(callback))

	C.g_dbus_address_get_stream(_arg1, _arg2, _arg3, _arg4)
}

// DBusAddressGetStreamFinish finishes an operation started with
// g_dbus_address_get_stream().
//
// A server is not required to set a GUID, so out_guid may be set to NULL even
// on success.
func DBusAddressGetStreamFinish(res AsyncResulter) (string, IOStreamer, error) {
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.gchar        // in
	var _cret *C.GIOStream    // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((res).(gextras.Nativer).Native()))

	_cret = C.g_dbus_address_get_stream_finish(_arg1, &_arg2, &_cerr)

	var _outGuid string      // out
	var _ioStream IOStreamer // out
	var _goerr error         // out

	_outGuid = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_ioStream = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(IOStreamer)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _outGuid, _ioStream, _goerr
}

// DBusAddressGetStreamSync: synchronously connects to an endpoint specified by
// address and sets up the connection so it is in a state to run the client-side
// of the D-Bus authentication conversation. address must be in the D-Bus
// address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// A server is not required to set a GUID, so out_guid may be set to NULL even
// on success.
//
// This is a synchronous failable function. See g_dbus_address_get_stream() for
// the asynchronous version.
func DBusAddressGetStreamSync(ctx context.Context, address string) (string, IOStreamer, error) {
	var _arg3 *C.GCancellable // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // in
	var _cret *C.GIOStream    // in
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(address)))

	_cret = C.g_dbus_address_get_stream_sync(_arg1, &_arg2, _arg3, &_cerr)

	var _outGuid string      // out
	var _ioStream IOStreamer // out
	var _goerr error         // out

	_outGuid = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_ioStream = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(IOStreamer)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _outGuid, _ioStream, _goerr
}

// DBusIsAddress checks if string is a D-Bus address
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This doesn't check if string is actually supported by BusServer or
// BusConnection - use g_dbus_is_supported_address() to do more checks.
func DBusIsAddress(_string string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))

	_cret = C.g_dbus_is_address(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusIsSupportedAddress: like g_dbus_is_address() but also checks if the
// library supports the transports in string and that key/value pairs for each
// transport are valid. See the specification of the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
func DBusIsSupportedAddress(_string string) error {
	var _arg1 *C.gchar  // out
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))

	C.g_dbus_is_supported_address(_arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
