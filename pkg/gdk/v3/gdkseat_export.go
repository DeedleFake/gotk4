// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

//export _gotk4_gdk3_Seat_ConnectDeviceAdded
func _gotk4_gdk3_Seat_ConnectDeviceAdded(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

//export _gotk4_gdk3_Seat_ConnectDeviceRemoved
func _gotk4_gdk3_Seat_ConnectDeviceRemoved(arg0 C.gpointer, arg1 *C.GdkDevice, arg2 C.guintptr) {
	var f func(device Devicer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(device Devicer))
	}

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	f(_device)
}

//export _gotk4_gdk3_Seat_ConnectToolAdded
func _gotk4_gdk3_Seat_ConnectToolAdded(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

//export _gotk4_gdk3_Seat_ConnectToolRemoved
func _gotk4_gdk3_Seat_ConnectToolRemoved(arg0 C.gpointer, arg1 *C.GdkDeviceTool, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}
