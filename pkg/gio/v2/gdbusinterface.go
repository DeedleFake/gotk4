// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
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
		{T: externglib.Type(C.g_dbus_interface_get_type()), F: marshalDBusInterface},
	})
}

// DBusInterfaceOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusInterfaceOverrider interface {
	// DupObject gets the BusObject that @interface_ belongs to, if any.
	DupObject() *DBusObjectIface
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by @interface_.
	Info() *DBusInterfaceInfo
	// SetObject sets the BusObject for @interface_ to @object.
	//
	// Note that @interface_ will hold a weak reference to @object.
	SetObject(object DBusObject)
}

// DBusInterface: the BusInterface type is the base type for D-Bus interfaces
// both on the service side (see BusInterfaceSkeleton) and client side (see
// BusProxy).
type DBusInterface interface {
	gextras.Objector

	// DupObject gets the BusObject that @interface_ belongs to, if any.
	DupObject() *DBusObjectIface
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by @interface_.
	Info() *DBusInterfaceInfo
	// SetObject sets the BusObject for @interface_ to @object.
	//
	// Note that @interface_ will hold a weak reference to @object.
	SetObject(object DBusObject)
}

// DBusInterfaceIface implements the DBusInterface interface.
type DBusInterfaceIface struct {
	*externglib.Object
}

var _ DBusInterface = (*DBusInterfaceIface)(nil)

func wrapDBusInterface(obj *externglib.Object) DBusInterface {
	return &DBusInterfaceIface{
		Object: obj,
	}
}

func marshalDBusInterface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusInterface(obj), nil
}

// DupObject gets the BusObject that @interface_ belongs to, if any.
func (interface_ *DBusInterfaceIface) DupObject() *DBusObjectIface {
	var _arg0 *C.GDBusInterface // out
	var _cret *C.GDBusObject    // in

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_dup_object(_arg0)

	var _dBusObject *DBusObjectIface // out

	_dBusObject = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DBusObjectIface)

	return _dBusObject
}

// Info gets D-Bus introspection information for the D-Bus interface implemented
// by @interface_.
func (interface_ *DBusInterfaceIface) Info() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterface     // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(interface_.Native()))

	_cret = C.g_dbus_interface_get_info(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// SetObject sets the BusObject for @interface_ to @object.
//
// Note that @interface_ will hold a weak reference to @object.
func (interface_ *DBusInterfaceIface) SetObject(object DBusObject) {
	var _arg0 *C.GDBusInterface // out
	var _arg1 *C.GDBusObject    // out

	_arg0 = (*C.GDBusInterface)(unsafe.Pointer(interface_.Native()))
	_arg1 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))

	C.g_dbus_interface_set_object(_arg0, _arg1)
}
