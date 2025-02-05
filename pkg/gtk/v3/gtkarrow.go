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
	GTypeArrow = coreglib.Type(C.gtk_arrow_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeArrow, F: marshalArrow},
	})
}

// ArrowOverrides contains methods that are overridable.
type ArrowOverrides struct {
}

func defaultArrowOverrides(v *Arrow) ArrowOverrides {
	return ArrowOverrides{}
}

// Arrow should be used to draw simple arrows that need to point in one of the
// four cardinal directions (up, down, left, or right). The style of the arrow
// can be one of shadow in, shadow out, etched in, or etched out. Note that
// these directions and style types may be amended in versions of GTK+ to come.
//
// GtkArrow will fill any space alloted to it, but since it is inherited from
// Misc, it can be padded and/or aligned, to fill exactly the space the
// programmer desires.
//
// Arrows are created with a call to gtk_arrow_new(). The direction or style of
// an arrow can be changed after creation by using gtk_arrow_set().
//
// GtkArrow has been deprecated; you can simply use a Image with a suitable icon
// name, such as “pan-down-symbolic“. When replacing GtkArrow by an image, pay
// attention to the fact that GtkArrow is doing automatic flipping between
// K_ARROW_LEFT and K_ARROW_RIGHT, depending on the text direction. To get the
// same effect with an image, use the icon names “pan-start-symbolic“ and
// “pan-end-symbolic“, which react to the text direction.
type Arrow struct {
	_ [0]func() // equal guard
	Misc
}

var (
	_ Miscer = (*Arrow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Arrow, *ArrowClass, ArrowOverrides](
		GTypeArrow,
		initArrowClass,
		wrapArrow,
		defaultArrowOverrides,
	)
}

func initArrowClass(gclass unsafe.Pointer, overrides ArrowOverrides, classInitFunc func(*ArrowClass)) {
	if classInitFunc != nil {
		class := (*ArrowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapArrow(obj *coreglib.Object) *Arrow {
	return &Arrow{
		Misc: Misc{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalArrow(p uintptr) (interface{}, error) {
	return wrapArrow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewArrow creates a new Arrow widget.
//
// Deprecated: Use a Image with a suitable icon.
//
// The function takes the following parameters:
//
//    - arrowType: valid ArrowType.
//    - shadowType: valid ShadowType.
//
// The function returns the following values:
//
//    - arrow: new Arrow widget.
//
func NewArrow(arrowType ArrowType, shadowType ShadowType) *Arrow {
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out
	var _cret *C.GtkWidget    // in

	_arg1 = C.GtkArrowType(arrowType)
	_arg2 = C.GtkShadowType(shadowType)

	_cret = C.gtk_arrow_new(_arg1, _arg2)
	runtime.KeepAlive(arrowType)
	runtime.KeepAlive(shadowType)

	var _arrow *Arrow // out

	_arrow = wrapArrow(coreglib.Take(unsafe.Pointer(_cret)))

	return _arrow
}

// Set sets the direction and style of the Arrow, arrow.
//
// Deprecated: Use a Image with a suitable icon.
//
// The function takes the following parameters:
//
//    - arrowType: valid ArrowType.
//    - shadowType: valid ShadowType.
//
func (arrow *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	var _arg0 *C.GtkArrow     // out
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out

	_arg0 = (*C.GtkArrow)(unsafe.Pointer(coreglib.InternObject(arrow).Native()))
	_arg1 = C.GtkArrowType(arrowType)
	_arg2 = C.GtkShadowType(shadowType)

	C.gtk_arrow_set(_arg0, _arg1, _arg2)
	runtime.KeepAlive(arrow)
	runtime.KeepAlive(arrowType)
	runtime.KeepAlive(shadowType)
}

// ArrowClass: instance of this type is always passed by reference.
type ArrowClass struct {
	*arrowClass
}

// arrowClass is the struct that's finalized.
type arrowClass struct {
	native *C.GtkArrowClass
}

func (a *ArrowClass) ParentClass() *MiscClass {
	valptr := &a.native.parent_class
	var _v *MiscClass // out
	_v = (*MiscClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
