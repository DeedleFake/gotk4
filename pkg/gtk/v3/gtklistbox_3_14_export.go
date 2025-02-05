// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_ListBoxForEachFunc
func _gotk4_gtk3_ListBoxForEachFunc(arg1 *C.GtkListBox, arg2 *C.GtkListBoxRow, arg3 C.gpointer) {
	var fn ListBoxForEachFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxForEachFunc)
	}

	var _box *ListBox    // out
	var _row *ListBoxRow // out

	_box = wrapListBox(coreglib.Take(unsafe.Pointer(arg1)))
	_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg2)))

	fn(_box, _row)
}
