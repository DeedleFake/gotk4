// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_ThreadedSocketServiceClass_run(GThreadedSocketService*, GSocketConnection*, GObject*);
// extern gboolean _gotk4_gio2_ThreadedSocketService_ConnectRun(gpointer, GSocketConnection*, GObject, guintptr);
import "C"

// GType values.
var (
	GTypeThreadedSocketService = coreglib.Type(C.g_threaded_socket_service_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeThreadedSocketService, F: marshalThreadedSocketService},
	})
}

// ThreadedSocketServiceOverrider contains methods that are overridable.
type ThreadedSocketServiceOverrider interface {
	// The function takes the following parameters:
	//
	//    - connection
	//    - sourceObject
	//
	// The function returns the following values:
	//
	Run(connection *SocketConnection, sourceObject *coreglib.Object) bool
}

// ThreadedSocketService is a simple subclass of Service that handles incoming
// connections by creating a worker thread and dispatching the connection to it
// by emitting the SocketService::run signal in the new thread.
//
// The signal handler may perform blocking IO and need not return until the
// connection is closed.
//
// The service is implemented using a thread pool, so there is a limited amount
// of threads available to serve incoming requests. The service automatically
// stops the Service from accepting new connections when all threads are busy.
//
// As with Service, you may connect to SocketService::run, or subclass and
// override the default handler.
type ThreadedSocketService struct {
	_ [0]func() // equal guard
	SocketService
}

var (
	_ coreglib.Objector = (*ThreadedSocketService)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:        GTypeThreadedSocketService,
		GoType:       reflect.TypeOf((*ThreadedSocketService)(nil)),
		InitClass:    initClassThreadedSocketService,
		ClassSize:    uint32(unsafe.Sizeof(C.GThreadedSocketService{})),
		InstanceSize: uint32(unsafe.Sizeof(C.GThreadedSocketServiceClass{})),
	})
}

func initClassThreadedSocketService(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GThreadedSocketServiceClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface {
		Run(connection *SocketConnection, sourceObject *coreglib.Object) bool
	}); ok {
		pclass.run = (*[0]byte)(C._gotk4_gio2_ThreadedSocketServiceClass_run)
	}
}

//export _gotk4_gio2_ThreadedSocketServiceClass_run
func _gotk4_gio2_ThreadedSocketServiceClass_run(arg0 *C.GThreadedSocketService, arg1 *C.GSocketConnection, arg2 *C.GObject) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Run(connection *SocketConnection, sourceObject *coreglib.Object) bool
	})

	var _connection *SocketConnection  // out
	var _sourceObject *coreglib.Object // out

	_connection = wrapSocketConnection(coreglib.Take(unsafe.Pointer(arg1)))
	_sourceObject = coreglib.Take(unsafe.Pointer(arg2))

	ok := iface.Run(_connection, _sourceObject)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapThreadedSocketService(obj *coreglib.Object) *ThreadedSocketService {
	return &ThreadedSocketService{
		SocketService: SocketService{
			SocketListener: SocketListener{
				Object: obj,
			},
		},
	}
}

func marshalThreadedSocketService(p uintptr) (interface{}, error) {
	return wrapThreadedSocketService(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_ThreadedSocketService_ConnectRun
func _gotk4_gio2_ThreadedSocketService_ConnectRun(arg0 C.gpointer, arg1 *C.GSocketConnection, arg2 C.GObject, arg3 C.guintptr) (cret C.gboolean) {
	var f func(connection *SocketConnection, sourceObject *coreglib.Object) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(connection *SocketConnection, sourceObject *coreglib.Object) (ok bool))
	}

	var _connection *SocketConnection  // out
	var _sourceObject *coreglib.Object // out

	_connection = wrapSocketConnection(coreglib.Take(unsafe.Pointer(arg1)))
	_sourceObject = coreglib.Take(unsafe.Pointer(&arg2))

	ok := f(_connection, _sourceObject)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectRun signal is emitted in a worker thread in response to an incoming
// connection. This thread is dedicated to handling connection and may perform
// blocking IO. The signal handler need not return until the connection is
// closed.
func (service *ThreadedSocketService) ConnectRun(f func(connection *SocketConnection, sourceObject *coreglib.Object) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(service, "run", false, unsafe.Pointer(C._gotk4_gio2_ThreadedSocketService_ConnectRun), f)
}

// NewThreadedSocketService creates a new SocketService with no listeners.
// Listeners must be added with one of the Listener "add" methods.
//
// The function takes the following parameters:
//
//    - maxThreads: maximal number of threads to execute concurrently handling
//      incoming clients, -1 means no limit.
//
// The function returns the following values:
//
//    - threadedSocketService: new Service.
//
func NewThreadedSocketService(maxThreads int) *ThreadedSocketService {
	var _arg1 C.int             // out
	var _cret *C.GSocketService // in

	_arg1 = C.int(maxThreads)

	_cret = C.g_threaded_socket_service_new(_arg1)
	runtime.KeepAlive(maxThreads)

	var _threadedSocketService *ThreadedSocketService // out

	_threadedSocketService = wrapThreadedSocketService(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _threadedSocketService
}
