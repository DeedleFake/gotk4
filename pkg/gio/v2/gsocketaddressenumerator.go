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
		{T: externglib.Type(C.g_socket_address_enumerator_get_type()), F: marshalSocketAddressEnumeratorer},
	})
}

// SocketAddressEnumeratorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketAddressEnumeratorOverrider interface {
	// Next retrieves the next Address from enumerator. Note that this may block
	// for some amount of time. (Eg, a Address may need to do a DNS lookup
	// before it can return an address.) Use
	// g_socket_address_enumerator_next_async() if you need to avoid blocking.
	//
	// If enumerator is expected to yield addresses, but for some reason is
	// unable to (eg, because of a DNS error), then the first call to
	// g_socket_address_enumerator_next() will return an appropriate error in
	// *error. However, if the first call to g_socket_address_enumerator_next()
	// succeeds, then any further internal errors (other than cancellable being
	// triggered) will be ignored.
	Next(cancellable *Cancellable) (*SocketAddress, error)
	// NextAsync: asynchronously retrieves the next Address from enumerator and
	// then calls callback, which must call
	// g_socket_address_enumerator_next_finish() to get the result.
	//
	// It is an error to call this multiple times before the previous callback
	// has finished.
	NextAsync(cancellable *Cancellable, callback AsyncReadyCallback)
	// NextFinish retrieves the result of a completed call to
	// g_socket_address_enumerator_next_async(). See
	// g_socket_address_enumerator_next() for more information about error
	// handling.
	NextFinish(result AsyncResulter) (*SocketAddress, error)
}

// SocketAddressEnumeratorer describes SocketAddressEnumerator's methods.
type SocketAddressEnumeratorer interface {
	// Next retrieves the next Address from enumerator.
	Next(cancellable *Cancellable) (*SocketAddress, error)
	// NextAsync: asynchronously retrieves the next Address from enumerator and
	// then calls callback, which must call
	// g_socket_address_enumerator_next_finish() to get the result.
	NextAsync(cancellable *Cancellable, callback AsyncReadyCallback)
	// NextFinish retrieves the result of a completed call to
	// g_socket_address_enumerator_next_async().
	NextFinish(result AsyncResulter) (*SocketAddress, error)
}

// SocketAddressEnumerator is an enumerator type for Address instances. It is
// returned by enumeration functions such as g_socket_connectable_enumerate(),
// which returns a AddressEnumerator to list each Address which could be used to
// connect to that Connectable.
//
// Enumeration is typically a blocking operation, so the asynchronous methods
// g_socket_address_enumerator_next_async() and
// g_socket_address_enumerator_next_finish() should be used where possible.
//
// Each AddressEnumerator can only be enumerated once. Once
// g_socket_address_enumerator_next() has returned NULL, further enumeration
// with that AddressEnumerator is not possible, and it can be unreffed.
type SocketAddressEnumerator struct {
	*externglib.Object
}

var (
	_ SocketAddressEnumeratorer = (*SocketAddressEnumerator)(nil)
	_ gextras.Nativer           = (*SocketAddressEnumerator)(nil)
)

func wrapSocketAddressEnumerator(obj *externglib.Object) *SocketAddressEnumerator {
	return &SocketAddressEnumerator{
		Object: obj,
	}
}

func marshalSocketAddressEnumeratorer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketAddressEnumerator(obj), nil
}

// Next retrieves the next Address from enumerator. Note that this may block for
// some amount of time. (Eg, a Address may need to do a DNS lookup before it can
// return an address.) Use g_socket_address_enumerator_next_async() if you need
// to avoid blocking.
//
// If enumerator is expected to yield addresses, but for some reason is unable
// to (eg, because of a DNS error), then the first call to
// g_socket_address_enumerator_next() will return an appropriate error in
// *error. However, if the first call to g_socket_address_enumerator_next()
// succeeds, then any further internal errors (other than cancellable being
// triggered) will be ignored.
func (enumerator *SocketAddressEnumerator) Next(cancellable *Cancellable) (*SocketAddress, error) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GCancellable             // out
	var _cret *C.GSocketAddress           // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_socket_address_enumerator_next(_arg0, _arg1, &_cerr)

	var _socketAddress *SocketAddress // out
	var _goerr error                  // out

	_socketAddress = wrapSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketAddress, _goerr
}

// NextAsync: asynchronously retrieves the next Address from enumerator and then
// calls callback, which must call g_socket_address_enumerator_next_finish() to
// get the result.
//
// It is an error to call this multiple times before the previous callback has
// finished.
func (enumerator *SocketAddressEnumerator) NextAsync(cancellable *Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GCancellable             // out
	var _arg2 C.GAsyncReadyCallback       // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(gbox.AssignOnce(callback))

	C.g_socket_address_enumerator_next_async(_arg0, _arg1, _arg2, _arg3)
}

// NextFinish retrieves the result of a completed call to
// g_socket_address_enumerator_next_async(). See
// g_socket_address_enumerator_next() for more information about error handling.
func (enumerator *SocketAddressEnumerator) NextFinish(result AsyncResulter) (*SocketAddress, error) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GSocketAddress           // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_socket_address_enumerator_next_finish(_arg0, _arg1, &_cerr)

	var _socketAddress *SocketAddress // out
	var _goerr error                  // out

	_socketAddress = wrapSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketAddress, _goerr
}
