// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeTCPWrapperConnection returns the GType for the type TCPWrapperConnection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTCPWrapperConnection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "TcpWrapperConnection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTCPWrapperConnection)
	return gtype
}

// TCPWrapperConnectionOverrider contains methods that are overridable.
type TCPWrapperConnectionOverrider interface {
}

// TCPWrapperConnection can be used to wrap a OStream that is based on a
// #GSocket, but which is not actually a Connection. This is used by Client so
// that it can always return a Connection, even when the connection it has
// actually created is not directly a Connection.
type TCPWrapperConnection struct {
	_ [0]func() // equal guard
	TCPConnection
}

var (
	_ IOStreamer = (*TCPWrapperConnection)(nil)
)

func classInitTCPWrapperConnectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTCPWrapperConnection(obj *coreglib.Object) *TCPWrapperConnection {
	return &TCPWrapperConnection{
		TCPConnection: TCPConnection{
			SocketConnection: SocketConnection{
				IOStream: IOStream{
					Object: obj,
				},
			},
		},
	}
}

func marshalTCPWrapperConnection(p uintptr) (interface{}, error) {
	return wrapTCPWrapperConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTCPWrapperConnection wraps base_io_stream and socket together as a
// Connection.
//
// The function takes the following parameters:
//
//    - baseIoStream to wrap.
//    - socket associated with base_io_stream.
//
// The function returns the following values:
//
//    - tcpWrapperConnection: new Connection.
//
func NewTCPWrapperConnection(baseIoStream IOStreamer, socket *Socket) *TCPWrapperConnection {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseIoStream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(socket).Native()))

	_info := girepository.MustFind("Gio", "TcpWrapperConnection")
	_gret := _info.InvokeClassMethod("new_TcpWrapperConnection", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(socket)

	var _tcpWrapperConnection *TCPWrapperConnection // out

	_tcpWrapperConnection = wrapTCPWrapperConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _tcpWrapperConnection
}

// BaseIOStream gets conn's base OStream.
//
// The function returns the following values:
//
//    - ioStream conn's base OStream.
//
func (conn *TCPWrapperConnection) BaseIOStream() IOStreamer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_info := girepository.MustFind("Gio", "TcpWrapperConnection")
	_gret := _info.InvokeClassMethod("get_base_io_stream", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(conn)

	var _ioStream IOStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}

	return _ioStream
}
