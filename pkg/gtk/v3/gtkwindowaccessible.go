// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeWindowAccessible = coreglib.Type(C.gtk_window_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWindowAccessible, F: marshalWindowAccessible},
	})
}

// WindowAccessibleOverrides contains methods that are overridable.
type WindowAccessibleOverrides struct {
}

func defaultWindowAccessibleOverrides(v *WindowAccessible) WindowAccessibleOverrides {
	return WindowAccessibleOverrides{}
}

type WindowAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Window
}

var (
	_ coreglib.Objector = (*WindowAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WindowAccessible, *WindowAccessibleClass, WindowAccessibleOverrides](
		GTypeWindowAccessible,
		initWindowAccessibleClass,
		wrapWindowAccessible,
		defaultWindowAccessibleOverrides,
	)
}

func initWindowAccessibleClass(gclass unsafe.Pointer, overrides WindowAccessibleOverrides, classInitFunc func(*WindowAccessibleClass)) {
	if classInitFunc != nil {
		class := (*WindowAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWindowAccessible(obj *coreglib.Object) *WindowAccessible {
	return &WindowAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
		Window: atk.Window{
			AtkObject: atk.AtkObject{
				Object: obj,
			},
		},
	}
}

func marshalWindowAccessible(p uintptr) (interface{}, error) {
	return wrapWindowAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WindowAccessibleClass: instance of this type is always passed by reference.
type WindowAccessibleClass struct {
	*windowAccessibleClass
}

// windowAccessibleClass is the struct that's finalized.
type windowAccessibleClass struct {
	native *C.GtkWindowAccessibleClass
}

func (w *WindowAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &w.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
