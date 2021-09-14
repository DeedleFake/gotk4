// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_socket_address_get_type()), F: marshalSocketAddresser},
	})
}

// SocketAddressOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketAddressOverrider interface {
	// Family gets the socket family type of address.
	Family() SocketFamily
	// NativeSize gets the size of address's native struct sockaddr. You can use
	// this to allocate memory to pass to g_socket_address_to_native().
	NativeSize() int
	// ToNative converts a Address to a native struct sockaddr, which can be
	// passed to low-level functions like connect() or bind().
	//
	// If not enough space is available, a G_IO_ERROR_NO_SPACE error is
	// returned. If the address type is not known on the system then a
	// G_IO_ERROR_NOT_SUPPORTED error is returned.
	ToNative(dest cgo.Handle, destlen uint) error
}

// SocketAddress is the equivalent of struct sockaddr in the BSD sockets API.
// This is an abstract class; use SocketAddress for internet sockets, or
// SocketAddress for UNIX domain sockets.
type SocketAddress struct {
	*externglib.Object

	SocketConnectable
}

// SocketAddresser describes SocketAddress's abstract methods.
type SocketAddresser interface {
	externglib.Objector

	// Family gets the socket family type of address.
	Family() SocketFamily
	// NativeSize gets the size of address's native struct sockaddr.
	NativeSize() int
	// ToNative converts a Address to a native struct sockaddr, which can be
	// passed to low-level functions like connect() or bind().
	ToNative(dest cgo.Handle, destlen uint) error
}

var _ SocketAddresser = (*SocketAddress)(nil)

func wrapSocketAddress(obj *externglib.Object) *SocketAddress {
	return &SocketAddress{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalSocketAddresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketAddress(obj), nil
}

// NewSocketAddressFromNative creates a Address subclass corresponding to the
// native struct sockaddr native.
func NewSocketAddressFromNative(native cgo.Handle, len uint) *SocketAddress {
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (C.gpointer)(unsafe.Pointer(native))
	_arg2 = C.gsize(len)

	_cret = C.g_socket_address_new_from_native(_arg1, _arg2)
	runtime.KeepAlive(native)
	runtime.KeepAlive(len)

	var _socketAddress *SocketAddress // out

	_socketAddress = wrapSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketAddress
}

// Family gets the socket family type of address.
func (address *SocketAddress) Family() SocketFamily {
	var _arg0 *C.GSocketAddress // out
	var _cret C.GSocketFamily   // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_socket_address_get_family(_arg0)
	runtime.KeepAlive(address)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// NativeSize gets the size of address's native struct sockaddr. You can use
// this to allocate memory to pass to g_socket_address_to_native().
func (address *SocketAddress) NativeSize() int {
	var _arg0 *C.GSocketAddress // out
	var _cret C.gssize          // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_socket_address_get_native_size(_arg0)
	runtime.KeepAlive(address)

	var _gssize int // out

	_gssize = int(_cret)

	return _gssize
}

// ToNative converts a Address to a native struct sockaddr, which can be passed
// to low-level functions like connect() or bind().
//
// If not enough space is available, a G_IO_ERROR_NO_SPACE error is returned. If
// the address type is not known on the system then a G_IO_ERROR_NOT_SUPPORTED
// error is returned.
func (address *SocketAddress) ToNative(dest cgo.Handle, destlen uint) error {
	var _arg0 *C.GSocketAddress // out
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(dest))
	_arg2 = C.gsize(destlen)

	C.g_socket_address_to_native(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(address)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(destlen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
