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
		{T: externglib.Type(C.g_dbus_object_skeleton_get_type()), F: marshalDBusObjectSkeletonner},
	})
}

// DBusObjectSkeletonnerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectSkeletonnerOverrider interface {
	AuthorizeMethod(interface_ DBusInterfaceSkeletonner, invocation DBusMethodInvocationer) bool
}

// DBusObjectSkeletonner describes DBusObjectSkeleton's methods.
type DBusObjectSkeletonner interface {
	gextras.Objector

	AddInterface(interface_ DBusInterfaceSkeletonner)
	Flush()
	RemoveInterface(interface_ DBusInterfaceSkeletonner)
	RemoveInterfaceByName(interfaceName string)
	SetObjectPath(objectPath string)
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

var _ DBusObjectSkeletonner = (*DBusObjectSkeleton)(nil)

func wrapDBusObjectSkeletonner(obj *externglib.Object) DBusObjectSkeletonner {
	return &DBusObjectSkeleton{
		Object: obj,
		DBusObject: DBusObject{
			Object: obj,
		},
	}
}

func marshalDBusObjectSkeletonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObjectSkeletonner(obj), nil
}

// NewDBusObjectSkeleton creates a new BusObjectSkeleton.
func NewDBusObjectSkeleton(objectPath string) *DBusObjectSkeleton {
	var _arg1 *C.gchar               // out
	var _cret *C.GDBusObjectSkeleton // in

	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_skeleton_new(_arg1)

	var _dBusObjectSkeleton *DBusObjectSkeleton // out

	_dBusObjectSkeleton = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DBusObjectSkeleton)

	return _dBusObjectSkeleton
}

// AddInterface adds @interface_ to @object.
//
// If @object already contains a BusInterfaceSkeleton with the same interface
// name, it is removed before @interface_ is added.
//
// Note that @object takes its own reference on @interface_ and holds it until
// removed.
func (object *DBusObjectSkeleton) AddInterface(interface_ DBusInterfaceSkeletonner) {
	var _arg0 *C.GDBusObjectSkeleton    // out
	var _arg1 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_object_skeleton_add_interface(_arg0, _arg1)
}

// Flush: this method simply calls g_dbus_interface_skeleton_flush() on all
// interfaces belonging to @object. See that method for when flushing is useful.
func (object *DBusObjectSkeleton) Flush() {
	var _arg0 *C.GDBusObjectSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_skeleton_flush(_arg0)
}

// RemoveInterface removes @interface_ from @object.
func (object *DBusObjectSkeleton) RemoveInterface(interface_ DBusInterfaceSkeletonner) {
	var _arg0 *C.GDBusObjectSkeleton    // out
	var _arg1 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_object_skeleton_remove_interface(_arg0, _arg1)
}

// RemoveInterfaceByName removes the BusInterface with @interface_name from
// @object.
//
// If no D-Bus interface of the given interface exists, this function does
// nothing.
func (object *DBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	var _arg0 *C.GDBusObjectSkeleton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_object_skeleton_remove_interface_by_name(_arg0, _arg1)
}

// SetObjectPath sets the object path for @object.
func (object *DBusObjectSkeleton) SetObjectPath(objectPath string) {
	var _arg0 *C.GDBusObjectSkeleton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_object_skeleton_set_object_path(_arg0, _arg1)
}
