// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tcp_connection_get_type()), F: marshalTCPConnection},
	})
}

// TCPConnection: this is the subclass of Connection that is created for TCP/IP
// sockets.
type TCPConnection interface {
	SocketConnection

	// AsSocketConnection casts the class to the SocketConnection interface.
	AsSocketConnection() SocketConnection

	// Connect @connection to the specified remote address.
	//
	// This method is inherited from SocketConnection
	Connect(address SocketAddress, cancellable Cancellable) error
	// ConnectAsync: asynchronously connect @connection to the specified remote
	// address.
	//
	// This clears the #GSocket:blocking flag on @connection's underlying socket
	// if it is currently set.
	//
	// Use g_socket_connection_connect_finish() to retrieve the result.
	//
	// This method is inherited from SocketConnection
	ConnectAsync(address SocketAddress, cancellable Cancellable, callback AsyncReadyCallback)
	// ConnectFinish gets the result of a g_socket_connection_connect_async()
	// call.
	//
	// This method is inherited from SocketConnection
	ConnectFinish(result AsyncResult) error
	// GetLocalAddress: try to get the local address of a socket connection.
	//
	// This method is inherited from SocketConnection
	GetLocalAddress() (SocketAddress, error)
	// GetRemoteAddress: try to get the remote address of a socket connection.
	//
	// Since GLib 2.40, when used with g_socket_client_connect() or
	// g_socket_client_connect_async(), during emission of
	// G_SOCKET_CLIENT_CONNECTING, this function will return the remote address
	// that will be used for the connection. This allows applications to print
	// e.g. "Connecting to example.com (10.42.77.3)...".
	//
	// This method is inherited from SocketConnection
	GetRemoteAddress() (SocketAddress, error)
	// GetSocket gets the underlying #GSocket object of the connection. This can
	// be useful if you want to do something unusual on it not supported by the
	// Connection APIs.
	//
	// This method is inherited from SocketConnection
	GetSocket() Socket
	// IsConnected checks if @connection is connected. This is equivalent to
	// calling g_socket_is_connected() on @connection's underlying #GSocket.
	//
	// This method is inherited from SocketConnection
	IsConnected() bool
	// ClearPending clears the pending flag on @stream.
	//
	// This method is inherited from IOStream
	ClearPending()
	// Close closes the stream, releasing resources related to it. This will
	// also close the individual input and output streams, if they are not
	// already closed.
	//
	// Once the stream is closed, all other operations will return
	// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
	// error.
	//
	// Closing a stream will automatically flush any outstanding buffers in the
	// stream.
	//
	// Streams will be automatically closed when the last reference is dropped,
	// but you might want to call this function to make sure resources are
	// released as early as possible.
	//
	// Some streams might keep the backing store of the stream (e.g. a file
	// descriptor) open after the stream is closed. See the documentation for
	// the individual stream for details.
	//
	// On failure the first error that happened will be reported, but the close
	// operation will finish as much as possible. A stream that failed to close
	// will still return G_IO_ERROR_CLOSED for all operations. Still, it is
	// important to check and report the error to the user, otherwise there
	// might be a loss of data as all data might not be written.
	//
	// If @cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	// Cancelling a close will still leave the stream closed, but some streams
	// can use a faster close that doesn't block to e.g. check errors.
	//
	// The default implementation of this method just calls close on the
	// individual input/output streams.
	//
	// This method is inherited from IOStream
	Close(cancellable Cancellable) error
	// CloseAsync requests an asynchronous close of the stream, releasing
	// resources related to it. When the operation is finished @callback will be
	// called. You can then call g_io_stream_close_finish() to get the result of
	// the operation.
	//
	// For behaviour details see g_io_stream_close().
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// This method is inherited from IOStream
	CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// CloseFinish closes a stream.
	//
	// This method is inherited from IOStream
	CloseFinish(result AsyncResult) error
	// GetInputStream gets the input stream for this object. This is used for
	// reading.
	//
	// This method is inherited from IOStream
	GetInputStream() InputStream
	// GetOutputStream gets the output stream for this object. This is used for
	// writing.
	//
	// This method is inherited from IOStream
	GetOutputStream() OutputStream
	// HasPending checks if a stream has pending actions.
	//
	// This method is inherited from IOStream
	HasPending() bool
	// IsClosed checks if a stream is closed.
	//
	// This method is inherited from IOStream
	IsClosed() bool
	// SetPending sets @stream to have actions pending. If the pending flag is
	// already set or @stream is closed, it will return false and set @error.
	//
	// This method is inherited from IOStream
	SetPending() error
	// SpliceAsync: asynchronously splice the output stream of @stream1 to the
	// input stream of @stream2, and splice the output stream of @stream2 to the
	// input stream of @stream1.
	//
	// When the operation is finished @callback will be called. You can then
	// call g_io_stream_splice_finish() to get the result of the operation.
	//
	// This method is inherited from IOStream
	SpliceAsync(stream2 IOStream, flags IOStreamSpliceFlags, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)

	// GracefulDisconnect checks if graceful disconnects are used. See
	// g_tcp_connection_set_graceful_disconnect().
	GracefulDisconnect() bool
	// SetGracefulDisconnect: this enables graceful disconnects on close. A
	// graceful disconnect means that we signal the receiving end that the
	// connection is terminated and wait for it to close the connection before
	// closing the connection.
	//
	// A graceful disconnect means that we can be sure that we successfully sent
	// all the outstanding data to the other end, or get an error reported.
	// However, it also means we have to wait for all the data to reach the
	// other side and for it to acknowledge this by closing the socket, which
	// may take a while. For this reason it is disabled by default.
	SetGracefulDisconnect(gracefulDisconnect bool)
}

// tcpConnection implements the TCPConnection interface.
type tcpConnection struct {
	*externglib.Object
}

var _ TCPConnection = (*tcpConnection)(nil)

// WrapTCPConnection wraps a GObject to a type that implements
// interface TCPConnection. It is primarily used internally.
func WrapTCPConnection(obj *externglib.Object) TCPConnection {
	return tcpConnection{obj}
}

func marshalTCPConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTCPConnection(obj), nil
}

func (t tcpConnection) AsSocketConnection() SocketConnection {
	return WrapSocketConnection(gextras.InternObject(t))
}

func (c tcpConnection) Connect(address SocketAddress, cancellable Cancellable) error {
	return WrapSocketConnection(gextras.InternObject(c)).Connect(address, cancellable)
}

func (c tcpConnection) ConnectAsync(address SocketAddress, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapSocketConnection(gextras.InternObject(c)).ConnectAsync(address, cancellable, callback)
}

func (c tcpConnection) ConnectFinish(result AsyncResult) error {
	return WrapSocketConnection(gextras.InternObject(c)).ConnectFinish(result)
}

func (c tcpConnection) GetLocalAddress() (SocketAddress, error) {
	return WrapSocketConnection(gextras.InternObject(c)).GetLocalAddress()
}

func (c tcpConnection) GetRemoteAddress() (SocketAddress, error) {
	return WrapSocketConnection(gextras.InternObject(c)).GetRemoteAddress()
}

func (c tcpConnection) GetSocket() Socket {
	return WrapSocketConnection(gextras.InternObject(c)).GetSocket()
}

func (c tcpConnection) IsConnected() bool {
	return WrapSocketConnection(gextras.InternObject(c)).IsConnected()
}

func (s tcpConnection) ClearPending() {
	WrapIOStream(gextras.InternObject(s)).ClearPending()
}

func (s tcpConnection) Close(cancellable Cancellable) error {
	return WrapIOStream(gextras.InternObject(s)).Close(cancellable)
}

func (s tcpConnection) CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapIOStream(gextras.InternObject(s)).CloseAsync(ioPriority, cancellable, callback)
}

func (s tcpConnection) CloseFinish(result AsyncResult) error {
	return WrapIOStream(gextras.InternObject(s)).CloseFinish(result)
}

func (s tcpConnection) GetInputStream() InputStream {
	return WrapIOStream(gextras.InternObject(s)).GetInputStream()
}

func (s tcpConnection) GetOutputStream() OutputStream {
	return WrapIOStream(gextras.InternObject(s)).GetOutputStream()
}

func (s tcpConnection) HasPending() bool {
	return WrapIOStream(gextras.InternObject(s)).HasPending()
}

func (s tcpConnection) IsClosed() bool {
	return WrapIOStream(gextras.InternObject(s)).IsClosed()
}

func (s tcpConnection) SetPending() error {
	return WrapIOStream(gextras.InternObject(s)).SetPending()
}

func (s tcpConnection) SpliceAsync(stream2 IOStream, flags IOStreamSpliceFlags, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapIOStream(gextras.InternObject(s)).SpliceAsync(stream2, flags, ioPriority, cancellable, callback)
}

func (c tcpConnection) GracefulDisconnect() bool {
	var _arg0 *C.GTcpConnection // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GTcpConnection)(unsafe.Pointer(c.Native()))

	_cret = C.g_tcp_connection_get_graceful_disconnect(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c tcpConnection) SetGracefulDisconnect(gracefulDisconnect bool) {
	var _arg0 *C.GTcpConnection // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GTcpConnection)(unsafe.Pointer(c.Native()))
	if gracefulDisconnect {
		_arg1 = C.TRUE
	}

	C.g_tcp_connection_set_graceful_disconnect(_arg0, _arg1)
}
