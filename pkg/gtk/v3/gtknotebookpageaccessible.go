// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeNotebookPageAccessible = coreglib.Type(C.gtk_notebook_page_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNotebookPageAccessible, F: marshalNotebookPageAccessible},
	})
}

// NotebookPageAccessibleOverrides contains methods that are overridable.
type NotebookPageAccessibleOverrides struct {
}

func defaultNotebookPageAccessibleOverrides(v *NotebookPageAccessible) NotebookPageAccessibleOverrides {
	return NotebookPageAccessibleOverrides{}
}

type NotebookPageAccessible struct {
	_ [0]func() // equal guard
	atk.AtkObject

	*coreglib.Object
	atk.Component
}

var (
	_ coreglib.Objector = (*NotebookPageAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*NotebookPageAccessible, *NotebookPageAccessibleClass, NotebookPageAccessibleOverrides](
		GTypeNotebookPageAccessible,
		initNotebookPageAccessibleClass,
		wrapNotebookPageAccessible,
		defaultNotebookPageAccessibleOverrides,
	)
}

func initNotebookPageAccessibleClass(gclass unsafe.Pointer, overrides NotebookPageAccessibleOverrides, classInitFunc func(*NotebookPageAccessibleClass)) {
	if classInitFunc != nil {
		class := (*NotebookPageAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapNotebookPageAccessible(obj *coreglib.Object) *NotebookPageAccessible {
	return &NotebookPageAccessible{
		AtkObject: atk.AtkObject{
			Object: obj,
		},
		Object: obj,
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalNotebookPageAccessible(p uintptr) (interface{}, error) {
	return wrapNotebookPageAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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

	_arg1 = (*C.GtkNotebookAccessible)(unsafe.Pointer(coreglib.InternObject(notebook).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.gtk_notebook_page_accessible_new(_arg1, _arg2)
	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _notebookPageAccessible *NotebookPageAccessible // out

	_notebookPageAccessible = wrapNotebookPageAccessible(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _notebookPageAccessible
}

func (page *NotebookPageAccessible) Invalidate() {
	var _arg0 *C.GtkNotebookPageAccessible // out

	_arg0 = (*C.GtkNotebookPageAccessible)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	C.gtk_notebook_page_accessible_invalidate(_arg0)
	runtime.KeepAlive(page)
}

// NotebookPageAccessibleClass: instance of this type is always passed by
// reference.
type NotebookPageAccessibleClass struct {
	*notebookPageAccessibleClass
}

// notebookPageAccessibleClass is the struct that's finalized.
type notebookPageAccessibleClass struct {
	native *C.GtkNotebookPageAccessibleClass
}

func (n *NotebookPageAccessibleClass) ParentClass() *atk.ObjectClass {
	valptr := &n.native.parent_class
	var _v *atk.ObjectClass // out
	_v = (*atk.ObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
