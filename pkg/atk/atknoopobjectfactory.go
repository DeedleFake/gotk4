// Code generated by girgen. DO NOT EDIT.

package atk

import (
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
	GTypeNoOpObjectFactory = coreglib.Type(C.atk_no_op_object_factory_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNoOpObjectFactory, F: marshalNoOpObjectFactory},
	})
}

// NoOpObjectFactoryOverrides contains methods that are overridable.
type NoOpObjectFactoryOverrides struct {
}

func defaultNoOpObjectFactoryOverrides(v *NoOpObjectFactory) NoOpObjectFactoryOverrides {
	return NoOpObjectFactoryOverrides{}
}

// NoOpObjectFactory: atkObjectFactory which creates an AtkNoOpObject. An
// instance of this is created by an AtkRegistry if no factory type has not been
// specified to create an accessible object of a particular type.
type NoOpObjectFactory struct {
	_ [0]func() // equal guard
	ObjectFactory
}

var (
	_ coreglib.Objector = (*NoOpObjectFactory)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*NoOpObjectFactory, *NoOpObjectFactoryClass, NoOpObjectFactoryOverrides](
		GTypeNoOpObjectFactory,
		initNoOpObjectFactoryClass,
		wrapNoOpObjectFactory,
		defaultNoOpObjectFactoryOverrides,
	)
}

func initNoOpObjectFactoryClass(gclass unsafe.Pointer, overrides NoOpObjectFactoryOverrides, classInitFunc func(*NoOpObjectFactoryClass)) {
	if classInitFunc != nil {
		class := (*NoOpObjectFactoryClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapNoOpObjectFactory(obj *coreglib.Object) *NoOpObjectFactory {
	return &NoOpObjectFactory{
		ObjectFactory: ObjectFactory{
			Object: obj,
		},
	}
}

func marshalNoOpObjectFactory(p uintptr) (interface{}, error) {
	return wrapNoOpObjectFactory(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNoOpObjectFactory creates an instance of an ObjectFactory which generates
// primitive (non-functioning) Objects.
//
// The function returns the following values:
//
//    - noOpObjectFactory: instance of an ObjectFactory.
//
func NewNoOpObjectFactory() *NoOpObjectFactory {
	var _cret *C.AtkObjectFactory // in

	_cret = C.atk_no_op_object_factory_new()

	var _noOpObjectFactory *NoOpObjectFactory // out

	_noOpObjectFactory = wrapNoOpObjectFactory(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noOpObjectFactory
}

// NoOpObjectFactoryClass: instance of this type is always passed by reference.
type NoOpObjectFactoryClass struct {
	*noOpObjectFactoryClass
}

// noOpObjectFactoryClass is the struct that's finalized.
type noOpObjectFactoryClass struct {
	native *C.AtkNoOpObjectFactoryClass
}

func (n *NoOpObjectFactoryClass) ParentClass() *ObjectFactoryClass {
	valptr := &n.native.parent_class
	var _v *ObjectFactoryClass // out
	_v = (*ObjectFactoryClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
