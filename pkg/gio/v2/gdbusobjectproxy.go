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
		{T: externglib.Type(C.g_dbus_object_proxy_get_type()), F: marshalDBusObjectProxy},
	})
}

// DBusObjectProxy is an object used to represent a remote object with one or
// more D-Bus interfaces. Normally, you don't instantiate a BusObjectProxy
// yourself - typically BusObjectManagerClient is used to obtain it.
type DBusObjectProxy interface {
	gextras.Objector

	// Connection gets the connection that @proxy is for.
	Connection() *DBusConnectionClass
}

// DBusObjectProxyClass implements the DBusObjectProxy interface.
type DBusObjectProxyClass struct {
	*externglib.Object
	DBusObjectIface
}

var _ DBusObjectProxy = (*DBusObjectProxyClass)(nil)

func wrapDBusObjectProxy(obj *externglib.Object) DBusObjectProxy {
	return &DBusObjectProxyClass{
		Object: obj,
		DBusObjectIface: DBusObjectIface{
			Object: obj,
		},
	}
}

func marshalDBusObjectProxy(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObjectProxy(obj), nil
}

// NewDBusObjectProxy creates a new BusObjectProxy for the given connection and
// object path.
func NewDBusObjectProxy(connection DBusConnection, objectPath string) *DBusObjectProxyClass {
	var _arg1 *C.GDBusConnection  // out
	var _arg2 *C.gchar            // out
	var _cret *C.GDBusObjectProxy // in

	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	_arg2 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dbus_object_proxy_new(_arg1, _arg2)

	var _dBusObjectProxy *DBusObjectProxyClass // out

	_dBusObjectProxy = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DBusObjectProxyClass)

	return _dBusObjectProxy
}

// Connection gets the connection that @proxy is for.
func (proxy *DBusObjectProxyClass) Connection() *DBusConnectionClass {
	var _arg0 *C.GDBusObjectProxy // out
	var _cret *C.GDBusConnection  // in

	_arg0 = (*C.GDBusObjectProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_object_proxy_get_connection(_arg0)

	var _dBusConnection *DBusConnectionClass // out

	_dBusConnection = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*DBusConnectionClass)

	return _dBusConnection
}
