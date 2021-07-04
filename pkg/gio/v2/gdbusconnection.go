// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// DBusInterfaceGetPropertyFunc: the type of the @get_property function in
// BusInterfaceVTable.
type DBusInterfaceGetPropertyFunc func(connection DBusConnection, sender string, objectPath string, interfaceName string, propertyName string, err *error, variant *glib.Variant)

//export gotk4_DBusInterfaceGetPropertyFunc
func gotk4_DBusInterfaceGetPropertyFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 **C.GError, arg6 C.gpointer) *C.GVariant {
	v := box.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection // out
	var sender string             // out
	var objectPath string         // out
	var interfaceName string      // out
	var propertyName string       // out
	var err *error                // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	sender = C.GoString(arg1)
	objectPath = C.GoString(arg2)
	interfaceName = C.GoString(arg3)
	propertyName = C.GoString(arg4)
	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *arg5

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		err = refTmpOut
	}

	fn := v.(DBusInterfaceGetPropertyFunc)
	variant := fn(connection, sender, objectPath, interfaceName, propertyName, err)

	var cret *C.GVariant // out

	cret = (*C.GVariant)(unsafe.Pointer(variant.Native()))

	return cret
}

// DBusInterfaceMethodCallFunc: the type of the @method_call function in
// BusInterfaceVTable.
type DBusInterfaceMethodCallFunc func(connection DBusConnection, sender string, objectPath string, interfaceName string, methodName string, parameters *glib.Variant, invocation DBusMethodInvocation)

//export gotk4_DBusInterfaceMethodCallFunc
func gotk4_DBusInterfaceMethodCallFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.GVariant, arg6 *C.GDBusMethodInvocation, arg7 C.gpointer) {
	v := box.Get(uintptr(arg7))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection       // out
	var sender string                   // out
	var objectPath string               // out
	var interfaceName string            // out
	var methodName string               // out
	var parameters *glib.Variant        // out
	var invocation DBusMethodInvocation // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	sender = C.GoString(arg1)
	objectPath = C.GoString(arg2)
	interfaceName = C.GoString(arg3)
	methodName = C.GoString(arg4)
	parameters = (*glib.Variant)(unsafe.Pointer(arg5))
	invocation = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(arg6))).(DBusMethodInvocation)

	fn := v.(DBusInterfaceMethodCallFunc)
	fn(connection, sender, objectPath, interfaceName, methodName, parameters, invocation)
}

// DBusInterfaceSetPropertyFunc: the type of the @set_property function in
// BusInterfaceVTable.
type DBusInterfaceSetPropertyFunc func(connection DBusConnection, sender string, objectPath string, interfaceName string, propertyName string, value *glib.Variant, err *error, ok bool)

//export gotk4_DBusInterfaceSetPropertyFunc
func gotk4_DBusInterfaceSetPropertyFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.GVariant, arg6 **C.GError, arg7 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg7))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection // out
	var sender string             // out
	var objectPath string         // out
	var interfaceName string      // out
	var propertyName string       // out
	var value *glib.Variant       // out
	var err *error                // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	sender = C.GoString(arg1)
	objectPath = C.GoString(arg2)
	interfaceName = C.GoString(arg3)
	propertyName = C.GoString(arg4)
	value = (*glib.Variant)(unsafe.Pointer(arg5))
	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *arg6

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		err = refTmpOut
	}

	fn := v.(DBusInterfaceSetPropertyFunc)
	ok := fn(connection, sender, objectPath, interfaceName, propertyName, value, err)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// DBusMessageFilterFunction: signature for function used in
// g_dbus_connection_add_filter().
//
// A filter function is passed a BusMessage and expected to return a BusMessage
// too. Passive filter functions that don't modify the message can simply return
// the @message object:
//
//    static GDBusMessage *
//    passive_filter (GDBusConnection *connection
//                    GDBusMessage    *message,
//                    gboolean         incoming,
//                    gpointer         user_data)
//    {
//      // inspect @message
//      return message;
//    }
//
// Filter functions that wants to drop a message can simply return nil:
//
//    static GDBusMessage *
//    drop_filter (GDBusConnection *connection
//                 GDBusMessage    *message,
//                 gboolean         incoming,
//                 gpointer         user_data)
//    {
//      if (should_drop_message)
//        {
//          g_object_unref (message);
//          message = NULL;
//        }
//      return message;
//    }
//
// Finally, a filter function may modify a message by copying it:
//
//    static GDBusMessage *
//    modifying_filter (GDBusConnection *connection
//                      GDBusMessage    *message,
//                      gboolean         incoming,
//                      gpointer         user_data)
//    {
//      GDBusMessage *copy;
//      GError *error;
//
//      error = NULL;
//      copy = g_dbus_message_copy (message, &error);
//      // handle @error being set
//      g_object_unref (message);
//
//      // modify @copy
//
//      return copy;
//    }
//
// If the returned BusMessage is different from @message and cannot be sent on
// @connection (it could use features, such as file descriptors, not compatible
// with @connection), then a warning is logged to standard error. Applications
// can check this ahead of time using g_dbus_message_to_blob() passing a
// BusCapabilityFlags value obtained from @connection.
type DBusMessageFilterFunction func(connection DBusConnection, message DBusMessage, incoming bool, dBusMessage DBusMessage)

//export gotk4_DBusMessageFilterFunction
func gotk4_DBusMessageFilterFunction(arg0 *C.GDBusConnection, arg1 *C.GDBusMessage, arg2 C.gboolean, arg3 C.gpointer) *C.GDBusMessage {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection // out
	var message DBusMessage       // out
	var incoming bool             // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	message = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(arg1))).(DBusMessage)
	if arg2 != 0 {
		incoming = true
	}

	fn := v.(DBusMessageFilterFunction)
	dBusMessage := fn(connection, message, incoming)

	var cret *C.GDBusMessage // out

	cret = (*C.GDBusMessage)(unsafe.Pointer(dBusMessage.Native()))

	return cret
}

// DBusSignalCallback: signature for callback function used in
// g_dbus_connection_signal_subscribe().
type DBusSignalCallback func(connection DBusConnection, senderName string, objectPath string, interfaceName string, signalName string, parameters *glib.Variant)

//export gotk4_DBusSignalCallback
func gotk4_DBusSignalCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.GVariant, arg6 C.gpointer) {
	v := box.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection // out
	var senderName string         // out
	var objectPath string         // out
	var interfaceName string      // out
	var signalName string         // out
	var parameters *glib.Variant  // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	senderName = C.GoString(arg1)
	objectPath = C.GoString(arg2)
	interfaceName = C.GoString(arg3)
	signalName = C.GoString(arg4)
	parameters = (*glib.Variant)(unsafe.Pointer(arg5))

	fn := v.(DBusSignalCallback)
	fn(connection, senderName, objectPath, interfaceName, signalName, parameters)
}

// DBusSubtreeDispatchFunc: the type of the @dispatch function in
// BusSubtreeVTable.
//
// Subtrees are flat. @node, if non-nil, is always exactly one segment of the
// object path (ie: it never contains a slash).
type DBusSubtreeDispatchFunc func(connection DBusConnection, sender string, objectPath string, interfaceName string, node string, outUserData *interface{}, dBusInterfaceVTable *DBusInterfaceVTable)

//export gotk4_DBusSubtreeDispatchFunc
func gotk4_DBusSubtreeDispatchFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 *C.gchar, arg5 *C.gpointer, arg6 C.gpointer) *C.GDBusInterfaceVTable {
	v := box.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection // out
	var sender string             // out
	var objectPath string         // out
	var interfaceName string      // out
	var node string               // out
	var outUserData *interface{}  // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	sender = C.GoString(arg1)
	objectPath = C.GoString(arg2)
	interfaceName = C.GoString(arg3)
	node = C.GoString(arg4)
	outUserData = box.Get(uintptr(arg5))

	fn := v.(DBusSubtreeDispatchFunc)
	dBusInterfaceVTable := fn(connection, sender, objectPath, interfaceName, node, outUserData)

	var cret *C.GDBusInterfaceVTable // out

	cret = (*C.GDBusInterfaceVTable)(unsafe.Pointer(dBusInterfaceVTable.Native()))

	return cret
}

// DBusSubtreeEnumerateFunc: the type of the @enumerate function in
// BusSubtreeVTable.
//
// This function is called when generating introspection data and also when
// preparing to dispatch incoming messages in the event that the
// G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES flag is not specified
// (ie: to verify that the object path is valid).
//
// Hierarchies are not supported; the items that you return should not contain
// the `/` character.
//
// The return value will be freed with g_strfreev().
type DBusSubtreeEnumerateFunc func(connection DBusConnection, sender string, objectPath string, utf8s []string)

//export gotk4_DBusSubtreeEnumerateFunc
func gotk4_DBusSubtreeEnumerateFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 C.gpointer) **C.gchar {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection // out
	var sender string             // out
	var objectPath string         // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	sender = C.GoString(arg1)
	objectPath = C.GoString(arg2)

	fn := v.(DBusSubtreeEnumerateFunc)
	utf8s := fn(connection, sender, objectPath)

	var cret **C.gchar

	cret = (**C.gchar)(C.malloc(C.ulong(len(utf8s)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	{
		out := unsafe.Slice(cret, len(utf8s))
		for i := range utf8s {
			out[i] = (*C.gchar)(C.CString(utf8s[i]))
		}
	}

	return cret
}

// DBusSubtreeIntrospectFunc: the type of the @introspect function in
// BusSubtreeVTable.
//
// Subtrees are flat. @node, if non-nil, is always exactly one segment of the
// object path (ie: it never contains a slash).
//
// This function should return nil to indicate that there is no object at this
// node.
//
// If this function returns non-nil, the return value is expected to be a
// nil-terminated array of pointers to BusInterfaceInfo structures describing
// the interfaces implemented by @node. This array will have
// g_dbus_interface_info_unref() called on each item before being freed with
// g_free().
//
// The difference between returning nil and an array containing zero items is
// that the standard DBus interfaces will returned to the remote introspector in
// the empty array case, but not in the nil case.
type DBusSubtreeIntrospectFunc func(connection DBusConnection, sender string, objectPath string, node string, dBusInterfaceInfos []*DBusInterfaceInfo)

//export gotk4_DBusSubtreeIntrospectFunc
func gotk4_DBusSubtreeIntrospectFunc(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 *C.gchar, arg4 C.gpointer) **C.GDBusInterfaceInfo {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var connection DBusConnection // out
	var sender string             // out
	var objectPath string         // out
	var node string               // out

	connection = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusConnection)
	sender = C.GoString(arg1)
	objectPath = C.GoString(arg2)
	node = C.GoString(arg3)

	fn := v.(DBusSubtreeIntrospectFunc)
	dBusInterfaceInfos := fn(connection, sender, objectPath, node)

	var cret **C.GDBusInterfaceInfo

	cret = (**C.GDBusInterfaceInfo)(C.malloc(C.ulong(len(dBusInterfaceInfos)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	{
		out := unsafe.Slice(cret, len(dBusInterfaceInfos))
		for i := range dBusInterfaceInfos {
			out[i] = (*C.GDBusInterfaceInfo)(unsafe.Pointer(dBusInterfaceInfos[i].Native()))
		}
	}

	return cret
}

// BusGet: asynchronously connects to the message bus specified by @bus_type.
//
// When the operation is finished, @callback will be invoked. You can then call
// g_bus_get_finish() to get the result of the operation.
//
// This is an asynchronous failable function. See g_bus_get_sync() for the
// synchronous version.
func BusGet(busType BusType, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg1 C.GBusType            // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg1 = C.GBusType(busType)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_bus_get(_arg1, _arg2, _arg3, _arg4)
}

// BusGetFinish finishes an operation started with g_bus_get().
//
// The returned object is a singleton, that is, shared with other callers of
// g_bus_get() and g_bus_get_sync() for @bus_type. In the event that you need a
// private message bus connection, use g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned BusConnection object will (usually) have the
// BusConnection:exit-on-close property set to true.
func BusGetFinish(res AsyncResult) (DBusConnection, error) {
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GDBusConnection // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	_cret = C.g_bus_get_finish(_arg1, &_cerr)

	var _dBusConnection DBusConnection // out
	var _goerr error                   // out

	_dBusConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(DBusConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusConnection, _goerr
}

// BusGetSync: synchronously connects to the message bus specified by @bus_type.
// Note that the returned object may shared with other callers, e.g. if two
// separate parts of a process calls this function with the same @bus_type, they
// will share the same object.
//
// This is a synchronous failable function. See g_bus_get() and
// g_bus_get_finish() for the asynchronous version.
//
// The returned object is a singleton, that is, shared with other callers of
// g_bus_get() and g_bus_get_sync() for @bus_type. In the event that you need a
// private message bus connection, use g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned BusConnection object will (usually) have the
// BusConnection:exit-on-close property set to true.
func BusGetSync(busType BusType, cancellable Cancellable) (DBusConnection, error) {
	var _arg1 C.GBusType         // out
	var _arg2 *C.GCancellable    // out
	var _cret *C.GDBusConnection // in
	var _cerr *C.GError          // in

	_arg1 = C.GBusType(busType)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_bus_get_sync(_arg1, _arg2, &_cerr)

	var _dBusConnection DBusConnection // out
	var _goerr error                   // out

	_dBusConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(DBusConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusConnection, _goerr
}

// DBusInterfaceVTable: virtual table for handling properties and method calls
// for a D-Bus interface.
//
// Since 2.38, if you want to handle getting/setting D-Bus properties
// asynchronously, give nil as your get_property() or set_property() function.
// The D-Bus call will be directed to your @method_call function, with the
// provided @interface_name set to "org.freedesktop.DBus.Properties".
//
// Ownership of the BusMethodInvocation object passed to the method_call()
// function is transferred to your handler; you must call one of the methods of
// BusMethodInvocation to return a reply (possibly empty), or an error. These
// functions also take ownership of the passed-in invocation object, so unless
// the invocation object has otherwise been referenced, it will be then be
// freed. Calling one of these functions may be done within your method_call()
// implementation but it also can be done at a later point to handle the method
// asynchronously.
//
// The usual checks on the validity of the calls is performed. For `Get` calls,
// an error is automatically returned if the property does not exist or the
// permissions do not allow access. The same checks are performed for `Set`
// calls, and the provided value is also checked for being the correct type.
//
// For both `Get` and `Set` calls, the BusMethodInvocation passed to the
// @method_call handler can be queried with
// g_dbus_method_invocation_get_property_info() to get a pointer to the
// BusPropertyInfo of the property.
//
// If you have readable properties specified in your interface info, you must
// ensure that you either provide a non-nil @get_property() function or provide
// implementations of both the `Get` and `GetAll` methods on
// org.freedesktop.DBus.Properties interface in your @method_call function. Note
// that the required return type of the `Get` call is `(v)`, not the type of the
// property. `GetAll` expects a return value of type `a{sv}`.
//
// If you have writable properties specified in your interface info, you must
// ensure that you either provide a non-nil @set_property() function or provide
// an implementation of the `Set` call. If implementing the call, you must
// return the value of type G_VARIANT_TYPE_UNIT.
type DBusInterfaceVTable C.GDBusInterfaceVTable

// WrapDBusInterfaceVTable wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusInterfaceVTable(ptr unsafe.Pointer) *DBusInterfaceVTable {
	return (*DBusInterfaceVTable)(ptr)
}

// Native returns the underlying C source pointer.
func (d *DBusInterfaceVTable) Native() unsafe.Pointer {
	return unsafe.Pointer(d)
}

// DBusSubtreeVTable: virtual table for handling subtrees registered with
// g_dbus_connection_register_subtree().
type DBusSubtreeVTable C.GDBusSubtreeVTable

// WrapDBusSubtreeVTable wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusSubtreeVTable(ptr unsafe.Pointer) *DBusSubtreeVTable {
	return (*DBusSubtreeVTable)(ptr)
}

// Native returns the underlying C source pointer.
func (d *DBusSubtreeVTable) Native() unsafe.Pointer {
	return unsafe.Pointer(d)
}
