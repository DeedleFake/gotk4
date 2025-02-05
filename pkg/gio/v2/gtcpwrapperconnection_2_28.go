// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeTCPWrapperConnection = coreglib.Type(C.g_tcp_wrapper_connection_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTCPWrapperConnection, F: marshalTCPWrapperConnection},
	})
}

// TCPWrapperConnectionOverrides contains methods that are overridable.
type TCPWrapperConnectionOverrides struct {
}

func defaultTCPWrapperConnectionOverrides(v *TCPWrapperConnection) TCPWrapperConnectionOverrides {
	return TCPWrapperConnectionOverrides{}
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

func init() {
	coreglib.RegisterClassInfo[*TCPWrapperConnection, *TCPWrapperConnectionClass, TCPWrapperConnectionOverrides](
		GTypeTCPWrapperConnection,
		initTCPWrapperConnectionClass,
		wrapTCPWrapperConnection,
		defaultTCPWrapperConnectionOverrides,
	)
}

func initTCPWrapperConnectionClass(gclass unsafe.Pointer, overrides TCPWrapperConnectionOverrides, classInitFunc func(*TCPWrapperConnectionClass)) {
	if classInitFunc != nil {
		class := (*TCPWrapperConnectionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
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
	var _arg1 *C.GIOStream         // out
	var _arg2 *C.GSocket           // out
	var _cret *C.GSocketConnection // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(coreglib.InternObject(baseIoStream).Native()))
	_arg2 = (*C.GSocket)(unsafe.Pointer(coreglib.InternObject(socket).Native()))

	_cret = C.g_tcp_wrapper_connection_new(_arg1, _arg2)
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
	var _arg0 *C.GTcpWrapperConnection // out
	var _cret *C.GIOStream             // in

	_arg0 = (*C.GTcpWrapperConnection)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_cret = C.g_tcp_wrapper_connection_get_base_io_stream(_arg0)
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
