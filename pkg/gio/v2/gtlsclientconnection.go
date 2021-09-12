// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gio.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_client_connection_get_type()), F: marshalTLSClientConnectioner},
	})
}

// TLSClientConnectionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSClientConnectionOverrider interface {
	// CopySessionState: possibly copies session state from one connection to
	// another, for use in TLS session resumption. This is not normally needed,
	// but may be used when the same session needs to be used between different
	// endpoints, as is required by some protocols, such as FTP over TLS. source
	// should have already completed a handshake and, since TLS 1.3, it should
	// have been used to read data at least once. conn should not have completed
	// a handshake.
	//
	// It is not possible to know whether a call to this function will actually
	// do anything. Because session resumption is normally used only for
	// performance benefit, the TLS backend might not implement this function.
	// Even if implemented, it may not actually succeed in allowing conn to
	// resume source's TLS session, because the server may not have sent a
	// session resumption token to source, or it may refuse to accept the token
	// from conn. There is no way to know whether a call to this function is
	// actually successful.
	//
	// Using this function is not required to benefit from session resumption.
	// If the TLS backend supports session resumption, the session will be
	// resumed automatically if it is possible to do so without weakening the
	// privacy guarantees normally provided by TLS, without need to call this
	// function. For example, with TLS 1.3, a session ticket will be
	// automatically copied from any ClientConnection that has previously
	// received session tickets from the server, provided a ticket is available
	// that has not previously been used for session resumption, since session
	// ticket reuse would be a privacy weakness. Using this function causes the
	// ticket to be copied without regard for privacy considerations.
	CopySessionState(source TLSClientConnectioner)
}

// TLSClientConnection is the client-side subclass of Connection, representing a
// client-side TLS connection.
type TLSClientConnection struct {
	TLSConnection
}

// TLSClientConnectioner describes TLSClientConnection's abstract methods.
type TLSClientConnectioner interface {
	externglib.Objector

	// CopySessionState: possibly copies session state from one connection to
	// another, for use in TLS session resumption.
	CopySessionState(source TLSClientConnectioner)
	// ServerIdentity gets conn's expected server identity
	ServerIdentity() SocketConnectabler
	// UseSSL3: SSL 3.0 is no longer supported.
	UseSSL3() bool
	// ValidationFlags gets conn's validation flags
	ValidationFlags() TLSCertificateFlags
	// SetServerIdentity sets conn's expected server identity, which is used
	// both to tell servers on virtual hosts which certificate to present, and
	// also to let conn know what name to look for in the certificate when
	// performing G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
	SetServerIdentity(identity SocketConnectabler)
	// SetUseSSL3: since GLib 2.42.1, SSL 3.0 is no longer supported.
	SetUseSSL3(useSsl3 bool)
	// SetValidationFlags sets conn's validation flags, to override the default
	// set of checks performed when validating a server certificate.
	SetValidationFlags(flags TLSCertificateFlags)
}

var _ TLSClientConnectioner = (*TLSClientConnection)(nil)

func wrapTLSClientConnection(obj *externglib.Object) *TLSClientConnection {
	return &TLSClientConnection{
		TLSConnection: TLSConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTLSClientConnectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTLSClientConnection(obj), nil
}

// CopySessionState: possibly copies session state from one connection to
// another, for use in TLS session resumption. This is not normally needed, but
// may be used when the same session needs to be used between different
// endpoints, as is required by some protocols, such as FTP over TLS. source
// should have already completed a handshake and, since TLS 1.3, it should have
// been used to read data at least once. conn should not have completed a
// handshake.
//
// It is not possible to know whether a call to this function will actually do
// anything. Because session resumption is normally used only for performance
// benefit, the TLS backend might not implement this function. Even if
// implemented, it may not actually succeed in allowing conn to resume source's
// TLS session, because the server may not have sent a session resumption token
// to source, or it may refuse to accept the token from conn. There is no way to
// know whether a call to this function is actually successful.
//
// Using this function is not required to benefit from session resumption. If
// the TLS backend supports session resumption, the session will be resumed
// automatically if it is possible to do so without weakening the privacy
// guarantees normally provided by TLS, without need to call this function. For
// example, with TLS 1.3, a session ticket will be automatically copied from any
// ClientConnection that has previously received session tickets from the
// server, provided a ticket is available that has not previously been used for
// session resumption, since session ticket reuse would be a privacy weakness.
// Using this function causes the ticket to be copied without regard for privacy
// considerations.
func (conn *TLSClientConnection) CopySessionState(source TLSClientConnectioner) {
	var _arg0 *C.GTlsClientConnection // out
	var _arg1 *C.GTlsClientConnection // out

	_arg0 = (*C.GTlsClientConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = (*C.GTlsClientConnection)(unsafe.Pointer(source.Native()))

	C.g_tls_client_connection_copy_session_state(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(source)
}

// ServerIdentity gets conn's expected server identity
func (conn *TLSClientConnection) ServerIdentity() SocketConnectabler {
	var _arg0 *C.GTlsClientConnection // out
	var _cret *C.GSocketConnectable   // in

	_arg0 = (*C.GTlsClientConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_client_connection_get_server_identity(_arg0)
	runtime.KeepAlive(conn)

	var _socketConnectable SocketConnectabler // out

	if _cret != nil {
		_socketConnectable = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(SocketConnectabler)
	}

	return _socketConnectable
}

// UseSSL3: SSL 3.0 is no longer supported. See
// g_tls_client_connection_set_use_ssl3() for details.
//
// Deprecated: SSL 3.0 is insecure.
func (conn *TLSClientConnection) UseSSL3() bool {
	var _arg0 *C.GTlsClientConnection // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GTlsClientConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_client_connection_get_use_ssl3(_arg0)
	runtime.KeepAlive(conn)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ValidationFlags gets conn's validation flags
func (conn *TLSClientConnection) ValidationFlags() TLSCertificateFlags {
	var _arg0 *C.GTlsClientConnection // out
	var _cret C.GTlsCertificateFlags  // in

	_arg0 = (*C.GTlsClientConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tls_client_connection_get_validation_flags(_arg0)
	runtime.KeepAlive(conn)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}

// SetServerIdentity sets conn's expected server identity, which is used both to
// tell servers on virtual hosts which certificate to present, and also to let
// conn know what name to look for in the certificate when performing
// G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
func (conn *TLSClientConnection) SetServerIdentity(identity SocketConnectabler) {
	var _arg0 *C.GTlsClientConnection // out
	var _arg1 *C.GSocketConnectable   // out

	_arg0 = (*C.GTlsClientConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))

	C.g_tls_client_connection_set_server_identity(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(identity)
}

// SetUseSSL3: since GLib 2.42.1, SSL 3.0 is no longer supported.
//
// From GLib 2.42.1 through GLib 2.62, this function could be used to force use
// of TLS 1.0, the lowest-supported TLS protocol version at the time. In the
// past, this was needed to connect to broken TLS servers that exhibited
// protocol version intolerance. Such servers are no longer common, and using
// TLS 1.0 is no longer considered acceptable.
//
// Since GLib 2.64, this function does nothing.
//
// Deprecated: SSL 3.0 is insecure.
func (conn *TLSClientConnection) SetUseSSL3(useSsl3 bool) {
	var _arg0 *C.GTlsClientConnection // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GTlsClientConnection)(unsafe.Pointer(conn.Native()))
	if useSsl3 {
		_arg1 = C.TRUE
	}

	C.g_tls_client_connection_set_use_ssl3(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(useSsl3)
}

// SetValidationFlags sets conn's validation flags, to override the default set
// of checks performed when validating a server certificate. By default,
// G_TLS_CERTIFICATE_VALIDATE_ALL is used.
func (conn *TLSClientConnection) SetValidationFlags(flags TLSCertificateFlags) {
	var _arg0 *C.GTlsClientConnection // out
	var _arg1 C.GTlsCertificateFlags  // out

	_arg0 = (*C.GTlsClientConnection)(unsafe.Pointer(conn.Native()))
	_arg1 = C.GTlsCertificateFlags(flags)

	C.g_tls_client_connection_set_validation_flags(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(flags)
}

// NewTLSClientConnection creates a new ClientConnection wrapping base_io_stream
// (which must have pollable input and output streams) which is assumed to
// communicate with the server identified by server_identity.
//
// See the documentation for Connection:base-io-stream for restrictions on when
// application code can run operations on the base_io_stream after this function
// has returned.
func NewTLSClientConnection(baseIoStream IOStreamer, serverIdentity SocketConnectabler) (TLSClientConnectioner, error) {
	var _arg1 *C.GIOStream          // out
	var _arg2 *C.GSocketConnectable // out
	var _cret *C.GIOStream          // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(baseIoStream.Native()))
	if serverIdentity != nil {
		_arg2 = (*C.GSocketConnectable)(unsafe.Pointer(serverIdentity.Native()))
	}

	_cret = C.g_tls_client_connection_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(serverIdentity)

	var _tlsClientConnection TLSClientConnectioner // out
	var _goerr error                               // out

	_tlsClientConnection = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(TLSClientConnectioner)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsClientConnection, _goerr
}
