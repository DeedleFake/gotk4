// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_object_skeleton_get_type()), F: marshalDBusObjectSkeletoner},
	})
}

// DBusObjectSkeletonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectSkeletonOverrider interface {
	AuthorizeMethod(interface_ DBusInterfaceSkeletoner, invocation *DBusMethodInvocation) bool
}

// DBusObjectSkeleton instance is essentially a group of D-Bus interfaces. The
// set of exported interfaces on the object may be dynamic and change at
// runtime.
//
// This type is intended to be used with BusObjectManager.
type DBusObjectSkeleton struct {
	*externglib.Object

	DBusObject
}

var _ gextras.Nativer = (*DBusObjectSkeleton)(nil)

func wrapDBusObjectSkeleton(obj *externglib.Object) *DBusObjectSkeleton {
	return &DBusObjectSkeleton{
		Object: obj,
		DBusObject: DBusObject{
			Object: obj,
		},
	}
}

func marshalDBusObjectSkeletoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObjectSkeleton(obj), nil
}

// NewDBusObjectSkeleton creates a new BusObjectSkeleton.
func NewDBusObjectSkeleton(objectPath string) *DBusObjectSkeleton {
	var _arg1 *C.gchar               // out
	var _cret *C.GDBusObjectSkeleton // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))

	_cret = C.g_dbus_object_skeleton_new(_arg1)

	var _dBusObjectSkeleton *DBusObjectSkeleton // out

	_dBusObjectSkeleton = wrapDBusObjectSkeleton(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusObjectSkeleton
}

// AddInterface adds interface_ to object.
//
// If object already contains a BusInterfaceSkeleton with the same interface
// name, it is removed before interface_ is added.
//
// Note that object takes its own reference on interface_ and holds it until
// removed.
func (object *DBusObjectSkeleton) AddInterface(interface_ DBusInterfaceSkeletoner) {
	var _arg0 *C.GDBusObjectSkeleton    // out
	var _arg1 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer((interface_).(gextras.Nativer).Native()))

	C.g_dbus_object_skeleton_add_interface(_arg0, _arg1)
}

// Flush: this method simply calls g_dbus_interface_skeleton_flush() on all
// interfaces belonging to object. See that method for when flushing is useful.
func (object *DBusObjectSkeleton) Flush() {
	var _arg0 *C.GDBusObjectSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_skeleton_flush(_arg0)
}

// RemoveInterface removes interface_ from object.
func (object *DBusObjectSkeleton) RemoveInterface(interface_ DBusInterfaceSkeletoner) {
	var _arg0 *C.GDBusObjectSkeleton    // out
	var _arg1 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer((interface_).(gextras.Nativer).Native()))

	C.g_dbus_object_skeleton_remove_interface(_arg0, _arg1)
}

// RemoveInterfaceByName removes the BusInterface with interface_name from
// object.
//
// If no D-Bus interface of the given interface exists, this function does
// nothing.
func (object *DBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	var _arg0 *C.GDBusObjectSkeleton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))

	C.g_dbus_object_skeleton_remove_interface_by_name(_arg0, _arg1)
}

// SetObjectPath sets the object path for object.
func (object *DBusObjectSkeleton) SetObjectPath(objectPath string) {
	var _arg0 *C.GDBusObjectSkeleton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))

	C.g_dbus_object_skeleton_set_object_path(_arg0, _arg1)
}
