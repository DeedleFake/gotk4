// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// PaperSizeGetDefault returns the name of the default paper size, which depends
// on the current locale.
//
// The function returns the following values:
//
//    - utf8: name of the default paper size. The string is owned by GTK+ and
//      should not be modified.
//
func PaperSizeGetDefault() string {
	var _cret *C.gchar // in

	_cret = C.gtk_paper_size_get_default()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
