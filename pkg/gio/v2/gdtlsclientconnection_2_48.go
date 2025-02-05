// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeDTLSClientConnection = coreglib.Type(C.g_dtls_client_connection_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDTLSClientConnection, F: marshalDTLSClientConnection},
	})
}

// DTLSClientConnectionOverrider contains methods that are overridable.
type DTLSClientConnectionOverrider interface {
}

// DTLSClientConnection is the client-side subclass of Connection, representing
// a client-side DTLS connection.
//
// DTLSClientConnection wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DTLSClientConnection struct {
	_ [0]func() // equal guard
	DTLSConnection
}

var ()

// DTLSClientConnectioner describes DTLSClientConnection's interface methods.
type DTLSClientConnectioner interface {
	coreglib.Objector

	// ServerIdentity gets conn's expected server identity.
	ServerIdentity() *SocketConnectable
	// ValidationFlags gets conn's validation flags.
	ValidationFlags() TLSCertificateFlags
	// SetServerIdentity sets conn's expected server identity, which is used
	// both to tell servers on virtual hosts which certificate to present, and
	// also to let conn know what name to look for in the certificate when
	// performing G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
	SetServerIdentity(identity SocketConnectabler)
	// SetValidationFlags sets conn's validation flags, to override the default
	// set of checks performed when validating a server certificate.
	SetValidationFlags(flags TLSCertificateFlags)
}

var _ DTLSClientConnectioner = (*DTLSClientConnection)(nil)

func ifaceInitDTLSClientConnectioner(gifacePtr, data C.gpointer) {
}

func wrapDTLSClientConnection(obj *coreglib.Object) *DTLSClientConnection {
	return &DTLSClientConnection{
		DTLSConnection: DTLSConnection{
			DatagramBased: DatagramBased{
				Object: obj,
			},
		},
	}
}

func marshalDTLSClientConnection(p uintptr) (interface{}, error) {
	return wrapDTLSClientConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ServerIdentity gets conn's expected server identity.
//
// The function returns the following values:
//
//    - socketConnectable describing the expected server identity, or NULL if the
//      expected identity is not known.
//
func (conn *DTLSClientConnection) ServerIdentity() *SocketConnectable {
	var _arg0 *C.GDtlsClientConnection // out
	var _cret *C.GSocketConnectable    // in

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_cret = C.g_dtls_client_connection_get_server_identity(_arg0)
	runtime.KeepAlive(conn)

	var _socketConnectable *SocketConnectable // out

	_socketConnectable = wrapSocketConnectable(coreglib.Take(unsafe.Pointer(_cret)))

	return _socketConnectable
}

// ValidationFlags gets conn's validation flags.
//
// The function returns the following values:
//
//    - tlsCertificateFlags: validation flags.
//
func (conn *DTLSClientConnection) ValidationFlags() TLSCertificateFlags {
	var _arg0 *C.GDtlsClientConnection // out
	var _cret C.GTlsCertificateFlags   // in

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(coreglib.InternObject(conn).Native()))

	_cret = C.g_dtls_client_connection_get_validation_flags(_arg0)
	runtime.KeepAlive(conn)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}

// SetServerIdentity sets conn's expected server identity, which is used both to
// tell servers on virtual hosts which certificate to present, and also to let
// conn know what name to look for in the certificate when performing
// G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
//
// The function takes the following parameters:
//
//    - identity describing the expected server identity.
//
func (conn *DTLSClientConnection) SetServerIdentity(identity SocketConnectabler) {
	var _arg0 *C.GDtlsClientConnection // out
	var _arg1 *C.GSocketConnectable    // out

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(identity).Native()))

	C.g_dtls_client_connection_set_server_identity(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(identity)
}

// SetValidationFlags sets conn's validation flags, to override the default set
// of checks performed when validating a server certificate. By default,
// G_TLS_CERTIFICATE_VALIDATE_ALL is used.
//
// The function takes the following parameters:
//
//    - flags to use.
//
func (conn *DTLSClientConnection) SetValidationFlags(flags TLSCertificateFlags) {
	var _arg0 *C.GDtlsClientConnection // out
	var _arg1 C.GTlsCertificateFlags   // out

	_arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(coreglib.InternObject(conn).Native()))
	_arg1 = C.GTlsCertificateFlags(flags)

	C.g_dtls_client_connection_set_validation_flags(_arg0, _arg1)
	runtime.KeepAlive(conn)
	runtime.KeepAlive(flags)
}

// NewDTLSClientConnection creates a new ClientConnection wrapping base_socket
// which is assumed to communicate with the server identified by
// server_identity.
//
// The function takes the following parameters:
//
//    - baseSocket to wrap.
//    - serverIdentity (optional): expected identity of the server.
//
// The function returns the following values:
//
//    - dtlsClientConnection: new ClientConnection, or NULL on error.
//
func NewDTLSClientConnection(baseSocket DatagramBasedder, serverIdentity SocketConnectabler) (*DTLSClientConnection, error) {
	var _arg1 *C.GDatagramBased     // out
	var _arg2 *C.GSocketConnectable // out
	var _cret *C.GDatagramBased     // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(baseSocket).Native()))
	if serverIdentity != nil {
		_arg2 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(serverIdentity).Native()))
	}

	_cret = C.g_dtls_client_connection_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(baseSocket)
	runtime.KeepAlive(serverIdentity)

	var _dtlsClientConnection *DTLSClientConnection // out
	var _goerr error                                // out

	_dtlsClientConnection = wrapDTLSClientConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dtlsClientConnection, _goerr
}

// DTLSClientConnectionInterface: vtable for a ClientConnection implementation.
//
// An instance of this type is always passed by reference.
type DTLSClientConnectionInterface struct {
	*dtlsClientConnectionInterface
}

// dtlsClientConnectionInterface is the struct that's finalized.
type dtlsClientConnectionInterface struct {
	native *C.GDtlsClientConnectionInterface
}
