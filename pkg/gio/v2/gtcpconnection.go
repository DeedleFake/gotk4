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
		{T: externglib.Type(C.g_tcp_connection_get_type()), F: marshalTCPConnection},
	})
}

// TCPConnection: this is the subclass of Connection that is created for TCP/IP
// sockets.
type TCPConnection interface {
	gextras.Objector

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

// TCPConnectionClass implements the TCPConnection interface.
type TCPConnectionClass struct {
	SocketConnectionClass
}

var _ TCPConnection = (*TCPConnectionClass)(nil)

func wrapTCPConnection(obj *externglib.Object) TCPConnection {
	return &TCPConnectionClass{
		SocketConnectionClass: SocketConnectionClass{
			IOStreamClass: IOStreamClass{
				Object: obj,
			},
		},
	}
}

func marshalTCPConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTCPConnection(obj), nil
}

// GracefulDisconnect checks if graceful disconnects are used. See
// g_tcp_connection_set_graceful_disconnect().
func (connection *TCPConnectionClass) GracefulDisconnect() bool {
	var _arg0 *C.GTcpConnection // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GTcpConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_tcp_connection_get_graceful_disconnect(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetGracefulDisconnect: this enables graceful disconnects on close. A graceful
// disconnect means that we signal the receiving end that the connection is
// terminated and wait for it to close the connection before closing the
// connection.
//
// A graceful disconnect means that we can be sure that we successfully sent all
// the outstanding data to the other end, or get an error reported. However, it
// also means we have to wait for all the data to reach the other side and for
// it to acknowledge this by closing the socket, which may take a while. For
// this reason it is disabled by default.
func (connection *TCPConnectionClass) SetGracefulDisconnect(gracefulDisconnect bool) {
	var _arg0 *C.GTcpConnection // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GTcpConnection)(unsafe.Pointer(connection.Native()))
	if gracefulDisconnect {
		_arg1 = C.TRUE
	}

	C.g_tcp_connection_set_graceful_disconnect(_arg0, _arg1)
}
