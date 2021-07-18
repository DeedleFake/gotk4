// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_connectable_get_type()), F: marshalSocketConnectabler},
	})
}

// SocketConnectableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketConnectableOverrider interface {
	// Enumerate creates a AddressEnumerator for connectable.
	Enumerate() SocketAddressEnumeratorrer
	// ProxyEnumerate creates a AddressEnumerator for connectable that will
	// return a Address for each of its addresses that you must connect to via a
	// proxy.
	//
	// If connectable does not implement g_socket_connectable_proxy_enumerate(),
	// this will fall back to calling g_socket_connectable_enumerate().
	ProxyEnumerate() SocketAddressEnumeratorrer
	// String: format a Connectable as a string. This is a human-readable format
	// for use in debugging output, and is not a stable serialization format. It
	// is not suitable for use in user interfaces as it exposes too much
	// information for a user.
	//
	// If the Connectable implementation does not support string formatting, the
	// implementation’s type name will be returned as a fallback.
	String() string
}

// SocketConnectable objects that describe one or more potential socket
// endpoints implement Connectable. Callers can then use
// g_socket_connectable_enumerate() to get a AddressEnumerator to try out each
// socket address in turn until one succeeds, as shown in the sample code below.
//
//    MyConnectionType *
//    connect_to_host (const char    *hostname,
//                     guint16        port,
//                     GCancellable  *cancellable,
//                     GError       **error)
//    {
//      MyConnection *conn = NULL;
//      GSocketConnectable *addr;
//      GSocketAddressEnumerator *enumerator;
//      GSocketAddress *sockaddr;
//      GError *conn_error = NULL;
//
//      addr = g_network_address_new (hostname, port);
//      enumerator = g_socket_connectable_enumerate (addr);
//      g_object_unref (addr);
//
//      // Try each sockaddr until we succeed. Record the first connection error,
//      // but not any further ones (since they'll probably be basically the same
//      // as the first).
//      while (!conn && (sockaddr = g_socket_address_enumerator_next (enumerator, cancellable, error))
//        {
//          conn = connect_to_sockaddr (sockaddr, conn_error ? NULL : &conn_error);
//          g_object_unref (sockaddr);
//        }
//      g_object_unref (enumerator);
//
//      if (conn)
//        {
//          if (conn_error)
//            {
//              // We couldn't connect to the first address, but we succeeded
//              // in connecting to a later address.
//              g_error_free (conn_error);
//            }
//          return conn;
//        }
//      else if (error)
//        {
//          /// Either initial lookup failed, or else the caller cancelled us.
//          if (conn_error)
//            g_error_free (conn_error);
//          return NULL;
//        }
//      else
//        {
//          g_error_propagate (error, conn_error);
//          return NULL;
//        }
//    }
type SocketConnectable struct {
	*externglib.Object
}

var _ gextras.Nativer = (*SocketConnectable)(nil)

// SocketConnectabler describes SocketConnectable's abstract methods.
type SocketConnectabler interface {
	// Enumerate creates a AddressEnumerator for connectable.
	Enumerate() SocketAddressEnumeratorrer
	// ProxyEnumerate creates a AddressEnumerator for connectable that will
	// return a Address for each of its addresses that you must connect to via a
	// proxy.
	ProxyEnumerate() SocketAddressEnumeratorrer
	// String: format a Connectable as a string.
	String() string
}

var _ SocketConnectabler = (*SocketConnectable)(nil)

func wrapSocketConnectable(obj *externglib.Object) *SocketConnectable {
	return &SocketConnectable{
		Object: obj,
	}
}

func marshalSocketConnectabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketConnectable(obj), nil
}

// Enumerate creates a AddressEnumerator for connectable.
func (connectable *SocketConnectable) Enumerate() SocketAddressEnumeratorrer {
	var _arg0 *C.GSocketConnectable       // out
	var _cret *C.GSocketAddressEnumerator // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))

	_cret = C.g_socket_connectable_enumerate(_arg0)

	var _socketAddressEnumerator SocketAddressEnumeratorrer // out

	_socketAddressEnumerator = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(SocketAddressEnumeratorrer)

	return _socketAddressEnumerator
}

// ProxyEnumerate creates a AddressEnumerator for connectable that will return a
// Address for each of its addresses that you must connect to via a proxy.
//
// If connectable does not implement g_socket_connectable_proxy_enumerate(),
// this will fall back to calling g_socket_connectable_enumerate().
func (connectable *SocketConnectable) ProxyEnumerate() SocketAddressEnumeratorrer {
	var _arg0 *C.GSocketConnectable       // out
	var _cret *C.GSocketAddressEnumerator // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))

	_cret = C.g_socket_connectable_proxy_enumerate(_arg0)

	var _socketAddressEnumerator SocketAddressEnumeratorrer // out

	_socketAddressEnumerator = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(SocketAddressEnumeratorrer)

	return _socketAddressEnumerator
}

// String: format a Connectable as a string. This is a human-readable format for
// use in debugging output, and is not a stable serialization format. It is not
// suitable for use in user interfaces as it exposes too much information for a
// user.
//
// If the Connectable implementation does not support string formatting, the
// implementation’s type name will be returned as a fallback.
func (connectable *SocketConnectable) String() string {
	var _arg0 *C.GSocketConnectable // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))

	_cret = C.g_socket_connectable_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
