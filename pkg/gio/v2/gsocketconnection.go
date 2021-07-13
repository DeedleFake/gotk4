// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
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
// #include <glib-object.h>
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_connection_get_type()), F: marshalSocketConnectioner},
	})
}

// SocketConnectioner describes SocketConnection's methods.
type SocketConnectioner interface {
	// ConnectSocketConnectioner: connect connection to the specified remote
	// address.
	ConnectSocketConnectioner(address SocketAddresser, cancellable *Cancellable) error
	// ConnectAsync: asynchronously connect connection to the specified remote
	// address.
	ConnectAsync(address SocketAddresser, cancellable *Cancellable, callback AsyncReadyCallback)
	// ConnectFinish gets the result of a g_socket_connection_connect_async()
	// call.
	ConnectFinish(result AsyncResulter) error
	// LocalAddress: try to get the local address of a socket connection.
	LocalAddress() (*SocketAddress, error)
	// RemoteAddress: try to get the remote address of a socket connection.
	RemoteAddress() (*SocketAddress, error)
	// Socket gets the underlying #GSocket object of the connection.
	Socket() *Socket
	// IsConnected checks if connection is connected.
	IsConnected() bool
}

// SocketConnection is a OStream for a connected socket. They can be created
// either by Client when connecting to a host, or by Listener when accepting a
// new client.
//
// The type of the Connection object returned from these calls depends on the
// type of the underlying socket that is in use. For instance, for a TCP/IP
// connection it will be a Connection.
//
// Choosing what type of object to construct is done with the socket connection
// factory, and it is possible for 3rd parties to register custom socket
// connection types for specific combination of socket family/type/protocol
// using g_socket_connection_factory_register_type().
//
// To close a Connection, use g_io_stream_close(). Closing both substreams of
// the OStream separately will not close the underlying #GSocket.
type SocketConnection struct {
	IOStream
}

var (
	_ SocketConnectioner = (*SocketConnection)(nil)
	_ gextras.Nativer    = (*SocketConnection)(nil)
)

func wrapSocketConnection(obj *externglib.Object) *SocketConnection {
	return &SocketConnection{
		IOStream: IOStream{
			Object: obj,
		},
	}
}

func marshalSocketConnectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketConnection(obj), nil
}

// ConnectSocketConnectioner: connect connection to the specified remote
// address.
func (connection *SocketConnection) ConnectSocketConnectioner(address SocketAddresser, cancellable *Cancellable) error {
	var _arg0 *C.GSocketConnection // out
	var _arg1 *C.GSocketAddress    // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))
	_arg1 = (*C.GSocketAddress)(unsafe.Pointer((address).(gextras.Nativer).Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_socket_connection_connect(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ConnectAsync: asynchronously connect connection to the specified remote
// address.
//
// This clears the #GSocket:blocking flag on connection's underlying socket if
// it is currently set.
//
// Use g_socket_connection_connect_finish() to retrieve the result.
func (connection *SocketConnection) ConnectAsync(address SocketAddresser, cancellable *Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketConnection  // out
	var _arg1 *C.GSocketAddress     // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))
	_arg1 = (*C.GSocketAddress)(unsafe.Pointer((address).(gextras.Nativer).Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(gbox.AssignOnce(callback))

	C.g_socket_connection_connect_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// ConnectFinish gets the result of a g_socket_connection_connect_async() call.
func (connection *SocketConnection) ConnectFinish(result AsyncResulter) error {
	var _arg0 *C.GSocketConnection // out
	var _arg1 *C.GAsyncResult      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_socket_connection_connect_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// LocalAddress: try to get the local address of a socket connection.
func (connection *SocketConnection) LocalAddress() (*SocketAddress, error) {
	var _arg0 *C.GSocketConnection // out
	var _cret *C.GSocketAddress    // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_get_local_address(_arg0, &_cerr)

	var _socketAddress *SocketAddress // out
	var _goerr error                  // out

	_socketAddress = wrapSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketAddress, _goerr
}

// RemoteAddress: try to get the remote address of a socket connection.
//
// Since GLib 2.40, when used with g_socket_client_connect() or
// g_socket_client_connect_async(), during emission of
// G_SOCKET_CLIENT_CONNECTING, this function will return the remote address that
// will be used for the connection. This allows applications to print e.g.
// "Connecting to example.com (10.42.77.3)...".
func (connection *SocketConnection) RemoteAddress() (*SocketAddress, error) {
	var _arg0 *C.GSocketConnection // out
	var _cret *C.GSocketAddress    // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_get_remote_address(_arg0, &_cerr)

	var _socketAddress *SocketAddress // out
	var _goerr error                  // out

	_socketAddress = wrapSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketAddress, _goerr
}

// Socket gets the underlying #GSocket object of the connection. This can be
// useful if you want to do something unusual on it not supported by the
// Connection APIs.
func (connection *SocketConnection) Socket() *Socket {
	var _arg0 *C.GSocketConnection // out
	var _cret *C.GSocket           // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_get_socket(_arg0)

	var _socket *Socket // out

	_socket = wrapSocket(externglib.Take(unsafe.Pointer(_cret)))

	return _socket
}

// IsConnected checks if connection is connected. This is equivalent to calling
// g_socket_is_connected() on connection's underlying #GSocket.
func (connection *SocketConnection) IsConnected() bool {
	var _arg0 *C.GSocketConnection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GSocketConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_socket_connection_is_connected(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SocketConnectionFactoryLookupType looks up the #GType to be used when
// creating socket connections on sockets with the specified family, type and
// protocol_id.
//
// If no type is registered, the Connection base type is returned.
func SocketConnectionFactoryLookupType(family SocketFamily, typ SocketType, protocolId int) externglib.Type {
	var _arg1 C.GSocketFamily // out
	var _arg2 C.GSocketType   // out
	var _arg3 C.gint          // out
	var _cret C.GType         // in

	_arg1 = C.GSocketFamily(family)
	_arg2 = C.GSocketType(typ)
	_arg3 = C.gint(protocolId)

	_cret = C.g_socket_connection_factory_lookup_type(_arg1, _arg2, _arg3)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SocketConnectionFactoryRegisterType looks up the #GType to be used when
// creating socket connections on sockets with the specified family, type and
// protocol.
//
// If no type is registered, the Connection base type is returned.
func SocketConnectionFactoryRegisterType(gType externglib.Type, family SocketFamily, typ SocketType, protocol int) {
	var _arg1 C.GType         // out
	var _arg2 C.GSocketFamily // out
	var _arg3 C.GSocketType   // out
	var _arg4 C.gint          // out

	_arg1 = C.GType(gType)
	_arg2 = C.GSocketFamily(family)
	_arg3 = C.GSocketType(typ)
	_arg4 = C.gint(protocol)

	C.g_socket_connection_factory_register_type(_arg1, _arg2, _arg3, _arg4)
}
