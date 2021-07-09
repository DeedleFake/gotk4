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
		{T: externglib.Type(C.g_socket_listener_get_type()), F: marshalSocketListener},
	})
}

// SocketListenerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketListenerOverrider interface {
	Changed()
}

// SocketListener is an object that keeps track of a set of server sockets and
// helps you accept sockets from any of the socket, either sync or async.
//
// Add addresses and ports to listen on using g_socket_listener_add_address()
// and g_socket_listener_add_inet_port(). These will be listened on until
// g_socket_listener_close() is called. Dropping your final reference to the
// Listener will not cause g_socket_listener_close() to be called implicitly, as
// some references to the Listener may be held internally.
//
// If you want to implement a network server, also look at Service and
// SocketService which are subclasses of Listener that make this even easier.
type SocketListener interface {
	gextras.Objector

	// Accept blocks waiting for a client to connect to any of the sockets added
	// to the listener. Returns a Connection for the socket that was accepted.
	//
	// If @source_object is not nil it will be filled out with the source object
	// specified when the corresponding socket or address was added to the
	// listener.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	Accept(cancellable Cancellable) (*externglib.Object, *SocketConnectionClass, error)
	// AcceptAsync: this is the asynchronous version of
	// g_socket_listener_accept().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_socket_listener_accept_socket() to get the result of the
	// operation.
	AcceptAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// AcceptFinish finishes an async accept operation. See
	// g_socket_listener_accept_async()
	AcceptFinish(result AsyncResult) (*externglib.Object, *SocketConnectionClass, error)
	// AcceptSocket blocks waiting for a client to connect to any of the sockets
	// added to the listener. Returns the #GSocket that was accepted.
	//
	// If you want to accept the high-level Connection, not a #GSocket, which is
	// often the case, then you should use g_socket_listener_accept() instead.
	//
	// If @source_object is not nil it will be filled out with the source object
	// specified when the corresponding socket or address was added to the
	// listener.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	AcceptSocket(cancellable Cancellable) (*externglib.Object, *SocketClass, error)
	// AcceptSocketAsync: this is the asynchronous version of
	// g_socket_listener_accept_socket().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_socket_listener_accept_socket_finish() to get the result of the
	// operation.
	AcceptSocketAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// AcceptSocketFinish finishes an async accept operation. See
	// g_socket_listener_accept_socket_async()
	AcceptSocketFinish(result AsyncResult) (*externglib.Object, *SocketClass, error)
	// AddAnyInetPort listens for TCP connections on any available port number
	// for both IPv6 and IPv4 (if each is available).
	//
	// This is useful if you need to have a socket for incoming connections but
	// don't care about the specific port number.
	//
	// @source_object will be passed out in the various calls to accept to
	// identify this particular source, which is useful if you're listening on
	// multiple addresses and do different things depending on what address is
	// connected to.
	AddAnyInetPort(sourceObject gextras.Objector) (uint16, error)
	// AddInetPort: helper function for g_socket_listener_add_address() that
	// creates a TCP/IP socket listening on IPv4 and IPv6 (if supported) on the
	// specified port on all interfaces.
	//
	// @source_object will be passed out in the various calls to accept to
	// identify this particular source, which is useful if you're listening on
	// multiple addresses and do different things depending on what address is
	// connected to.
	//
	// Call g_socket_listener_close() to stop listening on @port; this will not
	// be done automatically when you drop your final reference to @listener, as
	// references may be held internally.
	AddInetPort(port uint16, sourceObject gextras.Objector) error
	// AddSocket adds @socket to the set of sockets that we try to accept new
	// clients from. The socket must be bound to a local address and listened
	// to.
	//
	// @source_object will be passed out in the various calls to accept to
	// identify this particular source, which is useful if you're listening on
	// multiple addresses and do different things depending on what address is
	// connected to.
	//
	// The @socket will not be automatically closed when the @listener is
	// finalized unless the listener held the final reference to the socket.
	// Before GLib 2.42, the @socket was automatically closed on finalization of
	// the @listener, even if references to it were held elsewhere.
	AddSocket(socket Socket, sourceObject gextras.Objector) error
	// Close closes all the sockets in the listener.
	Close()
	// SetBacklog sets the listen backlog on the sockets in the listener. This
	// must be called before adding any sockets, addresses or ports to the
	// Listener (for example, by calling g_socket_listener_add_inet_port()) to
	// be effective.
	//
	// See g_socket_set_listen_backlog() for details
	SetBacklog(listenBacklog int)
}

// SocketListenerClass implements the SocketListener interface.
type SocketListenerClass struct {
	*externglib.Object
}

var _ SocketListener = (*SocketListenerClass)(nil)

func wrapSocketListener(obj *externglib.Object) SocketListener {
	return &SocketListenerClass{
		Object: obj,
	}
}

func marshalSocketListener(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketListener(obj), nil
}

// NewSocketListener creates a new Listener with no sockets to listen for. New
// listeners can be added with e.g. g_socket_listener_add_address() or
// g_socket_listener_add_inet_port().
func NewSocketListener() *SocketListenerClass {
	var _cret *C.GSocketListener // in

	_cret = C.g_socket_listener_new()

	var _socketListener *SocketListenerClass // out

	_socketListener = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SocketListenerClass)

	return _socketListener
}

// Accept blocks waiting for a client to connect to any of the sockets added to
// the listener. Returns a Connection for the socket that was accepted.
//
// If @source_object is not nil it will be filled out with the source object
// specified when the corresponding socket or address was added to the listener.
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (l *SocketListenerClass) Accept(cancellable Cancellable) (*externglib.Object, *SocketConnectionClass, error) {
	var _arg0 *C.GSocketListener   // out
	var _arg1 *C.GObject           // in
	var _arg2 *C.GCancellable      // out
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer((&Cancellable).Native()))

	_cret = C.g_socket_listener_accept(_arg0, &_arg1, _arg2, &_cerr)

	var _sourceObject *externglib.Object         // out
	var _socketConnection *SocketConnectionClass // out
	var _goerr error                             // out

	_sourceObject = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_arg1))).(*externglib.Object)
	_socketConnection = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SocketConnectionClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socketConnection, _goerr
}

// AcceptAsync: this is the asynchronous version of g_socket_listener_accept().
//
// When the operation is finished @callback will be called. You can then call
// g_socket_listener_accept_socket() to get the result of the operation.
func (l *SocketListenerClass) AcceptAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((&Cancellable).Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_socket_listener_accept_async(_arg0, _arg1, _arg2, _arg3)
}

// AcceptFinish finishes an async accept operation. See
// g_socket_listener_accept_async()
func (l *SocketListenerClass) AcceptFinish(result AsyncResult) (*externglib.Object, *SocketConnectionClass, error) {
	var _arg0 *C.GSocketListener   // out
	var _arg1 *C.GAsyncResult      // out
	var _arg2 *C.GObject           // in
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&AsyncResult).Native()))

	_cret = C.g_socket_listener_accept_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _sourceObject *externglib.Object         // out
	var _socketConnection *SocketConnectionClass // out
	var _goerr error                             // out

	_sourceObject = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_arg2))).(*externglib.Object)
	_socketConnection = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SocketConnectionClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socketConnection, _goerr
}

// AcceptSocket blocks waiting for a client to connect to any of the sockets
// added to the listener. Returns the #GSocket that was accepted.
//
// If you want to accept the high-level Connection, not a #GSocket, which is
// often the case, then you should use g_socket_listener_accept() instead.
//
// If @source_object is not nil it will be filled out with the source object
// specified when the corresponding socket or address was added to the listener.
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
func (l *SocketListenerClass) AcceptSocket(cancellable Cancellable) (*externglib.Object, *SocketClass, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GObject         // in
	var _arg2 *C.GCancellable    // out
	var _cret *C.GSocket         // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer((&Cancellable).Native()))

	_cret = C.g_socket_listener_accept_socket(_arg0, &_arg1, _arg2, &_cerr)

	var _sourceObject *externglib.Object // out
	var _socket *SocketClass             // out
	var _goerr error                     // out

	_sourceObject = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_arg1))).(*externglib.Object)
	_socket = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SocketClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socket, _goerr
}

// AcceptSocketAsync: this is the asynchronous version of
// g_socket_listener_accept_socket().
//
// When the operation is finished @callback will be called. You can then call
// g_socket_listener_accept_socket_finish() to get the result of the operation.
func (l *SocketListenerClass) AcceptSocketAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((&Cancellable).Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_socket_listener_accept_socket_async(_arg0, _arg1, _arg2, _arg3)
}

// AcceptSocketFinish finishes an async accept operation. See
// g_socket_listener_accept_socket_async()
func (l *SocketListenerClass) AcceptSocketFinish(result AsyncResult) (*externglib.Object, *SocketClass, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GAsyncResult    // out
	var _arg2 *C.GObject         // in
	var _cret *C.GSocket         // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&AsyncResult).Native()))

	_cret = C.g_socket_listener_accept_socket_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _sourceObject *externglib.Object // out
	var _socket *SocketClass             // out
	var _goerr error                     // out

	_sourceObject = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_arg2))).(*externglib.Object)
	_socket = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SocketClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _sourceObject, _socket, _goerr
}

// AddAnyInetPort listens for TCP connections on any available port number for
// both IPv6 and IPv4 (if each is available).
//
// This is useful if you need to have a socket for incoming connections but
// don't care about the specific port number.
//
// @source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
func (l *SocketListenerClass) AddAnyInetPort(sourceObject gextras.Objector) (uint16, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GObject         // out
	var _cret C.guint16          // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer((&gextras.Objector).Native()))

	_cret = C.g_socket_listener_add_any_inet_port(_arg0, _arg1, &_cerr)

	var _guint16 uint16 // out
	var _goerr error    // out

	_guint16 = uint16(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint16, _goerr
}

// AddInetPort: helper function for g_socket_listener_add_address() that creates
// a TCP/IP socket listening on IPv4 and IPv6 (if supported) on the specified
// port on all interfaces.
//
// @source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// Call g_socket_listener_close() to stop listening on @port; this will not be
// done automatically when you drop your final reference to @listener, as
// references may be held internally.
func (l *SocketListenerClass) AddInetPort(port uint16, sourceObject gextras.Objector) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.guint16          // out
	var _arg2 *C.GObject         // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = C.guint16(port)
	_arg2 = (*C.GObject)(unsafe.Pointer((&gextras.Objector).Native()))

	C.g_socket_listener_add_inet_port(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// AddSocket adds @socket to the set of sockets that we try to accept new
// clients from. The socket must be bound to a local address and listened to.
//
// @source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// The @socket will not be automatically closed when the @listener is finalized
// unless the listener held the final reference to the socket. Before GLib 2.42,
// the @socket was automatically closed on finalization of the @listener, even
// if references to it were held elsewhere.
func (l *SocketListenerClass) AddSocket(socket Socket, sourceObject gextras.Objector) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GSocket         // out
	var _arg2 *C.GObject         // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = (*C.GSocket)(unsafe.Pointer((&Socket).Native()))
	_arg2 = (*C.GObject)(unsafe.Pointer((&gextras.Objector).Native()))

	C.g_socket_listener_add_socket(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Close closes all the sockets in the listener.
func (l *SocketListenerClass) Close() {
	var _arg0 *C.GSocketListener // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))

	C.g_socket_listener_close(_arg0)
}

// SetBacklog sets the listen backlog on the sockets in the listener. This must
// be called before adding any sockets, addresses or ports to the Listener (for
// example, by calling g_socket_listener_add_inet_port()) to be effective.
//
// See g_socket_set_listen_backlog() for details
func (l *SocketListenerClass) SetBacklog(listenBacklog int) {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.int              // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer((&SocketListener).Native()))
	_arg1 = C.int(listenBacklog)

	C.g_socket_listener_set_backlog(_arg0, _arg1)
}
