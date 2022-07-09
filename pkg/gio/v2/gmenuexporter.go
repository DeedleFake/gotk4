// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// ExportMenuModel exports menu on connection at object_path.
//
// The implemented D-Bus API should be considered private. It is subject to
// change in the future.
//
// An object path can only have one menu model exported on it. If this
// constraint is violated, the export will fail and 0 will be returned (with
// error set accordingly).
//
// You can unexport the menu model using g_dbus_connection_unexport_menu_model()
// with the return value of this function.
//
// The function takes the following parameters:
//
//    - objectPath d-Bus object path.
//    - menu: Model.
//
// The function returns the following values:
//
//    - guint: ID of the export (never zero), or 0 in case of failure.
//
func (connection *DBusConnection) ExportMenuModel(objectPath string, menu MenuModeller) (uint32, error) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_info := girepository.MustFind("Gio", "DBusConnection")
	_gret := _info.InvokeClassMethod("export_menu_model", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(connection)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(menu)

	var _guint uint32 // out
	var _goerr error  // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// UnexportMenuModel reverses the effect of a previous call to
// g_dbus_connection_export_menu_model().
//
// It is an error to call this function with an ID that wasn't returned from
// g_dbus_connection_export_menu_model() or to call it with the same ID more
// than once.
//
// The function takes the following parameters:
//
//    - exportId: ID from g_dbus_connection_export_menu_model().
//
func (connection *DBusConnection) UnexportMenuModel(exportId uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(exportId)

	_info := girepository.MustFind("Gio", "DBusConnection")
	_info.InvokeClassMethod("unexport_menu_model", _args[:], nil)

	runtime.KeepAlive(connection)
	runtime.KeepAlive(exportId)
}
