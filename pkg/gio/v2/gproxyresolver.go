// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_ProxyResolverInterface_is_supported(GProxyResolver*);
// extern gchar** _gotk4_gio2_ProxyResolverInterface_lookup(GProxyResolver*, gchar*, GCancellable*, GError**);
// extern gchar** _gotk4_gio2_ProxyResolverInterface_lookup_finish(GProxyResolver*, GAsyncResult*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// glib.Type values for gproxyresolver.go.
var GTypeProxyResolver = externglib.Type(C.g_proxy_resolver_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeProxyResolver, F: marshalProxyResolver},
	})
}

// PROXY_RESOLVER_EXTENSION_POINT_NAME: extension point for proxy resolving
// functionality. See [Extending GIO][extending-gio].
const PROXY_RESOLVER_EXTENSION_POINT_NAME = "gio-proxy-resolver"

// ProxyResolver provides synchronous and asynchronous network proxy resolution.
// Resolver is used within Client through the method
// g_socket_connectable_proxy_enumerate().
//
// Implementations of Resolver based on libproxy and GNOME settings can be found
// in glib-networking. GIO comes with an implementation for use inside Flatpak
// portals.
type ProxyResolver struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*ProxyResolver)(nil)
)

// ProxyResolverer describes ProxyResolver's interface methods.
type ProxyResolverer interface {
	externglib.Objector

	// IsSupported checks if resolver can be used on this system.
	IsSupported() bool
	// Lookup looks into the system proxy configuration to determine what proxy,
	// if any, to use to connect to uri.
	Lookup(ctx context.Context, uri string) ([]string, error)
	// LookupAsync asynchronous lookup of proxy.
	LookupAsync(ctx context.Context, uri string, callback AsyncReadyCallback)
	// LookupFinish: call this function to obtain the array of proxy URIs when
	// g_proxy_resolver_lookup_async() is complete.
	LookupFinish(result AsyncResulter) ([]string, error)
}

var _ ProxyResolverer = (*ProxyResolver)(nil)

func wrapProxyResolver(obj *externglib.Object) *ProxyResolver {
	return &ProxyResolver{
		Object: obj,
	}
}

func marshalProxyResolver(p uintptr) (interface{}, error) {
	return wrapProxyResolver(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// IsSupported checks if resolver can be used on this system. (This is used
// internally; g_proxy_resolver_get_default() will only return a proxy resolver
// that returns TRUE for this method.).
//
// The function returns the following values:
//
//    - ok: TRUE if resolver is supported.
//
func (resolver *ProxyResolver) IsSupported() bool {
	var _arg0 *C.GProxyResolver // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(externglib.InternObject(resolver).Native()))

	_cret = C.g_proxy_resolver_is_supported(_arg0)
	runtime.KeepAlive(resolver)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Lookup looks into the system proxy configuration to determine what proxy, if
// any, to use to connect to uri. The returned proxy URIs are of the form
// <protocol>://[user[:password]@]host:port or direct://, where <protocol> could
// be http, rtsp, socks or other proxying protocol.
//
// If you don't know what network protocol is being used on the socket, you
// should use none as the URI protocol. In this case, the resolver might still
// return a generic proxy type (such as SOCKS), but would not return
// protocol-specific proxy types (such as http).
//
// direct:// is used when no proxy is needed. Direct connection should not be
// attempted unless it is part of the returned array of proxies.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - uri: URI representing the destination to connect to.
//
// The function returns the following values:
//
//    - utf8s: a NULL-terminated array of proxy URIs. Must be freed with
//      g_strfreev().
//
func (resolver *ProxyResolver) Lookup(ctx context.Context, uri string) ([]string, error) {
	var _arg0 *C.GProxyResolver // out
	var _arg2 *C.GCancellable   // out
	var _arg1 *C.gchar          // out
	var _cret **C.gchar         // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(externglib.InternObject(resolver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_proxy_resolver_lookup(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)

	var _utf8s []string // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8s, _goerr
}

// LookupAsync asynchronous lookup of proxy. See g_proxy_resolver_lookup() for
// more details.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - uri: URI representing the destination to connect to.
//    - callback (optional) to call after resolution completes.
//
func (resolver *ProxyResolver) LookupAsync(ctx context.Context, uri string, callback AsyncReadyCallback) {
	var _arg0 *C.GProxyResolver     // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(externglib.InternObject(resolver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_proxy_resolver_lookup_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(callback)
}

// LookupFinish: call this function to obtain the array of proxy URIs when
// g_proxy_resolver_lookup_async() is complete. See g_proxy_resolver_lookup()
// for more details.
//
// The function takes the following parameters:
//
//    - result passed to your ReadyCallback.
//
// The function returns the following values:
//
//    - utf8s: a NULL-terminated array of proxy URIs. Must be freed with
//      g_strfreev().
//
func (resolver *ProxyResolver) LookupFinish(result AsyncResulter) ([]string, error) {
	var _arg0 *C.GProxyResolver // out
	var _arg1 *C.GAsyncResult   // out
	var _cret **C.gchar         // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GProxyResolver)(unsafe.Pointer(externglib.InternObject(resolver).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.g_proxy_resolver_lookup_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(result)

	var _utf8s []string // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8s, _goerr
}

// ProxyResolverGetDefault gets the default Resolver for the system.
//
// The function returns the following values:
//
//    - proxyResolver: default Resolver, which will be a dummy object if no proxy
//      resolver is available.
//
func ProxyResolverGetDefault() ProxyResolverer {
	var _cret *C.GProxyResolver // in

	_cret = C.g_proxy_resolver_get_default()

	var _proxyResolver ProxyResolverer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ProxyResolverer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(ProxyResolverer)
			return ok
		})
		rv, ok := casted.(ProxyResolverer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.ProxyResolverer")
		}
		_proxyResolver = rv
	}

	return _proxyResolver
}
