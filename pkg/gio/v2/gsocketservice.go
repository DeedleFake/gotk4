// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_SocketServiceClass_incoming(void*, void*, void*);
import "C"

// GTypeSocketService returns the GType for the type SocketService.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSocketService() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "SocketService").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSocketService)
	return gtype
}

// SocketServiceOverrider contains methods that are overridable.
type SocketServiceOverrider interface {
	// The function takes the following parameters:
	//
	//    - connection
	//    - sourceObject
	//
	// The function returns the following values:
	//
	Incoming(connection *SocketConnection, sourceObject *coreglib.Object) bool
}

// SocketService is an object that represents a service that is provided to the
// network or over local sockets. When a new connection is made to the service
// the Service::incoming signal is emitted.
//
// A Service is a subclass of Listener and you need to add the addresses you
// want to accept connections on with the Listener APIs.
//
// There are two options for implementing a network service based on Service.
// The first is to create the service using g_socket_service_new() and to
// connect to the Service::incoming signal. The second is to subclass Service
// and override the default signal handler implementation.
//
// In either case, the handler must immediately return, or else it will block
// additional incoming connections from being serviced. If you are interested in
// writing connection handlers that contain blocking code then see
// SocketService.
//
// The socket service runs on the main loop of the [thread-default
// context][g-main-context-push-thread-default-context] of the thread it is
// created in, and is not threadsafe in general. However, the calls to start and
// stop the service are thread-safe so these can be used from threads that
// handle incoming clients.
type SocketService struct {
	_ [0]func() // equal guard
	SocketListener
}

var (
	_ coreglib.Objector = (*SocketService)(nil)
)

func classInitSocketServicer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "SocketServiceClass")

	if _, ok := goval.(interface {
		Incoming(connection *SocketConnection, sourceObject *coreglib.Object) bool
	}); ok {
		o := pclass.StructFieldOffset("incoming")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_SocketServiceClass_incoming)
	}
}

//export _gotk4_gio2_SocketServiceClass_incoming
func _gotk4_gio2_SocketServiceClass_incoming(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Incoming(connection *SocketConnection, sourceObject *coreglib.Object) bool
	})

	var _connection *SocketConnection  // out
	var _sourceObject *coreglib.Object // out

	_connection = wrapSocketConnection(coreglib.Take(unsafe.Pointer(arg1)))
	_sourceObject = coreglib.Take(unsafe.Pointer(arg2))

	ok := iface.Incoming(_connection, _sourceObject)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapSocketService(obj *coreglib.Object) *SocketService {
	return &SocketService{
		SocketListener: SocketListener{
			Object: obj,
		},
	}
}

func marshalSocketService(p uintptr) (interface{}, error) {
	return wrapSocketService(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSocketService creates a new Service with no sockets to listen for. New
// listeners can be added with e.g. g_socket_listener_add_address() or
// g_socket_listener_add_inet_port().
//
// New services are created active, there is no need to call
// g_socket_service_start(), unless g_socket_service_stop() has been called
// before.
//
// The function returns the following values:
//
//    - socketService: new Service.
//
func NewSocketService() *SocketService {
	_info := girepository.MustFind("Gio", "SocketService")
	_gret := _info.InvokeClassMethod("new_SocketService", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _socketService *SocketService // out

	_socketService = wrapSocketService(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketService
}

// IsActive: check whether the service is active or not. An active service will
// accept new clients that connect, while a non-active service will let
// connecting clients queue up until the service is started.
//
// The function returns the following values:
//
//    - ok: TRUE if the service is active, FALSE otherwise.
//
func (service *SocketService) IsActive() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(service).Native()))

	_info := girepository.MustFind("Gio", "SocketService")
	_gret := _info.InvokeClassMethod("is_active", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(service)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Start restarts the service, i.e. start accepting connections from the added
// sockets when the mainloop runs. This only needs to be called after the
// service has been stopped from g_socket_service_stop().
//
// This call is thread-safe, so it may be called from a thread handling an
// incoming client request.
func (service *SocketService) Start() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(service).Native()))

	_info := girepository.MustFind("Gio", "SocketService")
	_info.InvokeClassMethod("start", _args[:], nil)

	runtime.KeepAlive(service)
}

// Stop stops the service, i.e. stops accepting connections from the added
// sockets when the mainloop runs.
//
// This call is thread-safe, so it may be called from a thread handling an
// incoming client request.
//
// Note that this only stops accepting new connections; it does not close the
// listening sockets, and you can call g_socket_service_start() again later to
// begin listening again. To close the listening sockets, call
// g_socket_listener_close(). (This will happen automatically when the Service
// is finalized.)
//
// This must be called before calling g_socket_listener_close() as the socket
// service will start accepting connections immediately when a new socket is
// added.
func (service *SocketService) Stop() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(service).Native()))

	_info := girepository.MustFind("Gio", "SocketService")
	_info.InvokeClassMethod("stop", _args[:], nil)

	runtime.KeepAlive(service)
}
