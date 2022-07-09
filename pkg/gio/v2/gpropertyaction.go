// Code generated by girgen. DO NOT EDIT.

package gio

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

// NewPropertyAction creates a #GAction corresponding to the value of property
// property_name on object.
//
// The property must be existent and readable and writable (and not
// construct-only).
//
// This function takes a reference on object and doesn't release it until the
// action is destroyed.
//
// The function takes the following parameters:
//
//    - name of the action to create.
//    - object that has the property to wrap.
//    - propertyName: name of the property.
//
// The function returns the following values:
//
//    - propertyAction: new Action.
//
func NewPropertyAction(name string, object *coreglib.Object, propertyName string) *PropertyAction {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[0]))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(object.Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_args[2]))

	_info := girepository.MustFind("Gio", "PropertyAction")
	_gret := _info.InvokeClassMethod("new_PropertyAction", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)
	runtime.KeepAlive(object)
	runtime.KeepAlive(propertyName)

	var _propertyAction *PropertyAction // out

	_propertyAction = wrapPropertyAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _propertyAction
}
