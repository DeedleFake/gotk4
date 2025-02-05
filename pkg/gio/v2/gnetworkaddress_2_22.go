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
import "C"

// NetworkAddressParse creates a new Connectable for connecting to the given
// hostname and port. May fail and return NULL in case parsing host_and_port
// fails.
//
// host_and_port may be in any of a number of recognised formats; an IPv6
// address, an IPv4 address, or a domain name (in which case a DNS lookup is
// performed). Quoting with [] is supported for all address types. A port
// override may be specified in the usual way with a colon.
//
// If no port is specified in host_and_port then default_port will be used as
// the port number to connect to.
//
// In general, host_and_port is expected to be provided by the user (allowing
// them to give the hostname, and a port override if necessary) and default_port
// is expected to be provided by the application.
//
// (The port component of host_and_port can also be specified as a service name
// rather than as a numeric port, but this functionality is deprecated, because
// it depends on the contents of /etc/services, which is generally quite sparse
// on platforms other than Linux.).
//
// The function takes the following parameters:
//
//    - hostAndPort: hostname and optionally a port.
//    - defaultPort: default port if not in host_and_port.
//
// The function returns the following values:
//
//    - networkAddress: new Address, or NULL on error.
//
func NetworkAddressParse(hostAndPort string, defaultPort uint16) (*NetworkAddress, error) {
	var _arg1 *C.gchar              // out
	var _arg2 C.guint16             // out
	var _cret *C.GSocketConnectable // in
	var _cerr *C.GError             // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostAndPort)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)

	_cret = C.g_network_address_parse(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(hostAndPort)
	runtime.KeepAlive(defaultPort)

	var _networkAddress *NetworkAddress // out
	var _goerr error                    // out

	_networkAddress = wrapNetworkAddress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _networkAddress, _goerr
}
