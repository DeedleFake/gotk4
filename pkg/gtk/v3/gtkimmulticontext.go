// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeIMMulticontext returns the GType for the type IMMulticontext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeIMMulticontext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "IMMulticontext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalIMMulticontext)
	return gtype
}

// IMMulticontextOverrider contains methods that are overridable.
type IMMulticontextOverrider interface {
}

type IMMulticontext struct {
	_ [0]func() // equal guard
	IMContext
}

var (
	_ IMContexter = (*IMMulticontext)(nil)
)

func classInitIMMulticontexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapIMMulticontext(obj *coreglib.Object) *IMMulticontext {
	return &IMMulticontext{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMMulticontext(p uintptr) (interface{}, error) {
	return wrapIMMulticontext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIMMulticontext creates a new IMMulticontext.
//
// The function returns the following values:
//
//    - imMulticontext: new IMMulticontext.
//
func NewIMMulticontext() *IMMulticontext {
	_info := girepository.MustFind("Gtk", "IMMulticontext")
	_gret := _info.InvokeClassMethod("new_IMMulticontext", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _imMulticontext *IMMulticontext // out

	_imMulticontext = wrapIMMulticontext(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imMulticontext
}

// AppendMenuitems: add menuitems for various available input methods to a menu;
// the menuitems, when selected, will switch the input method for the context
// and the global default input method.
//
// Deprecated: It is better to use the system-wide input method framework for
// changing input methods. Modern desktop shells offer on-screen displays for
// this that can triggered with a keyboard shortcut, e.g. Super-Space.
//
// The function takes the following parameters:
//
//    - menushell: MenuShell.
//
func (context *IMMulticontext) AppendMenuitems(menushell MenuSheller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menushell).Native()))

	_info := girepository.MustFind("Gtk", "IMMulticontext")
	_info.InvokeClassMethod("append_menuitems", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(menushell)
}

// ContextID gets the id of the currently active slave of the context.
//
// The function returns the following values:
//
//    - utf8: id of the currently active slave.
//
func (context *IMMulticontext) ContextID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gtk", "IMMulticontext")
	_gret := _info.InvokeClassMethod("get_context_id", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetContextID sets the context id for context.
//
// This causes the currently active slave of context to be replaced by the slave
// corresponding to the new context id.
//
// The function takes the following parameters:
//
//    - contextId: id to use.
//
func (context *IMMulticontext) SetContextID(contextId string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(contextId)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "IMMulticontext")
	_info.InvokeClassMethod("set_context_id", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(contextId)
}
