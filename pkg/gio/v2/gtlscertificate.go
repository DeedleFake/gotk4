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
		{T: externglib.Type(C.g_tls_certificate_get_type()), F: marshalTLSCertificate},
	})
}

// TLSCertificate: a certificate used for TLS authentication and encryption.
// This can represent either a certificate only (eg, the certificate received by
// a client from a server), or the combination of a certificate and a private
// key (which is needed when acting as a ServerConnection).
type TLSCertificate interface {
	gextras.Objector

	Issuer() TLSCertificate

	IsSameTLSCertificate(certTwo TLSCertificate) bool

	VerifyTLSCertificate(identity SocketConnectable, trustedCa TLSCertificate) TLSCertificateFlags
}

// tlsCertificate implements the TLSCertificate class.
type tlsCertificate struct {
	gextras.Objector
}

// WrapTLSCertificate wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSCertificate(obj *externglib.Object) TLSCertificate {
	return tlsCertificate{
		Objector: obj,
	}
}

func marshalTLSCertificate(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSCertificate(obj), nil
}

func NewTLSCertificateFromFile(file string) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(file))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_tls_certificate_new_from_file(_arg1, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func NewTLSCertificateFromFiles(certFile string, keyFile string) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(certFile))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(keyFile))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_certificate_new_from_files(_arg1, _arg2, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func NewTLSCertificateFromPem(data string, length int) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 C.gssize           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(data))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.g_tls_certificate_new_from_pem(_arg1, _arg2, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func NewTLSCertificateFromPkcs11Uris(pkcs11Uri string, privateKeyPkcs11Uri string) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(pkcs11Uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(privateKeyPkcs11Uri))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_certificate_new_from_pkcs11_uris(_arg1, _arg2, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func (c tlsCertificate) Issuer() TLSCertificate {
	var _arg0 *C.GTlsCertificate // out
	var _cret *C.GTlsCertificate // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(c.Native()))

	_cret = C.g_tls_certificate_get_issuer(_arg0)

	var _tlsCertificate TLSCertificate // out

	_tlsCertificate = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TLSCertificate)

	return _tlsCertificate
}

func (c tlsCertificate) IsSameTLSCertificate(certTwo TLSCertificate) bool {
	var _arg0 *C.GTlsCertificate // out
	var _arg1 *C.GTlsCertificate // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certTwo.Native()))

	_cret = C.g_tls_certificate_is_same(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c tlsCertificate) VerifyTLSCertificate(identity SocketConnectable, trustedCa TLSCertificate) TLSCertificateFlags {
	var _arg0 *C.GTlsCertificate     // out
	var _arg1 *C.GSocketConnectable  // out
	var _arg2 *C.GTlsCertificate     // out
	var _cret C.GTlsCertificateFlags // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(trustedCa.Native()))

	_cret = C.g_tls_certificate_verify(_arg0, _arg1, _arg2)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}
