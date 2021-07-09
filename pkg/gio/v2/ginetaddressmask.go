// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_inet_address_mask_get_type()), F: marshalInetAddressMask},
	})
}

// InetAddressMask represents a range of IPv4 or IPv6 addresses described by a
// base address and a length indicating how many bits of the base address are
// relevant for matching purposes. These are often given in string form. Eg,
// "10.0.0.0/8", or "fe80::/10".
type InetAddressMask interface {
	gextras.Objector

	// Equal tests if @mask and @mask2 are the same mask.
	Equal(mask2 InetAddressMask) bool
	// Address gets @mask's base address
	Address() *InetAddressClass
	// Family gets the Family of @mask's address
	Family() SocketFamily
	// Length gets @mask's length
	Length() uint
	// Matches tests if @address falls within the range described by @mask.
	Matches(address InetAddress) bool
	// String converts @mask back to its corresponding string form.
	String() string
}

// InetAddressMaskClass implements the InetAddressMask interface.
type InetAddressMaskClass struct {
	*externglib.Object
	InitableInterface
}

var _ InetAddressMask = (*InetAddressMaskClass)(nil)

func wrapInetAddressMask(obj *externglib.Object) InetAddressMask {
	return &InetAddressMaskClass{
		Object: obj,
		InitableInterface: InitableInterface{
			Object: obj,
		},
	}
}

func marshalInetAddressMask(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInetAddressMask(obj), nil
}

// NewInetAddressMask creates a new AddressMask representing all addresses whose
// first @length bits match @addr.
func NewInetAddressMask(addr InetAddress, length uint) (*InetAddressMaskClass, error) {
	var _arg1 *C.GInetAddress     // out
	var _arg2 C.guint             // out
	var _cret *C.GInetAddressMask // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GInetAddress)(unsafe.Pointer((&InetAddress).Native()))
	_arg2 = C.guint(length)

	_cret = C.g_inet_address_mask_new(_arg1, _arg2, &_cerr)

	var _inetAddressMask *InetAddressMaskClass // out
	var _goerr error                           // out

	_inetAddressMask = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*InetAddressMaskClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _inetAddressMask, _goerr
}

// NewInetAddressMaskFromString parses @mask_string as an IP address and
// (optional) length, and creates a new AddressMask. The length, if present, is
// delimited by a "/". If it is not present, then the length is assumed to be
// the full length of the address.
func NewInetAddressMaskFromString(maskString string) (*InetAddressMaskClass, error) {
	var _arg1 *C.gchar            // out
	var _cret *C.GInetAddressMask // in
	var _cerr *C.GError           // in

	_arg1 = (*C.gchar)(C.CString(maskString))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_inet_address_mask_new_from_string(_arg1, &_cerr)

	var _inetAddressMask *InetAddressMaskClass // out
	var _goerr error                           // out

	_inetAddressMask = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*InetAddressMaskClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _inetAddressMask, _goerr
}

// Equal tests if @mask and @mask2 are the same mask.
func (m *InetAddressMaskClass) Equal(mask2 InetAddressMask) bool {
	var _arg0 *C.GInetAddressMask // out
	var _arg1 *C.GInetAddressMask // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer((&InetAddressMask).Native()))
	_arg1 = (*C.GInetAddressMask)(unsafe.Pointer((&InetAddressMask).Native()))

	_cret = C.g_inet_address_mask_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Address gets @mask's base address
func (m *InetAddressMaskClass) Address() *InetAddressClass {
	var _arg0 *C.GInetAddressMask // out
	var _cret *C.GInetAddress     // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer((&InetAddressMask).Native()))

	_cret = C.g_inet_address_mask_get_address(_arg0)

	var _inetAddress *InetAddressClass // out

	_inetAddress = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*InetAddressClass)

	return _inetAddress
}

// Family gets the Family of @mask's address
func (m *InetAddressMaskClass) Family() SocketFamily {
	var _arg0 *C.GInetAddressMask // out
	var _cret C.GSocketFamily     // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer((&InetAddressMask).Native()))

	_cret = C.g_inet_address_mask_get_family(_arg0)

	var _socketFamily SocketFamily // out

	_socketFamily = (SocketFamily)(C.GSocketFamily)

	return _socketFamily
}

// Length gets @mask's length
func (m *InetAddressMaskClass) Length() uint {
	var _arg0 *C.GInetAddressMask // out
	var _cret C.guint             // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer((&InetAddressMask).Native()))

	_cret = C.g_inet_address_mask_get_length(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Matches tests if @address falls within the range described by @mask.
func (m *InetAddressMaskClass) Matches(address InetAddress) bool {
	var _arg0 *C.GInetAddressMask // out
	var _arg1 *C.GInetAddress     // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer((&InetAddressMask).Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer((&InetAddress).Native()))

	_cret = C.g_inet_address_mask_matches(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String converts @mask back to its corresponding string form.
func (m *InetAddressMaskClass) String() string {
	var _arg0 *C.GInetAddressMask // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer((&InetAddressMask).Native()))

	_cret = C.g_inet_address_mask_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
