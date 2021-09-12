// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gio.h"
import "C"

// DBusMenuModelGet obtains a BusMenuModel for the menu model which is exported
// at the given bus_name and object_path.
//
// The thread default main context is taken at the time of this call. All
// signals on the menu model (and any linked models) are reported with respect
// to this context. All calls on the returned menu model (and linked models)
// must also originate from this same context, with the thread default main
// context unchanged.
func DBusMenuModelGet(connection *DBusConnection, busName string, objectPath string) *DBusMenuModel {
	var _arg1 *C.GDBusConnection // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.gchar           // out
	var _cret *C.GDBusMenuModel  // in

	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	if busName != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(busName)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_dbus_menu_model_get(_arg1, _arg2, _arg3)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(busName)
	runtime.KeepAlive(objectPath)

	var _dBusMenuModel *DBusMenuModel // out

	_dBusMenuModel = wrapDBusMenuModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusMenuModel
}
