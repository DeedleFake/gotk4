// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_CellRenderer_ConnectEditingCanceled
func _gotk4_gtk4_CellRenderer_ConnectEditingCanceled(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk4_CellRenderer_ConnectEditingStarted
func _gotk4_gtk4_CellRenderer_ConnectEditingStarted(arg0 C.gpointer, arg1 *C.GtkCellEditable, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(editable CellEditabler, path string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(editable CellEditabler, path string))
	}

	var _editable CellEditabler // out
	var _path string            // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellEditabler is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellEditabler)
			return ok
		})
		rv, ok := casted.(CellEditabler)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellEditabler")
		}
		_editable = rv
	}
	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_editable, _path)
}
