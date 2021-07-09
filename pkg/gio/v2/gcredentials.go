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
		{T: externglib.Type(C.g_credentials_get_type()), F: marshalCredentials},
	})
}

// Credentials: the #GCredentials type is a reference-counted wrapper for native
// credentials. This information is typically used for identifying,
// authenticating and authorizing other processes.
//
// Some operating systems supports looking up the credentials of the remote peer
// of a communication endpoint - see e.g. g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving credentials
// over a Unix Domain Socket, see CredentialsMessage,
// g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a `struct ucred` - see the unix(7)
// man page for details. This corresponds to G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On Apple operating systems (including iOS, tvOS, and macOS), the native
// credential type is a `struct xucred`. This corresponds to
// G_CREDENTIALS_TYPE_APPLE_XUCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native credential type is
// a `struct cmsgcred`. This corresponds to G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a `struct unpcbid`. This corresponds
// to G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a `struct sockpeercred`. This
// corresponds to G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native credential
// type is a `ucred_t`. This corresponds to G_CREDENTIALS_TYPE_SOLARIS_UCRED.
type Credentials interface {
	gextras.Objector

	// UnixPid tries to get the UNIX process identifier from @credentials. This
	// method is only available on UNIX platforms.
	//
	// This operation can fail if #GCredentials is not supported on the OS or if
	// the native credentials type does not contain information about the UNIX
	// process ID (for example this is the case for
	// G_CREDENTIALS_TYPE_APPLE_XUCRED).
	UnixPid() (int, error)
	// UnixUser tries to get the UNIX user identifier from @credentials. This
	// method is only available on UNIX platforms.
	//
	// This operation can fail if #GCredentials is not supported on the OS or if
	// the native credentials type does not contain information about the UNIX
	// user.
	UnixUser() (uint, error)
	// IsSameUser checks if @credentials and @other_credentials is the same
	// user.
	//
	// This operation can fail if #GCredentials is not supported on the the OS.
	IsSameUser(otherCredentials Credentials) error
	// SetUnixUser tries to set the UNIX user identifier on @credentials. This
	// method is only available on UNIX platforms.
	//
	// This operation can fail if #GCredentials is not supported on the OS or if
	// the native credentials type does not contain information about the UNIX
	// user. It can also fail if the OS does not allow the use of "spoofed"
	// credentials.
	SetUnixUser(uid uint) error
	// String creates a human-readable textual representation of @credentials
	// that can be used in logging and debug messages. The format of the
	// returned string may change in future GLib release.
	String() string
}

// CredentialsClass implements the Credentials interface.
type CredentialsClass struct {
	*externglib.Object
}

var _ Credentials = (*CredentialsClass)(nil)

func wrapCredentials(obj *externglib.Object) Credentials {
	return &CredentialsClass{
		Object: obj,
	}
}

func marshalCredentials(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCredentials(obj), nil
}

// NewCredentials creates a new #GCredentials object with credentials matching
// the the current process.
func NewCredentials() *CredentialsClass {
	var _cret *C.GCredentials // in

	_cret = C.g_credentials_new()

	var _credentials *CredentialsClass // out

	_credentials = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*CredentialsClass)

	return _credentials
}

// UnixPid tries to get the UNIX process identifier from @credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if the
// native credentials type does not contain information about the UNIX process
// ID (for example this is the case for G_CREDENTIALS_TYPE_APPLE_XUCRED).
func (c *CredentialsClass) UnixPid() (int, error) {
	var _arg0 *C.GCredentials // out
	var _cret C.pid_t         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer((&Credentials).Native()))

	_cret = C.g_credentials_get_unix_pid(_arg0, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

// UnixUser tries to get the UNIX user identifier from @credentials. This method
// is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if the
// native credentials type does not contain information about the UNIX user.
func (c *CredentialsClass) UnixUser() (uint, error) {
	var _arg0 *C.GCredentials // out
	var _cret C.uid_t         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer((&Credentials).Native()))

	_cret = C.g_credentials_get_unix_user(_arg0, &_cerr)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint, _goerr
}

// IsSameUser checks if @credentials and @other_credentials is the same user.
//
// This operation can fail if #GCredentials is not supported on the the OS.
func (c *CredentialsClass) IsSameUser(otherCredentials Credentials) error {
	var _arg0 *C.GCredentials // out
	var _arg1 *C.GCredentials // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer((&Credentials).Native()))
	_arg1 = (*C.GCredentials)(unsafe.Pointer((&Credentials).Native()))

	C.g_credentials_is_same_user(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetUnixUser tries to set the UNIX user identifier on @credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if the
// native credentials type does not contain information about the UNIX user. It
// can also fail if the OS does not allow the use of "spoofed" credentials.
func (c *CredentialsClass) SetUnixUser(uid uint) error {
	var _arg0 *C.GCredentials // out
	var _arg1 C.uid_t         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer((&Credentials).Native()))
	_arg1 = C.uid_t(uid)

	C.g_credentials_set_unix_user(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// String creates a human-readable textual representation of @credentials that
// can be used in logging and debug messages. The format of the returned string
// may change in future GLib release.
func (c *CredentialsClass) String() string {
	var _arg0 *C.GCredentials // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer((&Credentials).Native()))

	_cret = C.g_credentials_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
