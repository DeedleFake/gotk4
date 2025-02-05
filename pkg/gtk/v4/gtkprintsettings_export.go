// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_PrintSettingsFunc
func _gotk4_gtk4_PrintSettingsFunc(arg1 *C.char, arg2 *C.char, arg3 C.gpointer) {
	var fn PrintSettingsFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PrintSettingsFunc)
	}

	var _key string   // out
	var _value string // out

	_key = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_value = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn(_key, _value)
}
