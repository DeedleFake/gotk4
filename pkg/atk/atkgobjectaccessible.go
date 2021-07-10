// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_gobject_accessible_get_type()), F: marshalGObjectAccessible},
	})
}

// GObjectAccessible: this object class is derived from AtkObject. It can be
// used as a basis for implementing accessible objects for GObjects which are
// not derived from GtkWidget. One example of its use is in providing an
// accessible object for GnomeCanvasItem in the GAIL library.
type GObjectAccessible interface {
	gextras.Objector

	// Object gets the GObject for which @obj is the accessible object.
	Object() *externglib.Object
}

// GObjectAccessibleClass implements the GObjectAccessible interface.
type GObjectAccessibleClass struct {
	ObjectClass
}

var _ GObjectAccessible = (*GObjectAccessibleClass)(nil)

func wrapGObjectAccessible(obj *externglib.Object) GObjectAccessible {
	return &GObjectAccessibleClass{
		ObjectClass: ObjectClass{
			Object: obj,
		},
	}
}

func marshalGObjectAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGObjectAccessible(obj), nil
}

// Object gets the GObject for which @obj is the accessible object.
func (obj *GObjectAccessibleClass) Object() *externglib.Object {
	var _arg0 *C.AtkGObjectAccessible // out
	var _cret *C.GObject              // in

	_arg0 = (*C.AtkGObjectAccessible)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_gobject_accessible_get_object(_arg0)

	var _object *externglib.Object // out

	_object = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*externglib.Object)

	return _object
}
