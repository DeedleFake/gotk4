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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_address_get_type()), F: marshalSocketAddress},
	})
}

// SocketAddress is the equivalent of struct sockaddr in the BSD sockets API.
// This is an abstract class; use SocketAddress for internet sockets, or
// SocketAddress for UNIX domain sockets.
type SocketAddress interface {
	gextras.Objector

	// AsSocketConnectable casts the class to the SocketConnectable interface.
	AsSocketConnectable() SocketConnectable

	// Family gets the socket family type of @address.
	Family() SocketFamily
	// NativeSize gets the size of @address's native struct sockaddr. You can
	// use this to allocate memory to pass to g_socket_address_to_native().
	NativeSize() int
	// ToNativeSocketAddress converts a Address to a native struct sockaddr,
	// which can be passed to low-level functions like connect() or bind().
	//
	// If not enough space is available, a G_IO_ERROR_NO_SPACE error is
	// returned. If the address type is not known on the system then a
	// G_IO_ERROR_NOT_SUPPORTED error is returned.
	ToNativeSocketAddress(dest interface{}, destlen uint) error
}

// socketAddress implements the SocketAddress class.
type socketAddress struct {
	gextras.Objector
}

// WrapSocketAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketAddress(obj *externglib.Object) SocketAddress {
	return socketAddress{
		Objector: obj,
	}
}

func marshalSocketAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketAddress(obj), nil
}

// NewSocketAddressFromNative creates a Address subclass corresponding to the
// native struct sockaddr @native.
func NewSocketAddressFromNative(native interface{}, len uint) SocketAddress {
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (C.gpointer)(box.Assign(native))
	_arg2 = C.gsize(len)

	_cret = C.g_socket_address_new_from_native(_arg1, _arg2)

	var _socketAddress SocketAddress // out

	_socketAddress = WrapSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketAddress
}

func (a socketAddress) Family() SocketFamily {
	var _arg0 *C.GSocketAddress // out
	var _cret C.GSocketFamily   // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_socket_address_get_family(_arg0)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

func (a socketAddress) NativeSize() int {
	var _arg0 *C.GSocketAddress // out
	var _cret C.gssize          // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_socket_address_get_native_size(_arg0)

	var _gssize int // out

	_gssize = int(_cret)

	return _gssize
}

func (a socketAddress) ToNativeSocketAddress(dest interface{}, destlen uint) error {
	var _arg0 *C.GSocketAddress // out
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(a.Native()))
	_arg1 = (C.gpointer)(box.Assign(dest))
	_arg2 = C.gsize(destlen)

	C.g_socket_address_to_native(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s socketAddress) AsSocketConnectable() SocketConnectable {
	return WrapSocketConnectable(gextras.InternObject(s))
}
