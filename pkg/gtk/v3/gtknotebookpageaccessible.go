// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_notebook_page_accessible_get_type()), F: marshalNotebookPageAccessibler},
	})
}

type NotebookPageAccessible struct {
	atk.ObjectClass

	*externglib.Object
	atk.Component
}

var (
	_ externglib.Objector = (*NotebookPageAccessible)(nil)
)

func wrapNotebookPageAccessible(obj *externglib.Object) *NotebookPageAccessible {
	return &NotebookPageAccessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		Object: obj,
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalNotebookPageAccessibler(p uintptr) (interface{}, error) {
	return wrapNotebookPageAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
//    - notebook
//    - child
//
// The function returns the following values:
//
func NewNotebookPageAccessible(notebook *NotebookAccessible, child Widgetter) *NotebookPageAccessible {
	var _arg1 *C.GtkNotebookAccessible // out
	var _arg2 *C.GtkWidget             // out
	var _cret *C.AtkObject             // in

	_arg1 = (*C.GtkNotebookAccessible)(unsafe.Pointer(notebook.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_page_accessible_new(_arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _notebookPageAccessible *NotebookPageAccessible // out

	_notebookPageAccessible = wrapNotebookPageAccessible(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _notebookPageAccessible
}

func (page *NotebookPageAccessible) Invalidate() {
	var _arg0 *C.GtkNotebookPageAccessible // out

	_arg0 = (*C.GtkNotebookPageAccessible)(unsafe.Pointer(page.Native()))

	C.gtk_notebook_page_accessible_invalidate(_arg0)
	runtime.KeepAlive(page)
}
