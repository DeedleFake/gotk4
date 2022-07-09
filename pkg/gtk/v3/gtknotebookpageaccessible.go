// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeNotebookPageAccessible returns the GType for the type NotebookPageAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeNotebookPageAccessible() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "NotebookPageAccessible").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalNotebookPageAccessible)
	return gtype
}

// NotebookPageAccessibleOverrider contains methods that are overridable.
type NotebookPageAccessibleOverrider interface {
}

type NotebookPageAccessible struct {
	_ [0]func() // equal guard
	atk.ObjectClass

	*coreglib.Object
	atk.Component
}

var (
	_ coreglib.Objector = (*NotebookPageAccessible)(nil)
)

func classInitNotebookPageAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNotebookPageAccessible(obj *coreglib.Object) *NotebookPageAccessible {
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(notebook).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "NotebookPageAccessible")
	_gret := _info.InvokeClassMethod("new_NotebookPageAccessible", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _notebookPageAccessible *NotebookPageAccessible // out

	_notebookPageAccessible = wrapNotebookPageAccessible(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _notebookPageAccessible
}

func (page *NotebookPageAccessible) Invalidate() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "NotebookPageAccessible")
	_info.InvokeClassMethod("invalidate", _args[:], nil)

	runtime.KeepAlive(page)
}
