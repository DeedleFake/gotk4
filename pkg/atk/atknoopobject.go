// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeNoOpObject = coreglib.Type(C.atk_no_op_object_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNoOpObject, F: marshalNoOpObject},
	})
}

// NoOpObjectOverrides contains methods that are overridable.
type NoOpObjectOverrides struct {
}

func defaultNoOpObjectOverrides(v *NoOpObject) NoOpObjectOverrides {
	return NoOpObjectOverrides{}
}

// NoOpObject is an AtkObject which purports to implement all ATK interfaces. It
// is the type of AtkObject which is created if an accessible object is
// requested for an object type for which no factory type is specified.
type NoOpObject struct {
	_ [0]func() // equal guard
	AtkObject

	*coreglib.Object
	Action
	Component
	Document
	EditableText
	Hypertext
	Image
	Selection
	Table
	TableCell
	Text
	Value
	Window
}

var (
	_ coreglib.Objector = (*NoOpObject)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*NoOpObject, *NoOpObjectClass, NoOpObjectOverrides](
		GTypeNoOpObject,
		initNoOpObjectClass,
		wrapNoOpObject,
		defaultNoOpObjectOverrides,
	)
}

func initNoOpObjectClass(gclass unsafe.Pointer, overrides NoOpObjectOverrides, classInitFunc func(*NoOpObjectClass)) {
	if classInitFunc != nil {
		class := (*NoOpObjectClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapNoOpObject(obj *coreglib.Object) *NoOpObject {
	return &NoOpObject{
		AtkObject: AtkObject{
			Object: obj,
		},
		Object: obj,
		Action: Action{
			Object: obj,
		},
		Component: Component{
			Object: obj,
		},
		Document: Document{
			Object: obj,
		},
		EditableText: EditableText{
			Object: obj,
		},
		Hypertext: Hypertext{
			Object: obj,
		},
		Image: Image{
			Object: obj,
		},
		Selection: Selection{
			Object: obj,
		},
		Table: Table{
			Object: obj,
		},
		TableCell: TableCell{
			AtkObject: AtkObject{
				Object: obj,
			},
		},
		Text: Text{
			Object: obj,
		},
		Value: Value{
			Object: obj,
		},
		Window: Window{
			AtkObject: AtkObject{
				Object: obj,
			},
		},
	}
}

func marshalNoOpObject(p uintptr) (interface{}, error) {
	return wrapNoOpObject(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNoOpObject provides a default (non-functioning stub) Object. Application
// maintainers should not use this method.
//
// The function takes the following parameters:
//
//    - obj: #GObject.
//
// The function returns the following values:
//
//    - noOpObject: default (non-functioning stub) Object.
//
func NewNoOpObject(obj *coreglib.Object) *NoOpObject {
	var _arg1 *C.GObject   // out
	var _cret *C.AtkObject // in

	_arg1 = (*C.GObject)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_no_op_object_new(_arg1)
	runtime.KeepAlive(obj)

	var _noOpObject *NoOpObject // out

	_noOpObject = wrapNoOpObject(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noOpObject
}

// NoOpObjectClass: instance of this type is always passed by reference.
type NoOpObjectClass struct {
	*noOpObjectClass
}

// noOpObjectClass is the struct that's finalized.
type noOpObjectClass struct {
	native *C.AtkNoOpObjectClass
}

func (n *NoOpObjectClass) ParentClass() *ObjectClass {
	valptr := &n.native.parent_class
	var _v *ObjectClass // out
	_v = (*ObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
