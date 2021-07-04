// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_interface_skeleton_get_type()), F: marshalDBusInterfaceSkeleton},
	})
}

// DBusInterfaceSkeleton: abstract base class for D-Bus interfaces on the
// service side.
type DBusInterfaceSkeleton interface {
	DBusInterface

	ExportDBusInterfaceSkeleton(connection DBusConnection, objectPath string) error

	FlushDBusInterfaceSkeleton()

	Connection() DBusConnection

	Flags() DBusInterfaceSkeletonFlags

	GetInfo() *DBusInterfaceInfo

	ObjectPath() string

	Properties() *glib.Variant

	HasConnectionDBusInterfaceSkeleton(connection DBusConnection) bool

	SetFlagsDBusInterfaceSkeleton(flags DBusInterfaceSkeletonFlags)

	UnexportDBusInterfaceSkeleton()

	UnexportFromConnectionDBusInterfaceSkeleton(connection DBusConnection)
}

// dBusInterfaceSkeleton implements the DBusInterfaceSkeleton class.
type dBusInterfaceSkeleton struct {
	gextras.Objector
}

// WrapDBusInterfaceSkeleton wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusInterfaceSkeleton(obj *externglib.Object) DBusInterfaceSkeleton {
	return dBusInterfaceSkeleton{
		Objector: obj,
	}
}

func marshalDBusInterfaceSkeleton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusInterfaceSkeleton(obj), nil
}

func (i dBusInterfaceSkeleton) ExportDBusInterfaceSkeleton(connection DBusConnection, objectPath string) error {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out
	var _arg2 *C.gchar                  // out
	var _cerr *C.GError                 // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	_arg2 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_dbus_interface_skeleton_export(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (i dBusInterfaceSkeleton) FlushDBusInterfaceSkeleton() {
	var _arg0 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))

	C.g_dbus_interface_skeleton_flush(_arg0)
}

func (i dBusInterfaceSkeleton) Connection() DBusConnection {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GDBusConnection        // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))

	_cret = C.g_dbus_interface_skeleton_get_connection(_arg0)

	var _dBusConnection DBusConnection // out

	_dBusConnection = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(DBusConnection)

	return _dBusConnection
}

func (i dBusInterfaceSkeleton) Flags() DBusInterfaceSkeletonFlags {
	var _arg0 *C.GDBusInterfaceSkeleton     // out
	var _cret C.GDBusInterfaceSkeletonFlags // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))

	_cret = C.g_dbus_interface_skeleton_get_flags(_arg0)

	var _dBusInterfaceSkeletonFlags DBusInterfaceSkeletonFlags // out

	_dBusInterfaceSkeletonFlags = DBusInterfaceSkeletonFlags(_cret)

	return _dBusInterfaceSkeletonFlags
}

func (i dBusInterfaceSkeleton) GetInfo() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GDBusInterfaceInfo     // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))

	_cret = C.g_dbus_interface_skeleton_get_info(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))

	return _dBusInterfaceInfo
}

func (i dBusInterfaceSkeleton) ObjectPath() string {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))

	_cret = C.g_dbus_interface_skeleton_get_object_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i dBusInterfaceSkeleton) Properties() *glib.Variant {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _cret *C.GVariant               // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))

	_cret = C.g_dbus_interface_skeleton_get_properties(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_variant, func(v **glib.Variant) {
		C.free(unsafe.Pointer(v))
	})

	return _variant
}

func (i dBusInterfaceSkeleton) HasConnectionDBusInterfaceSkeleton(connection DBusConnection) bool {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))

	_cret = C.g_dbus_interface_skeleton_has_connection(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i dBusInterfaceSkeleton) SetFlagsDBusInterfaceSkeleton(flags DBusInterfaceSkeletonFlags) {
	var _arg0 *C.GDBusInterfaceSkeleton     // out
	var _arg1 C.GDBusInterfaceSkeletonFlags // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))
	_arg1 = C.GDBusInterfaceSkeletonFlags(flags)

	C.g_dbus_interface_skeleton_set_flags(_arg0, _arg1)
}

func (i dBusInterfaceSkeleton) UnexportDBusInterfaceSkeleton() {
	var _arg0 *C.GDBusInterfaceSkeleton // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))

	C.g_dbus_interface_skeleton_unexport(_arg0)
}

func (i dBusInterfaceSkeleton) UnexportFromConnectionDBusInterfaceSkeleton(connection DBusConnection) {
	var _arg0 *C.GDBusInterfaceSkeleton // out
	var _arg1 *C.GDBusConnection        // out

	_arg0 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))

	C.g_dbus_interface_skeleton_unexport_from_connection(_arg0, _arg1)
}

func (i dBusInterfaceSkeleton) DupObject() DBusObject {
	return WrapDBusInterface(gextras.InternObject(i)).DupObject()
}

func (i dBusInterfaceSkeleton) Info() *DBusInterfaceInfo {
	return WrapDBusInterface(gextras.InternObject(i)).Info()
}

func (i dBusInterfaceSkeleton) SetObject(object DBusObject) {
	WrapDBusInterface(gextras.InternObject(i)).SetObject(object)
}
