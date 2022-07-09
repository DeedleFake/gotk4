// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// BusAcquiredCallback: invoked when a connection to a message bus has been
// obtained.
type BusAcquiredCallback func(connection *DBusConnection, name string)

//export _gotk4_gio2_BusAcquiredCallback
func _gotk4_gio2_BusAcquiredCallback(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
	var fn BusAcquiredCallback
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(BusAcquiredCallback)
	}

	var _connection *DBusConnection // out
	var _name string                // out

	_connection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn(_connection, _name)
}

// BusNameAcquiredCallback: invoked when the name is acquired.
type BusNameAcquiredCallback func(connection *DBusConnection, name string)

//export _gotk4_gio2_BusNameAcquiredCallback
func _gotk4_gio2_BusNameAcquiredCallback(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
	var fn BusNameAcquiredCallback
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(BusNameAcquiredCallback)
	}

	var _connection *DBusConnection // out
	var _name string                // out

	_connection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn(_connection, _name)
}

// BusNameLostCallback: invoked when the name is lost or connection has been
// closed.
type BusNameLostCallback func(connection *DBusConnection, name string)

//export _gotk4_gio2_BusNameLostCallback
func _gotk4_gio2_BusNameLostCallback(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
	var fn BusNameLostCallback
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(BusNameLostCallback)
	}

	var _connection *DBusConnection // out
	var _name string                // out

	_connection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn(_connection, _name)
}

// BusUnownName stops owning a name.
//
// Note that there may still be D-Bus traffic to process (relating to owning and
// unowning the name) in the current thread-default Context after this function
// has returned. You should continue to iterate the Context until the Notify
// function passed to g_bus_own_name() is called, in order to avoid memory leaks
// through callbacks queued on the Context after it’s stopped being iterated.
//
// The function takes the following parameters:
//
//    - ownerId: identifier obtained from g_bus_own_name().
//
func BusUnownName(ownerId uint32) {
	var _args [1]girepository.Argument

	*(*C.guint)(unsafe.Pointer(&_args[0])) = C.guint(ownerId)

	_info := girepository.MustFind("Gio", "bus_unown_name")
	_info.Invoke(_args[:], nil)

	runtime.KeepAlive(ownerId)
}
